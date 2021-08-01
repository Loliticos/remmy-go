package main

import (
  "./src/client"
  "github.com/joho/godotenv"
)

func main() {
  err := godotenv.Load()

  if err != nil {
    panic("Something happened with the .env file. Couldn't open that shit")
  }

  client.Start()
}
