ByteSize
========

This is a fork of https://github.com/inhies/go-bytesize
This ByteSize fork uses an uint64 instead of a float64 for storing the byte sizes. As bytes are always natural numbers
an integer representation is desirable. You do not have to deal with floating point inaccuracy. The disadvantage
is that you are limited to sizes up to 2^64. So this fork does not support ZB and YB sizes. The biggest supported
size is 18.4EB.

Bytesize is a package for working with measurements of bytes. Using this package
allows you to easily add 100KB to 4928MB and get a nicely formatted string
representation of the result.

[![GoDoc](http://godoc.org/github.com/MalteJ/go-bytesize?status.png)](http://godoc.org/github.com/MalteJ/go-bytesize)
[![Build Status](https://travis-ci.org/MalteJ/go-bytesize.png)](https://travis-ci.org/MalteJ/go-bytesize)
[![Coverage Status](https://coveralls.io/repos/MalteJ/go-bytesize/badge.svg?branch=master&service=github)](https://coveralls.io/github/MalteJ/go-bytesize?branch=master)

Usage
-----

Check the built in documentation for examples using the godoc link above or by
running godoc locally. 
