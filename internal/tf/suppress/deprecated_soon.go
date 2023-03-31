package suppress

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// CaseDifference will be deprecated and removed in a future release.
// Rather than making the field case-insensitive (which will cause issues down the line)
// this issue can be fixed by normalizing the value being returned from the Azure API
// for example, either by using the `Parse{IDType}Insensitively` function, or by re-casing the value of the constant.
func CaseDifference(_, old, new string, _ *schema.ResourceData) bool {
	// fields should be case-sensitive, normalize the Azure Resource ID in the Read if required
	return strings.EqualFold(old, new)
}

// CaseDifferenceV2Only only suppress case difference for v2.0.
func CaseDifferenceV2Only(_, old, new string, _ *schema.ResourceData) bool {
	// FORK: Force attributes to maintain case-insensitivity to avoid breaking changes in Pulumi programs.
	return strings.EqualFold(old, new)
}
