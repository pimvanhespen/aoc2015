package aoc

//type PermuteCheckFunc[T any] func(v []T) int
//
//func Permute[T any](fn PermuteCheckFunc[T], items []T) [][]T {
//
//	var results [][]T
//
//	for i, t := range items {
//		n := fn([]T{t})
//		if n == 0 {
//			results = append(results, []T{t})
//		} else if n < 0 {
//			for _, p := range Permute(fn, items[i+1:]) {
//				results = append(results, append([]T{t}, p...))
//			}
//		}
//	}
//
//	return results
//}
