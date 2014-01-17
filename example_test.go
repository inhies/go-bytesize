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

func ExampleNew_math() {
	b1 := bytesize.New(1024)
	b2 := bytesize.New(4096)
	sum := b1 + b2
	fmt.Printf("%s", sum)

	// Output:
	// 5.00KB
}

func ExampleParse() {
	b, _ := bytesize.Parse("1024 GB")
	fmt.Printf("%s\n", b)

	b, _ = bytesize.Parse("3 petabytes")
	fmt.Printf("%s\n", b)

	bytesize.LongUnits = true
	bytesize.Format = "%.0f "
	fmt.Printf("%s\n", b)

	// Output:
	// 1.00TB
	// 3.00PB
	// 3 petabytes
}
