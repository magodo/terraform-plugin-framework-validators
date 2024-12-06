// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package float32validator_test

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	"github.com/hashicorp/terraform-plugin-framework-validators/float32validator"
)

func ExampleNoneOf() {
	// Used within a Schema method of a DataSource, Provider, or Resource
	_ = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"example_attr": schema.Float32Attribute{
				Required: true,
				Validators: []validator.Float32{
					// Validate floating point value must not be 1.2, 2.4, or 4.8
					float32validator.NoneOf([]float32{1.2, 2.4, 4.8}...),
				},
			},
		},
	}
}

func ExampleNoneOf_function() {
	_ = function.Definition{
		Parameters: []function.Parameter{
			function.Float32Parameter{
				Name: "example_param",
				Validators: []function.Float32ParameterValidator{
					// Validate floating point value must not be 1.2, 2.4, or 4.8
					float32validator.NoneOf([]float32{1.2, 2.4, 4.8}...),
				},
			},
		},
	}
}