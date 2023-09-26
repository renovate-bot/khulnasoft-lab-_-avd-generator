package main

import (
	"encoding/json"
	"fmt"
	"strings"

	tracker "github.com/khulnasoft-lab/tracker/tracker-ebpf/external"
)

type connectAddrData struct {
	SaFamily string `json:"sa_family"`
	SinPort  string `json:"sin_port"`
	SinAddr  string `json:"sin_addr"`
	SinPort6 string `json:"sin6_port"`
	SinAddr6 string `json:"sin6_addr"`
}

func GetTrackerArgumentByName(event tracker.Event, argName string) (tracker.Argument, error) {
	for _, arg := range event.Args {
		if arg.Name == argName {
			return arg, nil
		}
	}
	return tracker.Argument{}, fmt.Errorf("argument %s not found", argName)
}

func GetAddrStructFromArg(addrArg tracker.Argument, connectData *connectAddrData) error {
	addrStr := strings.Replace(addrArg.Value.(string), "'", "\"", -1)
	err := json.Unmarshal([]byte(addrStr), &connectData)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	return nil
}
