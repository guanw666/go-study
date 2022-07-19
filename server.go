package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/index", index)
	http.HandleFunc("/upload", upload)

	err := http.ListenAndServe(":8811", nil)
	if err != nil {
		return
	}
}

func index(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/html;charset=utf-8")
	file, err := ioutil.ReadFile("./html/index.html")
	if err != nil {
		return
	}
	_, err = writer.Write(file)
	if err != nil {
		return
	}
}

func upload(writer http.ResponseWriter, request *http.Request) {
	// get image dstFile from form
	pictureFile, fileHeader, err := request.FormFile("dstFile")
	if err != nil {
		return
	}
	// upload path
	uploadFullPath := "./upload/" + fileHeader.Filename
	dstFile, err := os.Create(uploadFullPath)
	if err != nil {
		return
	}
	defer dstFile.Close()
	// save
	_, err = io.Copy(dstFile, pictureFile)
	if err != nil {
		return
	}
	//dstFile, err = ioutil.ReadFile("./img/Snipaste_2022-07-19_16-18-12.jpg")
	//if err != nil {
	//	return
	//}
	//writer.Header().Set("Content-Type", "image/png")
	//_, err = writer.Write(dstFile)
	//if err != nil {
	//	return
	//}
}
