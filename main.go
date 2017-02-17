package main

import (
	"github.com/Sirupsen/logrus"
	"MicroFilm/route"
)

func init() {

	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func main() {

	router := route.Init()
	router.Start(":8888")
}
