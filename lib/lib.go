package lib

import "strings"

func In(item interface{}, collection interface{}) bool {
	switch collection.(type){
	case []string:
		if item, ok := item.(string); ok {
			for _, e := range (collection.([]string)) {
				if strings.Compare(item, e) == 0 {
					return true
				}
			}
		}
	case []int:
		if item, ok := item.(int); ok {
			for _, e := range (collection.([]int)) {
				if item == e {
					return true
				}
			}
		}
	case map[string]interface{}:
		if item, ok := item.(string); ok {
			var m  map[string]interface{}
			m = collection.(map[string]interface{})
			for k,_:= range m {
				if item == k {
					return true
				}
			}
		}
	}
	return false
}

