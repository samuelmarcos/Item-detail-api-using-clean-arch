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

### 1. Listar Todos os Produtos
```
GET /api/v1/item
```

Retorna uma lista de todos os produtos cadastrados.

**Resposta de Sucesso (200 OK)**
```json
[
    {
        "id": 1,
        "product_id": "MLB123456",
        "name": "Samsung Galaxy A55 5G",
        "description": "Smartphone com alto desempenho",
        "brand": "Samsung",
        "model": "Galaxy A55 5G",
        "color": "Azul escuro",
        "category": "Celulares",
        "images": ["https://example.com/img1.jpg"],
        "price": 439.00,
        "original_price": 499.00,
        "discount_percent": 12.0,
        "stock": 50,
        "seller": {
            "name": "Samsung",
            "type": "Oficial",
            "reputation": "Ótima",
            "sales": 5000,
            "official": true
        },
        "warranty": "1 ano de garantia",
        "payment_options": ["Cartão de crédito", "Boleto"],
        "specs": {
            "Tela": "6.6\"",
            "Memória": "256 GB",
            "RAM": "8 GB"
        },
        "related_products": ["MLB654321", "MLB789012"],
        "rating": 4.8,
        "review_count": 769,
        "free_shipping": true,
        "purchase_options": ["Novo"],
        "highlights": ["Mais vendido", "Promoção"],
        "created_at": "2024-03-14T10:00:00Z",
        "updated_at": "2024-03-14T10:00:00Z"
    }
]
```

**Resposta de Erro (500 Internal Server Error)**
```json
{
    "error": "Failed to fetch products"
}
```

### 2. Buscar Produto por ID
```
GET /api/v1/item/{id}
```

Retorna os detalhes de um produto específico.

**Parâmetros de URL**
- `id` (obrigatório): ID do produto (ex: MLB123456)

**Resposta de Sucesso (200 OK)**
```json
{
    "id": "MLB123456",
    "name": "Samsung Galaxy A55 5G",
    // ... mesmo formato do item acima ...
}
```

**Respostas de Erro**
- 400 Bad Request
```json
{
    "error": "Product ID is required"
}
```
- 404 Not Found
```json
{
    "error": "Product not found"
}
```

### 3. Criar Novo Produto
```
POST /api/v1/item
```

Cria um novo produto no sistema.

**Headers**
- `Content-Type: application/json`

**Corpo da Requisição**
```json
{
    "id": "MLB123456",           // obrigatório
    "name": "Samsung Galaxy A55 5G", // obrigatório
    "description": "Smartphone com alto desempenho",
    "brand": "Samsung",
    "model": "Galaxy A55 5G",
    "color": "Azul escuro",
    "category": "Celulares",
    "images": ["https://example.com/img1.jpg"],
    "price": 439.00,             // obrigatório
    "original_price": 499.00,
    "discount_percent": 12.0,
    "stock": 50,                 // obrigatório
    "seller": {                  // obrigatório
        "name": "Samsung",
        "type": "Oficial",
        "reputation": "Ótima",
        "sales": 5000,
        "official": true
    },
    "warranty": "1 ano de garantia",
    "payment_options": ["Cartão de crédito", "Boleto"],
    "specs": {
        "Tela": "6.6\"",
        "Memória": "256 GB",
        "RAM": "8 GB"
    },
    "related_products": ["MLB654321", "MLB789012"],
    "rating": 4.8,
    "review_count": 769,
    "free_shipping": true,
    "purchase_options": ["Novo"],
    "highlights": ["Mais vendido", "Promoção"]
}
```

**Resposta de Sucesso (201 Created)**
```json
{
    "id": "MLB123456",
    "name": "Samsung Galaxy A55 5G",
    // ... mesmo formato do item criado ...
}
```

**Respostas de Erro**
- 400 Bad Request
```json
{
    "error": "Invalid product data"
}
```
- 500 Internal Server Error
```json
{
    "error": "Failed to create product"
}
```

## Campos Obrigatórios
- `id`: Identificador único do produto
- `name`: Nome do produto
- `price`: Preço atual do produto
- `stock`: Quantidade em estoque
- `seller`: Informações do vendedor (objeto com campos obrigatórios)
  - `name`: Nome do vendedor
  - `type`: Tipo do vendedor
  - `reputation`: Reputação do vendedor
  - `sales`: Número de vendas
  - `official`: Se é vendedor oficial

## Exemplos de Uso

### Usando cURL

1. Listar todos os produtos:
```bash
curl http://localhost:8080/api/v1/item
```

2. Buscar produto específico:
```bash
curl http://localhost:8080/api/v1/item/MLB123456
```

3. Criar novo produto:
```bash
curl -X POST http://localhost:8080/api/v1/item \
  -H "Content-Type: application/json" \
  -d '{
    "id": "MLB123456",
    "name": "Samsung Galaxy A55 5G",
    "price": 439.00,
    "stock": 50,
    "seller": {
        "name": "Samsung",
        "type": "Oficial",
        "reputation": "Ótima",
        "sales": 5000,
        "official": true
    }
  }'
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