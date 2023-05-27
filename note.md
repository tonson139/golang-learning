### go ran and build
```bash
# Cannot use `go run .` because go run . is looking for main module and main function inside main file but in /go-learning doesn't have mod
➜  golang-learning git:(master) ✗ go run .
go: cannot find main module, but found .git/config in /Users/pataninphueanmuenwai/Desktop/golang-learning
        to create a module there, run:
        go mod init

# use `go run basic.go` to run a file
➜  golang-learning git:(master) ✗ go run basic.go
This is basic go
# build binary file
➜  golang-learning git:(master) ✗ go build basic.go
# run build
➜  golang-learning git:(master) ✗ ./basic     
This is basic go
➜  golang-learning git:(master) ✗ 
```

### godoc
```bash
➜  ~ godoc
zsh: command not found: godoc
➜  ~ GOPATH=$HOME/go
➜  ~ PATH=$PATH:$GOPATH/bin
➜  ~ godoc
using module mode; GOMOD=/dev/null

# godoc is on localhost:6060
```