package objectvalidator_test

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/objectvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func ExampleAnyWithAllWarnings() {
	// Used within a Schema method of a DataSource, Provider, or Resource
	_ = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"example_attr": schema.ObjectAttribute{
				Required: true,
				Validators: []validator.Object{
					objectvalidator.AnyWithAllWarnings( /* ... */ ),
				},
			},
		},
	}
}
