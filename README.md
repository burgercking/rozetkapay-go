[Public API Gateway RozetkaPay](https://docs.google.com/document/d/1AbNXlJlPdzjZcpotd83Qb7GWXt78UhYGRY-GQRWI35M)

## Installation

Make sure your project is using Go Modules (it will have a `go.mod` file in its
root if it already is):

```sh
go mod init
```

Then, reference rozetkapay-go in a Go program with `import`:

```go
import (
	"github.com/burgercking/rozetkapay-go"
)
```

Run any of the normal `go` commands (`build`/`install`/`test`). The Go
toolchain will resolve and fetch the rozetkapay-go module automatically.

Alternatively, you can also explicitly `go get` the package into a project:

```bash
go get -u github.com/burgercking/rozetkapay-go
```

### Payments

- [x] (Create payment)
- [x] (Confirm payment)
- [x] (Cancel payment)
- [x] (Refund payment)
- [x] (Payment info)
- [x] (Resend callback)

### Customers

- [x] (Delete customer payment from wallet)
- [x] (Get customer information & waller)
- [x] (Add customer payment to wallet)
