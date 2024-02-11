package ardor

import (
	"fmt"
)

var (
	appName     = "ardor-pkg"
	appVersion  = "0.0.4"
	httpTimeout = "5"
)

func (a *Ardor) Init(node string) {
	a.Endpoint = node
}

func (a *Ardor) buildURL(path string) string {
	return fmt.Sprintf("%s%s", a.Endpoint, path)
}
