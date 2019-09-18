package _5_test

import (
	. "github.com/tfbrother/leetcode/15"
	"github.com/tfbrother/leetcode/util"
	"testing"
)

func bench(b *testing.B, size int, algo func(nums []int) [][]int, name string) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		for n := size - 3; n <= size+3; n++ {
			nums := util.GenrateRandomArray(50, -100000, 100000)
			b.StartTimer()
			ret := algo(nums)
			b.StopTimer()
			if len(ret) < 0 {
				b.Errorf("%s did not sort %d ints", name, n)
			}
		}
	}
}

func BenchmarkThreeSum21e2(b *testing.B)   { bench(b, 1e2, ThreeSum2, "ThreeSum2") }
func BenchmarkThreeSum2_11e2(b *testing.B) { bench(b, 1e2, ThreeSum2_1, "ThreeSum2_1") }
func BenchmarkThreeSum2_21e2(b *testing.B) { bench(b, 1e2, ThreeSum2_2, "ThreeSum2_2") }
func BenchmarkThreeSum_en1e2(b *testing.B) { bench(b, 1e2, ThreeSum_en, "ThreeSum_en") }

/*
# go test -bench=. -benchmem
goos: darwin
goarch: amd64
pkg: github.com/tfbrother/leetcode/15
BenchmarkThreeSum21e2-4     	   20000	     67328 ns/op
BenchmarkThreeSum2_11e2-4   	   20000	     66810 ns/op
BenchmarkThreeSum2_21e2-4   	   30000	     62176 ns/op
BenchmarkThreeSum_en1e2-4   	   20000	     65414 ns/op
PASS
ok      github.com/tfbrother/leetcode/15        18.860s
*/
