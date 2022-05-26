package suppress

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func CaseDifference(_, old, new string, _ *schema.ResourceData) bool {
	return strings.EqualFold(old, new)
}

// CaseDifferenceV2Only only suppress case difference for v2.0.
func CaseDifferenceV2Only(_, old, new string, _ *schema.ResourceData) bool {
	// FORK: Force attributes to maintain case-insensitivity to avoid breaking changes in Pulumi programs.
	return strings.EqualFold(old, new)
}
