/**
* (c) Cloud Science LLC
* Contributor: Adam Crosby (adam@)
* See LICENSE file for full license information
**/
package main

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"hash"
	"io"
	"os"
	"path/filepath"
)

// Fileinfo holds info on the files as it's collected.
type Fileinfo struct {
	SHA256hash string
	SHA512hash string
	MD5hash    string
	Filename   string
	Filepath   string
}

// Constants for use in hash function
const (
	SHA512 = "sha512"
	SHA256 = "sha256"
	MD5    = "md5"
)

func main() {
	if len(os.Args) == 0 {
		fmt.Println("No arguments found")
		os.Exit(255) // kills xarg from further processing
	}
	var hashes []Fileinfo
	files := os.Args[1:] // chop of the program name itself from the list of files passed in via xargs
	for file := range files {
		infile, inerr := os.Open(files[file])
		if inerr != nil {
			fmt.Println(inerr)
			os.Exit(1)
		}
		// Build struct for file, fill it with info
		var f Fileinfo
		f.Filepath, f.Filename = filepath.Split(files[file])
		f.MD5hash = hasher(infile, MD5)
		f.SHA256hash = hasher(infile, SHA256)
		f.SHA512hash = hasher(infile, SHA512)

		hashes = append(hashes, f)
	}
	// Pretty print it as JSON
	j, _ := json.MarshalIndent(hashes, "", "  ")
	fmt.Println(string(j))

}

func hasher(fileptr *os.File, hashtype string) string {
	// Make sure streaming file read is done from the beginning, every time.
	_, err := fileptr.Seek(0, os.SEEK_SET)
	if err != nil {
		fmt.Println(err)
	}

	var hasheng hash.Hash
	switch hashtype {
	case SHA512:
		hasheng = sha512.New()
	case SHA256:
		hasheng = sha256.New()
	case MD5:
		hasheng = md5.New()
	default:
		panic("Wrong kind of hash function called.")
	}

	io.Copy(hasheng, fileptr)
	return hex.EncodeToString(hasheng.Sum(nil))
}
