package main

import "github.com/geruz/info"

func main() {
	info.RegApp("simple", "1.0.0")
	println(info.App.String())
	println(info.AppInstanceId)
}
