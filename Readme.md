Fabric error
============

Documenting an error.

- `Error` can be found in the `cmd/error` - `main` branch
- `Fix` can be found in the `cmd/fix` - `fix` branch

Error
-----

```
# github.com/hyperledger/fabric-sdk-go/internal/github.com/hyperledger/fabric/core/operations
vendor/github.com/hyperledger/fabric-sdk-go/internal/github.com/hyperledger/fabric/core/operations/system.go:225:49: not enough arguments in call to s.statsd.SendLoop
        have (<-chan time.Time, string, string)
        want (context.Context, <-chan time.Time, string, string)
make: *** [Makefile:5: run-error] Error 2
```