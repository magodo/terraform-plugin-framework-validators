// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package int32validator

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework-validators/helpers/validatordiag"
	"github.com/hashicorp/terraform-plugin-framework-validators/helpers/validatorfuncerr"
)

var _ validator.Int32 = oneOfValidator{}
var _ function.Int32ParameterValidator = oneOfValidator{}

type oneOfValidator struct {
	values []types.Int32
}

func (v oneOfValidator) Description(ctx context.Context) string {
	return v.MarkdownDescription(ctx)
}

func (v oneOfValidator) MarkdownDescription(_ context.Context) string {
	return fmt.Sprintf("value must be one of: %q", v.values)
}

func (v oneOfValidator) ValidateInt32(ctx context.Context, request validator.Int32Request, response *validator.Int32Response) {
	if request.ConfigValue.IsNull() || request.ConfigValue.IsUnknown() {
		return
	}

	value := request.ConfigValue

	for _, otherValue := range v.values {
		if value.Equal(otherValue) {
			return
		}
	}

	response.Diagnostics.Append(validatordiag.InvalidAttributeValueMatchDiagnostic(
		request.Path,
		v.Description(ctx),
		value.String(),
	))
}

func (v oneOfValidator) ValidateParameterInt32(ctx context.Context, request function.Int32ParameterValidatorRequest, response *function.Int32ParameterValidatorResponse) {
	if request.Value.IsNull() || request.Value.IsUnknown() {
		return
	}

	value := request.Value

	for _, otherValue := range v.values {
		if value.Equal(otherValue) {
			return
		}
	}

	response.Error = validatorfuncerr.InvalidParameterValueMatchFuncError(
		request.ArgumentPosition,
		v.Description(ctx),
		value.String(),
	)
}

// OneOf checks that the Int32 held in the attribute or function parameter
// is one of the given `values`.
func OneOf(values ...int32) oneOfValidator {
	frameworkValues := make([]types.Int32, 0, len(values))

	for _, value := range values {
		frameworkValues = append(frameworkValues, types.Int32Value(value))
	}

	return oneOfValidator{
		values: frameworkValues,
	}
}