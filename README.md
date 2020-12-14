# go-dirk

## Run from source

You can run dirk from source e.g. like this:

```bash
go run dirk/main.go init <project-folder>
```

## Build

You can build dirk like this:

```bash
go build -o dirk/build/dirk dirk/main.go
```

or this:

```bash
export GOOS=darwin
export GOARCH=amd64
go build -o dirk/build/dirk_${GOOS}_${GOARCH} dirk/main.go
```

## Run tests

You can run all of the tests by 

```bash
go test ./...
```
