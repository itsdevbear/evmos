package solidity

import (
	"bytes"
	_ "embed" // embed compiled smart contract
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	evmtypes "github.com/tharsis/ethermint/x/evm/types"
	erc20types "github.com/tharsis/evmos/x/erc20/types"
)

type Contract struct {
	JSON     *[]byte
	Bytecode *evmtypes.CompiledContract
}

var AllContracts []Contract

var (
	// TODO: Fix Forge Issue with Payable
	// //go:embed out/WBERA-PAYABLE.sol/WBERA.json
	// WBERAJson     []byte // nolint: golint
	// WBERAContract evmtypes.CompiledContract

	//go:embed out/BGT.sol/BGT.json
	BGTJson     []byte // nolint: golint
	BGTContract evmtypes.CompiledContract

	//go:embed out/HONEY.sol/HONEY.json
	HONEYJson     []byte // nolint: golint
	HONEYContract evmtypes.CompiledContract

	//go:embed out/CosmosStaking.sol/CosmosStaking.json
	CosmosStakingJson     []byte // nolint: golint
	CosmosStakingContract evmtypes.CompiledContract

	//go:embed out/CosmosStaking.sol/CosmosStaking.json
	CosmosGovernanceJson     []byte // nolint: golint
	CosmosGovernanceContract evmtypes.CompiledContract

	//go:embed out/CosmosSynapse.sol/CosmosSynapse.json
	CosmosSynapseJson     []byte // nolint: golint
	CosmosSynapseContract evmtypes.CompiledContract

	//go:embed out/CosmosNativeERC20.sol/CosmosNativeERC20.json
	CosmosNativeERC20JSON     []byte // nolint: golint
	CosmosNativeERC20Contract evmtypes.CompiledContract

	ERC20ModuleAddress common.Address
)

func init() {
	ERC20ModuleAddress = erc20types.ModuleAddress

	AllContracts = []Contract{
		// {
		// 	JSON:     &WBERAJson,
		// 	Bytecode: &WBERAContract,
		// },
		{
			JSON:     &BGTJson,
			Bytecode: &BGTContract,
		},
		{
			JSON:     &HONEYJson,
			Bytecode: &HONEYContract,
		},
		{
			JSON:     &CosmosStakingJson,
			Bytecode: &CosmosStakingContract,
		},
		{
			JSON:     &CosmosGovernanceJson,
			Bytecode: &CosmosGovernanceContract,
		},
		{
			JSON:     &CosmosSynapseJson,
			Bytecode: &CosmosSynapseContract,
		},
		{
			JSON:     &CosmosNativeERC20JSON,
			Bytecode: &CosmosNativeERC20Contract,
		},
	}

	for _, c := range AllContracts {
		var objmap map[string]interface{}
		if err := json.Unmarshal(*c.JSON, &objmap); err != nil {
			log.Fatal(err)
		}

		// Hacky fix to remove the "0x" from the foundry bytecode output
		x, err := json.Marshal(objmap["abi"])
		if err != nil {
			fmt.Println(err)
		}
		objmap["abi"] = strings.ReplaceAll(string(x), "\"", "\"")
		objmap["bin"] = string(bytes.TrimPrefix([]byte(fmt.Sprintf("%v", objmap["bin"])), []byte("0x")))
		temp, err := json.Marshal(objmap)
		if err != nil {
			fmt.Println(err)
		}
		*c.JSON = []byte(temp)
		if err := json.Unmarshal(*c.JSON, c.Bytecode); err != nil {
			fmt.Println(err)
			panic(c.JSON)
		}
	}
}
