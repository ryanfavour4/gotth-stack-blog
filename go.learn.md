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

**### `go install github.com/air-verse/air@latest`**

- Installs Air, a Go development tool that automatically rebuilds and restarts your application when you save changes.
- Similar in purpose to using `nodemon` in a Node.js project.
- Installs the `air` command-line executable, usually in Go's `bin` directory.
- If the `air` command isn't recognized, check that Go's `bin` directory is included in your system's `PATH`.

**### `air init`**

- Creates an `.air.toml` configuration file in the current project directory.
- This file tells Air how to build and run your application and which files or directories to watch or ignore.
- Run it from your project's root directory, where your `go.mod` file is located.
- You normally only need to run it once. Running it again when `.air.toml` already exists will report that the configuration already exists.

**### `.air.toml`**

- Air's configuration file, roughly comparable to a `nodemon.json` configuration file in a Node.js project.
- Controls the build command, the executable Air runs, watched file extensions, and excluded directories.
- You can configure it to ignore folders such as `node_modules` and `tmp` to avoid unnecessary file watching.
- You can also configure additional commands or build rules for tools such as Templ.

**### `air`**

- Starts your application in development mode using the settings in `.air.toml`.
- Air builds your Go application and runs it if the build succeeds.
- When you modify and save a watched Go file, Air rebuilds and restarts the application automatically.
- This removes the need to stop and rerun `go run .` manually after every change.
- If the application fails to compile, Air reports the build error and cannot successfully start the updated application.

**### How Air works with Templ**

- `templ fmt views` formats the `.templ` files in the `views` directory.
- `templ generate` converts `.templ` files into generated Go source files.
- Air can be configured to run these commands when template files change.
- Once the generated Go files change, Air can rebuild and restart the application.
- Templ formatting, Templ code generation, and Go compilation are separate steps, even when configured to run automatically during development.

**### Typical development setup**

Run these commands from your project's root directory:

```bash
# 1. Install Air
go install github.com/air-verse/air@latest

# 2. Create Air's configuration file (once)
air init

# 3. Start the development server
air
```

After the initial setup, you generally only need to run:

```bash
air
```

Keep that terminal open while developing. Air will watch for changes and rebuild your application according to your configuration.
