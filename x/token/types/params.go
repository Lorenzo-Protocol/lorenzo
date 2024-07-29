package types

// NewParams create Params object for token module.
func NewParams(enableConversion, enableEVMHook bool) Params {
	return Params{
		EnableConversion: enableConversion,
		EnableEVMHook:    enableEVMHook,
	}
}

// DefaultParams returns default Params for token module.
func DefaultParams() Params {
	return Params{
		EnableConversion: true,
		EnableEVMHook:    true,
	}
}
