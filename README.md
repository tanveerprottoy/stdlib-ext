## stdlib-ext
The stdlib-ext library provides functional enhancements for various standard library packages and other features.

Install
```
go get github.com/tanveerprottoy/stdlib-ext
```

## testing
unit test:
go test ./...

go test -v ./<pathToPackage>

benchmark:
go test ./<pathToPackage> -bench=. -benchmem