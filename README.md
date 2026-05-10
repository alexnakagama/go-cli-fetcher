# go-cli-fetcher

## Overview

go-cli-fetcher is a command-line tool written in Go for fetching data from GitHub repositories. It provides commands to interact with the GitHub API, retrieve repository information, and manage authentication securely from your terminal.

## Features
- Fetch general information about GitHub repositories
- Retrieve detailed information for a specific repository
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

### Fetch Repository List
Fetch information about your repositories:

```sh
go run cmd/main.go repo
```

### Fetch Repository Details
Fetch detailed information about a specific repository:

```sh
go run cmd/main.go repo detail <owner> <repo>
```
Replace `<owner>` and `<repo>` with the GitHub username and repository name.

## Environment Variables
- `GITHUB_TOKEN`: Your GitHub personal access token. Set via the `auth` command or manually in a `.env` file.

## Development
- All commands are implemented using [Cobra](https://github.com/spf13/cobra).
- The codebase is organized under `cmd/` for the entrypoint and `internal/` for command and API logic.

## Contributing
Contributions are welcome. Please open issues or submit pull requests for improvements or bug fixes.

## License
This project is licensed under the MIT License.
