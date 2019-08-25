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
# go test -bench=. -benchmem
goos: darwin
goarch: amd64
pkg: github.com/tfbrother/leetcode/451
 Benchmark 名字 - CPU核数				   循环次数          平均每次执行时间		内存占用字节数/每次	内存分配数/每次
BenchmarkFrequencySort11e4-4               10000            144521 ns/op           75328 B/op        904 allocs/op
BenchmarkFrequencySort1_11e4-4             10000            135440 ns/op           49226 B/op        897 allocs/op
BenchmarkFrequencySort1_21e4-4             10000            117860 ns/op           52476 B/op        814 allocs/op
BenchmarkFrequencySort1_31e4-4             10000            101549 ns/op           26840 B/op        142 allocs/op
BenchmarkFrequencySort1_41e4-4             20000             88255 ns/op           16645 B/op        100 allocs/op
BenchmarkFrequencySort_en1e4-4             20000             98461 ns/op           60668 B/op        763 allocs/op
PASS
ok      github.com/tfbrother/leetcode/451       24.357s
*/
