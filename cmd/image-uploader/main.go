package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/0xdod/imageuploadservice/internal/grpc"
	"github.com/spf13/viper"
)

var (
	cfgFile = flag.String("config", ".config.yaml", "The config file used")
)

// init reads in config file and ENV variables if set.
func init() {
	flag.Parse()

	if *cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(*cfgFile)
	} else {
		// Search config in current directory with name ".config" (without extension).
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		viper.SetConfigName(".config")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

func main() {
	s := grpc.NewServer()
	if err := s.Start(viper.GetString("port")); err != nil {
		log.Fatalf("cannot start server: %v", err)
	}
}
