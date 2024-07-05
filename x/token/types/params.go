package types

// NewParams create Params object for token module.
func NewParams(enableConvert, enableEVMHook bool) Params {
	return Params{
		EnableConvert: enableConvert,
		EnableEVMHook: enableEVMHook,
	}
}

// DefaultParams returns default Params for token module.
func DefaultParams() Params {
	return Params{
		EnableConvert: true,
		EnableEVMHook: true,
	}
}
