# go-cli-fetcher

## Take Control of GitHub from Your Terminal

go-cli-fetcher is a powerful, extensible command-line interface for developers who want to interact with GitHub efficiently—no browser required. Built in Go, it brings the most common repository, issue, and user operations to your fingertips, making automation and scripting seamless for power users and teams.

## Why use go-cli-fetcher?
- **Speed:** Instantly fetch, create, or update GitHub data without leaving your terminal.
- **Automation:** Integrate with scripts and CI/CD pipelines for advanced workflows.
- **Security:** Manage your GitHub token securely with local `.env` storage.
- **Extensible:** Clean codebase, easy to add new commands.

## Key Features
- List, create, and delete repositories
- Get detailed repository information
- List branches and commits (with SHA, author, message, date, etc.)
- Create issues directly from the CLI
- Fetch user info, followers, and following lists
- Secure authentication flow

## Quickstart
1. **Clone and install dependencies:**
   ```sh
   git clone https://github.com/alexnakagama/go-cli-fetcher.git
   cd go-cli-fetcher
   go mod download
   ```
2. **Authenticate:**
   ```sh
   go run cmd/main.go auth
   ```
   Enter your GitHub token when prompted. It will be stored in `.env`.

## Usage Examples
- **List your repositories:**
  ```sh
  go run cmd/main.go repo list
  ```
- **Get repository details:**
  ```sh
  go run cmd/main.go repo detail <owner> <repo>
  ```
- **Create a repository:**
  ```sh
  go run cmd/main.go repo create <owner> <repoName>
  ```
- **Delete a repository:**
  ```sh
  go run cmd/main.go repo delete <owner> <repoName>
  ```
- **List branches:**
  ```sh
  go run cmd/main.go repo list-branches <owner> <repoName>
  ```
- **List commits:**
  ```sh
  go run cmd/main.go repo list-commits <owner> <repoName>
  ```
- **Create an issue:**
  ```sh
  go run cmd/main.go issue create <owner> <repoName> "Title" "Body"
  ```
- **Show who a user is following:**
  ```sh
  go run cmd/main.go user following <username>
  ```

## Environment Variables
- `GITHUB_TOKEN`: Your GitHub personal access token. Set via the `auth` command or manually in a `.env` file.

## Project Structure
- `cmd/` — CLI entrypoint
- `internal/cmd/` — Command definitions
- `internal/github/` — GitHub API logic

## Contributing
We welcome contributions! Open an issue or submit a pull request to add features, fix bugs, or improve documentation.
