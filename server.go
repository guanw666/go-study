package main

import (
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
		return
	}
}

func index(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/html;charset=utf-8")
	file, err := ioutil.ReadFile("./html/index.html")
	if err != nil {
		log.Println(err)
		return
	}
	_, err = writer.Write(file)
	if err != nil {
		log.Println(err)
		return
	}
}

func upload(writer http.ResponseWriter, request *http.Request) {
	// get image dstFile from form
	pictureFile, fileHeader, err := request.FormFile("picture")
	if err != nil {
		log.Println(err)
		return
	}
	// create file
	dstFile, err2 := os.Create("upload/" + fileHeader.Filename)
	if err2 != nil {
		log.Println(err)
		return
	}
	// close file after return
	defer func(dstFile *os.File) {
		err := dstFile.Close()
		if err != nil {
			log.Println(err)
		}
	}(dstFile)
	// copy
	_, err = io.Copy(dstFile, pictureFile)
	if err != nil {
		log.Println(err)
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
