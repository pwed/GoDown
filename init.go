package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func init() {
	pflag.Bool("RestoreAssets", false, "Do you want to unpack static files to modify")
	pflag.Bool("Unpack", false, "Do you want to unpack static files to modify")
	pflag.Bool("Dev", false, "Do you want to automatically build angular dependencies")
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
	viper.SetDefault("Port", ":8080")
	viper.SetDefault("DownloadFolder", filepath.FromSlash("Downloads/"))
	viper.RegisterAlias("RestoreAssets", "Unpack")
	viper.SetConfigName("config")        // name of config file (without extension)
	viper.AddConfigPath("/etc/godown/")  // path to look for the config file in
	viper.AddConfigPath("$HOME/.godown") // call multiple times to add many search paths
	viper.AddConfigPath(".")             // optionally look for config in the working directory
	err := viper.ReadInConfig()          // Find and read the config file
	if err != nil {                      // Handle errors reading the config file
		fmt.Printf("Error in config file: %s \n", err)
	}
	os.Mkdir(viper.GetString("DownloadFolder"), 0777)
}
