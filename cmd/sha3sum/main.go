package main

import (
	"bufio"
	"encoding/hex"
	"flag"
	"fmt"
	"hash"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/crypto/sha3"
)

var (
	bits      = flag.Int("b", 256, "Bits: 224, 256, 384 and 512.")
	check     = flag.String("c", "", "Check hashsum file.")
	recursive = flag.Bool("r", false, "Process directories recursively.")
)

func main() {
	flag.Parse()

	if (len(os.Args) < 2) || (*bits != 224 && *bits != 256 && *bits != 384 && *bits != 512) {
		fmt.Println("SHA3 Hashsum Tool - ALBANESE Lab (c) 2020-2021\n")
		fmt.Println("Usage of", os.Args[0]+":")
		fmt.Printf("%s [-v] [-b N] [-c <hash.ext>] [-r] <file.ext>\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}

	Files := strings.Join(flag.Args(), " ")

	if Files == "-" {
		var h hash.Hash
		if *bits == 224 {
			h = sha3.New224()
		} else if *bits == 256 {
			h = sha3.New256()
		} else if *bits == 384 {
			h = sha3.New384()
		} else if *bits == 512 {
			h = sha3.New512()
		}
		io.Copy(h, os.Stdin)
		fmt.Println(hex.EncodeToString(h.Sum(nil)) + " (stdin)")
		os.Exit(0)
	}

	if *check == "" && *recursive == false {
		for _, wildcard := range flag.Args() {
			files, err := filepath.Glob(wildcard)
			if err != nil {
				log.Fatal(err)
			}
			for _, match := range files {
				var h hash.Hash
				if *bits == 224 {
					h = sha3.New224()
				} else if *bits == 256 {
					h = sha3.New256()
				} else if *bits == 384 {
					h = sha3.New384()
				} else if *bits == 512 {
					h = sha3.New512()
				}
				f, err := os.Open(match)
				if err != nil {
					log.Fatal(err)
				}
				file, err := os.Stat(match)
				if file.IsDir() {
				} else {
					if _, err := io.Copy(h, f); err != nil {
						log.Fatal(err)
					}
					fmt.Println(hex.EncodeToString(h.Sum(nil)), "*"+f.Name())
				}
				f.Close()
			}
		}
		os.Exit(0)
	}

	if *check == "" && *recursive == true {
		err := filepath.Walk(filepath.Dir(Files),
			func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				file, err := os.Stat(path)
				if file.IsDir() {
				} else {
					for _, match := range flag.Args() {
						filename := filepath.Base(path)
						pattern := filepath.Base(match)
						matched, err := filepath.Match(pattern, filename)
						if err != nil {
							fmt.Println(err)
						}
						if matched {
							var h hash.Hash
							if *bits == 224 {
								h = sha3.New224()
							} else if *bits == 256 {
								h = sha3.New256()
							} else if *bits == 384 {
								h = sha3.New384()
							} else if *bits == 512 {
								h = sha3.New512()
							}
							f, err := os.Open(path)
							if err != nil {
								log.Fatal(err)
							}
							if _, err := io.Copy(h, f); err != nil {
								log.Fatal(err)
							}
							f.Close()
							fmt.Println(hex.EncodeToString(h.Sum(nil)), "*"+f.Name())
						}
					}
				}
				return nil
			})
		if err != nil {
			log.Println(err)
		}
	}

	if *check != "" {
		var file io.Reader
		var err error
		if *check == "-" {
			file = os.Stdin
		} else {
			file, err = os.Open(*check)
			if err != nil {
				log.Fatalf("failed opening file: %s", err)
			}
		}
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		var txtlines []string

		for scanner.Scan() {
			txtlines = append(txtlines, scanner.Text())
		}
		var exit int
		for _, eachline := range txtlines {
			lines := strings.Split(string(eachline), " *")
			if strings.Contains(string(eachline), " *") {
				var h hash.Hash
				if len(lines[0])*4 == 224 {
					h = sha3.New224()
				} else if len(lines[0])*4 == 256 {
					h = sha3.New256()
				} else if len(lines[0])*4 == 384 {
					h = sha3.New384()
				} else if len(lines[0])*4 == 512 {
					h = sha3.New512()
				}
				_, err := os.Stat(lines[1])
				if err == nil {
					f, err := os.Open(lines[1])
					if err != nil {
						log.Fatal(err)
					}
					io.Copy(h, f)
					f.Close()

					if hex.EncodeToString(h.Sum(nil)) == lines[0] {
						fmt.Println(lines[1]+"\t", "OK")
					} else {
						fmt.Println(lines[1]+"\t", "FAILED")
						exit = 1
					}
				} else {
					fmt.Println(lines[1]+"\t", "Not found!")
					exit = 1
				}
			}
		}
		os.Exit(exit)
	}
}
