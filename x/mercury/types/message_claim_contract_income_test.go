package types

import (
	fmt "fmt"
	"mercury/common"
	"mercury/common/cosmos"

	. "gopkg.in/check.v1"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	cKeys "github.com/cosmos/cosmos-sdk/crypto/keyring"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/module"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

type MsgClaimContractIncomeSuite struct{}

var _ = Suite(&MsgClaimContractIncomeSuite{})

func (MsgClaimContractIncomeSuite) TestValidateBasic(c *C) {
	// setup
	interfaceRegistry := codectypes.NewInterfaceRegistry()
	std.RegisterInterfaces(interfaceRegistry)
	module.NewBasicManager().RegisterInterfaces(interfaceRegistry)
	types.RegisterInterfaces(interfaceRegistry)
	cdc := codec.NewProtoCodec(interfaceRegistry)

	pubkey := GetRandomPubKey()
	acct, err := pubkey.GetMyAddress()
	c.Assert(err, IsNil)
	kb := cKeys.NewInMemory(cdc)
	info, _, err := kb.NewMnemonic("whatever", cKeys.English, `m/44'/931'/0'/0/0`, "", hd.Secp256k1)
	c.Assert(err, IsNil)
	pub, err := info.GetPubKey()
	c.Assert(err, IsNil)
	pk, err := common.NewPubKeyFromCrypto(pub)
	c.Assert(err, IsNil)

	// invalid address
	msg := MsgClaimContractIncome{
		Creator: "invalid address",
	}
	err = msg.ValidateBasic()
	c.Check(err, ErrIs, sdkerrors.ErrInvalidAddress)

	msg = MsgClaimContractIncome{
		Creator: acct.String(),
		PubKey:  pubkey,
		Client:  pk,
		Height:  12,
		Nonce:   24,
	}
	err = msg.ValidateBasic()
	c.Check(err, ErrIs, ErrInvalidChain)

	msg.Chain = common.BTCChain
	message := []byte(fmt.Sprintf("%s:%s:%s:%d:%d", msg.PubKey, msg.Chain, msg.Client, msg.Height, msg.Nonce))
	msg.Signature, _, err = kb.Sign("whatever", message)
	c.Assert(err, IsNil)
	err = msg.ValidateBasic()
	c.Assert(err, IsNil)

	// check bad client
	msg.Client = common.PubKey("bogus")
	err = msg.ValidateBasic()
	c.Check(err, ErrIs, sdkerrors.ErrInvalidPubKey)

	// check auth to cancel a specific contract
	msg = MsgClaimContractIncome{
		Creator: GetRandomBech32Addr().String(),
		PubKey:  pubkey,
		Client:  GetRandomPubKey(),
		Chain:   common.BTCChain,
	}
	err = msg.ValidateBasic()
	c.Check(err, ErrIs, ErrProviderBadSigner)
}

func (MsgClaimContractIncomeSuite) TestValidateSignature(c *C) {
	c.Skip("kept here for archival / 'working with pubkeys' purposes")
	// setup
	pubkey := GetRandomPubKey()
	acct, err := pubkey.GetMyAddress()
	c.Assert(err, IsNil)

	interfaceRegistry := codectypes.NewInterfaceRegistry()
	std.RegisterInterfaces(interfaceRegistry)
	module.NewBasicManager().RegisterInterfaces(interfaceRegistry)
	types.RegisterInterfaces(interfaceRegistry)
	cdc := codec.NewProtoCodec(interfaceRegistry)

	msg := MsgClaimContractIncome{
		Creator: acct.String(),
		PubKey:  pubkey,
		Client:  pubkey,
		Chain:   common.BTCChain,
		Height:  100,
		Nonce:   48,
	}
	err = msg.ValidateBasic()
	c.Assert(err, IsNil)

	// setup
	var pub cryptotypes.PubKey
	kb := cKeys.NewInMemory(cdc)
	_, _, err = kb.NewMnemonic("whatever", cKeys.English, `m/44'/931'/0'/0/0`, "", hd.Secp256k1)
	c.Assert(err, IsNil)

	message := []byte(fmt.Sprintf("%s:%s:%s:%d:%d", msg.PubKey, msg.Chain, msg.Client, msg.Height, msg.Nonce))
	msg.Signature, pub, err = kb.Sign("whatever", message)
	c.Assert(err, IsNil)

	c.Check(pub.VerifySignature(message, msg.Signature), Equals, true)

	pk, err := common.NewPubKeyFromCrypto(pub)
	c.Assert(err, IsNil)
	fmt.Println(pk)

	pk2, err := cosmos.GetPubKeyFromBech32(cosmos.Bech32PubKeyTypeAccPub, pk.String())
	c.Assert(err, IsNil)
	c.Check(pk2.Equals(pub), Equals, true)

	acc, err := pk.GetMyAddress()
	c.Assert(err, IsNil)
	account := authtypes.NewBaseAccountWithAddress(acc)
	c.Check(pk2.Equals(account.GetPubKey()), Equals, true)
}