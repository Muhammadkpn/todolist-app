package util

import "math"

func CountPage(count int64, page_total int) (page int) {
	page = int(math.Ceil(float64(count) / float64(page_total)))
	return
}
