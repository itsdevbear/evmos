package contracts

import (
	_ "embed" // embed compiled smart contract
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
	evmtypes "github.com/tharsis/ethermint/x/evm/types"
	"github.com/tharsis/evmos/x/erc20/types"
)

var (

	//go:embed compiled_contracts/ERC20DirectBalanceManipulation.json
	ERC20DirectBalanceManipulationJSON []byte // nolint: golint

	// ERC20DirectBalanceManipulationContract is the compiled erc20 contract
	ERC20DirectBalanceManipulationContract evmtypes.CompiledContract

	// ERC20DirectBalanceManipulationAddress is the erc20 module address
	ERC20DirectBalanceManipulationAddress common.Address

	//go:embed compiled_contracts/ERC20MaliciousDelayed.json
	ERC20MaliciousDelayedJSON []byte // nolint: golint

	// ERC20MaliciousDelayedContract is the compiled erc20 contract
	ERC20MaliciousDelayedContract evmtypes.CompiledContract

	// ERC20MaliciousDelayedAddress is the erc20 module address
	ERC20MaliciousDelayedAddress common.Address
)

func init() {
	ERC20DirectBalanceManipulationAddress = types.ModuleAddress

	err := json.Unmarshal(ERC20DirectBalanceManipulationJSON, &ERC20DirectBalanceManipulationContract)
	if err != nil {
		panic(err)
	}

	if len(ERC20DirectBalanceManipulationContract.Bin) == 0 {
		panic("load contract failed")
	}

	ERC20MaliciousDelayedAddress = types.ModuleAddress

	err = json.Unmarshal(ERC20MaliciousDelayedJSON, &ERC20MaliciousDelayedContract)
	if err != nil {
		panic(err)
	}

	if len(ERC20MaliciousDelayedContract.Bin) == 0 {
		panic("load contract failed")
	}
}
