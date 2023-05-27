```bash
# Not use work space
➜  go-workspace git:(go-workspace) ✗ go run test1
package test1 is not in GOROOT (/usr/local/go/src/test1)

# Init workspace use only test1
➜  go-workspace git:(go-workspace) ✗ go work init ./test1
➜  go-workspace git:(go-workspace) ✗ go run test1
this is module: test1
➜  go-workspace git:(go-workspace) ✗ go run test2
package test2 is not in GOROOT (/usr/local/go/src/test2)

# Add test to workspace
➜  go-workspace git:(go-workspace) ✗ go work use ./test2
➜  go-workspace git:(go-workspace) ✗ go run test2       
this is module: test2
```