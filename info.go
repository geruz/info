package info

import (
	"os"

	"errors"
	"strconv"
	"strings"
	"time"
)

var (
	App           Project
	Libs          []Project
	AppInstanceId string
)

func init() {
	RegLib("info", "1.0.0")
}

type Project struct {
	Name    string
	Version string
}

func (p Project) String() string {
	return p.Name + ":" + p.Version
}

func RegApp(name, version string) {
	App = Project{
		Name:    name,
		Version: version,
	}
	AppInstanceId = initAppInstanceId()
	checkVersion()
}

func RegLib(name, version string) {
	Libs = append(Libs, Project{name, version})
}

func checkVersion() {
	for _, arg := range os.Args {
		if arg == "--version" {
			println(App.Version)
			os.Exit(0)
		}
		if arg == "--version-instance" {
			println(AppInstanceId)
			os.Exit(0)
		}
		if arg == "--version-all" {
			println("app " + App.String())
			for _, lib := range Libs {
				println("lib " + lib.String())
			}
			os.Exit(0)
		}
	}
}

type Version struct {
	Major int
	Minor int
	Patch int
}

func (v Version) String() string {
	return strconv.Itoa(v.Major) + "." + strconv.Itoa(v.Minor) + "." + strconv.Itoa(v.Patch)
}

func ParseVersion(str string) (v Version, err error) {
	vs := strings.Split(str, ".")
	if len(vs) != 3 {
		err = errors.New("bad version string: " + str)
		return
	}
	v.Major, err = strconv.Atoi(vs[0])
	if err != nil {
		return
	}
	v.Minor, err = strconv.Atoi(vs[1])
	if err != nil {
		return
	}
	v.Patch, err = strconv.Atoi(vs[2])
	return
}

func initAppInstanceId() string {
	if hostname, ok := os.LookupEnv("INFO_APP_ID"); ok {
		return hostname
	}
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}
	t := time.Now().UnixNano()
	return App.Name + ":" + App.Version + ":" + hostname + ":" + strconv.FormatInt(t, 10)
}
