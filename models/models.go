package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type YtVideoInfo struct {
    Key string `json:"k"`
    Quality string `json:"q"`
    Format string `json:"f"`
    Size string `json:"size"`
    Vid string `json:"vid"`
}

type GetInfoResponse struct {
	Status  string `json:"status"`
    Message string `json:"mess"`
    VideoId string `json:"vid"`
    Links   struct {
        Mp4 map[string]YtVideoInfo `json:"mp4"`
        Mp3 map[string]YtVideoInfo `json:"mp3"`
        Other map[string]YtVideoInfo `json:"other"`
    } `json:"links"`
}

func GetVideoInfo(VideoUrl string) ([]YtVideoInfo,error) {
    if VideoUrl == "" {
        return nil, fmt.Errorf("No video URL provided")
    }


    parsedURL, err := url.Parse(VideoUrl)
    if err != nil {
        fmt.Println("Error parsing URL:", err)
        return nil,err
    }

    if !strings.Contains(parsedURL.Host, "youtube.com") {
        return nil, fmt.Errorf("Invalid video URL")
    }

    requestBody := []byte(fmt.Sprintf("k_query=%s&k_page=home&hl=en&q_auto=0",VideoUrl))

    requestUrl := "https://www.y2mate.com/mates/en948/analyzeV2/ajax"

    req, err := http.NewRequest("POST",requestUrl,bytes.NewBuffer(requestBody))
    if err != nil {
        fmt.Println("Error creating request:", err)
        return nil,err
    }

    req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:126.0) Gecko/20100101 Firefox/126.0")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Set("Accept-Encoding", "")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Origin", "https://www.y2mate.com")
	req.Header.Set("Referer", "https://www.y2mate.com/en948")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Te", "trailers")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return nil,err
    }
    defer resp.Body.Close()


    var response GetInfoResponse
    err = json.NewDecoder(resp.Body).Decode(&response)
    if err != nil {
        return nil,err
    }

    var responseVids []YtVideoInfo

    if response.Status == "ok" {
        for _, vid := range response.Links.Mp4 {
            vid.Vid = response.VideoId
            responseVids = append(responseVids,vid)
        }
        for _, vid := range response.Links.Mp3 {
            vid.Vid = response.VideoId
            responseVids = append(responseVids,vid)
        }
        return responseVids,nil
	} else {
        return nil,nil
	}
}

type GetDownloadLinkResponse struct {
    Status  string `json:"status"`
    Message string `json:"mess"`
    Title string `json:"title"`
    Link string `json:"dlink"`
}

func (s * YtVideoInfo) GetDownloadLink() (string,error) {
    if s.Key == "" {
        return "", fmt.Errorf("No key provided")
    }
    
    requestBody := []byte(fmt.Sprintf("vid=%s&k=%s",s.Vid,s.Key))

    requestUrl := "https://www.y2mate.com/mates/convertV2/index"

    req, err := http.NewRequest("POST",requestUrl,bytes.NewBuffer(requestBody))
    if err != nil {
        return "",err
    }

    req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:126.0) Gecko/20100101 Firefox/126.0")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Set("Accept-Encoding", "")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Origin", "https://www.y2mate.com")
	req.Header.Set("Referer", "https://www.y2mate.com")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Te", "trailers")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return "",err
    }
    defer resp.Body.Close()

    var response GetDownloadLinkResponse
    err = json.NewDecoder(resp.Body).Decode(&response)
    if err != nil {
        return "",err
    }

    if response.Status == "ok" {
        return response.Link,nil
    } else {
        return "",fmt.Errorf(response.Message)
    }
    
}
