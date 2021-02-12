# testprotosize

A temporary repo for reproducing https://github.com/mwitkow/go-proto-validators/issues/51.

## Generate proto files

Run `make proto`

## Run test

Run `make run`

Expected output: `x=size:1` (see `fmt.Printf(...)` in `cmd/testsize/main.go`).

Actual output: compilation fails with:
```
proto/testsize.validator.pb.go:20:11: this.Size_ undefined (type *ThingWithSize has no field or method Size_)
proto/testsize.validator.pb.go:21:123: this.Size_ undefined (type *ThingWithSize has no field or method Size_)
```

If `Size_` is renamed to `Size`, compilation succeeds.
