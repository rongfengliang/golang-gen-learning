package main

import (
	"log"
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
		log.Printf("%s", item)
	}
	myUserSets := NewMyUserSet([]MyUser{
		{
			Name: "dalong",
			Age:  333,
		},
		{
			Name: "dalong2",
			Age:  44,
		},
	}...)

	lists := myUserSets.ToSlice()
	for _, item := range lists {
		log.Printf("%s", item)
	}
}
