package main

import "log"

const (
	// The single letters are the abbreviations
	// used by the String method's formatting.
	ModeDir       = 1 << iota // d: is a directory
	ModeAppend                // a: append-only
	ModeExclusive             // l: exclusive use
	ModeNamedPipe             // T: temporary file; Plan 9 only
	ModeSymlink               // L: symbolic link
)

func main() {
	var Flags uint32 = ModeDir | ModeDir | ModeSymlink

	log.Printf("%d", Flags)

	log.Printf("%b", ModeDir)
	log.Printf("%b", ModeAppend)
	log.Printf("%b", ModeExclusive)
	log.Printf("%b", ModeNamedPipe)
	log.Printf("%b", ModeSymlink)

	log.Printf("%t", Flags&ModeDir > 0)
	log.Printf("%t", Flags&ModeDir > 0)
	log.Printf("%t", Flags&ModeExclusive > 0)
	log.Printf("%t", Flags&ModeNamedPipe > 0)
	log.Printf("%t", Flags&ModeSymlink > 0)
}
