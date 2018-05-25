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

	// Mask for the type bits. For regular files, none will be set.

	CForWB           = 1 << iota // for wb co-worker
	CIsMulti                     // a reusable coupon or not
	CAllowIcon                   // show a icon in the directory
	CIsMultiOneOrder             // the coupon can be applied by any client, but only once
	CIsMinSumPerItem             // the coupon for one position in bucket or for all bucket
	CInCatalog                   // display the coupon in the directory (including Auto apply)
	CInOffer                     // display the coupon in the gooditem card
	CAutoApply                   // automatic application in the basket

	CIsOneOrderExist // the order with this coupon already exist. excessive field
)

func main() {
	var Flags uint32 = CIsMulti | CInCatalog | CInOffer

	log.Printf("%b", Flags)

	log.Printf("%b", ModeDir)
	log.Printf("%b", ModeAppend)
	log.Printf("%b", ModeExclusive)
	log.Printf("%b", ModeNamedPipe)
	log.Printf("%b", ModeSymlink)

	log.Printf("%t", Flags&ModeDir > 0)
	log.Printf("%t", Flags&ModeAppend > 0)
	log.Printf("%t", Flags&ModeExclusive > 0)
	log.Printf("%t", Flags&ModeNamedPipe > 0)
	log.Printf("%t", Flags&ModeSymlink > 0)
}
