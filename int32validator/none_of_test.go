// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package int32validator_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework-validators/int32validator"
)

func TestNoneOfValidator(t *testing.T) {
	t.Parallel()

	type testCase struct {
		in           types.Int32
		noneOfValues []int32
		expectError  bool
	}

	testCases := map[string]testCase{
		"simple-match": {
			in: types.Int32Value(123),
			noneOfValues: []int32{
				123,
				234,
				8910,
				1213,
			},
			expectError: true,
		},
		"simple-mismatch": {
			in: types.Int32Value(123),
			noneOfValues: []int32{
				234,
				8910,
				1213,
			},
		},
		"skip-validation-on-null": {
			in: types.Int32Null(),
			noneOfValues: []int32{
				234,
				8910,
				1213,
			},
		},
		"skip-validation-on-unknown": {
			in: types.Int32Unknown(),
			noneOfValues: []int32{
				234,
				8910,
				1213,
			},
		},
	}

	for name, test := range testCases {
		name, test := name, test

		t.Run(fmt.Sprintf("ValidateInt32 - %s", name), func(t *testing.T) {
			t.Parallel()
			req := validator.Int32Request{
				ConfigValue: test.in,
			}
			res := validator.Int32Response{}
			int32validator.NoneOf(test.noneOfValues...).ValidateInt32(context.TODO(), req, &res)

			if !res.Diagnostics.HasError() && test.expectError {
				t.Fatal("expected error, got no error")
			}

			if res.Diagnostics.HasError() && !test.expectError {
				t.Fatalf("got unexpected error: %s", res.Diagnostics)
			}
		})

		t.Run(fmt.Sprintf("ValidateParameterInt32 - %s", name), func(t *testing.T) {
			t.Parallel()
			req := function.Int32ParameterValidatorRequest{
				Value: test.in,
			}
			res := function.Int32ParameterValidatorResponse{}
			int32validator.NoneOf(test.noneOfValues...).ValidateParameterInt32(context.TODO(), req, &res)

			if res.Error == nil && test.expectError {
				t.Fatal("expected error, got no error")
			}

			if res.Error != nil && !test.expectError {
				t.Fatalf("got unexpected error: %s", res.Error)
			}
		})
	}
}