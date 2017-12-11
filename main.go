package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"os"
	"io/ioutil"
	"github.com/cavaliercoder/grab"
	"fmt"
	"time"
	"net/url"
	"encoding/hex"
	"crypto/sha256"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.AddConfigPath("/etc/appname/")   // path to look for the config file in
	viper.AddConfigPath("$HOME/.appname")  // call multiple times to add many search paths
	viper.AddConfigPath(".")               // optionally look for config in the working directory
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func main() {

	//RestoreAssets("./", "static")

	fmt.Println("Starting Server")

	m := mux.NewRouter()

	m.HandleFunc("/download", downloadHandler)

	m.PathPrefix("/").Handler(http.FileServer(assetFS()))

	//m.PathPrefix("/").Handler(http.FileServer(http.Dir("static")))

	fmt.Println("Server running on port 8080")

	http.ListenAndServe(":8080", m)
}

func downloadHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Request Recieved")

	body, _ := ioutil.ReadAll(r.Body)

	downloadURL := string(body)

	_, err := url.ParseRequestURI(downloadURL)

	if err == nil {

		filePath := viper.GetString("DownloadFolder")

		client := grab.NewClient()
		req, _ := grab.NewRequest(filePath, downloadURL)

		sum, err := hex.DecodeString("12767bda45b430d66e538a8780587260427935f7513479371dc2a884723ae410")
		if err != nil {
			panic(err)
		}


		req.SetChecksum(sha256.New(), sum, true)

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
