package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func main()  {
	err := godotenv.Load();
	if err != nil {
		fmt.Print("Error loading .env file: %s\n",err.Error())
	}

	fmt.Println(os.Getenv("APP_MESSAGE"))
}