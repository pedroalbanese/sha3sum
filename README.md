# SHA3SUM(2)
[![ISC License](http://img.shields.io/badge/license-ISC-blue.svg)](https://github.com/pedroalbanese/sha3sum/blob/master/LICENSE.md) 
[![GoDoc](https://godoc.org/github.com/pedroalbanese/sha3sum?status.png)](http://godoc.org/github.com/pedroalbanese/sha3sum)
[![Go Report Card](https://goreportcard.com/badge/github.com/pedroalbanese/sha3sum)](https://goreportcard.com/report/github.com/pedroalbanese/sha3sum)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/pedroalbanese/sha3sum)](https://golang.org)
[![GitHub release (latest by date)](https://img.shields.io/github/v/release/pedroalbanese/sha3sum)](https://github.com/pedroalbanese/sha3sum/releases)
### sha3sum Parallel Implementation written in Go

<PRE>Usage of sha3sum:
sha3sum [-v] [-b N] [-c &lt;hash.ext&gt;] [-r] &lt;file.ext&gt;
  -b int
        Bits: 224, 256, 384 and 512. (default 256)
  -c string
        Check hashsum file.
  -r    Process directories recursively.
  -v    Verbose mode. (The exit code is always 0 in this mode)</PRE>
  
### Examples:

#### Generate hashsum list:
<pre>
$ ./sha3sum [-r] "*.*" > hash.txt
</pre>
##### Always works in binary mode. 

#### Check hashsum file:
<pre>
$ ./sha3sum [-v] -c hash.txt
</pre>
##### Exit code is always 0 in verbose mode. 

## License

This project is licensed under the ISC License.
##### Copyright (c) 2020-2021 Pedro Albanese - ALBANESE Lab.
