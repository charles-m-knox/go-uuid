# uuid4

My own minimalist Go library that generates uuid's that implement v4, variant 10.

Securely uses `crypto/rand` for random number generation.

Using this library instead of typical uuid libraries (that package other uuid versions/features) can reduce a compiled binary size by as much as 7% according to my experiments.

## Usage

First, get it:

```bash
go get -v github.com/charles-m-knox/go-uuid
```

Now, in a `main.go` file:

```go
package main

import (
    "log"

    uuid "github.com/charles-m-knox/go-uuid"
)

func main() {
    u := uuid.New() // generates the uuid
    log.Println(u)  // prints the uuid to console
}
```

You can also check out [`uuid_test.go`](./uuid_test.go) to see the unit tests.
