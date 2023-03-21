package main

import (
	"fmt"
	"testing"

	"github.com/spf13/viper"
)

func TestDiff(t *testing.T) {
	viper.SetConfigFile("config.yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
	Diff(nil, nil)

}
