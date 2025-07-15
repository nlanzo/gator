# gator

RSS feed aggregator

## Requirements
- **PostgreSQL**: You need a running Postgres instance for data storage.
- **Go**: Version 1.18 or newer is recommended.

## Installation
Install the gator CLI using Go:

```sh
go install github.com/nlanzo/gator@latest
```

This will place the `gator` binary in your `$GOPATH/bin` or `$HOME/go/bin` directory. Make sure this directory is in your `PATH`.

## Configuration
The configuration file is named `.gatorconfig.json` and should be located in your home directory at `~/.gatorconfig.json`.

The config file must include your database URL. The program will automatically set up the `current_user_name` for you.

Example `~/.gatorconfig.json`:

```json
{
  "db_url": "connection_string_goes_here",
  "current_user_name": "username_goes_here"
}
```

- Replace `db_url` with your actual Postgres connection string.
- You do not need to set `current_user_name` manually; the program will handle it.

## Available Commands
- `login <username>` — Log in to your account.
- `register <username>` — Register a new user account.
- `users` — List all users.
- `agg <interval>` — Start aggregating feeds every specified interval (e.g., `agg 10m`).
- `addfeed <feed name> <feed_url>` — Add a new feed to the aggregator.
- `feeds` — List all available feeds.
- `follow <feed_url>` — Follow a feed.
- `following` — List feeds you are following.
- `unfollow <feed_url>` — Unfollow a feed.
- `browse [number of posts]` — Browse posts from feeds you follow.


---

