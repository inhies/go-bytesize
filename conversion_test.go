package bytesize

import (
	"testing"
)

func Test_Overflow(t *testing.T) {
	b, err := Parse("1797693134862315708145274237317043567981000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000B")
	if err == nil || b != 0 {
		t.Fatal("Max float64 test did not fail")
	}
}

var formatTable = []struct {
	Bytes  float64
	Result string
}{
	{1, "1 byte"},
	{1024, "1 kilobyte"},
}

func Test_Format(t *testing.T) {
	for _, v := range formatTable {
		bSize := New(v.Bytes)
		b := bSize.Format("%.0f ", true)
		if b != v.Result {
			t.Fatalf("Expected %s, received %s", v.Result, b)
		}
	}
}

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

var globalFormatTable = []struct {
	Bytes  float64
	Result string
}{
	{1, "1 byte"},
	{1023, "1023 bytes"},
	{1024, "1 kilobyte"},
	{1048576, "1 megabyte"},
	{1073741824, "1 gigabyte"},
	{1099511627776, "1 terabyte"},
	{1125899906842624, "1 petabyte"},
	{1152921504606846976, "1 exabyte"},
	{1180591620717411303424, "1 zettabyte"},
	{1208925819614629174706176, "1 yottabyte"},
	{2 * 1, "2 bytes"},
	{2 * 1024, "2 kilobytes"},
	{2 * 1048576, "2 megabytes"},
	{2 * 1073741824, "2 gigabytes"},
	{2 * 1099511627776, "2 terabytes"},
	{2 * 1125899906842624, "2 petabytes"},
	{2 * 1152921504606846976, "2 exabytes"},
	{2 * 1180591620717411303424, "2 zettabytes"},
	{2 * 1208925819614629174706176, "2 yottabytes"},
}

func Test_GlobalFormat(t *testing.T) {
	Format = "%.0f "
	LongUnits = true
	for _, v := range globalFormatTable {
		b := New(v.Bytes)
		if b.String() != v.Result {
			t.Fatalf("Expected %s, received %s", v.Result, b)
		}
	}
	Format = "%.2f"
	LongUnits = false
}

var parseTable = []struct {
	Input  string
	Result string
	Fail   bool
}{
	{"1B", "1.00B", false},
	{"1 B", "1.00B", false},
	{"1 byte", "1.00B", false},
	{"2 bytes", "2.00B", false},
	{"1B ", "1.00B", false},
	{" 1 B ", "1.00B", false},
	{"1023B", "1023.00B", false},
	{"1024B", "1.00KB", false},
	{"1KB 1023B", "", true},
	{"1", "", true},
}

func Test_Parse(t *testing.T) {
	for _, v := range parseTable {
		b, err := Parse(v.Input)
		if err != nil && !v.Fail {
			t.Fatal(err)
		}
		if b.String() != v.Result && !v.Fail {
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
