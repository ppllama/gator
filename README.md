# gator
Boot.dev Guided Project- Blog AggreGator

Gator is a command-line RSS feed aggregator written in Go. It continuously fetches posts from configured feeds, stores them in Postgres, and lets you browse and manage feeds from the terminal.

---

## Prerequisites

To run **gator**, you must have the following installed:

- **Go** (1.21 or newer recommended)
- **PostgreSQL** (14 or newer recommended)

Make sure both `go` and `psql` are available in your PATH.

---

## Installation

Install the gator CLI using `go install`:

```bash
go install github.com/ppllama/gator@latest
```

After installation, ensure `$GOPATH/bin` (or `$HOME/go/bin`) is in your PATH:

```bash
export PATH="$PATH:$HOME/go/bin"
```

You should now be able to run:

```bash
gator
```

---

## Database Setup

Create a Postgres database for gator:

```sql
CREATE DATABASE gator;
```

---

## Configuration

Gator uses a config file located at:

```text
~/.gatorconfig.json
```

Example config:

```json
{
  "db_url": "postgres://user:password@localhost:5432/gator?sslmode=disable",
  "current_user": "alice"
}
```

- `db_url` – Postgres connection string
- `current_user` – the active gator user

---

## Usage

### Create a user

```bash
gator register alice
```

### Add a feed

```bash
gator addfeed "Hacker News" https://news.ycombinator.com/rss
```

### Follow a feed

```bash
gator follow https://news.ycombinator.com/rss
```

### Fetch posts (runs continuously)

```bash
gator agg 1h
```

### Browse posts

```bash
gator browse 10
```

---

## Available Commands

- register
- login
- users
- addfeed
- feeds
- follow
- unfollow
- following
- agg
- browse
- reset (dangerous- resets the database)

---