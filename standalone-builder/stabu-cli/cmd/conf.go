/*
Copyright © 2023 NAME HERE sudogauss
*/
package cmd

import (
	// "fmt"
	"errors"
	"os"
	// "gopkg.in/yaml.v3"
)

type Config struct {
	Path struct {
		Workspace            string   `yaml:"workspace"`
		CompilersProj        string   `yaml:"compilers_projects"`
		ExecutablesProj      string   `yaml:"executables_projects"`
		AvailableCompilers   []string `yaml:"available_compilers"`
		AvailableExecutables []string `yaml:"available_executables"`
	} `yaml:"path"`
}

func createConfig() {
	config_path := os.Getenv("HOME") + ".local/stabu"

	_, err := os.Stat(config_path)
	if err != nil {

	}

	if _, err := os.Stat(config_path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(config_path, os.ModePerm)
		if err != nil {

		}
	}
}
