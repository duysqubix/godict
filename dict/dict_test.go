package dict

import (
	"math/rand"
	"testing"
	"time"
)

func BenchmarkCopyDict(b *testing.B) {
	d := Dict()
	
	for i := 0; i < b.N; i++ {
		d := Dict()
	}
}

func BenchmarkRandInt(b *testing.B) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	for i := 0; i < b.N; i++ {
		r1.Intn(1)
	}
}
