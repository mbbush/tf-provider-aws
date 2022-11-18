package enum

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func Validate[T valueser[T]]() schema.SchemaValidateDiagFunc {
	return validation.ToDiagFunc(validation.StringInSlice(Values[T](), false))
}

func FrameworkValidate[T valueser[T]]() tfsdk.AttributeValidator {
	return stringvalidator.OneOf(Values[T]()...)
}
