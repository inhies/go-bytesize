package bytesize_test

import (
	"fmt"
	"github.com/inhies/go-bytesize"
)

func ExampleNew() {
	b := bytesize.New(1024)
	fmt.Printf("%s", b)

	// Output:
	// 1.00KB
}

func ExampleParse() {
	b, _ := bytesize.Parse("1024 GB")
	fmt.Printf("%s", b)

	// Output:
	// 1.00TB
}
