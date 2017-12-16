package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/cavaliercoder/grab"
	"github.com/spf13/viper"
)

func downloadHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Request Recieved")

	body, _ := ioutil.ReadAll(r.Body)

	downloadURL := string(body)

	_, err := url.ParseRequestURI(downloadURL)

	if err == nil {

		filePath := viper.GetString("DownloadFolder")

		client := grab.NewClient()
		req, _ := grab.NewRequest(filePath, downloadURL)

		// sum, err := hex.DecodeString("12767bda45b430d66e538a8780587260427935f7513479371dc2a884723ae410")
		// if err != nil {
		// 	panic(err)
		// }
		//
		// req.SetChecksum(sha256.New(), sum, true)

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

		// check for errors
		if err := resp.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "Download failed: %v\n", err)
			//os.Exit(1)
		}

		fmt.Printf("Download saved to %v \n", resp.Filename)
	}
}
