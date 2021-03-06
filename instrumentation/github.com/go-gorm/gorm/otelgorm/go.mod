module go.opentelemetry.io/contrib/instrumentation/github.com/go-gorm/gorm/otelgorm

go 1.14

replace go.opentelemetry.io/contrib => ../../../../..

require (
	github.com/mattn/go-sqlite3 v1.14.6 // indirect
	github.com/stretchr/testify v1.6.1
	go.opentelemetry.io/contrib v0.16.0
	go.opentelemetry.io/otel v0.16.0
	go.opentelemetry.io/otel/exporters/stdout v0.16.0
	go.opentelemetry.io/otel/sdk v0.16.0
	gorm.io/driver/sqlite v1.1.4
	gorm.io/gorm v1.20.12
)
