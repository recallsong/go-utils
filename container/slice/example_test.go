package slice_test

import (
	"fmt"
	"sort"
	"testing"

	"github.com/recallsong/go-utils/container/slice"
)

func ExampleInterfaces() {
	s := slice.Interfaces{}
	s.Append(1, 2)
	s.Prepend(-4, -3, -2, -1, 0)
	s.Insert(1, 3, 4, 5, 6)
	fmt.Println(s, s.Len())

	s1 := s.Copy()
	removed := s1.Remove(3, 6)
	fmt.Println(s1, s1.Len(), removed)
	// Output:
	// [-4 3 4 5 6 -3 -2 -1 0 1 2] 11
	// [-4 3 4 1 2] 5 6
}

type Persion struct {
	Name string
}

func ExampleSort() {
	ps := []Persion{
		Persion{Name: "333"},
		Persion{Name: "111"},
		Persion{Name: "444"},
		Persion{Name: "222"},
	}
	slice.Sort(ps, func(i, j int) bool { return ps[i].Name < ps[j].Name })
	fmt.Println(ps)

	// Output:
	// [{111} {222} {333} {444}]
}

func BenchmarkSort(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ps := []Persion{
			Persion{Name: "333"},
			Persion{Name: "111"},
			Persion{Name: "444"},
			Persion{Name: "222"},
		}
		b.StartTimer()
		slice.Sort(ps, func(i, j int) bool { return ps[i].Name < ps[j].Name })
		b.StopTimer()
	}
}

func BenchmarkSort2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ps := []Persion{
			Persion{Name: "333"},
			Persion{Name: "111"},
			Persion{Name: "444"},
			Persion{Name: "222"},
		}
		b.StartTimer()
		slice.Sort2(ps,
			func(i, j int) bool { return ps[i].Name < ps[j].Name },
			func(i, j int) { ps[i], ps[j] = ps[j], ps[i] })
		b.StopTimer()
	}
}

func BenchmarkSort_persions(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ps := []Persion{
			Persion{Name: "333"},
			Persion{Name: "111"},
			Persion{Name: "444"},
			Persion{Name: "222"},
		}
		b.StartTimer()
		sort.Sort(Persions(ps))
		b.StopTimer()
	}
}

type Persions []Persion

func (s Persions) Len() int {
	return len(s)
}
func (s Persions) Less(i, j int) bool {
	return s[i].Name < s[j].Name
}

func (s Persions) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
