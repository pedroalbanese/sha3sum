//go:generate goversioninfo -manifest=testdata/resource/goversioninfo.exe.manifest
package main
import (
	"bufio"
	"encoding/hex"
	"flag"
	"fmt"
	"golang.org/x/crypto/sha3"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

	var check = flag.String("c", "", "Check hashsum file.")
	var bits = flag.Int("b", 224, "Bits: 224, 256, 384 and 512.")
	var target = flag.String("t", "", "Target file/wildcard to generate hashsum list.")
	var verbose = flag.Bool("v", false, "Verbose mode. (The exit code is always 0 in this mode)")

func main() {
    flag.Parse()

        if (len(os.Args) < 2) {
	fmt.Println("SHA3 Keccak Hashsum Tool - ALBANESE Lab (c) 2020-2021\n")
	fmt.Println("Usage of",os.Args[0]+":")
        fmt.Printf("%s [-v] [-b N] [-c <hash.ext>] -t <file.ext>\n\n", os.Args[0])
        flag.PrintDefaults()
        os.Exit(1)
        }


        if *target != "" && *bits == 224 {
	files, err := filepath.Glob(*target)
	if err != nil {
	    log.Fatal(err)
	}

	for _, match := range files {
	h := sha3.New224()
        f, err := os.Open(match)
        if err != nil {
            log.Fatal(err)
        }
        if _, err := io.Copy(h, f); err != nil {
            log.Fatal(err)
        }
    	fmt.Println(hex.EncodeToString(h.Sum(nil)), "*" + f.Name())
	}
	}


        if *target != "" && *bits == 256 {
	files, err := filepath.Glob(*target)
	if err != nil {
	    log.Fatal(err)
	}

	for _, match := range files {
	h := sha3.New256()
        f, err := os.Open(match)
        if err != nil {
            log.Fatal(err)
        }
        if _, err := io.Copy(h, f); err != nil {
            log.Fatal(err)
        }
    	fmt.Println(hex.EncodeToString(h.Sum(nil)), "*" + f.Name())
	}
	}


        if *target != "" && *bits == 384 {
	files, err := filepath.Glob(*target)
	if err != nil {
	    log.Fatal(err)
	}

	for _, match := range files {
	h := sha3.New384()
        f, err := os.Open(match)
        if err != nil {
            log.Fatal(err)
        }
        if _, err := io.Copy(h, f); err != nil {
            log.Fatal(err)
        }
    	fmt.Println(hex.EncodeToString(h.Sum(nil)), "*" + f.Name())
	}
	}


        if *target != "" && *bits == 512 {
	files, err := filepath.Glob(*target)
	if err != nil {
	    log.Fatal(err)
	}

	for _, match := range files {
	h := sha3.New512()
        f, err := os.Open(match)
        if err != nil {
            log.Fatal(err)
        }
        if _, err := io.Copy(h, f); err != nil {
            log.Fatal(err)
        }
    	fmt.Println(hex.EncodeToString(h.Sum(nil)), "*" + f.Name())
	}
	}


        if *check != "" && *bits == 224 {
	file, err := os.Open(*check)
 
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string
 
	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}
	file.Close()
	for _, eachline := range txtlines {
	lines := strings.Split(string(eachline), " *")
	if strings.Contains(string(eachline), " *") {
	h := sha3.New224()
	_, err := os.Stat(lines[1])
	if err == nil {
		f, err := os.Open(lines[1])
		if err != nil {
		     log.Fatal(err)
		}
		io.Copy(h, f)
		
		if *verbose {
			if hex.EncodeToString(h.Sum(nil)) == lines[0] {
				fmt.Println(lines[1] + "\t", "OK")
			} else {
				fmt.Println(lines[1] + "\t", "FAILED")
			}
		} else {
			if hex.EncodeToString(h.Sum(nil)) == lines[0] {
			} else {
				os.Exit(1)
			}
		}
	} else {
		if *verbose {
			fmt.Println(lines[1] + "\t", "Not found!")
		} else {
			os.Exit(1)	
		}	
	}
	}
	}

	}


        if *check != "" && *bits == 256 {
	file, err := os.Open(*check)
 
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string
 
	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}
	file.Close()
	for _, eachline := range txtlines {
	lines := strings.Split(string(eachline), " *")
	if strings.Contains(string(eachline), " *") {
	h := sha3.New256()
	_, err := os.Stat(lines[1])
	if err == nil {
		f, err := os.Open(lines[1])
		if err != nil {
		     log.Fatal(err)
		}
		io.Copy(h, f)
		
		if *verbose {
			if hex.EncodeToString(h.Sum(nil)) == lines[0] {
				fmt.Println(lines[1] + "\t", "OK")
			} else {
				fmt.Println(lines[1] + "\t", "FAILED")
			}
		} else {
			if hex.EncodeToString(h.Sum(nil)) == lines[0] {
			} else {
				os.Exit(1)
			}
		}
	} else {
		if *verbose {
			fmt.Println(lines[1] + "\t", "Not found!")
		} else {
			os.Exit(1)	
		}	
	}
	}
	}

	}


        if *check != "" && *bits == 384 {
	file, err := os.Open(*check)
 
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string
 
	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}
	file.Close()
	for _, eachline := range txtlines {
	lines := strings.Split(string(eachline), " *")
	if strings.Contains(string(eachline), " *") {
	h := sha3.New384()
	_, err := os.Stat(lines[1])
	if err == nil {
		f, err := os.Open(lines[1])
		if err != nil {
		     log.Fatal(err)
		}
		io.Copy(h, f)
		
		if *verbose {
			if hex.EncodeToString(h.Sum(nil)) == lines[0] {
				fmt.Println(lines[1] + "\t", "OK")
			} else {
				fmt.Println(lines[1] + "\t", "FAILED")
			}
		} else {
			if hex.EncodeToString(h.Sum(nil)) == lines[0] {
			} else {
				os.Exit(1)
			}
		}
	} else {
		if *verbose {
			fmt.Println(lines[1] + "\t", "Not found!")
		} else {
			os.Exit(1)	
		}	
	}
	}
	}

	}


        if *check != "" && *bits == 512 {
	file, err := os.Open(*check)
 
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string
 
	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}
	file.Close()
	for _, eachline := range txtlines {
	lines := strings.Split(string(eachline), " *")
	if strings.Contains(string(eachline), " *") {
	h := sha3.New512()
	_, err := os.Stat(lines[1])
	if err == nil {
		f, err := os.Open(lines[1])
		if err != nil {
		     log.Fatal(err)
		}
		io.Copy(h, f)
		
		if *verbose {
			if hex.EncodeToString(h.Sum(nil)) == lines[0] {
				fmt.Println(lines[1] + "\t", "OK")
			} else {
				fmt.Println(lines[1] + "\t", "FAILED")
			}
		} else {
			if hex.EncodeToString(h.Sum(nil)) == lines[0] {
			} else {
				os.Exit(1)
			}
		}
	} else {
		if *verbose {
			fmt.Println(lines[1] + "\t", "Not found!")
		} else {
			os.Exit(1)	
		}	
	}
	}
	}

	}
}
