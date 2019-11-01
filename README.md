dally
===

To run this test, first run `$ go get https://github.com/MarAvFe/dally`, then run `$ go run main.go` on:

```go
package main

import (
	"fmt"

	d "github.com/MarAvFe/dally"
)

func main() {
	db, e := d.NewMongoDAL("localhost", "awesomeDB")
	fmt.Println(db, e)

	fmt.Println(d.InsertTime(db))
}

```

Based on [Harry Gogonis](https://medium.com/@harrygogonis)'s [awesome tutorial](https://medium.com/@harrygogonis/testing-go-mocking-third-party-dependancies-4ab4e1c9bd3f).
