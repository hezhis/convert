package convert

import (
	"log"
	"testing"
)

func TestSplitAsSlice(t *testing.T) {
	log.Println(SplitAsSlice[int64]("1, 2, 3, 4,5,6,7", ","))
}
