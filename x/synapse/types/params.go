package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/ethereum/go-ethereum/common"
	"gopkg.in/yaml.v2"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	KeyKmsAddress           = []byte("KmsAddress")
	KeyKmsPendingAddress    = []byte("KmsPendingAddress")
	KeyEvmMiddlewareAddress = []byte("EvmMiddlewareAddress")
	// TODO: Determine the default value
	DefaultKmsAddress           string = "evmos1ssfzktgqclfjtq7aey08pg38w2v6sas3qmkhnu"
	DefaultKmsPendingAddress    string = "evmos1ssfzktgqclfjtq7aey08pg38w2v6sas3qmkhnu"
	DefaultEvmMiddlewareAddress string = "0x0000000000000000000000000000000000000000"
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	KmsAddress string,
	KmsPendingAddress string,
	EvmMiddlewareAddress string,
) Params {
	return Params{
		KmsAddress:           KmsAddress,
		KmsPendingAddress:    KmsPendingAddress,
		EvmMiddlewareAddress: EvmMiddlewareAddress,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultKmsAddress,
		DefaultKmsPendingAddress,
		DefaultEvmMiddlewareAddress,
	)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyKmsAddress, &p.KmsAddress, validateKmsAddress),
		paramtypes.NewParamSetPair(KeyKmsPendingAddress, &p.KmsPendingAddress, validateKmsPendingAddress),
		paramtypes.NewParamSetPair(KeyEvmMiddlewareAddress, &p.EvmMiddlewareAddress, validateHexAddress),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateKmsAddress(p.KmsAddress); err != nil {
		return err
	}

	// PendingAddress is allowed to be empty
	if err := validateKmsPendingAddress(p.KmsPendingAddress); err != nil {
		return err
	}

	if err := validateHexAddress(p.EvmMiddlewareAddress); err != nil {
		return err
	}

	return nil
}

// validateKmsAddress validates the KmsAddress param
func validateKmsAddress(v interface{}) error {
	KmsAddress, ok := v.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	if err := sdk.VerifyAddressFormat([]byte(KmsAddress)); err != nil {
		return fmt.Errorf("invalid address: %s", KmsAddress)
	}

	return nil
}

// validateKmsPendingAddress validates the KmsPendingAddress param
func validateKmsPendingAddress(v interface{}) error {
	pendingKmsAddress, ok := v.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// Pending Address is allowed to be empty
	if pendingKmsAddress == "" {
		return nil
	}

	if err := sdk.VerifyAddressFormat([]byte(pendingKmsAddress)); err != nil {
		return fmt.Errorf("invalid address: %s", pendingKmsAddress)
	}

	return nil
}

// validateHexAddress validates the EvmMiddlewareAddress param
func validateHexAddress(v interface{}) error {
	HexAddress, ok := v.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	if !common.IsHexAddress(HexAddress) {
		return fmt.Errorf("invalid address: %s", HexAddress)
	}
	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}
