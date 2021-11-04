# SHA3SUM(2)
[![ISC License](http://img.shields.io/badge/license-ISC-blue.svg)](https://github.com/pedroalbanese/sha3sum/blob/master/LICENSE.md) 
[![GoDoc](https://godoc.org/github.com/pedroalbanese/whirlpoolsum?status.png)](http://godoc.org/github.com/pedroalbanese/sha3sum)
[![Go Report Card](https://goreportcard.com/badge/github.com/pedroalbanese/sha3sum)](https://goreportcard.com/report/github.com/pedroalbanese/sha3sum)
### sha3sum Parallel Implementation written in Go

<PRE>
SHA3 Hashsum Tool - ALBANESE Lab (c) 2020-2021

Usage of sha3sum:
sha3sum [-v] [-b N] [-c &lt;hash.ext&gt;] [-r] -t &lt;file.ext&gt;

  -b int
        Bits: 224, 256, 384 and 512. (default 256)
  -c string
        Check hashsum file.
  -r    Process directories recursively.
  -t string
        Target file/wildcard to generate hashsum list.
  -v    Verbose mode. (The exit code is always 0 in this mode)
  </PRE>
  
### Examples:

#### Generate hashsum list:
<pre>
$ ./sha3sum [-r] -t "*.*" > hash.txt
</pre>
##### Always works in binary mode. 

#### Check hashsum file:
<pre>
$ ./sha3sum [-v] -c hash.txt
</pre>
##### Exit code is always 0 in vebose mode. 

## License

This project is licensed under the ISC License.
##### Copyright (c) 2020-2021 Pedro Albanese - ALBANESE Lab.
