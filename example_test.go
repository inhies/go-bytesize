package bytesize

import (
	"fmt"
)

func ExampleNew() {
	b := bytesize.New(1024)
	fmt.Printf("%s", b)

	// Output:
	// 1.00KB
}

func ExampleParse() {
	b, _ := Parse("1024 GB")
	fmt.Printf("%s", b)

	// Output:
	// 1.00TB
}
