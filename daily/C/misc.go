package main

//#include "common.h"
import "C"
import (
	"runtime"
	"runtime/debug"

	"github.com/dailycrypto-me/daily-evm/daily/util/bin"

	"github.com/dailycrypto-me/daily-evm/core"

	"github.com/dailycrypto-me/daily-evm/daily/util/keccak256"
)

//export go_set_gc_percent
func go_set_gc_percent(pct C.int) {
	debug.SetGCPercent(int(pct))
}

//export go_gc
func go_gc() {
	runtime.GC()
}

//export go_gc_async
func go_gc_async() {
	go runtime.GC()
}

//export daily_evm_keccak256_init_pool
func daily_evm_keccak256_init_pool(size C.uint64_t) {
	keccak256.InitPool(uint64(size))
}

//export daily_evm_mainnet_initial_balances
func daily_evm_mainnet_initial_balances() C.daily_evm_Bytes {
	return go_bytes_to_c(bin.BytesView(core.MainnetAllocData))
}

//export daily_evm_traceback
func daily_evm_traceback(cb C.daily_evm_BytesCallback) {
	call_bytes_cb(debug.Stack(), cb)
}

func main() {

}
