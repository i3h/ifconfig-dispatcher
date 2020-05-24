package main

import (
	"fmt"
	"os"
	"reflect"
	"strconv"

	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

var (
	cyan   = color.New(color.FgCyan).SprintfFunc()
	blue   = color.New(color.FgBlue).SprintfFunc()
	green  = color.New(color.FgGreen).SprintfFunc()
	yellow = color.New(color.FgYellow).SprintfFunc()
	red    = color.New(color.FgRed).SprintfFunc()
)

func init_log() {
	formatter := &logrus.TextFormatter{
		DisableTimestamp: true,
	}
	log.SetFormatter(formatter)
	log.Out = os.Stdout
}

func getMaxNameLength(names []string) int {
	var length int
	for _, val := range names {
		if len(val) > length {
			length = len(val)
		}
	}
	return length
}

func GetPrettyJson(info interface{}) string {
	var pjs string
	pjs += "{\n"

	v := reflect.ValueOf(info)
	names := make([]string, v.NumField())
	values := make([]string, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		names[i] = v.Type().Field(i).Name
		t := v.Field(i).Type().Name()
		if t == "string" {
			values[i] = v.Field(i).Interface().(string)
		} else if t == "bool" {
			values[i] = strconv.FormatBool(v.Field(i).Interface().(bool))
		} else if t == "float64" {
			values[i] = fmt.Sprintf("%f", v.Field(i).Interface().(float64))
		} else if t == "uint" {
			values[i] = fmt.Sprint(v.Field(i).Interface().(uint))
		}
	}
	l := getMaxNameLength(names) + 2
	for i := 0; i < v.NumField(); i++ {
		if values[i] != "" {
			t := v.Field(i).Type().Name()
			if t == "string" {
				pjs += fmt.Sprintf("  %s:  %s,\n", cyan("%-*s", l, "\""+names[i]+"\""), green("\""+values[i]+"\""))
			} else if t == "bool" {
				pjs += fmt.Sprintf("  %s:  %s,\n", cyan("%-*s", l, "\""+names[i]+"\""), blue(values[i]))
			} else if t == "float64" {
				pjs += fmt.Sprintf("  %s:  %s,\n", cyan("%-*s", l, "\""+names[i]+"\""), red(values[i]))
			} else if t == "uint" {
				pjs += fmt.Sprintf("  %s:  %s,\n", cyan("%-*s", l, "\""+names[i]+"\""), red(values[i]))
			}
		}
	}
	pjs = pjs[:len(pjs)-2]
	pjs += "\n}\n"
	return pjs
}
