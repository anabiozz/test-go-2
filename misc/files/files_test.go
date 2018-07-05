package files

import "testing"

func BenchmarkOs(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Os()
	}
}

func BenchmarkBufio(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Bufio()
	}
}

func BenchmarkIoutil(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Ioutil()
	}
}

func BenchmarkWrite(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Write()
	}
}

func TestWrite(t *testing.T) {
	Write()
	t.Error("as")
}
