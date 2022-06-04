# SHA3SUM(2)
[![ISC License](http://img.shields.io/badge/license-ISC-blue.svg)](https://github.com/pedroalbanese/sha3sum/blob/master/LICENSE.md) 
[![GoDoc](https://godoc.org/github.com/pedroalbanese/sha3sum?status.png)](http://godoc.org/github.com/pedroalbanese/sha3sum)
[![GitHub downloads](https://img.shields.io/github/downloads/pedroalbanese/sha3sum/total.svg?logo=github&logoColor=white)](https://github.com/pedroalbanese/sha3sum/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/pedroalbanese/sha3sum)](https://goreportcard.com/report/github.com/pedroalbanese/sha3sum)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/pedroalbanese/sha3sum)](https://golang.org)
[![GitHub release (latest by date)](https://img.shields.io/github/v/release/pedroalbanese/sha3sum)](https://github.com/pedroalbanese/sha3sum/releases)
### sha3sum Parallel Implementation written in Go

<PRE>Usage of sha3sum:
sha3sum [-c &lt;hash.ext&gt;] [-b N] [-r] &lt;file.ext&gt;
  -b int
        Bits: 224, 256, 384 and 512. (default 256)
  -c string
        Check hashsum file.
  -r    Process directories recursively.</PRE>
  
### Examples:

#### Generate hashsum list:
<pre>
$ ./sha3sum [-r] "*.*" > hash.txt
</pre>
##### Always works in binary mode. 

#### Check hashsum file:
<pre>
$ ./sha3sum -c hash.txt
$ echo $?
</pre>

## License

This project is licensed under the ISC License.
##### Copyright (c) 2020-2022 Pedro F. Albanese - ALBANESE Research Lab.
