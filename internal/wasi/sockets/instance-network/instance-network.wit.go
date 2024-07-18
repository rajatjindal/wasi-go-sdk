// Code generated by wit-bindgen-go. DO NOT EDIT.

//go:build !wasip1

// Package instancenetwork represents the imported interface "wasi:sockets/instance-network@0.2.0".
//
// This interface provides a value-export of the default network handle..
package instancenetwork

import (
	"github.com/rajatjindal/wasi/internal/wasi/sockets/network"
	"github.com/ydnar/wasm-tools-go/cm"
)

// InstanceNetwork represents the imported function "instance-network".
//
// Get a handle to the default network.
//
//	instance-network: func() -> network
//
//go:nosplit
func InstanceNetwork() (result network.Network) {
	result0 := wasmimport_InstanceNetwork()
	result = cm.Reinterpret[network.Network]((uint32)(result0))
	return
}

//go:wasmimport wasi:sockets/instance-network@0.2.0 instance-network
//go:noescape
func wasmimport_InstanceNetwork() (result0 uint32)
