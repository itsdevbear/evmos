package ante

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type SynapseKeeper interface {
	GetKmsAddress(ctx sdk.Context) (res string)
}
