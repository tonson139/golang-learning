```bash
➜  golang-learning git:(master) ✗ cd go-module 
➜  go-module git:(master) ✗ go run .
this is main
Meow
Meow !!!
 _______ 
< Hello >
 ------- 
        \   ^__^
         \  (oo)\_______
            (__)\       )\/\
                ||----w |
                ||     ||

➜  go-module git:(master) ✗ go build -o dist
➜  go-module git:(master) ✗ ./dist 
this is main
Meow
Meow !!!
 _______ 
< Hello >
 ------- 
        \   ^__^
         \  (oo)\_______
            (__)\       )\/\
                ||----w |
                ||     ||
➜  go-module git:(master) ✗ 
```

### testing 
```bash
➜  go-module git:(go-tool) ✗ go test -v
=== RUN   TestSayThisIsBasic
    main_test.go:12: SUCCESS: expect `This is basic go`
--- PASS: TestSayThisIsBasic (0.00s)
PASS
ok      soda-meow       0.463s
```