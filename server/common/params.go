package common

import (
	"errors"
	"net/url"
)

// GetRequiredParam gets a query parameter and errors if nothing is found.
func GetRequiredParam(params url.Values, paramName string) (val string, err error) {
	return getParam(params, paramName, true, "")
}

// GetDefaultParam gets a query parameter and returns a specified default if nothing is found.
func GetDefaultParam(params url.Values, paramName string, defaultVal string) (val string) {
	val, _ = getParam(params, paramName, false, defaultVal)
	return
}

func getParam(params url.Values, paramName string, required bool, defaultVal string) (val string, err error) {
	val = params.Get(paramName)
	if val == "" {
		if required {
			err = errors.New("missing required " + paramName + " param")
			return
		}
		val = defaultVal
	}
	return
}
