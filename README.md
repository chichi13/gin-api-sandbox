# Go

By default, `Go` has a linter.

If we don't use a variable or a package, the application won't be able to run. A compilation error will occur:

```bash
./main.go:9:6: conferenceName declared but not used
```

## Go CLI

### Initialize a project

```bash
go mod init gin-api
```

### Run the application

```bash
go run main.go
```

### Compile the application

```bash
go build main.go
```

### Install a package and save it in the `go.mod` file

```
go get -u github.com/gin-gonic/gin
go install github.com/gin-gonic/gin
```

### Remove unused packages

```bash
go mod tidy
```

### Generate a swagger documentation

```bash
swag init
```

### Format the code

```bash
go fmt
swag fmt # for the format of the swagger documentation
```
