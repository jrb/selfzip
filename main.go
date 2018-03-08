package main

import (
	"archive/zip"
	"io"
	"log"
	"os"
)

func main() {
	path := os.Args[0]
	zipfile, err := os.Create(path + ".zip")
	if err != nil { log.Fatal(err) }
	defer zipfile.Close()

	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	info, err := os.Stat(path)
	if err != nil { log.Fatal(err) }

	header, err := zip.FileInfoHeader(info)
	if err != nil { log.Fatal(err) }

	header.Method = zip.Deflate

	writer, err := archive.CreateHeader(header)
	if err != nil { log.Fatal(err) }

	file, err := os.Open(path)
	if err != nil { log.Fatal(err) }
	defer file.Close()

	_, err = io.Copy(writer, file)
}
