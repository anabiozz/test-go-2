package binary

import "strconv"

// Binary ..
type Binary uint64

func (i Binary) String() string {
	return strconv.FormatUint(i.Get(), 10)
}

// Get ..
func (i Binary) Get() uint64 {
	return uint64(i)
}
