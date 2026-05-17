# go-cli-fetcher

## Effortlessly Manage GitHub from Your Terminal

**go-cli-fetcher** is a fast, extensible command-line tool written in Go, designed for developers who want to interact with GitHub directly from the terminal. Perform common repository, issue, and user operations without ever opening a browser—perfect for automation, scripting, and power users.

### Features

- List, create, and delete repositories
- View detailed repository information
- List branches and commits (with SHA, author, message, date, etc.)
- Create issues from the CLI
- Fetch user info, followers, and following lists
- Secure authentication with local `.env` storage

### Why Choose go-cli-fetcher?

- **Speed:** Instantly fetch, create, or update GitHub data from your terminal.
- **Automation:** Integrate seamlessly with scripts and CI/CD pipelines.
- **Security:** Manage your GitHub token securely.
- **Extensibility:** Easily add new commands thanks to a clean codebase.

### Quickstart

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

### Usage Examples

- List your repositories:
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
- Create an issue:
  ```sh
  go run cmd/main.go issue create <owner> <repoName> "Title" "Body"
  ```
- Show who a user is following:
  ```sh
  go run cmd/main.go user following <username>
  ```

### Environment Variables

- `GITHUB_TOKEN`: Your GitHub personal access token. Set via the `auth` command or manually in a `.env` file.

### Project Structure

- `cmd/` — CLI entrypoint
- `internal/cmd/` — Command definitions
- `internal/github/` — GitHub API logic

### Contributing

Contributions are welcome! Open an issue or submit a pull request to add features, fix bugs, or improve documentation.
