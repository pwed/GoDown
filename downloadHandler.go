package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/cavaliercoder/grab"
	"github.com/spf13/viper"
)

type DownloadRequest struct {
	DownloadURL      string `json:"downloadURL"`
	DownloadChecksum string `json:"downloadChecksum"`
	HashType         string `json:"hashType"`
}

func downloadHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Request Recieved")

	body, _ := ioutil.ReadAll(r.Body)

	fmt.Printf("Body %v", string(body))

	var dr DownloadRequest

	if err := json.Unmarshal(body, &dr); err != nil {
		log.Fatal(err)
	}

	_, err := url.ParseRequestURI(dr.DownloadURL)

	if err == nil {

		filePath := viper.GetString("DownloadFolder")

		client := grab.NewClient()
		req, _ := grab.NewRequest(filePath, dr.DownloadURL)

		if dr.DownloadChecksum != "" {
			sum, err := hex.DecodeString(dr.DownloadChecksum)
			if err != nil {
				panic(err)
			}

			req.SetChecksum(md5.New(), sum, true)
		}
		fmt.Printf("Downloading %v...\n", req.URL())
		resp := client.Do(req)
		fmt.Printf("  %v\n", resp.HTTPResponse.Status)

		// start UI loop
		t := time.NewTicker(500 * time.Millisecond)
		defer t.Stop()

	Loop:
		for {
			select {
			case <-t.C:
				fmt.Printf("  transferred %v / %v bytes (%.2f%%)\n",
					resp.BytesComplete(),
					resp.Size,
					100*resp.Progress())

			case <-resp.Done:
				// download is complete
				break Loop
			}
		}
		response, _ := json.Marshal(resp)
		fmt.Printf("response %v", response)

		// check for errors
		if err := resp.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "Download failed: %v\n", err)
			//os.Exit(1)
		}

	}
}
