# go-cli-fetcher

## Overview

go-cli-fetcher is a command-line tool written in Go for interacting with GitHub repositories. It allows you to list, create, delete, and get details about repositories, as well as list branches and commits, all from your terminal.

## Features
- List repositories of the authenticated user
- Get detailed information about a specific repository
- Create and delete repositories
- List branches of a repository
- List commits of a repository (with SHA, author, message, date, etc.)
- Configure and store your GitHub token securely

## Requirements
- Go 1.25 or higher
- A valid GitHub personal access token (with appropriate scopes)

## Installation
Clone the repository and install dependencies:

```sh
git clone https://github.com/alexnakagama/go-cli-fetcher.git
cd go-cli-fetcher
go mod download
```

## Usage

### Authentication
Before using the CLI, configure your GitHub token:

```sh
go run cmd/main.go auth
```
You will be prompted to enter your token. It will be saved in a `.env` file.

### Repository Commands
All repository operations are grouped under the `repo` command:

- List repositories:
  ```sh
  go run cmd/main.go repo list
  ```
- Get repository details:
  ```sh
  go run cmd/main.go repo detail <owner> <repo>
  ```
- Create a repository:
  ```sh
  go run cmd/main.go repo create <owner> <repoName>
  ```
- Delete a repository:
  ```sh
  go run cmd/main.go repo delete <owner> <repoName>
  ```
- List branches:
  ```sh
  go run cmd/main.go repo list-branches <owner> <repoName>
  ```
- List commits:
  ```sh
  go run cmd/main.go repo list-commits <owner> <repoName>
  ```

### Example Output for Commits
Each commit will show:
- SHA
- URL
- Author (login)
- Email
- Message
- Date

## Environment Variables
- `GITHUB_TOKEN`: Your GitHub personal access token. Set via the `auth` command or manually in a `.env` file.

## Development
- All commands are implemented using [Cobra](https://github.com/spf13/cobra).
- The codebase is organized under `cmd/` for the entrypoint and `internal/` for command and API logic.

## Contributing
Contributions are welcome. Please open issues or submit pull requests for improvements or bug fixes.

## License
This project is licensed under the MIT License.
