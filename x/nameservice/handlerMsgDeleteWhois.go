package nameservice

import (
	"github.com/PiotrTrzpil/nameservice/x/nameservice/keeper"
	"github.com/PiotrTrzpil/nameservice/x/nameservice/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Handle a message to delete name
func handleMsgDeleteWhois(ctx sdk.Context, k keeper.Keeper, msg types.MsgDeleteWhois) (*sdk.Result, error) {
	//if !k.WhoisExists(ctx, msg.ID) {
	//	// replace with ErrKeyNotFound for 0.39+
	//	return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, msg.ID)
	//}
	//if !msg.Creator.Equals(k.GetWhoisOwner(ctx, msg.ID)) {
	//	return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
	//}
	//
	//k.DeleteWhois(ctx, msg.ID)
	return &sdk.Result{}, nil
}
