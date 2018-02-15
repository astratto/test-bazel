package main

import (
	stamping "github.com/astratto/test-bazel"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.WithFields(log.Fields{
		"test": "bazel",
	}).Info("Exposed value is ", stamping.ExposedValue)
}
