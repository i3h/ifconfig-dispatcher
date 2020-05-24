package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

var Port string
var API string

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
	Port = os.Getenv("IFCONFIG_DISPATCHER_PORT")
	if Port == "" {
		Port = "3080"
	}
	// Set api endpoint
	API = os.Getenv("IFCONFIG_API")
	if API == "" {
		API = "http://localhost:3000"
	}

	r := gin.Default()

	r.Use(Dispatcher(API))

	r.Run(":" + Port)
}
