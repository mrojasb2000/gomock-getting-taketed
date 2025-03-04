# GoMock

1. Add command generate GoMock to interface files

```sh
//go:generate mockgen -source=<PATH>/interface.go -destination=mocks/mocks.go -package=<Package> <Repository>,<Cases>
```

1. Run GoMock generator
   
```sh
go generate <PATH>/interface.go
```

1. Running Unit Testing

```sh
go test ./...
```
