package main

import (
	"flag"
	"github.com/ImMin5/api-test/gateway"
	"github.com/golang/glog"
)

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := gateway.Run(); err != nil {
		glog.Fatal(err)
	}
}
