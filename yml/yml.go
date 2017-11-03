package yml

import (
	"github.com/geruz/info"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func LoadInfo(name string) {
	b, err := ioutil.ReadFile(name)
	if err != nil {
		return
	}
	var p struct {
		Project info.Project
	}
	err = yaml.Unmarshal(b, &p)
	if err != nil {
		return
	}
	info.RegApp(p.Project.Name, p.Project.Version)
}
