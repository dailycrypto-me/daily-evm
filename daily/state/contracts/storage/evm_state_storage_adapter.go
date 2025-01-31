package contract_storage

import (
	"math/big"

	"github.com/dailycrypto-me/daily-evm/core/vm"
	"github.com/dailycrypto-me/daily-evm/daily/state/state_evm"

	"github.com/dailycrypto-me/daily-evm/common"
)

type EVMStateStorage struct{ *state_evm.EVMState }

func (self EVMStateStorage) SubBalance(address *common.Address, b *big.Int) bool {
	if acc := self.GetAccountConcrete(address); vm.BalanceGTE(acc, b) {
		acc.SubBalance(b)
		return true
	}
	return false
}

func (self EVMStateStorage) AddBalance(address *common.Address, b *big.Int) {
	self.GetAccountConcrete(address).AddBalance(b)
}

func (self EVMStateStorage) Put(address *common.Address, k *common.Hash, v []byte) {
	self.GetAccountConcrete(address).SetStateRawIrreversibly(k, v)
}

func (self EVMStateStorage) GetAccountStorage(address *common.Address, k *common.Hash, cb func([]byte)) {
	self.GetAccountStorageFromDB(address, k, cb)
}

func (self EVMStateStorage) IncrementNonce(address *common.Address) {
	self.GetAccountConcrete(address).IncrementNonce()
}

func (self EVMStateStorage) GetNonce(address *common.Address) *big.Int {
	return self.GetAccountConcrete(address).GetNonce()
}

