package main

import (
	"fmt"
	"mt-mad/yt-downloader/models"
)

func main() {
    fmt.Println("hello nigga, its codeium")


    vidInfo := models.YtVideo{
        Url: "https://www.youtube.com/watch?v=lcpAXz0ixuI",
        Title: "tetepkq",
        Note: "360p",
    }

    test,err := models.Downloadyt(vidInfo)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(test)
}
