package main

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/jessevdk/go-flags"
	"github.com/relops/cqlc/generator"
	"io/ioutil"
	"log"
	"os"
)

var opts generator.Options
var parser = flags.NewParser(&opts, flags.Default)

var VERSION string = "0.9.1"

func init() {
	opts.Version = printVersionAndExit
}

func main() {

	if _, err := parser.Parse(); err != nil {
		os.Exit(1)
	}

	if err := generator.Generate(&opts); err != nil {
		if err == generator.ErrInvalidOptions {
			parser.WriteHelp(os.Stderr)
			os.Exit(1)
		} else {
			log.Fatalln(err)
		}
	}
}

func printVersionAndExit() {

	if VERSION == "" {

		buf, err := ioutil.ReadFile(".goxc.json")

		if err != nil {
			panic(err)
		}

		json, err := simplejson.NewJson(buf)

		if err != nil {
			panic(err)
		}

		VERSION = json.Get("PackageVersion").MustString()
	}

	fmt.Fprintf(os.Stderr, "%s %s\n", "cqlc", VERSION)
	os.Exit(0)
}
