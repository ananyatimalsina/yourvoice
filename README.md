# YourVoice

**Anonymous eVoting & Feedback Platform**

YourVoice is a cryptographically secure anonymous voting platform that uses RSA blind signatures to ensure voter privacy while preventing double-voting. Built with Go, PostgreSQL, and modern web technologies.

## 🔐 Features

- **Anonymous Voting**: RSA blind signatures ensure votes cannot be traced to voters
- **Double-Vote Prevention**: Storage of both anonymous secrets as well as identifiable spent events prevent multiple votes per person
- **Network Privacy**: Tor network support for complete anonymity
- **Secure Architecture**: Mathematical privacy guarantees through cryptography
- **Open Source**: Fully auditable implementation

## 🏗️ Architecture

### Cryptographic Protocol

YourVoice implements a 4-step RSA blind signature protocol:

1. **AUTH**: Voter proves identity without revealing their vote
2. **BLIND**: Voter blinds their randomly generated secret using cryptographic blinding factor
3. **SIGN**: Authority signs the blinded secret without knowing its content
4. **VOTE**: Voter unblinds the signature and submits anonymously via Tor

```
unblind(sign(blind(message))) = sign(message)
```

The authority never sees the original secret or the voters decision, and the server cannot link secrets to voter identity.

## 🛠️ Tech Stack

- **Go** - Backend server with stdlib HTTP routing
- **PostgreSQL** - Database with GORM ORM
- **Tailwind CSS** - Utility-first CSS framework
- **HTMX** - Dynamic web applications without complex JavaScript

## 📁 Project Structure

```
├── internal/
│   ├── database/           # Database configuration and models
│   │   ├── models/
│   │   │   ├── Candidate.go
│   │   │   ├── Message.go
│   │   │   ├── MessageEvent.go
│   │   │   ├── Party.go
│   │   │   ├── Vote.go
│   │   │   └── VoteEvent.go
│   │   └── database.go
│   ├── handlers/           # HTTP handlers and routing
│   │   ├── routes/
│   │   │   ├── expression/ # Vote and message submission
│   │   │   │   ├── message.go
│   │   │   │   └── vote.go
│   │   │   └── identity/   # Cryptographic verification
│   │   │       └── verify.go
│   │   └── handlers.go
│   ├── middleware/         # HTTP middleware
│   │   ├── contentTypeJson.go
│   │   ├── logging.go
│   │   └── middleware.go
│   └── utils/              # Cryptographic utilities
│       ├── rss.go
│       └── types.go
├── web/
│   ├── static/             # CSS and assets
│   │   └── main.css
│   └── templates/          # HTML templates
│       └── index.html      # Landing page with API docs
├── docker-compose.yml      # PostgreSQL database
├── .env                    # Environment configuration
├── go.mod                  # Go dependencies
└── main.go                 # Application entrypoint
```

## 🚀 Quick Start

### Prerequisites

- Go
- Docker & Docker Compose
- Node.js (for Tailwind CSS)
- Optional: Nix (for reproducible development environment)

### 0. Nix Development Environment (Optional)

```bash
nix-shell
# Now you have Go and Node.js available
```

### 1. Clone Repository

```bash
git clone https://github.com/ananyatimalsina/yourvoice
cd yourvoice
```

### 2. Install Dependencies

```bash
# Go dependencies
go mod tidy

# Node.js dependencies (for Tailwind CSS)
npm install
```

### 3. Environment Configuration

The `.env` file is pre-configured for development:

```bash
SERVER_PORT=3000
TZ=Europe/Berlin

# PostgreSQL configuration
POSTGRES_HOST=localhost
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_DB=postgres
POSTGRES_PORT=5432
```

### 4. Run Development Server

```bash
air
```

Visit `http://localhost:3000` to see the platform.

## 📡 API Endpoints

### Identity Verification

Get blind signatures for voting or messaging:

#### POST `/api/identity/verifyVote`
```json
// Request
{
  "digest": "blinded_secret",
  "event_id": 1
}

// Response
{
  "signature": "signed_digest"
}
```

#### POST `/api/identity/verifyMessage`
```json
// Request
{
  "digest": "blinded_secret", 
  "event_id": 1
}

// Response
{
  "signature": "signed_digest"
}
```

### Expression Submission

Submit anonymous votes and messages:

#### POST `/api/expression/vote`
```json
// Request
{
  "data": "signed_data",
  "digest": "unsigned_data",
  "vote_event_id": 1,
  "candidate_id": 2
}

// Response
"Vote received successfully"
```

#### POST `/api/expression/message`
```json
// Request
{
  "data": "signed_data",
  "digest": "unsigned_data", 
  "message_event_id": 1,
  "message": "Your feedback text"
}

// Response
"Vote received successfully"
```

## 🔧 Development

### Adding New Models

Create models in `internal/database/models/`:

```go
// internal/database/models/NewModel.go
package models

import "yourvoice/internal/utils"

type NewModel struct {
    utils.Expression
    // Add your fields
}
```

### Adding New Routes

1. Create handler in `internal/handlers/routes/`
2. Register route in `internal/handlers/handlers.go`

### Database Migrations

GORM auto-migration runs on startup. Models are automatically migrated when the server starts.

## Production Deployment

### Manual Build

```bash
# Build CSS for production
npx @tailwindcss/cli -i "./web/static/main.css" -o "./web/static/output.css" --minify

# Build Go binary
go build -o yourvoice main.go

# Deploy binary + web/ directory + .env
```

### Environment Variables

Configure production environment in `.env`:

```bash
SERVER_PORT=8080
POSTGRES_HOST=your-db-host
POSTGRES_USER=your-db-user
POSTGRES_PASSWORD=your-secure-password
POSTGRES_DB=yourvoice_prod
POSTGRES_PORT=5432
```

### Security Considerations

- **HTTPS**: Always use HTTPS in production
- **Database**: Use secure PostgreSQL credentials
- **Tor Integration**: Consider Tor hidden service deployment
- **Rate Limiting**: Implement API rate limiting
- **Logging**: Monitor for suspicious voting patterns

## 🔒 Cryptographic Security

### RSA Blind Signatures

The platform implements David Chaum's blind signature protocol:

1. **Blinding Factor**: Random value `r` is used to blind the message
2. **Blind Message**: `blind = message * r^e mod n`
3. **Sign Blind**: Authority signs without seeing original: `sig = blind^d mod n`
4. **Unblind**: Voter recovers signature: `real_sig = sig / r mod n`

### Security Properties

- **Anonymity**: Authority cannot link signatures to voters
- **Unforgeability**: Only authority can create valid signatures
- **Single-Use**: Secrets are stored to prevent double-voting
- **Unlinkability**: Submitted votes cannot be traced to sign requests

## 🧪 Testing

```bash
# Run tests
go test ./...

# Test with verbose output
go test -v ./...

# Test specific package
go test ./internal/utils
```

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 🔗 Links

- **GitHub**: [https://github.com/ananyatimalsina/yourvoice](https://github.com/ananyatimalsina/yourvoice)
- **Live Demo**: [Coming Soon]
- **Documentation**: [This README]

## ⚠️ Disclaimer

This is experimental software. While implementing well-established cryptographic protocols, it should be thoroughly audited before use in production voting systems. The authors are not responsible for any security vulnerabilities or election integrity issues.

---

**Built with privacy in mind. Express your voice, protect your identity.**
