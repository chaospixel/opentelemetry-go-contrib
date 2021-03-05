// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Based on opentelemetry-go-contrib/instrumentation/gihub.com/gin-gonic/gin/otelgin

package otelswag

import (
	"net/http"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/semconv"
)

const (
	tracerKey  = "otel-go-contrib-tracer"
	tracerName = "go.opentelemetry.io/contrib/instrumentation/github.com/go-swagger/go-swagger/otelgo-swagger"
)

// Middleware returns middleware that will trace incoming requests.
// The service parameter should describe the name of the (virtual)
// server handling the request.
func Middleware(service string, opts ...Option) func(http.Handler) http.Handler {
	cfg := config{}
	for _, opt := range opts {
		opt(&cfg)
	}
	if cfg.TracerProvider == nil {
		cfg.TracerProvider = otel.GetTracerProvider()
	}
	tracer := cfg.TracerProvider.Tracer(
		tracerName,
		oteltrace.WithInstrumentationVersion(otelcontrib.SemVersion()),
	)
	if cfg.Propagators == nil {
		cfg.Propagators = otel.GetTextMapPropagator()
	}
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// c.Set(tracerKey, tracer)

			savedCtx := r.Context()
			defer func() {
				r = r.WithContext(savedCtx)
			}()
			ctx := cfg.Propagators.Extract(savedCtx, r.Header)
			opts := []oteltrace.SpanOption{
				oteltrace.WithAttributes(semconv.NetAttributesFromHTTPRequest("tcp", r)...),
				oteltrace.WithAttributes(semconv.EndUserAttributesFromHTTPRequest(r)...),
				oteltrace.WithAttributes(semconv.HTTPServerAttributesFromHTTPRequest(service, r.URL.Path, r)...),
				oteltrace.WithSpanKind(oteltrace.SpanKindServer),
			}
			spanName := r.URL.Path
			ctx, span := tracer.Start(ctx, spanName, opts...)
			defer span.End()

			// pass the span through the request context
			r = r.WithContext(ctx)

			// serve the request to the next middleware
			next.ServeHTTP(w, r)

			// TODO Not really sure how I'll pick up the status outside of Gin...
			status := w.Status()
			attrs := semconv.HTTPAttributesFromHTTPStatusCode(status)
			spanStatus, spanMessage := semconv.SpanStatusFromHTTPStatusCode(status)
			span.SetAttributes(attrs...)
			span.SetStatus(spanStatus, spanMessage)
			// if len(c.Errors) > 0 {
			// 	span.SetAttributes(label.String("gin.errors", c.Errors.String()))
			// }
		})
	}
}
