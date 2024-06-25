package main

import "log"

func (config *Config) setSavePath(path string) {
	if path == "" {
		log.Println("Empty input, set path needs a path string")
		return
	}
	config.SAVEPATH = path

}
