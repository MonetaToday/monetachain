package types

import (
	fmt "fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	KeyCanAnyoneAddCoins     = []byte("CanAnyoneAddCoins")
	DefaultCanAnyoneAddCoins = false

	KeyAccessList     = []byte("AccessList")
	DefaultAccessList = []string{}
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	defaultCanAnyoneAddCoins bool,
	defaultAccessList []string,
) Params {
	return Params{
		defaultCanAnyoneAddCoins,
		defaultAccessList,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		false,
		[]string{},
	)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyCanAnyoneAddCoins, &p.CanAnyoneAddCoins, validateCanAnyoneAddCoins),
		paramtypes.NewParamSetPair(KeyAccessList, &p.AccessList, validateAccessList),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateCanAnyoneAddCoins(p.CanAnyoneAddCoins); err != nil {
		return err
	}

	if err := validateAccessList(p.AccessList); err != nil {
		return err
	}

	return nil
}

func validateCanAnyoneAddCoins(i interface{}) error {
	_, ok := i.(bool)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}

func validateAccessList(i interface{}) error {
	_, ok := i.([]string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}
