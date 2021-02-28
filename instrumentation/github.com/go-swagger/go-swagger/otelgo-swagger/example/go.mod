module go.opentelemetry.io/opentelemetry-go-contrib/instrumentation/github.com/go-swagger/go-swagger/otelgo-swagger/example

go 1.14

replace (
	go.opentelemetry.io/contrib => ../../../../../../
	go.opentelemetry.io/contrib/instrumentation/github.com/go-swagger/go-swagger/otelgo-swagger => ../
	go.opentelemetry.io/contrib/propagators => ../../../../../../propagators
)

require (
	github.com/go-openapi/errors v0.20.0
	github.com/go-openapi/loads v0.20.2
	github.com/go-openapi/runtime v0.19.26
	github.com/go-openapi/spec v0.20.3
	github.com/go-openapi/strfmt v0.20.0
	github.com/go-openapi/swag v0.19.14
	github.com/go-openapi/validate v0.20.2
	github.com/google/go-cmp v0.5.4 // indirect
	github.com/jessevdk/go-flags v1.4.0
	github.com/mailru/easyjson v0.7.7 // indirect
	golang.org/x/net v0.0.0-20210226172049-e18ecbb05110
)
