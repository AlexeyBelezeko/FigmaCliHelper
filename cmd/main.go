package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/AlexeyBelezeko/FigmaCliHelper/internal/figma"
)

var accessToken = flag.String("token", "", "figma personal access token")
var userName = flag.String("user", "", "figma account user name")
var fileKey = flag.String("file-key", "", "figma file key")

func main() {
	const baseURL = "https://api.figma.com/"
	flag.Parse()

	if accessToken == nil || *accessToken == "" {
		log.Fatalln("empty access token you can get it here https://www.figma.com/developers/api#authentication")
	}
	if fileKey == nil || *fileKey == "" {
		log.Fatalln("empty file key")
	}

	figmaClient := figma.NewClient(*accessToken, baseURL, &http.Client{})

	if userName == nil || *userName == "" {
		result, err := figmaClient.CountComments(*fileKey)
		if err != nil {
			log.Fatalln("failed to count comments: %w", err)
		}
		for k,v := range result {
			log.Println(k + ": ", v)
		}
	} else {
		counter, err := figmaClient.DeleteAllComments(*fileKey, *userName)
		if err != nil {
			log.Fatalln("failed to delete all comments: %w", err)
		}
		log.Printf("%d comments were successfully delted", counter)
	}
}