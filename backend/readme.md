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

### How to ping backend?

To ensure that the backend is working and available, use the `/ping` endpoint with the `GET` method. You can use tools like `curl`, Postman, or your preferred HTTP client.

**Expected Response:**

```plaintext
pong
```

### How to Log In to the System?

To log in to the system, use the `/api/login` endpoint with the `POST` method. The default credentials are:

- **Username:** `admin`
- **Password:** `admin`

**Expected Request:**

```json
{
  "username": "admin",
  "password": "admin"
}
```

**Success Response:**

```json
{
  "success": true,
  "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzY2MTYyMDksImlhdCI6MTczNjUyOTgwOSwidXNlcm5h8"
}
```

**Failed Response:**

```json
{
  "success": false,
  "error": {
    "code": "INVALID_AUTH_DATA",
    "message": "Username or password is invalid"
  }
}
```

**Possible Error Messages:**

- `INVALID_AUTH_DATA` - Occurs if the username or password is incorrect.
- `INCORRECT_PAYLOAD` - Occurs if the request payload is missing required fields (e.g., `username` or `password`).

#### Accessing Authorized Endpoints

Endpoints that require authorization will only respond to requests that include a valid access token.
For every request to an endpoint that requires authorization, include the following header:

- **Header Key:** Authorization
- **Header Value:** Bearer YOUR_ACCESS_TOKEN

Replace YOUR_ACCESS_TOKEN with the token you obtained.

**HTTP Header Example:**

```planetext
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6...
```

### Orders Endpoints

#### GET /api/orders

**Description:**  
Retrieves a list of all existing orders.  
**Note:** This endpoint requires authorization. Include an `Authorization` header with a valid Bearer token.

**Headers:**

- `Authorization: Bearer YOUR_ACCESS_TOKEN`

**Success Response:**

- **Status Code:** 200 OK

```json
{
  "orders": [
    {
      "id": 5,
      "reason": "Something went wrong",
      "defect": "Something broken",
      "item_name": "test",
      "total_price": 10,
      "prepayment": 5,
      "created_at": "2025-01-10T13:39:21.435678Z",
      "status": {
        "id": 5,
        "ready_at": null,
        "returned_at": null,
        "customer_notified_at": null,
        "is_outsourced": false,
        "is_recipient_lost": false
      },
      "type": {
        "id": 1,
        "full_name": "Tablet"
      },
      "customer": {
        "id": 5,
        "language_id": 3,
        "phone_number": "+37126578411",
        "language": null
      },
      "worker": {
        "id": 2,
        "first_name": "Andrew",
        "last_name": "Ponatovskis"
      }
    }
  ],
  "success": true
}
```

**Failed Response:**

- **Status Code:** 500 Internal Server Error (or appropriate error code)

```json
{
  "success": false,
  "error": {
    "code": "INTERNAL_ERROR",
    "message": "We're sorry, something went wrong on our end. Please try again later."
  }
}
```

#### GET /api/orders/active

**Description:**  
Retrieves a list of active orders (orders that are not completed).

**Note:** This endpoint requires authorization. Include an `Authorization` header with a valid Bearer token.

**Headers:**

- `Authorization: Bearer YOUR_ACCESS_TOKEN`

**Success Response:**

- **Status Code:** 200 OK

```json
{
  "orders": [
    {
      "id": 5,
      "reason": "Something went wrong",
      "defect": "Something broken",
      "item_name": "test",
      "total_price": 10,
      "prepayment": 5,
      "created_at": "2025-01-10T13:39:21.435678Z",
      "status": {
        "id": 5,
        "ready_at": null,
        "returned_at": null,
        "customer_notified_at": null,
        "is_outsourced": true,
        "is_recipient_lost": false
      },
      "type": {
        "id": 1,
        "full_name": "Tablet"
      },
      "customer": {
        "id": 5,
        "language_id": 3,
        "phone_number": "+37126578411",
        "language": null
      },
      "worker": {
        "id": 2,
        "first_name": "Andrew",
        "last_name": "Ponatovskis"
      }
    }
  ],
  "success": true
}
```

**Failed Response:**

- **Status Code:** 500 Internal Server Error (or appropriate error code)

```json
{
  "success": false,
  "error": {
    "code": "INTERNAL_ERROR",
    "message": "We're sorry, something went wrong on our end. Please try again later."
  }
}
```

#### GET /api/orders/completed

**Description:**  
Retrieves a list of completed orders.

**Note:** This endpoint requires authorization. Include an `Authorization` header with a valid Bearer token.

**Headers:**

- `Authorization: Bearer YOUR_ACCESS_TOKEN`

**Success Response:**

- **Status Code:** 200 OK

```json
{
  "orders": [
    {
      "id": 5,
      "reason": "Something went wrong",
      "defect": "Something broken",
      "item_name": "test",
      "total_price": 10,
      "prepayment": 5,
      "created_at": "2025-01-10T13:39:21.435678Z",
      "status": {
        "id": 5,
        "ready_at": "2025-01-20T18:12:45.435678Z",
        "returned_at": "2025-01-21T09:51:33.435678Z",
        "customer_notified_at": "2025-01-20T18:12:45.435678Z",
        "is_outsourced": true,
        "is_recipient_lost": false
      },
      "type": {
        "id": 1,
        "full_name": "Tablet"
      },
      "customer": {
        "id": 5,
        "language_id": 3,
        "phone_number": "+37126578411",
        "language": null
      },
      "worker": {
        "id": 2,
        "first_name": "Andrew",
        "last_name": "Ponatovskis"
      }
    }
  ],
  "success": true
}
```

**Failed Response:**

- **Status Code:** 500 Internal Server Error (or appropriate error code)

```json
{
  "success": false,
  "error": {
    "code": "INTERNAL_ERROR",
    "message": "We're sorry, something went wrong on our end. Please try again later."
  }
}
```

#### POST /api/orders

**Description:**  
Creates a new order with the specified details.

**Note:** This endpoint requires authorization. Include an `Authorization` header with a valid Bearer token.

**Headers:**

- `Authorization: Bearer YOUR_ACCESS_TOKEN`
- `Content-Type: application/json`

**Request:**

```json
{
  "order_type_id": 1,
  "worker_id": 2,
  "reason": "Something went wrong",
  "defect": "Something broken",
  "total_price": 10.0,
  "prepayment": 5.0,
  "customer": {
    "language_id": 3,
    "phone_number": "+37126578411"
  }
}
```

**Success Response:**

- **Status Code:** 200 OK

```json
{
  "success": true
}
```

**Failed Response:**

- **Status Code:** 500 Internal Server Error (or appropriate error code)

```json
{
  "success": false,
  "error": {
    "code": "INTERNAL_ERROR",
    "message": "We're sorry, something went wrong on our end. Please try again later."
  }
}
```

#### DELETE /api/order/{id}

**Description:**  
Deletes the specified order by its ID.

**Note:** This endpoint requires authorization. Include an `Authorization` header with a valid Bearer token.

**Headers:**

- `Authorization: Bearer YOUR_ACCESS_TOKEN`

**Request:**

- `{id}`: The ID of the order to be deleted (e.g., `/api/order/1`).

**Success Response:**

- **Status Code:** 200 OK

```json
{
  "success": true
}
```

**Failed Response:**

- **Status Code:** 500 Internal Server Error (or appropriate error code)

```json
{
  "success": false,
  "error": {
    "code": "INTERNAL_ERROR",
    "message": "We're sorry, something went wrong on our end. Please try again later."
  }
}
```

### Order Type Endpoints

#### GET /api/orders/types

**Description:**  
Retrieves a list of all available order types.

**Note:** This endpoint requires authorization. Include an `Authorization` header with a valid Bearer token.

**Headers:**

- `Authorization: Bearer YOUR_ACCESS_TOKEN`

**Success Response:**

- **Status Code:** 200 OK

```json
{
  "order_types": [
    {
      "id": 1,
      "full_name": "Tablet"
    },
    {
      "id": 2,
      "full_name": "Laptop"
    }
  ],
  "success": true
}
```

**Failed Response:**

- **Status Code:** 500 Internal Server Error (or appropriate error code)

```json
{
  "success": false,
  "error": {
    "code": "INTERNAL_ERROR",
    "message": "We're sorry, something went wrong on our end. Please try again later."
  }
}
```

#### POST /api/orders/types

**Description:**  
Creates a new order type.

**Note:** This endpoint requires authorization. Include an `Authorization` header with a valid Bearer token.

**Headers:**

- `Authorization: Bearer YOUR_ACCESS_TOKEN`

**Success Response:**

- **Status Code:** 200 OK

```json
{
  "success": true
}
```

**Failed Response:**

- **Status Code:** 500 Internal Server Error (or appropriate error code)

```json
{
  "success": false,
  "error": {
    "code": "INTERNAL_ERROR",
    "message": "We're sorry, something went wrong on our end. Please try again later."
  }
}
```

#### DELETE /api/orders/type/{id}

**Description:**  
Deletes the specified order type by its ID.

**Note:** This endpoint requires authorization. Include an `Authorization` header with a valid Bearer token.

**Headers:**

- `Authorization: Bearer YOUR_ACCESS_TOKEN`

**Request:**

- `{id}`: The ID of the order type to be deleted (e.g., `/api/orders/type/1`).

**Success Response:**

- **Status Code:** 200 OK

```json
{
  "success": true
}
```

**Failed Response:**

- **Status Code:** 500 Internal Server Error (or appropriate error code)

```json
{
  "success": false,
  "error": {
    "code": "INTERNAL_ERROR",
    "message": "We're sorry, something went wrong on our end. Please try again later."
  }
}
```

### Worker Endpoints

#### GET /api/workers

**Description:**  
Retrieves a list of all workers.

**Note:** This endpoint requires authorization. Include an `Authorization` header with a valid Bearer token.

**Headers:**

- `Authorization: Bearer YOUR_ACCESS_TOKEN`

**Success Response:**

- **Status Code:** 200 OK

```json
{
  "success": true,
  "workers": [
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
}
```

**Failed Response:**

- **Status Code:** 500 Internal Server Error (or appropriate error code)

```json
{
  "success": false,
  "error": {
    "code": "INTERNAL_ERROR",
    "message": "We're sorry, something went wrong on our end. Please try again later."
  }
}
```

#### POST /api/workers

**Description:**  
Creates a new worker.

**Note:** This endpoint requires authorization. Include an `Authorization` header with a valid Bearer token.

**Headers:**

- `Authorization: Bearer YOUR_ACCESS_TOKEN`
  
**Request:**

```json
{
  "first_name": "Gabe",
  "last_name": "Oldell"
}
```

**Success Response:**

- **Status Code:** 200 OK

```json
{
  "success": true
}
```

**Failed Response:**

- **Status Code:** 500 Internal Server Error (or appropriate error code)

```json
{
  "success": false,
  "error": {
    "code": "INTERNAL_ERROR",
    "message": "We're sorry, something went wrong on our end. Please try again later."
  }
}
```

#### DELETE /api/worker/{id}

**Description:**  
Deletes the specified worker by their ID.

**Note:** This endpoint requires authorization. Include an `Authorization` header with a valid Bearer token.

**Headers:**

- `Authorization: Bearer YOUR_ACCESS_TOKEN`

**Request:**

- `{id}`: The ID of the worker to be deleted (e.g., `/api/worker/1`).

**Success Response:**

- **Status Code:** 200 OK

```json
{
  "success": true
}
```

**Failed Response:**

- **Status Code:** 500 Internal Server Error (or appropriate error code)

```json
{
  "success": false,
  "error": {
    "code": "INTERNAL_ERROR",
    "message": "We're sorry, something went wrong on our end. Please try again later."
  }
}
```

### Language Endpoints

#### GET /api/languages

**Description:**  
Retrieves a list of all available languages.

**Note:** This endpoint requires authorization. Include an `Authorization` header with a valid Bearer token.

**Headers:**

- `Authorization: Bearer YOUR_ACCESS_TOKEN`

**Success Response:**

- **Status Code:** 200 OK

```json
{
  "languages": [
    {
      "id": 1,
      "short_name": "ru",
      "full_name": "Russian"
    },
    {
      "id": 2,
      "short_name": "lv",
      "full_name": "Latvian"
    },
    {
      "id": 3,
      "short_name": "en",
      "full_name": "English"
    }
  ],
  "success": true
}
```

**Failed Response:**

- **Status Code:** 500 Internal Server Error (or appropriate error code)

```json
{
  "success": false,
  "error": {
    "code": "INTERNAL_ERROR",
    "message": "We're sorry, something went wrong on our end. Please try again later."
  }
}
```

#### POST /api/languages

**Description:**  
Creates a new language.

**Note:** This endpoint requires authorization. Include an `Authorization` header with a valid Bearer token.

**Headers:**

- `Authorization: Bearer YOUR_ACCESS_TOKEN`

**Request:**

```json
{
  "short_name": "LV",
  "full_name": "Latvian"
}
```

**Success Response:**

- **Status Code:** 200 OK

```json
{
  "success": true
}
```

**Failed Response:**

- **Status Code:** 500 Internal Server Error (or appropriate error code)

```json
{
  "success": false,
  "error": {
    "code": "INTERNAL_ERROR",
    "message": "We're sorry, something went wrong on our end. Please try again later."
  }
}
```

#### DELETE /api/language/{id}

**Description:**  
Deletes the specified language by their ID.

**Note:** This endpoint requires authorization. Include an `Authorization` header with a valid Bearer token.

**Headers:**

- `Authorization: Bearer YOUR_ACCESS_TOKEN`

**Request:**

- `{id}`: The ID of the language to be deleted (e.g., `/api/language/1`).

**Success Response:**

- **Status Code:** 200 OK

```json
{
  "success": true
}
```

**Failed Response:**

- **Status Code:** 500 Internal Server Error (or appropriate error code)

```json
{
  "success": false,
  "error": {
    "code": "INTERNAL_ERROR",
    "message": "We're sorry, something went wrong on our end. Please try again later."
  }
}
```
