package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/skeptycal/resume"
)

func main() {
	r, err := resume.New("")
	if err != nil {
		log.Fatal(err)
	}
}
