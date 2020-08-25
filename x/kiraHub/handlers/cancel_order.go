package handlers

import (
	"github.com/KiraCore/sekai/x/kiraHub/keeper"
	"github.com/KiraCore/sekai/x/kiraHub/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// HandlerMsgCancelOrder is a function to handler cancel order message
func HandlerMsgCancelOrder(context sdk.Context, keeper keeper.Keeper, message *types.MsgCancelOrder) (*sdk.Result, error) {
	keeper.CancelOrder(context, message.OrderID)
	return &sdk.Result{}, nil
}