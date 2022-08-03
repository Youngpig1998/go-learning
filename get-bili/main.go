package main

import (
	"get-bili/downloader"
	myfmt "get-bili/fmt"
)

func main() {
	//fmt.Println("Hello world")

	//myfmt.Logger.Println("hello")

	request := downloader.InfoRequest{Bvids: []string{"BV1aY4y1P7Ej", "BV1Ga411S7mk"}}

	response, err := downloader.BatchDownloadVideoInfo(request)
	if err != nil {
		panic(err)
	}

	for _, info := range response.Infos {
		myfmt.Logger.Printf("title: %s \n desc: %s\n", info.Data.Title, info.Data.Desc)
	}

}
