package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/upload", ImageUpload)
	http.ListenAndServe("localhost:8080", nil)
}

func ImageUpload(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {

		file, header, err := req.FormFile("image")
		defer file.Close()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, "Bad request!")
			return
		}

		name := strings.Split(header.Filename, ".")
		extension := name[len(name)-1]
		if  extension != "JPEG" && extension != "jpeg" {
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, "Bad image extension!")
			return
		}
		fmt.Printf("File name %s\n", name[0])

		var Image bytes.Buffer
		fileSize, err := Image.ReadFrom(file)
		fmt.Printf("File size %d\n", fileSize)
		if fileSize > 8192 {
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, "Image is too large!")
			return
		}
		Image.Reset()
		io.WriteString(w, "Image uploaded successfully\n")
		return
	} else {
		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, "Not found!")
		return
	}
}
