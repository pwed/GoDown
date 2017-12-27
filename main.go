package main

import (
	"fmt"
	"net/http"
	"os/exec"

	"github.com/fsnotify/fsnotify"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func main() {

	if viper.GetBool("RestoreAssets") {
		RestoreAssets("./", "static")
		fmt.Println("Static folder unpacked")
	}

	if viper.GetBool("WatchConfig") {
		viper.WatchConfig()
		viper.OnConfigChange(func(e fsnotify.Event) {
			fmt.Println("Config file changed:", e.Name)
		})
	}

	if viper.GetBool("Dev") {
		ng := exec.Command("ng", "build", "--watch")
		ng.Dir = "gui"
		ng.Start()
		viper.Set("LocalStaticFiles", true)
		fmt.Println("Server is running in Development mode and will serve static files locally and rebuild them when changes are detected. \nFirst build may take a few seconds, please be patient on slow systems")
	}

	fmt.Println("Starting Server")

	m := mux.NewRouter()

	m.HandleFunc("/api/startDownload", downloadHandler)

	if viper.GetBool("LocalStaticFiles") {
		m.PathPrefix("/").Handler(http.FileServer(http.Dir("static")))
	} else {
		m.PathPrefix("/").Handler(http.FileServer(assetFS()))
	}

	fmt.Printf("Server running @ http//localhost%s\n", viper.GetString("Port"))

	http.ListenAndServe(viper.GetString("Port"), m)
}
