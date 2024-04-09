package uow

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type RepositoryFactory func(tx *sql.Tx) interface{}

type UOW interface {
	Register(name string, fc RepositoryFactory)
	Unregister(name string)
	GetRepository(ctx context.Context, name string) (interface{}, error)
	Do(ctx context.Context, fn func(uow UOW) error) error
}

var (
	ErrTransationAlreadyStarted = errors.New("transation already started")
	ErrNoTransactionToRollback  = errors.New("no transaction to rollback")
	ErrRepositoryNotRegistered  = errors.New("repository not registered")
)

type uow struct {
	db           *sql.DB
	tx           *sql.Tx
	repositories map[string]RepositoryFactory
}

func NewUOW(ctx context.Context, db *sql.DB) UOW {
	return &uow{
		db:           db,
		repositories: make(map[string]RepositoryFactory),
	}
}

func (u *uow) Register(name string, fc RepositoryFactory) {
	u.repositories[name] = fc
}

func (u *uow) Unregister(name string) {
	delete(u.repositories, name)
}

func (u *uow) GetRepository(ctx context.Context, name string) (interface{}, error) {
	if u.tx == nil {
		tx, err := u.db.BeginTx(ctx, nil)
		if err != nil {
			return nil, err
		}
		u.tx = tx
	}

	repo := u.repositories[name]
	if repo == nil {
		return nil, ErrRepositoryNotRegistered
	}

	return repo(u.tx), nil
}

func (u *uow) Do(ctx context.Context, fn func(uow UOW) error) error {
	if u.tx != nil {
		return ErrTransationAlreadyStarted
	}

	tx, err := u.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	u.tx = tx

	err = fn(u)
	if err != nil {
		return u.rollback(err)
	}

	return u.commitOrRollback()
}

func (u *uow) commitOrRollback() error {
	err := u.tx.Commit()
	if err != nil {
		return u.rollback(err)
	}

	u.tx = nil
	return nil
}

func (u *uow) rollback(originalError error) error {
	if u.tx == nil {
		if originalError == nil {
			return ErrNoTransactionToRollback
		}
		return fmt.Errorf("error on rollback: %v, original error: %w", ErrNoTransactionToRollback, originalError)
	}

	tx := u.tx
	u.tx = nil

	err := tx.Rollback()
	if err != nil {
		if originalError == nil {
			return err
		}
		return fmt.Errorf("error on rollback: %v, original error: %w", err, originalError)
	}

	return nil
}
