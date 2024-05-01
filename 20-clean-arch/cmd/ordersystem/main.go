package main

import (
	"database/sql"
	"fmt"
	"net"
	"net/http"

	gqlhandler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/allanmaral/go-expert/20-clean-arch/configs"
	"github.com/allanmaral/go-expert/20-clean-arch/internal/infra/eventhandlers"
	"github.com/allanmaral/go-expert/20-clean-arch/internal/infra/graphql"
	"github.com/allanmaral/go-expert/20-clean-arch/internal/infra/grpc/pb"
	"github.com/allanmaral/go-expert/20-clean-arch/internal/infra/grpc/service"
	"github.com/allanmaral/go-expert/20-clean-arch/internal/infra/web"
	"github.com/allanmaral/go-expert/20-clean-arch/pkg/events"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := sql.Open(config.DBDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	amqpchan := getRabbitMQChannel()

	eventDispatcher := events.NewEventDispatcher()
	eventDispatcher.Register("OrderCreated", eventhandlers.NewOrderCreatedHandlerAMQP(amqpchan))

	createOrderUseCase := NewCreateOrderUseCase(db, eventDispatcher)

	webserver := web.NewWebServer(config.WebServerPort)
	webOrderHandler := NewOrderHandlerWeb(db, eventDispatcher)
	webserver.AddHandler("/order", webOrderHandler.Create)
	fmt.Println("Starting web server on port", config.WebServerPort)
	go webserver.Start()

	grpcServer := grpc.NewServer()
	createOrderService := service.NewOrderService(createOrderUseCase)
	pb.RegisterOrderServiceServer(grpcServer, createOrderService)
	reflection.Register(grpcServer)

	fmt.Println("Starting gRPC server on port", config.GRPCServerPort)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", config.GRPCServerPort))
	if err != nil {
		panic(err)
	}
	go grpcServer.Serve(lis)

	srv := gqlhandler.NewDefaultServer(graphql.NewExecutableSchema(graphql.Config{
		Resolvers: &graphql.Resolver{CreateOrderUseCase: createOrderUseCase}}))
	http.Handle("/", playground.Handler("GraphiQL playground", "/query"))
	http.Handle("/query", srv)

	fmt.Println("Starting GraphQL server or port", config.GraphQLServerPort)
	http.ListenAndServe(":"+config.GraphQLServerPort, nil)
}

func getRabbitMQChannel() *amqp.Channel {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		panic(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	return ch
}
