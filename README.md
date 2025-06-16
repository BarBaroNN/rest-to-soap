# SOAP-to-REST Proxy Server

A high-performance SOAP-to-REST proxy server written in Go, designed for production use with minimal overhead and maximum throughput.

## Features

- JSON configuration for route mapping and server settings
- Quicktemplate-based SOAP request/response templating
- WSDL parsing and type generation via gowsdl
- Structured logging with zap
- High-performance request handling with worker pools
- Graceful shutdown support
- Production-ready error handling
- Prometheus metrics for monitoring
- Comprehensive benchmarking tools

## Requirements

- Go 1.24 or later
- Make (optional, for build scripts)
- gowsdl (for WSDL parsing)

## Quick Start

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/rest-to-soap.git
   cd rest-to-soap
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Build the server:
   ```bash
   go build -o rest-to-soap
   ```

4. Create a configuration file:
   ```bash
   cp config/config.example.json config/config.json
   # Edit config.json with your settings
   ```

5. Run the server:
   ```bash
   ./rest-to-soap -config config/config.json
   ```

## Configuration

The server is configured via a JSON file. See `config/config.schema.json` for the full schema and `config/config.example.json` for an example configuration.

Key configuration sections:
- `server`: Server settings (port, timeouts)
- `routes`: Route mappings (REST to SOAP)
- `logging`: Logging configuration

## Monitoring

The server exposes Prometheus metrics on port 9090 (configurable via `-metrics-port`). Available metrics:

- `soap_proxy_request_duration_seconds`: Request duration histogram
- `soap_proxy_requests_total`: Total request counter
- `soap_proxy_request_errors_total`: Error counter
- `soap_proxy_active_requests`: Active request gauge
- `soap_proxy_worker_pool_size`: Worker pool size
- `soap_proxy_worker_pool_usage`: Worker pool usage

## WSDL Support

The server can automatically parse WSDL definitions and generate Go types:

1. Add the WSDL URL to your route configuration:
   ```json
   {
     "path": "/api/soap/example",
     "wsdl_url": "http://example.com/service?wsdl"
   }
   ```

2. The server will:
   - Fetch and parse the WSDL
   - Generate Go types
   - Use type information for request/response validation

## Benchmarking

Run benchmarks to compare proxy vs direct calls:

```bash
go test -bench=. ./benchmarks
```

The benchmark will:
- Send concurrent requests to both proxy and direct endpoints
- Measure latency and throughput
- Calculate proxy overhead
- Generate detailed performance report

## Development

### Project Structure

```
.
├── benchmarks/        # Benchmarking tools
├── config/           # Configuration files and types
├── handlers/         # HTTP request handlers
├── metrics/          # Prometheus metrics
├── templates/        # Quicktemplate files
├── transport/        # HTTP client and transport
├── wsdl/            # WSDL parsing and types
├── main.go          # Application entry point
└── README.md        # This file
```

### Running Tests

```bash
go test ./...
```

### Running Benchmarks

```bash
go test -bench=. ./...
```

## Performance

The server is optimized for high throughput with:
- Worker pool for request handling
- Connection pooling
- Template caching
- Buffer pooling
- Efficient XML parsing
- WSDL type caching

## Contributing

1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

## License

MIT License - see LICENSE file for details
