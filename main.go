/*
Copyright © 2024 Alex Whitmore heyimalexw@gmail.com
*/
package main

import (
	"log"

	"github.com/joho/godotenv"

	"azgo/cmd"
	_ "azgo/cmd/auth"
	_ "azgo/cmd/secrets"
	_ "azgo/cmd/static"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	cmd.Execute()
}
