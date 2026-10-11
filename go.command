> go run . 
- Starts the go project looking for the main.go file in the current (.) directory. If nested in a folder then go run ./folder

! go.sum is almost like node_modules in js/ts world.

> go get github.com/a-h/templ
- "go get" is an equivalent of nodes "npm install" following the package or library source. But in most GO cases its a github url for some reason.

! go.mod Think of go.mod roughly like package.json in Node. It tells Go:
"This is the identity of my project, and these are the Go dependencies it uses."

! templ generate
This is like compiling your regular code to what Golang actually uses a little like part of the npm run dev process/step where it gives you a dist but this is just the annoying repetitive version. - .templ isn't directly executed by Go. Templ generates Go code from it.

