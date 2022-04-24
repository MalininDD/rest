package utils

import (
	"fmt"
	"github.com/pkg/errors"
	"strings"
)

func IntSliceSum(value []int) (result int) {
	for _, r := range value {
		result += r
	}
	return
}

func BetterFormat(num float64) string {
	s := fmt.Sprintf("%.10f", num)
	return strings.TrimRight(strings.TrimRight(s, "0"), ".")
}

func GetProxyHost(proxy string) (host string, err error) {
	proxySplited := strings.Split(strings.Replace(proxy, "@", ":", 1), ":")
	if len(proxySplited) != 4 {
		return host, errors.New("wrong proxy length")
	}
	return proxySplited[2], nil
}

func GetProxyHostPort(proxy string) (host string) {
	proxySplited := strings.Split(strings.Replace(proxy, "@", ":", 1), ":")
	if len(proxySplited) != 4 {
		return "wrong proxy length"
	}
	return fmt.Sprintf("%s:%s", proxySplited[2], proxySplited[3])
}
