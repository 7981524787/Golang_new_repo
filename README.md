Fast Compilation 
Simple Programming language(25)
CSP(Concurrency)-> Go routines and channels
no runtime to be installed
static binaries
static linking vs dynamic linking-->libc musl glibc
cross compilation
Go routines
Go echo system
Cloud native 
GC (Concurrent GC) Low Latency, concurrent GC
General purpose statically typed programming language






### go mod

```
go mod init demo
```

```
go tool compile -o output.o
```

```
go tool link -o demo output.o
```

```
go run hello.go
```

```
 go build -o hello1 hello.go
 ```


- all distributions

```
go tool dist list
```

- cross compilation and build

 ```
 GOOS=linux GOARCH=amd64 go build -o build/hello-linux-amd64.exe hello.go
 ```

 ## Go presentations link

 https://docs.google.com/presentation/d/1WVvsbvgHKBrNrKtnT4XWRfrsfkNlbw5u6L9O1DeVBn0/edit?usp=sharing


 ## Context implementation in Golang

 https://www.youtube.com/playlist?list=PLJE7PIP1qj_Rn9vq4V4jGJbj5KqEIWSUc