# SHA3SUM(2)
## sha3sum Parallel Implementation written in Go

<PRE>
SHA3 Keccak Hashsum Tool - ALBANESE Lab (c) 2020-2021

Usage of sha3sum:
sha3sum [-v] [-b N] [-c <hash.ext>] -t &lt;file.ext&gt;

  -b int
        Bits: 224, 256, 384 and 512. (default 224)
  -c string
        Check hashsum file.
  -t string
        Target file/wildcard to generate hashsum list.
  -v    Verbose mode. (The exit code is always 0 in this mode)
  </PRE>
  
### Examples:

#### Generate hashsum list:
<pre>
./gostsum -t "*.*" > hash.txt
</pre>

#### Generate recursive hashsum list:
<pre>
$ find . -type f -name "*.*" -exec ./gostsum -t '{}' \; > hash.txt 
</pre>
##### Always works in binary mode. 

#### Check hashsum file:
<pre>
./gostsum [-v] -c hash.txt
</pre>
##### Exit code is always 0 in vebose mode. 
