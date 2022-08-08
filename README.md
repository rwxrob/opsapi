# OpenAPI 3.0 Code Generation Experiment

This is just a working example of how to generate Go Gin server code
using latest OpenAPI and OpenAPITools (which is java). Templates have
been overridden in simple ways to update to the latest Go 1.18
requirements and enable clean separation of all business logic for the
API into the `internal/api` package --- specifically, the `api.Handle`
function.
