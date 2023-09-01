package main

import "hyphen-backend-hellog/configuration"

func main() {
	// setup configuration
	config := configuration.New()
	configuration.NewDatabase(config)
}
