# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Common Commands

### Build System
- `make generate` - Generate templates, types, and route handlers from WSDL files
- `make build` - Build the application binary (`./app`)
- `make run` - Run the built application 
- `make all` - Generate, build, and run in sequence
- `make clean` - Remove generated files and binary

### Direct Go Commands
- `go run ./cmd/build/main.go` - Run the code generator directly
- `go run ./cmd/server/main.go -config config/config.json` - Run server with config
- `go build -o app ./cmd/server/main.go` - Build server binary
- `go test ./...` - Run all tests
- `go mod download` - Download dependencies

### Development Flow
1. Always run `make generate` after modifying WSDL URLs or route configurations
2. Run `make build` to compile the application
3. The application requires a config file to run (default: `config/config.json`)

## Architecture Overview

### Core Components

**Build System (`core/build/`)**:
- `generators/wsdl2struct.go` - Extracts Go structs from WSDL files
- `generators/template_generator.go` - Generates parsing functions for each SOAP operation
- `generators/registry_generator.go` - Creates route handler registry from config
- `cmd/build/main.go` - Entry point for code generation

**Server (`core/server/`)**:
- `handler/handler.go` - Main HTTP handler that processes REST requests and forwards to SOAP
- `handler/pool.go` - Worker pool for concurrent request processing
- `soap/client.go` - HTTP client wrapper with logging for SOAP requests
- `wsdl/parser.go` - WSDL parsing utilities

**Configuration (`core/config/`)**:
- `config_types.go` - Configuration structs with JSON unmarshaling
- `config.schema.json` - JSON schema for configuration validation

**Generated Code (`pkg/generated/`)**:
- `*_parser.go` - Auto-generated SOAP response parsers for each operation
- `route_handler_registry.go` - Auto-generated route registry mapping paths to handlers

### Request Flow

1. HTTP request received by `handler.Handler.ServeHTTP`
2. Route lookup in generated registry (`pkg/generated/route_handler_registry.go`)
3. Request body parsed and applied to request template
4. SOAP request sent via `soap.Client` to configured endpoint
5. SOAP response parsed using generated parser functions
6. Response template applied to create JSON output
7. JSON response returned to client

### Template System

The application uses Go templates for both request and response transformation:
- **Request templates** (`config/templates/*-request.tmpl`): Transform JSON input to SOAP XML
- **Response templates** (`config/templates/*-response.tmpl`): Transform SOAP XML to JSON output
- Templates are compiled at generation time and embedded in generated parsers

### Code Generation Flow

1. `cmd/build/main.go` reads configuration file
2. For each route with `wsdl_url`:
   - Downloads and parses WSDL using `wsdl2struct.go`
   - Extracts Go struct definitions for SOAP types
   - Generates parser function using `template_generator.go`
3. Creates unified route registry using `registry_generator.go`
4. Generated files are placed in `pkg/generated/`

## Configuration

Routes are defined in `config/config.json` with these key fields:
- `path` - REST endpoint path
- `soap_endpoint` - Target SOAP service URL
- `soap_action` - SOAP action header value (used as operation name)
- `wsdl_url` - WSDL URL for type generation
- `request_template` - Path to request transformation template
- `response_template` - Path to response transformation template

## Key Dependencies

- `go.uber.org/zap` - Structured logging
- Standard library `encoding/xml` for SOAP parsing
- Standard library `text/template` for transformations
- Standard library `net/http` for server and client

## Development Notes

- The application must be built using `make generate` before running to ensure all generated code is up to date
- WSDL parsing and struct generation happens at build time, not runtime
- Each SOAP operation gets its own generated parser function
- Worker pools are used for concurrent request processing
- Graceful shutdown is implemented with 10-second timeout