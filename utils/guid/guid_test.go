package guid

import (
	"crypto/rand"
	"fmt"
	"testing"
)

func TestCorrect(t *testing.T) {
	var size int64 = 10000000

	var i int64
	for range make([]byte, size) {
		var b = new(GUID)
		rand.Read(b[:])
		b.setVarient()
		if b.Verify() {
			i++
		}
	}
	if i/size != 1 {
		t.Error("can't pass Verfy")
	}

	i = 0
	for range make([]byte, size) {
		var b = new(GUID)
		rand.Read(b[:])
		if b.Verify() {
			i++
		}
	}

	if rate := i / size; rate != 0 {
		fmt.Println("some id pass the Verify setVarient")
	}
}

func BenchmarkCreateCost(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var by = new(GUID)
		rand.Read(by[:])
		by.setVarient()
		by.Verify()
	}
}
