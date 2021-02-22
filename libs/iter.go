package libs

import (
	"log"
	"sort"
	"strings"
)

func reverse(x interface{}) interface{} {
	switch x.(type) {
	case string:
		ret := []rune(x.(string))
		for i := 0; i < len(ret)/2; i++ {
			ret[i], ret[len(ret)-i-1] = ret[len(ret)-i-1], ret[i]
		}
		return string(ret)
	case []string:
		ret := x.([]string)
		for i := 0; i < len(ret)/2; i++ {
			ret[i], ret[len(ret)-i-1] = ret[len(ret)-i-1], ret[i]
		}
		return ret
	case []int:
		ret := x.([]int)
		for i := 0; i < len(ret)/2; i++ {
			ret[i], ret[len(ret)-i-1] = ret[len(ret)-i-1], ret[i]
		}
		return ret
	case []float64:
		ret := x.([]float64)
		for i := 0; i < len(ret)/2; i++ {
			ret[i], ret[len(ret)-i-1] = ret[len(ret)-i-1], ret[i]
		}
		return ret
	default:
		log.Fatalf("Invalid input type for reverse: %T\n", x)
	}
	return nil
}

func lowerBound(x interface{}, key interface{}) int {
	switch x.(type) {
	case []int:
		x, key := x.([]int), key.(int)
		return sort.Search(len(x), func(i int) bool { return x[i] >= key })
	case []float64:
		x, key := x.([]float64), key.(float64)
		return sort.Search(len(x), func(i int) bool { return x[i] >= key })
	default:
		log.Fatalf("Invalid input type for reverse: %T\n", x)
	}
	return -1
}

func upperBound(x interface{}, key interface{}) int {
	switch x.(type) {
	case []int:
		x, key := x.([]int), key.(int)
		return sort.Search(len(x), func(i int) bool { return x[i] > key })
	case []float64:
		x, key := x.([]float64), key.(float64)
		return sort.Search(len(x), func(i int) bool { return x[i] > key })
	default:
		log.Fatalf("Invalid input type for reverse: %T\n", x)
	}
	return -1
}

func count(x interface{}, key interface{}) int {
	ret := 0
	switch x.(type) {
	case string:
		x, key := strings.Split(x.(string), ""), key.(string)
		for _, v := range x {
			if v == key {
				ret++
			}
		}
	case []string:
		x, key := x.([]string), key.(string)
		for _, v := range x {
			if v == key {
				ret++
			}
		}
	case []int:
		x, key := x.([]int), key.(int)
		for _, v := range x {
			if v == key {
				ret++
			}
		}
	case []float64:
		x, key := x.([]float64), key.(float64)
		for _, v := range x {
			if v == key {
				ret++
			}
		}
	default:
		log.Fatalf("Invalid input type for reverse: %T\n", x)
	}
	return ret
}

func find(x interface{}, key interface{}) int {
	ret := -1
	switch x.(type) {
	case string:
		x, key := strings.Split(x.(string), ""), key.(string)
		for i, v := range x {
			if v == key {
				ret = i
				break
			}
		}
		if ret < 0 {
			ret = len(x)
		}
	case []string:
		x, key := x.([]string), key.(string)
		for i, v := range x {
			if v == key {
				ret = i
				break
			}
		}
		if ret < 0 {
			ret = len(x)
		}
	case []int:
		x, key := x.([]int), key.(int)
		for i, v := range x {
			if v == key {
				ret = i
				break
			}
		}
		if ret < 0 {
			ret = len(x)
		}
	case []float64:
		x, key := x.([]float64), key.(float64)
		for i, v := range x {
			if v == key {
				ret = i
				break
			}
		}
		if ret < 0 {
			ret = len(x)
		}
	default:
		log.Fatalf("Invalid input type for reverse: %T\n", x)
	}
	return ret
}
