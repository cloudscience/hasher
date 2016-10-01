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

func main() {
	if len(os.Args) == 0 {
		fmt.Println("No arguments found")
		os.Exit(255) // kills xarg from further processing
	}
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
		f.MD5hash = getMD5(infile)
		f.SHA256hash = getSHA256(infile)
		f.SHA512hash = getSHA512(infile)

		// Pretty print it as JSON
		j, _ := json.MarshalIndent(f, "", "  ")
		fmt.Println(string(j))
	}

}

func getMD5(fileptr *os.File) string {
	// Make sure streaming file read is done from the beginning, every time.
	_, err := fileptr.Seek(0, os.SEEK_SET)
	if err != nil {
		fmt.Println(err)
	}
	md5h := md5.New()
	io.Copy(md5h, fileptr)
	return hex.EncodeToString(md5h.Sum(nil))
}

func getSHA256(fileptr *os.File) string {
	// Make sure streaming file read is done from the beginning, every time.
	_, err := fileptr.Seek(0, os.SEEK_SET)
	if err != nil {
		fmt.Println(err)
	}
	sha256hasher := sha256.New()
	io.Copy(sha256hasher, fileptr)
	return hex.EncodeToString(sha256hasher.Sum(nil))
}

func getSHA512(fileptr *os.File) string {
	// Make sure streaming file read is done from the beginning, every time.
	_, err := fileptr.Seek(0, os.SEEK_SET)
	if err != nil {
		fmt.Println(err)
	}
	sha512hasher := sha512.New()
	io.Copy(sha512hasher, fileptr)
	return hex.EncodeToString(sha512hasher.Sum(nil))
}
