package main

import (
	"fmt"
	"net/url"
	"time"
	"bytes"
	"io"
	"net/http"
	"encoding/json"

	"github.com/gin-gonic/gin"
) 

func YTExtractor(videoId string) []byte {
	formData := url.Values{}
	formData.Set("videoId", videoId) 

	for i := 0; i<5; i++ {
		req, err := http.NewRequest("POST", "https://codingindialab.com/mp3-converter-tools/youtube-to-mp3/output.php", bytes.NewBufferString(formData.Encode()))
		if err != nil {
			panic(err)
		}

		// Set headers
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Set("Accept", "*/*")  // Example header   
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/137.0.0.0 Safari/537.36 Edg/137.0.0.0")           // Another custom header


		// Make the request
		client := &http.Client{}
		resp, err := client.Do(req)

		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)  
		}

		fmt.Println(string(bodyBytes))

		var result map[string]interface{}
		err = json.Unmarshal(bodyBytes, &result)
		if err != nil {
			fmt.Println("JSON unmarshal error:", err)
			return []byte{}
		}

		link, ok := result["link"].(string)
		if ( !ok  ) {
			fmt.Println("couldnot find link")
			return []byte{}
		}

		fmt.Println("llii ", link)

		if link != "" {
			return bodyBytes
		}
		time.Sleep(500 * time.Millisecond)
	}
	
	return []byte{}
	
}


func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "room.html", nil)
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"currentTime": time.Now().UnixMilli(),
		})
	}) 

	r.GET("/ytextract/:videoId", func(c *gin.Context) {
		videoId := c.Param("videoId")
		c.Data(200, "application/json", YTExtractor(videoId))
	})

	// Start server on :8080
	fmt.Println("server running on https://localhost:8080/") 

	r.Run() // Defaults to localhost:8080
}
