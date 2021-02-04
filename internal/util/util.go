7package util

func Merge(target map[string]float64, source map[string]float64) {
	for key, val := range source {
		target[key] += val
 }
}