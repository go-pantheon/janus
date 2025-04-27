# Janus

Janus is a high-performance gateway service framework based on microservice architecture, developed in Go. This framework aims to provide a scalable and reliable connection layer between clients and backend services, supporting various protocols and efficient request routing. Janus is a core component of the go-pantheon ecosystem, responsible for implementing client connection handling and request forwarding.

For more, please check out: [deepwiki/go-pantheon/janus](https://deepwiki.com/go-pantheon/janus)

## go-pantheon Ecosystem

**go-pantheon** is an out-of-the-box game server framework providing high-performance, highly available game server cluster solutions based on microservices architecture. Janus, as the gateway service, works alongside other core services to form a complete game service ecosystem:

- **Roma**: Game core logic services
- **Janus**: Gateway service for client connection handling and request forwarding
- **Lares**: Account service for user authentication and account management
- **Senate**: Backend management service providing operational interfaces

### Core Features

- 🚀 High-performance TCP connection management and protocol handling
- 🔄 Efficient request routing and load balancing
- 🛡️ Token-based authentication and handshake verification
- 📊 Real-time monitoring and distributed tracing
- 🔌 Plugin-based architecture for extensibility
- 📈 Scalable design supporting millions of concurrent connections
- 🔍 Protocol conversion between TCP and gRPC

### Service Layer Features

- TCP server for client connections
- gRPC client for backend service communication
- Protocol conversion and transformation
- Connection pooling and management
- Load balancing across Roma service instances
- Health checking and circuit breaking
- Traffic throttling and rate limiting

## System Architecture

The relationship between Janus and other go-pantheon components is illustrated below:

```
   (1)HTTPS Token Request
┌────────┐                  ┌────────┐
│        │------------------>        │
│ Client │                  │ Lares  │
│        │<-----------------|        │
└────────┘   Return Token   └────────┘
    │
    │  (2)TCP Connection
    │  & Token Handshake
    │
    ▼                       ┌────────────────┐
┌────────┐                  │                │
│        │  (4)Game Protocol│     Roma       │
│ Janus  │<---------------->│    (Hidden)    │
│        │    gRPC Tunnel   │                │
└────────┘                  └────────────────┘
    ▲                              │
    │                              │
    │  (3)Game Protocol            │
    │     TCP                      ▼
    │                        ┌──────────────┐
└────────┘                   │ Senate Admin │
  Client                     └──────────────┘
```

### Authentication Flow Details

```
┌────────┐        ┌────────┐        ┌────────┐        ┌────────┐
│ Client │        │ Lares  │        │ Janus  │        │  Roma  │
└───┬────┘        └───┬────┘        └───┬────┘        └───┬────┘
    │                 │                 │                 │
    │ 1.HTTPS Token Request             │                 │
    │---------------->│                 │                 │
    │                 │                 │                 │
    │ Return Token    │                 │                 │
    │<----------------│                 │                 │
    │                 │                 │                 │
    │ 2.Establish TCP Connection        │                 │
    │---------------------------------->│                 │
    │                 │                 │                 │
    │ Send Token in Handshake           │                 │
    │---------------------------------->│                 │
    │                 │                 │                 │
    │                 │                 │ (Internal Token │
    │                 │                 │  Verification)  │
    │                 │                 │                 │
    │ TCP Handshake Success Response    │                 │
    │<----------------------------------│                 │
    │                 │                 │                 │
    │ 3.Send Login Protocol (TCP)       │                 │
    │---------------------------------->│                 │
    │                 │                 │                 │
    │                 │                 │ 4.Select Roma Service │
    │                 │                 │ Based on Token Info   │
    │                 │                 │                 │
    │                 │                 │ Establish gRPC Tunnel │
    │                 │                 │---------------->│
    │                 │                 │                 │
    │                 │                 │ Tunnel Established      
    │                 │                 │<----------------│
    │                 │                 │                 │
    │ Game Protocol Messages (TCP)      │ Forward as gRPC │
    │---------------------------------->│---------------->│
    │                 │                 │                 │
    │                 │                 │ Game Logic Processing │
    │                 │                 │<----------------│
    │ Response Messages (TCP)           │                 │
    │<----------------------------------│                 │
    │                 │                 │                 │
```

Janus internally adopts a modular architecture, with services communicating via gRPC:

```
                    ┌─────────────────────────────┐
                    │           Janus             │
                    │                             │
┌─────────────┐     │  ┌───────────┐ ┌─────────┐ │
│  Config     │◀───▶│  │TCP Server │ │Protocol │ │
│  (etcd)     │     │  │Module     │ │Handlers │ │
└─────────────┘     │  └───────────┘ └─────────┘ │
                    │        ▲            ▲      │
┌─────────────┐     │        │            │      │
│  Monitoring │◀───▶│        └─────┬──────┘      │
│(Prometheus) │     │              │             │
└─────────────┘     │         ┌────▼────┐        │
                    │         │ gRPC    │        │
┌─────────────┐     │         │ Client  │        │
│  Tracing    │◀───▶│         │ Pool    │        │
│     (OT)    │     │         └─────────┘        │
└─────────────┘     └─────────────────────────────┘
```

## Project Overview

The Janus framework is built on high-performance network libraries optimized for high concurrency, supporting both TCP and WebSocket protocols for client connections and gRPC for backend service communication. The framework design follows Domain-Driven Design (DDD) principles with a clean architecture approach, achieving high availability and scalability through stateless design and load balancing capabilities.

## Technology Stack

Janus utilizes the following core technologies:

| Technology/Component | Purpose | Version |
|---------|------|------|
| Go | Primary development language | 1.23+ |
| netpoll | High-performance network framework | Latest |
| gRPC | Backend service communication | v1.71.1 |
| Protobuf | Data serialization | v1.36.6 |
| etcd | Service discovery & registry | v3.5.21 |
| Redis | Connection tracking | v9.7.3 |
| OpenTelemetry | Distributed tracing | v1.35.0 |
| Prometheus | Monitoring system | v1.22.0 |
| Google Wire | Dependency injection | v0.6.0 |
| zap | Logging | v1.26.0 |
| Buf | API management | Latest |

## Key Features

- **High-Performance TCP Server**: Custom-built TCP server optimized for game traffic
- **Protocol Conversion**: Seamless conversion between TCP and gRPC protocols
- **Dynamic Service Discovery**: Integration with etcd for service discovery
- **Intelligent Routing**: Route requests to appropriate Roma services based on token and request information
- **Load Balancing**: Advanced load balancing algorithms for traffic distribution
- **Connection Management**: Efficient handling of long-lived TCP connections
- **Token Verification**: Secure token verification mechanism
- **Distributed Tracing**: OpenTelemetry integration for end-to-end request tracking
- **Metrics Collection**: Comprehensive metrics for monitoring and alerting
- **Hot Reload**: Support for configuration changes without service restart

## Core Components

### Application Services (app/)

- **server**: TCP server module for client connections
- **protocol**: Protocol handling and conversion
- **client**: gRPC client for backend service communication

### API Definitions (api/)

- **server**: Server-side internal API definitions
- **interface**: Client-facing interface definitions

### Common Libraries (pkg/)

- **network**: Network-related functionality
- **protocol**: Protocol definitions and handlers
- **middleware**: Common middleware components
- **util**: Utility functions

## Requirements

- Go 1.23+
- Protobuf
- etcd
- Redis

## Quick Start

### Initialize Environment

```bash
make init
```

### Generate API Code

```bash
make proto
make api
```

### Build Services

```bash
make build
```

### Start Services

```bash
# Start all services
make run

# Start with specific configuration
make run conf=dev
```

## Integration with go-pantheon Components

Integration of Janus with other go-pantheon components typically follows these steps:

### Integration with Lares Authentication

1. Configure token validation mechanism to verify tokens generated by Lares
2. Set up security parameters for token decryption and validation
3. Configure handshake protocol for client connection establishment

```yaml
# Janus configuration example for Lares integration
security:
  token:
    enable: true
    encryption:
      type: aes
      key: "${TOKEN_ENCRYPTION_KEY}"
    validation:
      issuer: "lares-auth"
      audience: "janus-gateway"
      expiration: 86400  # 24 hours in seconds
```

### Integration with Roma Game Services

1. Configure service discovery to locate Roma service instances
2. Set up gRPC client connection pool for efficient communication
3. Implement intelligent routing based on game service types

```yaml
# Roma service discovery configuration
service_discovery:
  type: etcd
  endpoints:
    - "127.0.0.1:2379"
  service_prefix: "/services/roma"
  watch_interval: 10s

# gRPC client pool configuration  
grpc_client:
  pool_size: 100
  timeout: 5s
  retry:
    max_attempts: 3
    backoff:
      initial: 100ms
      max: 1s
      multiplier: 1.5
```

## Project Structure

```
.
├── api/                # API definitions
│   ├── protocol/       # Game protocol definitions
│   └── server/         # Server-side API
├── app/                # Application services
│   ├── server/         # TCP server implementation
│   ├── protocol/       # Protocol handling
│   └── client/         # gRPC client implementation
├── cmd/                # Command-line entry points
│   └── server/         # Main server entry point
├── configs/            # Configuration files
├── internal/           # Internal packages
│   ├── biz/            # Business logic
│   ├── data/           # Data access
│   ├── server/         # Server implementations
│   └── service/        # Service implementations
├── pkg/                # Common libraries
│   ├── network/        # Network utilities
│   ├── protocol/       # Protocol definitions
│   └── util/           # Utility functions
└── third_party/        # Third-party dependencies
```

## Port Conventions

### Gateway Service

- TCP Ports:
  - Game Client: 8080
  - WebSocket: 8081
- HTTP Ports:
  - Management API: 8090
- gRPC Ports:
  - Internal API: 9090

## Development Guide

### Development Workflow

1. Define protocol formats and transformations
2. Implement TCP server handlers for new protocols
3. Implement gRPC client logic for backend service communication
4. Add routing logic based on protocol and token information
5. Write unit tests for all components
6. Implement metrics and tracing for new functionality
7. Test under load conditions

### Adding New Protocol Support

Steps to add support for a new game protocol:

1. Define protocol structure in `api/protocol/`
2. Generate protocol code with `make proto`
3. Implement protocol handler in `app/protocol/`
4. Add routing logic to direct traffic to appropriate Roma service
5. Update configuration to recognize the new protocol

### Performance Optimization

Janus is optimized for high throughput and low latency:

1. Connection pooling for gRPC clients
2. Zero-copy buffer management
3. Lock-free concurrency patterns
4. Protocol-specific optimizations
5. Efficient memory usage with object pooling

## Monitoring and Operations

### Metrics

Key metrics exposed by Janus:

- Connection counts (active, total, peak)
- Request throughput (per second)
- Response latency (p50, p95, p99)
- Error rates (by type)
- Protocol distribution
- Backend service call statistics

### Logging

Janus uses structured logging with different verbosity levels:

- ERROR: Critical issues requiring immediate attention
- WARN: Issues that need investigation but don't impact service
- INFO: Normal operational information
- DEBUG: Detailed information for troubleshooting
- TRACE: Very detailed protocol-level information (high volume)

## Troubleshooting

### 1. Connection Establishment Failures

**Issue**: Clients cannot establish connections with Janus

**Solution**:
- Check network firewall settings
- Verify that Janus is listening on the correct ports
- Check TLS configuration if using secure connections
- Review logs for handshake errors

### 2. Protocol Conversion Errors

**Issue**: Protocol conversion fails between TCP and gRPC

**Solution**:
- Check protocol version compatibility
- Review protocol handler implementation
- Enable debug logging for detailed protocol information
- Verify that protocol definitions are up to date

### 3. Backend Service Connection Issues

**Issue**: Janus cannot connect to Roma services

**Solution**:
- Check service discovery configuration
- Verify that Roma services are registered and healthy
- Review network connectivity between Janus and Roma services
- Check gRPC client configuration and timeout settings

## Contributing

1. Fork this repository
2. Create a feature branch
3. Submit changes
4. Ensure all tests pass
5. Submit a Pull Request

## License

This project is licensed under the terms specified in the LICENSE file.
