package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"os"
	"io"
	"io/ioutil"
	"github.com/cavaliercoder/grab"
	"fmt"
	"time"
	"net/url"
)

func main() {

	//RestoreAssets("./", "")

	m := mux.NewRouter()

	m.HandleFunc("/download", func(w http.ResponseWriter, r *http.Request) {
		body, _ := ioutil.ReadAll(r.Body)

		downloadURL := string(body)

		_, err := url.ParseRequestURI(downloadURL)

		if err == nil {

			filePath := "C:\\Data\\"

			client := grab.NewClient()
			req, _ := grab.NewRequest(filePath, downloadURL)

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

			fmt.Printf("Download saved to ./%v \n", resp.Filename)
		}
	})

	m.PathPrefix("/").Handler(http.FileServer(assetFS()))

	http.ListenAndServe(":8080", m)
}

func downloadFile(filepath string, url string) (err error) {

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}