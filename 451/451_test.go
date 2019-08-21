package _51_test

import (
	. "github.com/tfbrother/leetcode/451"
	"testing"
)

func bench(b *testing.B, size int, algo func(s1 string) string, name string) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		for n := size - 3; n <= size+3; n++ {
			s1 := "efefadfdsfertyhnhu877456354653fdfasdfefvevertgrgdgagsafaefwefsfsfjlkjoinihaonubwomenbuxoapqwefnbv,zjjhfndajn.,']]"
			b.StartTimer()
			ret := algo(s1)
			b.StopTimer()
			if len(ret) < 0 {
				b.Errorf("%s did not sort %d ints", name, n)
			}
		}
	}
}

func BenchmarkFrequencySort11e4(b *testing.B)   { bench(b, 1e4, FrequencySort1, "FrequencySort1") }
func BenchmarkFrequencySort1_11e4(b *testing.B) { bench(b, 1e4, FrequencySort1_1, "FrequencySort1_1") }
func BenchmarkFrequencySort1_21e4(b *testing.B) { bench(b, 1e4, FrequencySort1_2, "FrequencySort1_2") }
func BenchmarkFrequencySort1_31e4(b *testing.B) { bench(b, 1e4, FrequencySort1_3, "FrequencySort1_3") }
func BenchmarkFrequencySort1_41e4(b *testing.B) { bench(b, 1e4, FrequencySort1_4, "FrequencySort1_4") }
func BenchmarkFrequencySort_en1e4(b *testing.B) { bench(b, 1e4, FrequencySort_en, "FrequencySort_en") }

/*
BenchmarkFrequencySort11e4-4     	   10000	    146655 ns/op
BenchmarkFrequencySort1_11e4-4   	   10000	    138514 ns/op
BenchmarkFrequencySort1_21e4-4   	   10000	    120669 ns/op
BenchmarkFrequencySort1_31e4-4   	   20000	     98178 ns/op
BenchmarkFrequencySort1_41e4-4   	   20000	     89451 ns/op
BenchmarkFrequencySort_en1e4-4   	   20000	    100934 ns/op
PASS
*/
