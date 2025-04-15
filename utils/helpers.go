package utils

import "github.com/ethereum/go-ethereum/common"

func ValidateAddress(addr common.Address) bool {
	var hexStr = addr.Hex()
	return common.IsHexAddress(hexStr)
}
