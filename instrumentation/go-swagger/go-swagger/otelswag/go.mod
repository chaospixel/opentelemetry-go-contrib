module go.opentelemetry.io/contrib/instrumentation/github.com/go-swagger/go-swagger/otelswag

go 1.14

replace (
	go.opentelemetry.io/contrib => ../../../../../
	go.opentelemetry.io/contrib/propagators => ../../../../../propagators
)

require (
	go.opentelemetry.io/otel v0.17.0
	go.opentelemetry.io/otel/trace v0.17.0
)
