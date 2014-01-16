package bytesize

import (
	"testing"
)

var newTable = []struct {
	Bytes  int64
	Result string
}{
	{1, "1.00B"},
	{1023, "1023.00B"},
	{1024, "1.00KB"},
}

func Test_New(t *testing.T) {
	for _, v := range newTable {
		b := New(float64(v.Bytes))
		if b.String() != v.Result {
			t.Fatalf("Expected %s, received %s", v.Result, b)
		}
	}
}

var parseTable = []struct {
	Input  string
	Result string
}{
	{"1B", "1.00B"},
	{"1 B", "1.00B"},
	{"1B ", "1.00B"},
	{" 1 B ", "1.00B"},
	{"1023B", "1023.00B"},
	{"1024B", "1.00KB"},
}

func Test_Parse(t *testing.T) {
	for _, v := range parseTable {
		b, err := Parse(v.Input)
		if err != nil {
			t.Fatal(err)
		}
		if b.String() != v.Result {
			t.Fatalf("Expected %s, received %s", v.Result, b)
		}
	}
}

var mathTable = []struct {
	B1       ByteSize
	Function rune
	B2       ByteSize
	Result   string
}{
	{1024, '+', 1024, "2.00KB"},
	{1073741824, '+', 10485760, "1.01GB"},
	{1073741824, '-', 536870912, "512.00MB"},
}

func Test_Math(t *testing.T) {
	for _, v := range mathTable {
		switch v.Function {
		case '+':
			total := v.B1 + v.B2
			if total.String() != v.Result {
				t.Fatalf("Fail: %s + %s = %s, received %s", v.B1, v.B2, v.Result, total)
			}
		case '-':
			total := v.B1 - v.B2
			if total.String() != v.Result {
				t.Fatalf("Fail: %s - %s = %s, received %s", v.B1, v.B2, v.Result, total)
			}
		}
	}
}
