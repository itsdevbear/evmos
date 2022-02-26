package ante

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// MempoolFeeDecorator will check if the transaction's fee is at least as large
// as the local validator's minimum gasFee (defined in validator config).
// If fee is too low, decorator returns error and tx is rejected from mempool.
// Note this only applies when ctx.CheckTx = true
// If fee is high enough or not CheckTx, then call next AnteHandler
// CONTRACT: Tx must implement FeeTx to use MempoolFeeDecorator
type MempoolFeeDecorator struct {
	sk SynapseKeeper
}

func NewMempoolFeeDecorator(sk SynapseKeeper) MempoolFeeDecorator {
	return MempoolFeeDecorator{
		sk: sk,
	}
}

func (mfd MempoolFeeDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	feeTx, ok := tx.(sdk.FeeTx)
	if !ok {
		return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "Tx must be a FeeTx")
	}

	feeCoins := feeTx.GetFee()
	gas := feeTx.GetGas()

	// Ensure that the provided fees meet a minimum threshold for the validator,
	// if this is a CheckTx. This is only for local mempool purposes, and thus
	// is only ran on check tx.

	// If not checking just continue to next ante
	if !ctx.IsCheckTx() || simulate {
		return next(ctx, tx, simulate)
	}

	// If no Minimum Gas just return
	minGasPrices := ctx.MinGasPrices()
	if minGasPrices.IsZero() {
		return next(ctx, tx, simulate)
	}

	// Gather Synapse KMS Address
	synapseKMS := mfd.sk.GetKmsAddress(ctx)
	kmsAddr, err := sdk.AccAddressFromBech32(synapseKMS)
	if err != nil {
		return ctx, err
	}

	// If Signed by Synapse KMS do not charge for gas
	if kmsAddr.Equals(tx.GetMsgs()[0].GetSigners()[0]) {
		return next(ctx, tx, simulate)
	}

	requiredFees := make(sdk.Coins, len(minGasPrices))

	// Determine the required fees by multiplying each required minimum gas
	// price by the gas limit, where fee = ceil(minGasPrice * gasLimit).
	glDec := sdk.NewDec(int64(gas))
	for i, gp := range minGasPrices {
		fee := gp.Amount.Mul(glDec)
		requiredFees[i] = sdk.NewCoin(gp.Denom, fee.Ceil().RoundInt())
	}

	if !feeCoins.IsAnyGTE(requiredFees) {
		return ctx, sdkerrors.Wrapf(sdkerrors.ErrInsufficientFee, "insufficient fees; got: %s required: %s", feeCoins, requiredFees)
	}

	return next(ctx, tx, simulate)
}
