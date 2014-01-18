// Package bytesize provides functionality for measuring and formatting  byte
// sizes.
//
// You can also perfom mathmatical operation with ByteSize's and the result
// will be a valid ByteSize with the correct size suffix.
package bytesize

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// This code was originally based on http://golang.org/doc/progs/eff_bytesize.go
//
// Since then many improvements have been made. The following is the original
// copyright notice:

// Copyright 2009 The Go Authors. All rights reserved. Use of this source code
// is governed by a BSD-style license that can be found in the LICENSE file.

// ByteSize represents a number of bytes
type ByteSize float64

// Byte size size suffixes.
const (
	B  ByteSize = 1
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

// Used for returning long unit form of string representation.
var longUnitMap = map[ByteSize]string{
	B:  "byte",
	KB: "kilobyte",
	MB: "megabyte",
	GB: "gigabyte",
	TB: "terabyte",
	PB: "petabyte",
	EB: "exabyte",
	ZB: "zettabyte",
	YB: "yottabyte",
}

// Used for returning string representation.
var shortUnitMap = map[ByteSize]string{
	B:  "B",
	KB: "KB",
	MB: "MB",
	GB: "GB",
	TB: "TB",
	PB: "PB",
	EB: "EB",
	ZB: "ZB",
	YB: "YB",
}

// Used to convert user input to ByteSize
var unitMap = map[string]ByteSize{
	"B":     B,
	"BYTE":  B,
	"BYTES": B,

	"KB":        KB,
	"KILOBYTE":  KB,
	"KILOBYTES": KB,

	"MB":        MB,
	"MEGABYTE":  MB,
	"MEGABYTES": MB,

	"GB":        GB,
	"GIGABYTE":  GB,
	"GIGABYTES": GB,

	"TB":        TB,
	"TERABYTE":  TB,
	"TERABYTES": TB,

	"PB":        PB,
	"PETABYTE":  PB,
	"PETABYTES": PB,

	"EB":       EB,
	"EXABYTE":  EB,
	"EXABYTES": EB,

	"ZB":         ZB,
	"ZETTABYTE":  ZB,
	"ZETTABYTES": ZB,

	"YB":         YB,
	"YOTTABYTE":  YB,
	"YOTTABYTES": YB,
}

var (
	// Use long units, such as "megabytes" instead of "MB".
	LongUnits bool = false

	// String format of bytesize output. The unit of measure will be appended
	// to the end. Uses the same formatting options as the fmt package.
	Format string = "%.2f"
)

// Parse parses a byte size string. A byte size string is a number followed by
// a unit suffix, such as "1024B" or "1 MB". Valid byte units are "B", "KB",
// "MB", "GB", "TB", "PB", "EB", "ZB", and "YB". You can also use the long
// format of units, such as "kilobyte" or "kilobytes".
func Parse(s string) (ByteSize, error) {
	// Remove leading and trailing whitespace
	s = strings.TrimSpace(s)

	split := make([]string, 0)
	for i, r := range s {
		if !unicode.IsDigit(r) {
			// Split the string by digit and size designator, remove whitespace
			split = append(split, strings.TrimSpace(string(s[:i])))
			split = append(split, strings.TrimSpace(string(s[i:])))
			break
		}
	}

	// Check to see if we split successfully
	if len(split) != 2 {
		return 0, errors.New("Unrecognized size suffix")
	}

	// Check for MB, MEGABYTE, and MEGABYTES
	unit, ok := unitMap[strings.ToUpper(split[1])]
	if !ok {
		return 0, errors.New("Unrecognized size suffix " + split[1])

	}

	value, err := strconv.ParseFloat(split[0], 64)
	if err != nil {
		return 0, err
	}

	bytesize := ByteSize(value * float64(unit))
	return bytesize, nil

}

// New returns a new ByteSize type set to s.
func New(s float64) ByteSize {
	return ByteSize(s)
}

// Returns a string representation of b with the specified formatting and units.
func (b ByteSize) Format(format string, unit string, longUnits bool) string {
	return b.format(format, unit, longUnits)
}

// String returns the string form of b using the package global Format and
// LongUnits options.
func (b ByteSize) String() string {
	return b.format(Format, "", LongUnits)
}

func (b ByteSize) format(format string, unit string, longUnits bool) string {
	var unitSize ByteSize
	if unit != "" {
		var ok bool
		unitSize, ok = unitMap[strings.ToUpper(unit)]
		if !ok {
			return "Unrecognized unit: " + unit
		}
	} else {
		switch {
		case b >= YB:
			unitSize = YB
		case b >= ZB:
			unitSize = ZB
		case b >= EB:
			unitSize = EB
		case b >= PB:
			unitSize = PB
		case b >= TB:
			unitSize = TB
		case b >= GB:
			unitSize = GB
		case b >= MB:
			unitSize = MB
		case b >= KB:
			unitSize = KB
		default:
			unitSize = B
		}
	}

	if longUnits {
		var s string
		value := fmt.Sprintf(format, b/unitSize)
		if printS, _ := strconv.ParseFloat(strings.TrimSpace(value), 64); printS > 0 && printS != 1 {
			s = "s"
		}
		return fmt.Sprintf(format+longUnitMap[unitSize]+s, b/unitSize)
	}
	return fmt.Sprintf(format+shortUnitMap[unitSize], b/unitSize)
}
