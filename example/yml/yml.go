package main

import (
	"github.com/geruz/info/yml"
	"github.com/geruz/info"
)

func main() {
	yml.LoadInfo("./config.yml")
	println(info.App.String())
	println(info.AppInstanceId)

}
