package main

import (
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/index", index)
	http.HandleFunc("/upload", upload)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println(err)
	}
}

func index(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/html;charset=utf-8")
	file, err := ioutil.ReadFile("./html/index.html")
	if err != nil {
		log.Println(err)
	}
	_, err = writer.Write(file)
	if err != nil {
		log.Println(err)
	}
}

func upload(writer http.ResponseWriter, request *http.Request) {
	// get image dstFile from form
	pictureFile, fileHeader, err := request.FormFile("upload_image")
	if err != nil {
		log.Println(err)
		return
	}
	// upload path
	uploadPath := "uploadaaa"
	uploadFullPath := uploadPath + fileHeader.Filename
	if _, err := os.Stat(uploadPath); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(uploadPath, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}
	dstFile, err := os.Create(uploadFullPath)
	if err != nil {
		log.Println(err)
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
