# Journal CLI (Go version)

A simple command-line journaling application built with Go and SQLite.

## About

A Go version of [journal-cli](https://github.com/TaylorMichaelHall/journal-cli). Inspired by [jrnl](https://jrnl.sh/), which is a better app you should use if you're interested in journaling in your terminal.

## Features

- Add journal entries
- View entries (with options to show IDs and timestamps)
- Edit existing entries
- Delete entries
- Colorized output for improved readability

## Installation

### Prerequisites

- Go 1.16 or higher
- SQLite3

### How to build from source

1. Clone this repository:

   ```
   git clone https://github.com/taylormichaelhall/journal-cli-go.git
   cd journal-cli-go
   ```

2. Build the application:

   ```
   go build -o journal-cli
   ```

3. Move the binary to a directory in your PATH or run it from the current directory.

## Uninstallation

- If built from source, simply delete the binary.

- To remove the journal database, delete it from your user's home folder in `~/.journal-cli/journal.db`

## Usage

After installation, you can use the `journal-cli` command from anywhere:

- Show help:
  ```
  journal-cli --help
  ```

- Add an entry:
  ```
  journal-cli add Your journal entry goes here
  ```

  Note that this is a very simple terminal entry method so you will need to escape special characters like quote.

- Show entries:
  ```
  journal-cli show
  journal-cli show -id  # Show with IDs
  journal-cli show -timestamps  # Show with timestamps
  journal-cli show -id -timestamps  # Show with both IDs and timestamps
  ```

- Edit an entry:
  ```
  journal-cli edit <entry-id> Your updated journal entry
  ```

- Delete an entry:
  ```
  journal-cli delete <entry-id>
  ```

## Acknowledgements

- Inspired by [jrnl](https://jrnl.sh)
- Built with Go and SQLite
- Uses [fatih/color](https://github.com/fatih/color) for colored terminal output
- Uses [mattn/go-sqlite3](https://github.com/mattn/go-sqlite3) for SQLite database operations

