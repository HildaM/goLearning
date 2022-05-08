package performance

import "testing"

func TestMain(m *testing.M) {
	m.Run()
}

func BenchmarkBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Builder()
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add()
	}
}

func BenchmarkBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Buffer()
	}
}
