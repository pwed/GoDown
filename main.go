package main

import (
	"fmt"
	"net/http"

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

	fmt.Println("Starting Server")

	m := mux.NewRouter()

	m.HandleFunc("/api/startDownload", downloadHandler)

	if viper.GetBool("LocalStaticFiles") {
		m.PathPrefix("/").Handler(http.FileServer(http.Dir("Angular/GoDown/dist")))
	} else {
		m.PathPrefix("/").Handler(http.FileServer(assetFS()))
	}

	fmt.Printf("Server running @ http//localhost%s\n", viper.GetString("Port"))

	http.ListenAndServe(viper.GetString("Port"), m)
}
