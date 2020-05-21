package model

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func StartPicServer(async bool) {
	f := func() {
		http.HandleFunc("/upload/", getStickerFile) // 上传
		http.HandleFunc("/asset/", getStickerFile)  // //
		err := http.ListenAndServe(":12345", nil)
		fmt.Println(err)
	}

	if async {
		go f()
	} else {
		f()
	}
}

func getStickerFile(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r.URL.Path)
	file, err := os.Open("." + r.URL.Path)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	byte, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(byte)
}
