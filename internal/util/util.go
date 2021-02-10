package util

// MergeMap 合并map[string]float64, 将相同键的值相加
func MergeMap(target map[string]float64, source map[string]float64) {
	for key, val := range source {
		target[key] += val
	}
}

// GetStringMapKeys 获取map[string]interface{}的键值slice
func GetStringMapKeys(m map[string]interface{}) []string {
	keys := make([]string, len(m))
	i := 0
	for key := range m {
		keys[i] = key
		i++
	}

	return keys
}

// GetIntMapKeys 获取map[int]interface{}的键值slice
func GetIntMapKeys(m map[int]interface{}) []int {
	keys := make([]int, len(m))
	i := 0
	for key := range m {
		keys[i] = key
		i++
	}
	return keys
}
