package nameservice

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/PiotrTrzpil/nameservice/x/nameservice/keeper"
	"github.com/PiotrTrzpil/nameservice/x/nameservice/types"
)

func handleMsgCreateWhois(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateWhois) (*sdk.Result, error) {
	//TODO
	//var whois = types.Whois{
	//	Creator: msg.Creator,
	//	ID:      msg.ID,
	//  	Value: msg.Value,
	//  	Price: msg.Price,
	//}
	//k.CreateWhois(ctx, whois)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
