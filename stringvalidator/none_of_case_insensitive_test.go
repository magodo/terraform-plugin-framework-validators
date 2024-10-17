// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package stringvalidator_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
)

func TestNoneOfCaseInsensitiveValidator(t *testing.T) {
	t.Parallel()

	type testCase struct {
		in           types.String
		noneOfValues []string
		expectError  bool
	}

	testCases := map[string]testCase{
		"simple-match": {
			in: types.StringValue("foo"),
			noneOfValues: []string{
				"foo",
				"bar",
				"baz",
			},
			expectError: true,
		},
		"simple-match-case-insensitive": {
			in: types.StringValue("foo"),
			noneOfValues: []string{
				"FOO",
				"bar",
				"baz",
			},
			expectError: true,
		},
		"simple-mismatch": {
			in: types.StringValue("foz"),
			noneOfValues: []string{
				"foo",
				"bar",
				"baz",
			},
		},
		"skip-validation-on-null": {
			in: types.StringNull(),
			noneOfValues: []string{
				"foo",
				"bar",
				"baz",
			},
		},
		"skip-validation-on-unknown": {
			in: types.StringUnknown(),
			noneOfValues: []string{
				"foo",
				"bar",
				"baz",
			},
		},
	}

	for name, test := range testCases {
		name, test := name, test

		t.Run(fmt.Sprintf("ValidateString - %s", name), func(t *testing.T) {
			t.Parallel()
			req := validator.StringRequest{
				ConfigValue: test.in,
			}
			res := validator.StringResponse{}
			stringvalidator.NoneOfCaseInsensitive(test.noneOfValues...).ValidateString(context.TODO(), req, &res)

			if !res.Diagnostics.HasError() && test.expectError {
				t.Fatal("expected error, got no error")
			}

			if res.Diagnostics.HasError() && !test.expectError {
				t.Fatalf("got unexpected error: %s", res.Diagnostics)
			}
		})

		t.Run(fmt.Sprintf("ValidateParameterString - %s", name), func(t *testing.T) {
			t.Parallel()
			req := function.StringParameterValidatorRequest{
				Value: test.in,
			}
			res := function.StringParameterValidatorResponse{}
			stringvalidator.NoneOfCaseInsensitive(test.noneOfValues...).ValidateParameterString(context.TODO(), req, &res)

			if res.Error == nil && test.expectError {
				t.Fatal("expected error, got no error")
			}

			if res.Error != nil && !test.expectError {
				t.Fatalf("got unexpected error: %s", res.Error)
			}
		})
	}
}

func TestNoneOfCaseInsensitiveValidator_Description(t *testing.T) {
	t.Parallel()

	type testCase struct {
		in       []string
		expected string
	}

	testCases := map[string]testCase{
		"quoted-once": {
			in:       []string{"foo", "bar", "baz"},
			expected: `value must be none of: ["foo" "bar" "baz"]`,
		},
	}

	for name, test := range testCases {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			v := stringvalidator.NoneOfCaseInsensitive(test.in...)

			got := v.MarkdownDescription(context.Background())

			if diff := cmp.Diff(got, test.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
