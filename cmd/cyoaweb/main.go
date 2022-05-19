package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jamisonwilliams99/Exercise3_CYOA/cyoa"
)

func main() {
	port := flag.Int("port", 8080, "the port to start the CYOA application on")
	filename := flag.String("file", "gopher.json", "the JSON fiel with the CYOA story")
	flag.Parse()
	fmt.Printf("Using the story in %s.\n", *filename)

	f, err := os.Open(*filename)
	if err != nil {
		log.Fatal(err)
	}

	story, err := cyoa.JsonStory(f)
	if err != nil {
		log.Fatal(err)
	}

	h := cyoa.NewHandler(story)
	fmt.Printf("Starting the server at: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))

}
