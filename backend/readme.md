# Backend Docs

## Backend Structure

```text
├── cmd/
│   ├── app/
│       └── main.go    # Application entry point
├── internal/          # Private application and package code
│   ├── api/    # Package for handling HTTP requests and API logic
│   |   ├── handler/
│   │   |   ├── handlers.go      # Handler creation
│   │   │   ├── language.go      # Methods for handling languages
│   │   │   ├── order_type.go    # Methods for handling order types
│   │   │   ├── order.go         # Methods for handling orders
│   │   │   └── worker.go        # Methods for handling workers
|   │   ├── routes/
|   │       └──  routes.go  # HTTP routes configuration, mapping endpoints to handlers
│   ├── db/
│   │   ├── config.go       # Database configuration
│   │   ├── db.go           # Database initialization and connection pool setup
|   |   └── interface.go    # Database interface definition
│   ├── repository/   # Repository layer for abstracting database access logic
│   │   ├── admins.go           # Data access methods for managing admin users
│   │   ├── languages.go        # Data access methods for managing supported languages
│   │   ├── order_status.go     # Data access methods for managing order status
│   │   ├── order_type.go       # Data access methods for managing order types
│   │   ├── order.go            # Data access methods for managing orders
|   |   └── workers.go          # Data access methods for managing workers
│   ├── server/
│   │   ├── interface.go     # Server interface definitions
│       └── server.go        # Server configuration
├── utils/     # Utility functions and helpers
│   └── env.go/
├── build_docker_image.py    # Python script to automate the building of Docker images for the backend
├── Dockerfile               # Dockerfile to build the backend application image
├── go.mod                   # Go module file
├── go.sum                   # Go module dependencies file
└── README.md                # Project README
```

## API

### GET /api/orders

Retrieves a list of all existing orders.

#### Response

```json
[
  {
    "id": 16,
    "order_status_id": 13,
    "order_type_id": 1,
    "worker_id": 2,
    "customer_id": 3,
    "reason": "Something went wrong",
    "defect": "Something broke",
    "total_price": 10.0,
    "prepayment": 5.0,
    "status": {
      "id": 13,
      "ready_at": null,
      "returned_at": null,
      "customer_notified_at": null,
      "is_outsourced": false,
      "is_recipient_lost": false
    },
    "type": {
      "id": 1,
      "full_name": "Phone"
    },
    "customer": {
      "id": 3,
      "language_id": 3,
      "phone_number": "+37126578411"
    },
    "worker": {
      "id": 1,
      "first_name": "Andrew",
      "last_name": "Ponatovskis"
    }
  }
]
```

### POST /api/orders

Creates a new order with the specified details.

#### Request Body

```json
{
  "order_type_id": 1,
  "worker_id": 2,
  "customer_id": 3,
  "reason": "Something went wrong",
  "defect": "Something broken",
  "total_price": 10.0,
  "prepayment": 5.0
}
```

#### Response

```json
{
  Success
}
```

Note: The response structure will be refactored in future iterations.

### DELETE /api/order/{id}

Deletes the specified order by its ID.

#### Request

- `{id}`: The ID of the order to be deleted (e.g., `/api/order/1`).

#### Response

```json
{
  Success
}
```

Note: The response structure will be refactored in future iterations.

### GET /api/orders/types

Retrieves a list of all available order types.

#### Response

```json
[
  {
    "id": 1,
    "full_name": "Phone"
  },
  {
    "id": 2,
    "full_name": "Laptop"
  }
]
```

### POST /api/orders/types

Creates a new order type.

#### Request Body

```json
{
  "full_name": "Tablet"
}
```

#### Response

```json
{
  "Success"
}
```

Note: The response structure will be refactored in future iterations.

### DELETE /api/orders/type/{id}

Deletes the specified order type by its ID.

#### Request

- `{id}`: The ID of the order type to be deleted (e.g., `/api/orders/type/1`).

#### Response

```json
{
  "Success"
}
```

### GET /api/workers

Retrieves a list of all workers.

```json
[
  {
    "id": 1,
    "first_name": "Steve",
    "last_name": "Stew"
  },
  {
    "id": 2,
    "first_name": "Bob",
    "last_name": "Bobinskis"
  }
]
```

### POST /api/workers

Creates a new worker.

#### Request Body

```json
{
  "first_name": "Gabe",
  "last_name": "Oldell"
}
```

#### Response

```json
{
  "Success"
}
```

#### DELETE /api/worker/{id}

Deletes the specified worker by their ID.

#### Request

- `{id}`: The ID of the worker to be deleted (e.g., `/api/worker/1`).

#### Response

```json
{
  "Success"
}
```

Note: The response structure will be refactored in future iterations.

### GET /api/languages

Retrieves a list of all available languages.

#### Response

```json
[
  {
    "id": 1,
    "short_name": "EN",
    "full_name": "English"
  },
  {
    "id": 2,
    "short_name": "RU",
    "full_name": "Russian"
  }
]
```

### POST /api/languages

Creates a new language.

#### Request Body

```json
{
  "short_name": "LV",
  "full_name": "Latvian"
}
```

#### Response

```json
{
  "Success"
}
```

Note: The response structure will be refactored in future iterations.

### DELETE /api/language/{id}

#### Request

- `{id}`: The ID of the language to be deleted (e.g., `/api/language/1`).

#### Response

```json
{
  "Success"
}
```

Note: The response structure will be refactored in future iterations.
