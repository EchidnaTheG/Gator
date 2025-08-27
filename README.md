# Gator - RSS Feed Aggregator CLI

Gator is a command-line RSS feed aggregator that allows users to register, manage their feeds, and follow RSS content from various sources.

## Overview

Gator is a Go-based CLI application that provides users with a simple interface to manage and aggregate RSS feeds. It uses PostgreSQL for data storage and implements a command-based architecture for user interactions.

## Features

- **User Management**: Register, login, and manage user accounts
- **Feed Management**: Add, list, and follow RSS feeds
- **RSS Aggregation**: Fetch and display content from various RSS sources
- **Persistent Configuration**: Saves user preferences between sessions

## Tech Stack

- **Language**: Go 1.23.5+
- **Database**: PostgreSQL
- **ORM**: SQLC for type-safe SQL queries
- **Migration**: Goose for database migrations
- **Dependencies**:
  - github.com/google/uuid: For generating unique IDs
  - github.com/lib/pq: PostgreSQL driver for Go
  - modernc.org/libc: C library implementation

## Project Structure

```text
Gator/
├── internal/            # Application internals
│   ├── commands/        # Command handlers and interfaces
│   ├── config/          # Configuration management
│   ├── database/        # Database models and queries (generated)
│   └── rss/             # RSS feed fetching and parsing
├── sql/                 # SQL files
│   ├── schema/          # Database schema definitions
│   └── queries/         # SQL queries for SQLC
├── main.go              # Application entry point
├── go.mod               # Go module definition
├── sqlc.yaml            # SQLC configuration
└── README.md            # This documentation file
```

## Database Schema

The application uses three primary tables:

1. **users**: Stores user information
   - id (UUID, primary key)
   - created_at (TIMESTAMP)
   - updated_at (TIMESTAMP)
   - name (TEXT, unique)

2. **feeds**: Stores RSS feed information
   - id (UUID, primary key)
   - created_at (TIMESTAMP)
   - updated_at (TIMESTAMP)
   - name (TEXT)
   - URL (TEXT, unique)
   - userid (UUID, foreign key to users)

3. **feed_follows**: Tracks which users follow which feeds
   - id (UUID, primary key)
   - created_at (TIMESTAMP)
   - updated_at (TIMESTAMP)
   - userid (UUID, foreign key to users)
   - feedid (UUID, foreign key to feeds)
   - A unique constraint on (userid, feedid)

## Commands

Gator implements the following commands:

- `login`: Log in as an existing user
- `register`: Create a new user account
- `reset`: Reset the database (danger: removes all data)
- `users`: List all registered users
- `agg`: Aggregate and display RSS content
- `addfeed`: Add a new RSS feed
- `feeds`: List all available feeds
- `follow`: Follow a specific feed
- `following`: List feeds being followed

## Configuration

Gator stores its configuration in a JSON file at `~/.gatorconfig.json` with the following structure:

```json
{
  "db_url": "postgresql connection string",
  "current_user_name": "currently logged in user"
}
```

## RSS Functionality

Gator fetches RSS feeds using standard HTTP requests and parses the XML responses. It handles HTML entity decoding to ensure proper display of feed content.

## Usage Examples

### User Management

```bash
# Register a new user
gator register username

# Login as an existing user
gator login username

# List all users
gator users
```

### Feed Management

```bash
# Add a new RSS feed
gator addfeed https://example.com/rss

# List all available feeds
gator feeds

# Follow a feed
gator follow feed_id

# Show feeds you're following
gator following
```

### Content Aggregation

```bash
# View aggregated content from followed feeds
gator agg
```

## Development

### Prerequisites

- Go 1.23.5 or higher
- PostgreSQL database
- SQLC for SQL code generation
- Goose for database migrations

### Setup

1. Clone the repository
2. Install dependencies: `go get -u ./...`
3. Set up PostgreSQL and create a database
4. Configure the database URL in `~/.gatorconfig.json`
5. Run database migrations: `goose up`
6. Build the application: `go build -o gator`

## License

[License details would go here]

## Contributors

- [EchidnaTheG](https://github.com/EchidnaTheG)
