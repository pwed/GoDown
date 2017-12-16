package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func main() {

	if viper.GetBool("RestoreAssets") {
		RestoreAssets("./", "static")
		fmt.Println("Static folder unpacked")
	}

	fmt.Println("Starting Server")

	m := mux.NewRouter()

	m.HandleFunc("/download", downloadHandler)

	if viper.GetBool("LocalStaticFiles") {
		m.PathPrefix("/").Handler(http.FileServer(http.Dir("static")))
	} else {
		m.PathPrefix("/").Handler(http.FileServer(assetFS()))
	}

	fmt.Printf("Server running @ http//localhost%s\n", viper.GetString("Port"))

	http.ListenAndServe(viper.GetString("Port"), m)
}
