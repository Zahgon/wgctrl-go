//go:build freebsd
// +build freebsd

package nv

// #cgo LDFLAGS: -lnv
// #include <sys/nv.h>
import "C"

// Unmarshal decodes a FreeBSD name-value list (nv(9)) to a Go map
func Unmarshal(d []byte, out List) error { _ = "STUB: not implemented"; return nil }

func unmarshal(nvl *C.struct_nvlist, out List) error {
	_ = "STUB: not implemented"
	// For debugging
	// C.nvlist_dump(nvl, C.int(os.Stdout.Fd()))
	return nil
}
