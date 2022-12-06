package main

import "github.com/samber/lo"

func IsDistinct[T comparable](a []T) bool {
	return len(lo.Uniq[T](a)) == len(a)
}

func FirstDistinct[T comparable](width int, a []T) int {
	for i := width; i < len(a); i++ {
		if IsDistinct(a[i-width : i]) {
			return i
		}
	}
	return -1
}
