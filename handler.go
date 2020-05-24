package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

func IsBrowser(c *gin.Context) bool {
	r, _ := regexp.Compile("Gecko|WebKit|Presto|Trident|EdgeHTML|Blink")
	return r.MatchString(c.Request.UserAgent())
}

func Dispatcher(API string) gin.HandlerFunc {
	return func(c *gin.Context) {
		isBrowser := IsBrowser(c)
		path := c.Request.URL.String()
		r := strings.Split(path, "/")[1]

		if r == "gql" {
			PassToApi(c)
		} else if r == "json" {
			ReplyPrettyJson(c, API)
		} else if !isBrowser {
			PassToApi(c)
		} else {
			// Serve static resources for browser
			file := "./static/" + c.Request.URL.Path
			http.ServeFile(c.Writer, c.Request, file)
		}
	}
}

func PassToApi(c *gin.Context) {
	url, err := url.Parse(API)
	if err != nil {
		fmt.Println(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(url)
	proxy.ServeHTTP(c.Writer, c.Request)
}

func ReplyPrettyJson(c *gin.Context, API string) {
	var url string

	path := c.Request.URL.String()
	path = strings.TrimPrefix(path, "/json")
	path = strings.TrimPrefix(path, "/")

	if path == "" {
		url = API + "/json/" + c.ClientIP()
	} else {
		url = API + "/json/" + path
	}

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	js := SimpleData{}
	err = json.Unmarshal(body, &js)
	if err != nil {
		fmt.Println(err)
	}
	pjs := GetPrettyJson(js)

	// Write response
	c.String(200, pjs)
}
