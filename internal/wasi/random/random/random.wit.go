// Code generated by wit-bindgen-go. DO NOT EDIT.

//go:build !wasip1

// Package random represents the imported interface "wasi:random/random@0.2.0".
//
// WASI Random is a random data API.
//
// It is intended to be portable at least between Unix-family platforms and
// Windows.
package random

import (
	"github.com/ydnar/wasm-tools-go/cm"
)

// GetRandomBytes represents the imported function "get-random-bytes".
//
// Return `len` cryptographically-secure random or pseudo-random bytes.
//
// This function must produce data at least as cryptographically secure and
// fast as an adequately seeded cryptographically-secure pseudo-random
// number generator (CSPRNG). It must not block, from the perspective of
// the calling program, under any circumstances, including on the first
// request and on requests for numbers of bytes. The returned data must
// always be unpredictable.
//
// This function must always return fresh data. Deterministic environments
// must omit this function, rather than implementing it with deterministic
// data.
//
//	get-random-bytes: func(len: u64) -> list<u8>
//
//go:nosplit
func GetRandomBytes(len_ uint64) (result cm.List[uint8]) {
	len0 := (uint64)(len_)
	wasmimport_GetRandomBytes((uint64)(len0), &result)
	return
}

//go:wasmimport wasi:random/random@0.2.0 get-random-bytes
//go:noescape
func wasmimport_GetRandomBytes(len0 uint64, result *cm.List[uint8])

// GetRandomU64 represents the imported function "get-random-u64".
//
// Return a cryptographically-secure random or pseudo-random `u64` value.
//
// This function returns the same type of data as `get-random-bytes`,
// represented as a `u64`.
//
//	get-random-u64: func() -> u64
//
//go:nosplit
func GetRandomU64() (result uint64) {
	result0 := wasmimport_GetRandomU64()
	result = (uint64)((uint64)(result0))
	return
}

//go:wasmimport wasi:random/random@0.2.0 get-random-u64
//go:noescape
func wasmimport_GetRandomU64() (result0 uint64)
