package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/skeptycal/resume"
)

func main() {
	r, err := resume.New("")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(r)
}
