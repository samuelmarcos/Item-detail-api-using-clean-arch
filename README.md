# Item Detail API - Clean Architecture Implementation

## Overview
This project implements a RESTful API for fetching Mercado Libre item details using Go (Golang) with Clean Architecture principles. The application is designed to provide detailed information about products listed on Mercado Libre, including item specifications, pricing, seller information, and other relevant details. The API follows clean architecture principles to ensure maintainability, testability, and scalability while adhering to best practices in software development.

## Architecture

### Clean Architecture Layers
The project follows Clean Architecture principles with the following layers:

1. **Domain Layer** (`src/internal/domain/`)
   - Contains business entities and rules
   - Independent of any framework or external concerns
   - Defines interfaces for repositories and use cases

2. **Use Cases Layer** (`src/internal/application/`)
   - Implements business logic
   - Orchestrates the flow of data to and from entities
   - Contains application-specific business rules

3. **Interface Adapters Layer** (`src/internal/infra/`)
   - Converts data between the format most convenient for entities and use cases
   - Contains controllers, presenters, and gateways
   - Implements interfaces defined in the domain layer

4. **Frameworks & Drivers Layer** (`src/cmd/`)
   - Contains web frameworks, databases, and external interfaces
   - Implements the interfaces defined in the interface adapters layer

## Project Structure
```
.
├── cmd/
│   └── api/
│       └── main.go           # Application entry point
├── src/
│   ├── internal/
│   │   ├── domain/          # Enterprise business rules
│   │   ├── application/     # Application business rules
│   │   └── infra/          # Interface adapters
│   │       ├── persistence/ # Data persistence implementations
│   │       └── web/        # Web server and handlers
│   └── pkg/                # Shared utilities and helpers
├── tests/                  # Integration and e2e tests
├── go.mod
├── go.sum
└── README.md
```

## Features
- RESTful API for product details
- Clean Architecture implementation
- Local JSON file-based persistence
- Comprehensive error handling
- High test coverage (>80%)
- Swagger/OpenAPI documentation
- Dependency injection
- Middleware support (CORS, logging, etc.)

## API Endpoints

### Get Product Details
```
GET /api/v1/products/{id}
```
Returns detailed information about a specific product.

**Response Example:**
```json
{
    "id": "123",
    "name": "Product Name",
    "description": "Product Description",
    "price": 99.99,
    "category": "Category",
    "stock": 100
}
```

## Getting Started

### Prerequisites
- Go 1.21 or higher
- Git

### Installation
1. Clone the repository:
```bash
git clone https://github.com/samuelmarcos/item-detail-api-using-clean-arch.git
cd item-detail-api-using-clean-arch
```

2. Install dependencies:
```bash
go mod download
```

3. Run the application:
```bash
go run cmd/api/main.go
```

The server will start on `http://localhost:8080`

### Running Tests
```bash
go test ./... -cover
```

## Design Choices

### Why Clean Architecture?
- **Separation of Concerns**: Each layer has a specific responsibility
- **Testability**: Business logic can be tested independently
- **Maintainability**: Changes in one layer don't affect others
- **Flexibility**: Easy to swap implementations (e.g., database, web framework)

### Why Gin Framework?
- Lightweight and fast
- Middleware support
- Easy routing
- Good community support
- Built-in validation

### Why JSON File Storage?
- Simple to implement and maintain
- No database setup required
- Easy to version control
- Suitable for demonstration purposes

## Challenges and Solutions

### Challenge 1: Dependency Management
**Problem**: Managing dependencies between layers while maintaining clean architecture principles.
**Solution**: Used dependency injection and interface-based design to decouple layers.

### Challenge 2: Error Handling
**Problem**: Consistent error handling across all layers.
**Solution**: Implemented custom error types and centralized error handling middleware.

### Challenge 3: Testing
**Problem**: Achieving high test coverage while maintaining clean architecture.
**Solution**: Used interfaces and mocks to test each layer independently.

## Contributing
1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License
This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments
- Clean Architecture by Robert C. Martin
- Gin Web Framework
- Go Standard Library 