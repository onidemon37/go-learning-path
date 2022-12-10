# go-learning-path

## Install go lintint and pre-commit hooks
```
go install golang.org/x/lint/golint@latest
```

## Install goimports
```
go install golang.org/x/tools/cmd/goimports@latest
```

## Install gocyclo
```
go install github.com/fzipp/gocyclo/cmd/gocyclo@latest
```

## Install golangci-lint
```
go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.46.2
```

## Install ginkgo
```
go install github.com/onsi/ginkgo/ginkgo@latest
```

# Install Cobra cli
```
go get github.com/spf13/cobra/cobra
or
go get -u github.com/spf13/cobra@latest

go get github.com/spf13/cobra@upgrade

go install github.com/spf13/cobra-cli@latest
cd project
cobra-cli init
```

## go build

### Compile an hr Binary for Linux
Go into the GOPATH src directory by using cd
```
$GOPATH/src/project.
```

Compile a binary called hr.linux that can run on Linux systems like your workstation by using
```
go build
go build -o project.linux.
```

### Compile an hr Binary for macOS
Compile a binary that can run on macOS systems and call the binary hr.darwin by using 
```
$ GOOS=darwin GOARCH=amd64 go build -o project.darwin
```

### Compile an hr Binary for FreeBSD
Compile a binary that can run on macOS systems and call the binary hr.freebsd by using 
```
$ GOOS=freebsd GOARCH=amd64 go build -o project.freebsd.
```
