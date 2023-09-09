module github.com/allanmaral/go-expert/07-packaging/04-working-with-go-mod-replace/sys

go 1.20

replace github.com/allanmaral/go-expert/07-packaging/04-working-with-go-mod-replace/math => ../math

require github.com/allanmaral/go-expert/07-packaging/04-working-with-go-mod-replace/math v0.0.0-00010101000000-000000000000
