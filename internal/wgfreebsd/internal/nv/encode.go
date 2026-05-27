//go:build freebsd
// +build freebsd

package nv

/*
#cgo LDFLAGS: -lnv
#include <stdlib.h>
#include <sys/nv.h>

// For sizeof(*struct nvlist)
typedef struct nvlist *nvlist_ptr;
*/
import "C"

// Marshal encodes a Go map to a FreeBSD name-value list (nv(9))
func Marshal(m List) (*byte, int, error) { _ = "STUB: not implemented"; return nil, 0, nil }

// For debugging
// C.nvlist_dump(nvl, C.int(os.Stdout.Fd()))

func marshal(m List) (nvl *C.struct_nvlist, err error) { _ = "STUB: not implemented"; return nil, nil }
