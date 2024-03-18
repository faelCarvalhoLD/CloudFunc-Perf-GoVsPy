package function

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("sort_random_array", main)
}

func main(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inciando cloud function sort_random_array GOLANG!")

	var numbers []int

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 1000000; i++ {
		numbers = append(numbers, rand.Intn(1000000))
	}

	_ = mergeSort(numbers)
	fmt.Println("Fim Execução cloud function sort_random_array GOLANG!")
}

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	L := mergeSort(arr[:mid])
	R := mergeSort(arr[mid:])

	return merge(L, R)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))

	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(result, right...)
		}
		if len(right) == 0 {
			return append(result, left...)
		}
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	return result
}
