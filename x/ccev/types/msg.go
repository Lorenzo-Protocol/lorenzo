package types

import (
	"strings"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

var (
	_ sdk.Msg = (*MsgCreateClient)(nil)
	_ sdk.Msg = (*MsgUploadContract)(nil)
	_ sdk.Msg = (*MsgUpdateHeader)(nil)
	_ sdk.Msg = (*MsgUploadHeaders)(nil)
	_ sdk.Msg = (*MsgUpdateParams)(nil)
)

// ValidateBasic implements sdk.Msg interface
func (m *MsgCreateClient) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Sender); err != nil {
		return err
	}
	return ValidateClient(&m.Client)
}

// GetSigners implements sdk.Msg interface
func (m *MsgCreateClient) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(m.Sender)}
}

// ValidateBasic implements sdk.Msg interface
func (m *MsgUploadContract) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Sender); err != nil {
		return err
	}
	return ValidateContract(&Contract{
		ChainId:   m.ChainId,
		Address:   m.Address,
		EventName: m.EventName,
		Abi:       m.Abi,
	})
}

// GetSigners implements sdk.Msg interface
func (m *MsgUploadContract) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(m.Sender)}
}

// ValidateBasic implements sdk.Msg interface
func (m *MsgUpdateHeader) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Sender); err != nil {
		return err
	}
	return ValidateHeader(&m.Header)
}

// GetSigners implements sdk.Msg interface
func (m *MsgUpdateHeader) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(m.Sender)}
}

// ValidateBasic implements sdk.Msg interface
func (m *MsgUploadHeaders) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Sender); err != nil {
		return err
	}

	if len(m.Headers) == 0 {
		return errorsmod.Wrapf(ErrInvalidHeader, "headers cannot be empty")
	}

	for i := range m.Headers {
		header := m.Headers[i]
		if err := ValidateHeader(&header); err != nil {
			return err
		}
	}
	return nil
}

// GetSigners implements sdk.Msg interface
func (m *MsgUploadHeaders) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(m.Sender)}
}

// ValidateBasic implements sdk.Msg interface
func (m *MsgUpdateParams) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return err
	}
	return m.Params.Validate()
}

// GetSigners implements sdk.Msg interface
func (m *MsgUpdateParams) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(m.Authority)}
}

// ValidateClient validates a Client object. It checks that the Client is not
// nil, that the ChainId is not 0, that the ChainName is not empty, and that the
// InitialBlock is valid.
func ValidateClient(client *Client) error {
	if client == nil {
		return ErrInvalidClient
	}

	if client.ChainId == 0 {
		return errorsmod.Wrapf(ErrInvalidClient, "chain id cannot be 0")
	}

	if client.ChainName == "" {
		return errorsmod.Wrapf(ErrInvalidClient, "chain name cannot be empty")
	}
	return ValidateHeader(&client.InitialBlock)
}

// ValidateHeader validates a TinyHeader object. It checks that the TinyHeader is not
// nil, that the Number is not 0, and that the Hash is a valid hex string or a valid
// byte slice. If any of these conditions are not met, it returns an ErrInvalidHeader
// error with a descriptive message.
func ValidateHeader(header *TinyHeader) error {
	if header == nil {
		return ErrInvalidHeader
	}

	if header.Number == 0 {
		return errorsmod.Wrapf(ErrInvalidHeader, "header number cannot be 0")
	}

	_, err := common.ParseHexOrString(header.Hash)
	if err != nil {
		return errorsmod.Wrapf(ErrInvalidHeader, "header hash is invalid: %s", err)
	}
	return nil
}

// ValidateContract validates a Contract object. It checks that the Contract is not
// nil, that the ChainId is not 0, that the Address is a valid Ethereum hex address,
// that the EventName is not empty, that the Abi is not empty, and that the Abi is
// valid JSON. If any of these conditions are not met, it returns an ErrInvalidContract
// error with a descriptive message.
func ValidateContract(contract *Contract) error {
	if contract == nil {
		return ErrInvalidContract
	}

	if contract.ChainId == 0 {
		return errorsmod.Wrapf(ErrInvalidContract, "chain id cannot be 0")
	}

	if contract.Address == "" {
		return errorsmod.Wrapf(ErrInvalidContract, "address cannot be empty")
	}

	if !common.IsHexAddress(contract.Address) {
		return errorsmod.Wrapf(
			ErrInvalidContract, "address %s is not a valid ethereum hex address", contract.Address)
	}

	if contract.EventName == "" {
		return errorsmod.Wrapf(ErrInvalidContract, "event name cannot be empty")
	}

	if len(contract.Abi) == 0 {
		return errorsmod.Wrapf(ErrInvalidContract, "abi cannot be empty")
	}

	_, err := abi.JSON(strings.NewReader(string(contract.Abi)))
	if err != nil {
		return errorsmod.Wrapf(ErrInvalidContract, "abi is invalid: %s", err)
	}
	return nil
}
