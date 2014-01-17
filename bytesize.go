// Package bytesize provides functionality for measuring and displaying byte
// sizes.
//
// You can perfom mathmatical operation with ByteSize's and the result will be
// a valid ByteSize with the correct size suffix. See the tests for examples.
package bytesize

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// This code was originall based on http://golang.org/doc/progs/eff_bytesize.go
// Since then many improvements have been made. The following is the original
// copyright notice:

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

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

var unitMap = map[string]float64{
	"B":     float64(B),
	"BYTE":  float64(B),
	"BYTES": float64(B),

	"KB":        float64(KB),
	"KILOBYTE":  float64(KB),
	"KILOBYTES": float64(KB),

	"MB":        float64(MB),
	"MEGABYTE":  float64(MB),
	"MEGABYTES": float64(MB),

	"GB":        float64(GB),
	"GIGABYTE":  float64(GB),
	"GIGABYTES": float64(GB),

	"TB":        float64(TB),
	"TERABYTE":  float64(TB),
	"TERABYTES": float64(TB),

	"PB":        float64(PB),
	"PETABYTE":  float64(PB),
	"PETABYTES": float64(PB),

	"EB":       float64(EB),
	"EXABYTE":  float64(EB),
	"EXABYTES": float64(EB),

	"ZB":         float64(ZB),
	"ZETTABYTE":  float64(ZB),
	"ZETTABYTES": float64(ZB),

	"YB":         float64(YB),
	"YOTTABYTE":  float64(YB),
	"YOTTABYTES": float64(YB),
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
// "MB", "GB", "TB", "PB", "EB", "ZB", and "YB".
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

	bytesize := ByteSize(value * unit)
	return bytesize, nil

}

// New returns a new ByteSize type set to s.
func New(s float64) ByteSize {
	return ByteSize(s)
}

// Returns a string representation of b with the specified formatting and units.
func (b ByteSize) Format(format string, longUnits bool) string {
	return b.format(format, longUnits)
}

// String returns the string form of b using the package global Format and
// LongUnits options.
func (b ByteSize) String() string {
	return b.format(Format, LongUnits)
}

func (b ByteSize) format(format string, longUnits bool) string {
	switch {
	case b >= YB:
		if longUnits {
			var s string
			value := fmt.Sprintf(format, b/YB)
			if printS, _ := strconv.ParseFloat(strings.TrimSpace(value), 64); printS > 1 {
				s = "s"
			}
			return fmt.Sprintf(format+"yottabyte"+s, b/YB)
		}
		return fmt.Sprintf(format+"YB", b/YB)
	case b >= ZB:
		if longUnits {
			var s string
			value := fmt.Sprintf(format, b/ZB)
			if printS, _ := strconv.ParseFloat(strings.TrimSpace(value), 64); printS > 1 {
				s = "s"
			}
			return fmt.Sprintf(format+"zettabyte"+s, b/ZB)
		}
		return fmt.Sprintf(format+"ZB", b/ZB)
	case b >= EB:
		if longUnits {
			var s string
			value := fmt.Sprintf(format, b/EB)
			if printS, _ := strconv.ParseFloat(strings.TrimSpace(value), 64); printS > 1 {
				s = "s"
			}
			return fmt.Sprintf(format+"exabyte"+s, b/EB)
		}
		return fmt.Sprintf(format+"EB", b/EB)
	case b >= PB:
		if longUnits {
			var s string
			value := fmt.Sprintf(format, b/PB)
			if printS, _ := strconv.ParseFloat(strings.TrimSpace(value), 64); printS > 1 {
				s = "s"
			}
			return fmt.Sprintf(format+"petabyte"+s, b/PB)
		}
		return fmt.Sprintf(format+"PB", b/PB)
	case b >= TB:
		if longUnits {
			var s string
			value := fmt.Sprintf(format, b/TB)
			if printS, _ := strconv.ParseFloat(strings.TrimSpace(value), 64); printS > 1 {
				s = "s"
			}
			return fmt.Sprintf(format+"terabyte"+s, b/TB)
		}
		return fmt.Sprintf(format+"TB", b/TB)
	case b >= GB:
		if longUnits {
			var s string
			value := fmt.Sprintf(format, b/GB)
			if printS, _ := strconv.ParseFloat(strings.TrimSpace(value), 64); printS > 1 {
				s = "s"
			}
			return fmt.Sprintf(format+"gigabyte"+s, b/GB)
		}
		return fmt.Sprintf(format+"GB", b/GB)
	case b >= MB:
		if longUnits {
			var s string
			value := fmt.Sprintf(format, b/MB)
			if printS, _ := strconv.ParseFloat(strings.TrimSpace(value), 64); printS > 1 {
				s = "s"
			}
			return fmt.Sprintf(format+"megabyte"+s, b/MB)
		}
		return fmt.Sprintf(format+"MB", b/MB)
	case b >= KB:
		if longUnits {
			var s string
			value := fmt.Sprintf(format, b/KB)
			if printS, _ := strconv.ParseFloat(strings.TrimSpace(value), 64); printS > 1 {
				s = "s"
			}
			return fmt.Sprintf(format+"kilobyte"+s, b/KB)
		}
		return fmt.Sprintf(format+"KB", b/KB)
	}
	if longUnits {
		var s string
		if b > 1 {
			s = "s"
		}
		return fmt.Sprintf(format+"byte"+s, b)
	}
	return fmt.Sprintf(format+"B", b)
}
