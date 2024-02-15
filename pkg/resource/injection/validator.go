package injection

import (
	sdkValidatorV10 "gitlab.banksinarmas.com/go/sdkv2/validator/integrations/validatorV10"
	sdkValidator "gitlab.banksinarmas.com/go/sdkv2/validator/validator"
)

func NewValidator() (sdkValidator.Validator, error) {
	return sdkValidatorV10.New()
}
