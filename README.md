# Gorth Stack Template

A modern, full-stack web development template combining **Go**, **GORM**, **Tailwind CSS**, and **HTMX** for rapid, efficient web application development.

## Tech Stack

- **Go** - Backend server with stdlib HTTP routing
- **GORM** - Go ORM for database operations 
- **Tailwind CSS** - Utility-first CSS framework
- **HTMX** - Dynamic web applications without complex JavaScript

## Project Structure

```
├── internal/
│   ├── database/           # Database configuration and models
│   │   ├── models/
│   │   └── database.go
│   ├── handlers/           # HTTP handlers and routing
│   │   ├── routes/         # API route handlers
│   │   │   ├── greeting.go
│   │   │   ├── stats.go
│   │   │   └── time.go
│   │   └── handlers.go     # Main route registration
│   └── middleware/         # HTTP middleware stack
│       ├── logging.go      # Request logging
│       └── middleware.go   # Middleware composition
├── web/
│   ├── static/             # Static assets
│   │   └── main.css        #  input file
│   └── templates/          # HTML templates
│       ├── index.html      # Main page
│       ├── greeting.html   # HTMX fragment
│       ├── stats.html      # HTMX fragment  
│       └── time.html       # HTMX fragment
├── .env                    # Environment variables
├── .air.toml               # Air configuration
├── go.mod                  # Go dependencies
├── main.go                 # Application entrypoint
├── package.json            # Node.js dependencies ()
└── shell.nix               # Nix development environment
```

## Quick Start

### Prerequisites

- Go
- Node.js (for  CSS)
- Optional: Nix (for reproducible development environment)

### 1. Clone Template

```bash
# Use this repository as a GitHub template
git clone https://github.com/ananyatimalsina/gorth
cd gorth
```

### 2. Install Dependencies

```bash
# Go dependencies
go mod tidy

# Node.js dependencies (for )
npm install
```

### 3. Environment Setup

Modfiy `.env` as needed:

```bash
nano .env
```

### 4. Run Development Server

```bash
# Option 1: Use air live reload
air

# Option 2: Manual commands (no live reload)
npx @tailwindcss/cli -i "./web/static/main.css" -o "./web/static/output.css" --minify
go run main.go
```

Visit `http://localhost:8080` to see the demo page.

### 6. Nix Development Environment (Optional)

```bash
nix-shell
# Now you have Go and Node.js available
```

## Architecture

### HTTP Server

- **Standard Library Routing**: Uses Go's `http.ServeMux` for routing
- **Middleware Stack**: Composable middleware using functional composition
- **Static File Serving**: Serves CSS, JS, and other assets from `/static`

### Middleware

The middleware stack includes:

- **Logging**: Request/response logging with status codes
- **Custom**: Easy to add your own middleware

### Database Integration

GORM is pre-configured but commented out in `main.go`. To enable:

1. Uncomment database initialization in `main.go`
2. Configure your database connection in `internal/database/database.go`
3. Add models and migrations as needed

### HTMX Integration

The template demonstrates HTMX patterns:

- **Fragment Responses**: API endpoints return HTML fragments
- **Progressive Enhancement**: Works without JavaScript
- **Server-Side State**: All state management on the server

### CSS Architecture

- **Tailwind CSS**: Utility-first styling
- **Build Process**: CSS is compiled from `main.css` to `output.css`
- **Development**: Rebuild CSS when adding new Tailwind classes

## Development Workflow

### Adding New Routes

1. Create handler in `internal/handlers/routes/`
2. Register route in `internal/handlers/handlers.go`
3. Create template in `web/templates/` (if needed)

Example:

```go
// internal/handlers/routes/example.go
package routes

import (
    "html/template"
    "net/http"
)

func ExampleHandler(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("web/templates/example.html"))
    data := map[string]interface{}{
        "Message": "Hello from HTMX!",
    }
    tmpl.Execute(w, data)
}

// Register in handlers.go
apiRouter.HandleFunc("GET /example", routes.ExampleHandler)
```

### Adding Middleware

```go
// internal/middleware/auth.go
package middleware

import "net/http"

func Auth(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Auth logic here
        next.ServeHTTP(w, r)
    })
}

// Add to stack in main.go
stack := middleware.CreateStack(
    middleware.Logging,
    middleware.Auth,  // Add your middleware
)
```

### Working with HTMX

Templates in `web/templates/` are HTMX fragments. They should:

- Return pure HTML without `<html>` or `<head>` tags
- Use Tailwind classes for styling
- Include HTMX success indicators if desired

Example HTMX endpoint:

```html
<!-- In your main template -->
<button hx-get="/api/data" hx-target="#result" hx-swap="innerHTML">
    Load Data
</button>
<div id="result"></div>

<!-- web/templates/data.html -->
<div class="p-4 bg-green-50 rounded">
    <p>Data loaded: {{.Timestamp}}</p>
</div>
```

### CSS Development
When adding new Tailwind classes:

```bash
# Rebuild CSS
npx @tailwindcss/cli -i "./web/static/main.css" -o "./web/static/output.css" --minify

# Or use air which includes CSS build
air
```

### Database Models

Add models to `internal/database/models`:

```go
// internal/database/models/user.go
package database

import "gorm.io/gorm"

type User struct {
    gorm.Model
    Name  string
    Email string `gorm:"unique"`
}

// Add to migrateModels in database.go
func migrateModels(db *gorm.DB) {
    db.AutoMigrate(&User{})
}
```

## Production Deployment

### Build Process

```bash
# Build CSS
npx @tailwindcss/cli -i "./web/static/main.css" -o "./web/static/output.css" --minify

# Build Go binary
go build -o app main.go

# Deploy binary + web/ directory + .env
```

### Environment Variables

Configure production environment:

```bash
# .env
SERVER_PORT=8080
# Add your production configuration
```

### Deployment Considerations

- Serve static files through a CDN or reverse proxy
- Use a production database (PostgreSQL, MySQL, etc.)
- Implement proper logging and monitoring
- Add rate limiting and security headers
- Use HTTPS in production

## Features Demonstrated

The template includes working examples of:

- **Server-rendered HTML** with Go templates
- **HTMX interactions** with multiple endpoints
- **Tailwind CSS** styling with responsive design
- **Middleware composition** for cross-cutting concerns
- **Static file serving** for CSS, JS, and assets
- **Environment configuration** with `.env` files
- **Development tooling** with build scripts

## Extending the Template

This template provides a solid foundation for:

- **REST APIs** with JSON responses
- **Server-side rendered applications** with HTMX
- **Database-driven applications** with GORM
- **Real-time features** with WebSockets (add your own)
- **Authentication systems** (add middleware)
- **File uploads** and processing
- **Background jobs** and workers

## Contributing

Feel free to submit issues and pull requests to improve this template.
