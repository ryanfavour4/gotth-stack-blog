### `go run .`

- Compiles and runs the Go package in the current directory.
- `.` represents the current package, not specifically the `main.go` file.
- To run a package in a subdirectory, use something like `go run ./cmd/server`.
- Go compiles and runs the application without normally leaving a compiled executable in the project directory.

### `go.mod`

- Roughly equivalent to `package.json`.
- Identifies the Go module, declares its Go version, and records dependency requirements.
- It also helps Go resolve imports within the project.

### `go.sum`

- Not equivalent to `node_modules`.
- Records cryptographic checksums used to verify Go module downloads.
- The Go module cache is a closer comparison to `node_modules`, although Go typically stores dependencies outside the project directory.

### `go get github.com/a-h/templ`

- Adds or updates a Go module dependency in `go.mod`.
- The module path often resembles a GitHub URL, but Go can retrieve modules through proxies.
- `go install` is used separately when we need to install a command-line tool such as `templ`.

### `templ generate`

- Reads `.templ` template files and generates corresponding `.go` source files.
- Go compiles the generated Go code when we run the application.
- This is code generation, not the same thing as running the Go application or starting its development server.
