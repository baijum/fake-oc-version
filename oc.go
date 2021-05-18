package main

import (
	"fmt"
	"os"
)

func main() {
	withOpenShiftVersion := `clientVersion:
  buildDate: "2021-04-25T08:15:27Z"
  compiler: gc
  gitCommit: 95881afb5df065c250d98cf7f30ee4bb6d281acf
  gitTreeState: clean
  gitVersion: 4.7.0-202104250659.p0-95881af
  goVersion: go1.15.7
  major: ""
  minor: ""
  platform: linux/amd64
openshiftVersion: 4.7.8
releaseClientVersion: 4.7.9
serverVersion:
  buildDate: "2021-04-16T13:08:35Z"
  compiler: gc
  gitCommit: 7d0a2b269a27413f5f125d30c9d726684886c69a
  gitTreeState: clean
  gitVersion: v1.20.0+7d0a2b2
  goVersion: go1.15.7
  major: "1"
  minor: "20"
  platform: linux/amd64`

	withoutOpenShiftVersion := `clientVersion:
  buildDate: "2021-04-25T08:15:27Z"
  compiler: gc
  gitCommit: 95881afb5df065c250d98cf7f30ee4bb6d281acf
  gitTreeState: clean
  gitVersion: 4.7.0-202104250659.p0-95881af
  goVersion: go1.15.7
  major: ""
  minor: ""
  platform: linux/amd64
releaseClientVersion: 4.7.9
serverVersion:
  buildDate: "2021-04-16T13:08:35Z"
  compiler: gc
  gitCommit: 7d0a2b269a27413f5f125d30c9d726684886c69a
  gitTreeState: clean
  gitVersion: v1.20.0+7d0a2b2
  goVersion: go1.15.7
  major: "1"
  minor: "20"
  platform: linux/amd64`

	kind := os.Getenv("OC_OPENSHIFT_VERSION")
	if kind == "true" {
		fmt.Println(withOpenShiftVersion)
	} else {
		fmt.Println(withoutOpenShiftVersion)
	}
}
