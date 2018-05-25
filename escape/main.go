package main

type S struct {
	s1 int
}

func (s *S) M1(i int) {
	s.s1 = i
}

type I interface {
	M1(int)
}

func g() {
	var s1 S // this escapes
	var s2 S // this does not

	f1(&s1)
	f2(&s2)
}

func f1(s I) {
	s.M1(42)
}

func f2(s *S) {
	s.M1(42)
}

// go build -gcflags='-m'
func main() {
	g()
}
