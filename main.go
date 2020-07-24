package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

var (
	Port   string
	API    string
	Static string
)

type SimpleData struct {
	Continent string  `json:"Continent"`
	Country   string  `json:"Country"`
	City      string  `json:"City"`
	Latitude  float64 `json:"Latitude"`
	Longitude float64 `json:"Longitude"`
	TimeZone  string  `json:"TimeZone"`
	IsEU      bool    `json:"IsEU"`
	ASN       uint    `json:"ASN"`
	ORG       string  `json:"ORG"`
}

func init() {
	init_log()
}

func main() {
	// Set running port
	Port = os.Getenv("IFCONFIGIS_DISPATCHER_PORT")
	if Port == "" {
		Port = "5080"
	}
	// Set api endpoint
	API = os.Getenv("IFCONFIGIS_API")
	if API == "" {
		API = "http://localhost:5000"
	}
	// Set static path
	Static = os.Getenv("IFCONFIGIS_STATIC")
	if Static == "" {
		Static = "./static"
	}

	r := gin.Default()

	r.Use(Dispatcher(API))

	r.Run(":" + Port)
}
