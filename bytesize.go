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
	"B":  float64(B),
	"KB": float64(KB),
	"MB": float64(MB),
	"GB": float64(GB),
	"TB": float64(TB),
	"PB": float64(PB),
	"EB": float64(EB),
	"ZB": float64(ZB),
	"YB": float64(YB),
}

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

// String returns the string form of byte size with the prefix, such as
// "2.00MB".
func (b ByteSize) String() string {
	switch {
	case b >= YB:
		return fmt.Sprintf("%.2fYB", b/YB)
	case b >= ZB:
		return fmt.Sprintf("%.2fZB", b/ZB)
	case b >= EB:
		return fmt.Sprintf("%.2fEB", b/EB)
	case b >= PB:
		return fmt.Sprintf("%.2fPB", b/PB)
	case b >= TB:
		return fmt.Sprintf("%.2fTB", b/TB)
	case b >= GB:
		return fmt.Sprintf("%.2fGB", b/GB)
	case b >= MB:
		return fmt.Sprintf("%.2fMB", b/MB)
	case b >= KB:
		return fmt.Sprintf("%.2fKB", b/KB)
	}
	return fmt.Sprintf("%.2fB", b)
}
