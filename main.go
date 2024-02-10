package ardor

import (
	"fmt"
	"net/url"
)

var (
	appName     = "ardor-pkg"
	appVersion  = "0.0.1"
	httpTimeout = "5"
)

func (a *Ardor) New(prefix string) *Ardor {
	a.Endpoint = prefix
	return a
}

func (a *Ardor) buildURL(path string) string {
	s, err := url.JoinPath(a.Endpoint, path)
	if err != nil {
		return path
	}
	return fmt.Sprint(s)
}
