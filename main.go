package main

import (
	"fmt"
	"mt-mad/yt-downloader/models"
)

func main() {
    fmt.Println("hello nigga, its codeium")

    test,err := models.GetVideoInfo("https://www.youtube.com/shorts/kcPf7QgnENk")
    if err != nil {
        fmt.Println(err)
    }

    for i,item := range test {
        fmt.Println(i,item.Quality)
        if item.Quality == "1080p" {
            fmt.Println(item.GetDownloadLink())
        }
    }
}
