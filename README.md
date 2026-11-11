# Gator Aggregator

A command-line RSS feed aggregator built with Go and PostgreSQL. Gator allows users to register accounts, subscribe to RSS feeds, automatically collect posts, and browse aggregated content from their followed feeds.

## Motivation

I built Gator as part of the [Boot.dev](https://www.boot.dev/) backend development curriculum to gain practical experience working with Go and PostgreSQL.

The project brings together several backend concepts, including database design, SQL queries, HTTP requests, XML parsing, and command-line application development.

Some of the key features include:

- User registration and account switching.
- Adding and managing RSS feeds.
- Following and unfollowing feeds.
- Periodically fetching and storing feed posts.
- Browsing aggregated posts from followed feeds.

## Quick Start

### Prerequisites

- [Go](https://go.dev/dl/)
- [PostgreSQL](https://www.postgresql.org/download/)
- [Goose](https://github.com/pressly/goose) (database migrations)

### Installation

Clone the repository:

```bash
git clone https://github.com/tmcausla/gatoraggregator.git
cd gatoraggregator
```

Install dependencies:

```bash
go mod download
```

### Database Setup

Create a PostgreSQL database for Gator:

```sql
CREATE DATABASE gator;
```

Configure your database connection by creating a `.gatorconfig.json` file in your home directory:

```json
{
    "db_url": "postgres://username:password@localhost:5432/gator?sslmode=disable",
    "current_user_name": ""
}
```

Replace the connection details with your PostgreSQL credentials.

Run the database migrations:

```bash
goose -dir sql/schema postgres "postgres://username:password@localhost:5432/gator?sslmode=disable" up
```

### Build

Compile the application:

```bash
go build -o gatoraggregator .
```

You're ready to start using Gator! Commands can be entered in the following format:

```bash
./gatoraggregator <command> <parameter>
```

## Usage

Gator is operated entirely through the command line.

### User Management

Register a new user:

```bash
./gatoraggregator register <username>
```

Switch to an existing user:

```bash
./gatoraggregator login <username>
```

List registered users:

```bash
./gatoraggregator users
```

Reset the database:

```bash
./gatoraggregator reset
```

**Note:** Resetting the database deletes registered users and their associated data.

### Feed Management

Add a new RSS feed:

```bash
./gatoraggregator addfeed <name> <url>
```

Adding a feed automatically subscribes the current user to it.

List available feeds:

```bash
./gatoraggregator feeds
```

Follow an existing feed:

```bash
./gatoraggregator follow <url>
```

List feeds followed by the current user:

```bash
./gatoraggregator following
```

Unfollow a feed:

```bash
./gatoraggregator unfollow <url>
```

### Aggregation

Start collecting posts from registered feeds:

```bash
./gatoraggregator agg <duration>
```

The duration determines how frequently Gator fetches the next feed. Go's duration syntax is supported, for example:

```bash
./gatoraggregator agg 1m
```

This starts the aggregation process, fetching feeds continuously at the specified interval.

### Browsing Posts

Browse posts collected from feeds followed by the current user:

```bash
./gatoraggregator browse
```

By default, the command displays the two most recent posts. An optional argument specifies the number of posts to retrieve:

```bash
./gatoraggregator browse 10
```

## Database

Gator uses PostgreSQL for persistent storage, with [Goose](https://github.com/pressly/goose) managing schema migrations and [sqlc](https://sqlc.dev/) generating type-safe Go code from SQL queries.

The database is organized around three primary entities:

- **Users:** Registered accounts that can manage their own feed subscriptions.
- **Feeds:** RSS feeds added to the aggregator. Feeds can be shared between multiple users.
- **Posts:** Articles retrieved from RSS feeds, associated with the feed they originate from.

Users and feeds have a many-to-many relationship through feed subscriptions, allowing multiple users to follow the same feed without duplicating it in the database.

## Contributing

This project was built as a learning exercise, but suggestions and improvements are welcome!

If you'd like to contribute:

1. Fork the repository.
2. Create a branch for your changes.
3. Make your changes and test them.
4. Submit a pull request describing your changes.

