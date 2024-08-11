# uuid4

My own minimalist Go library that generates uuid's that implement v4, variant 10.

Securely uses `crypto/rand` for random number generation.

## Usage

First, get it:

```bash
go get -v git.cmcode.dev/cmcode/uuid
```

Now, in a `main.go` file:

```go
package main

import (
    "log"

    uuid "git.cmcode.dev/cmcode/uuid"
)

func main() {
    u := uuid.New() // generates the uuid
    log.Println(u)  // prints the uuid to console
}
```

You can also check out [`uuid_test.go`](./uuid_test.go) to see the unit tests.
