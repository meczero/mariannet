package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type MsgCreateMarian struct {
	Creator sdk.AccAddress `json:"creator"`
	Amount  int64          `json:"amount"`
}

func NewMsgCreateMarian(creator sdk.AccAddress, amount int64) MsgCreateMarian {
	return MsgCreateMarian{
		Creator: creator,
		Amount:  amount,
	}
}

func (msg MsgCreateMarian) Route() string {
	return RouterKey
}

func (msg MsgCreateMarian) Type() string {
	return "CreateMarian"
}

func (msg MsgCreateMarian) ValidateBasic() sdk.Error {
	if msg.Creator.Empty() {
		return sdk.ErrInvalidAddress(msg.Creator.String())
	}
	if msg.Amount <= 0 {
		return sdk.ErrInvalidCoins("Amount should be positive")
	}
	return nil
}

func (msg MsgCreateMarian) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgCreateMarian) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Creator}
}

