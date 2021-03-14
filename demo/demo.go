package main

import (
	"fmt"
	"os"

	"danilo/dotenv"
)

func main() {
	dotenv.Load()
	fmt.Println(os.Getenv("APP_USER"))
}
