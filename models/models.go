package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type YtVideo struct {
    Url string `json:"url"`
    Title string `json:"title"`
    Note string `json:"note"`
}

type Response struct {
	Status       string `json:"status"`
	DownloadUrlX string `json:"downloadUrlX,omitempty"`
	ID           string `json:"id,omitempty"`
	ErrorCode    string `json:"error_code,omitempty"`
}

func Downloadyt(vidInfo YtVideo) (string,error) {
    if vidInfo.Url == "" {
        return "", fmt.Errorf("No video URL provided")
    }


    parsedURL, err := url.Parse(vidInfo.Url)
    if err != nil {
        fmt.Println("Error parsing URL:", err)
        return "",err
    }

    if !strings.Contains(parsedURL.Host, "youtube.com") {
        return "", fmt.Errorf("Invalid video URL")
    }

    queryParams := parsedURL.Query()
    id := queryParams.Get("v")

    if id == "" {
        pathSegments := strings.Split(parsedURL.Path, "/")
		if len(pathSegments) > 0 {
		    id = pathSegments[len(pathSegments)-1]
		}
    }

    ext,note,format := "mp4","1080p","137"
    switch vidInfo.Note{
        case "1080p":
            ext = "mp4"
            note = "1080p"
            format = "137"
        case "720p":
            ext = "mp4"
            note = "720p"
            format = "136"
        case "360p":
            ext = "mp4"
            note = "360pp"
            format = "135"
        case "mp3":
            ext = "mp3"
            note = "128k"
            format = "128k"
    }
    fmt.Println(ext,note,format)
    requestBody := []byte(fmt.Sprintf(
    "platform=youtube&url=%s&title=%s&id=%s&ext=%s&note=%s&format=%s",
    vidInfo.Url,
    vidInfo.Title,
    id,
    ext,
    note,
    format,
    ))

    requestUrl := fmt.Sprintf("https://sss.instasaverpro.com/mates/en/convert?id=%s", id)

    req, err := http.NewRequest("POST",requestUrl,bytes.NewBuffer(requestBody))
    if err != nil {
        fmt.Println("Error creating request:", err)
        return "",err
    }

    req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:126.0) Gecko/20100101 Firefox/126.0")
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Set("Accept-Encoding", "")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Origin", "https://sss.instasaverpro.com")
	req.Header.Set("Referer", "https://sss.instasaverpro.com")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Te", "trailers")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("Error sending request:", err)
        return "",err
    }
    defer resp.Body.Close()

    fmt.Println("Response status:", resp.Status)

    var response Response
    err = json.NewDecoder(resp.Body).Decode(&response)
    if err != nil {
        fmt.Println("Error decoding response:", err)
        return "",err
    }

    if response.Status == "success" {
        return response.DownloadUrlX, nil
	} else {
        return "", fmt.Errorf("Error: %s", response.ErrorCode)
	}
}
