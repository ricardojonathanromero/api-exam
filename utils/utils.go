package utils

import "strings"

func StringToMap(val string) map[string]bool {
	reply := make(map[string]bool, 0)

	if strings.Contains(val, ",") {
		spt := strings.Split(val, ",")
		for _, item := range spt {
			reply[item] = true
		}
	} else {
		reply[val] = true
	}

	return reply
}
