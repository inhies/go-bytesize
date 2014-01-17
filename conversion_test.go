package bytesize

import (
	"testing"
)

var newTable = []struct {
	Bytes  float64
	Result string
}{
	{1, "1.00B"},
	{1023, "1023.00B"},
	{1024, "1.00KB"},
	{1048576, "1.00MB"},
	{1073741824, "1.00GB"},
	{1099511627776, "1.00TB"},
	{1125899906842624, "1.00PB"},
	{1152921504606846976, "1.00EB"},
	{1180591620717411303424, "1.00ZB"},
	{1208925819614629174706176, "1.00YB"},
}

func Test_New(t *testing.T) {
	for _, v := range newTable {
		b := New(v.Bytes)
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
