package formatters

import "sort"

func sortedKeys(m map[string]interface{}) (keys []string) {
	keys = []string{}
	for k, _ := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return
}
