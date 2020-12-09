package main

import (
	"fmt"
	"strings"
)

func main() {
	platforms := []Platform{
		{
			Name:    "first",
			Version: "v1",
			Config: map[string]interface{}{
				"name":     "v1",
				"register": true,
			},
		}, {
			Name:    "demosecond",
			Version: "v2",
			Config: map[string]interface{}{
				"name":     "v2",
				"register": false,
			},
		},
	}
	mPlatformSlice := PlatformSlice(platforms)
	newItems := mPlatformSlice.Where(func(p Platform) bool {
		if strings.HasPrefix(p.Name, "demo") {
			return true
		}
		return false
	})
	for _, item := range newItems {
		fmt.Printf("%s", item)
	}
}
