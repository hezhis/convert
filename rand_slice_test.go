/**
 * @Author: ChenJunJi
 * @Date: 2025/04/11
 * @Desc:
**/

package convert

import (
	"log"
	"testing"
)

func TestRandSlice(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	original := make([]int, len(s))
	copy(original, s)

	for i := 0; i < 20; i++ {
		RandSlice(s)

		log.Println(s)
	}
}
