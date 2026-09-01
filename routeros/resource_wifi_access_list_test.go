package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestResourceWifiAccessListHasSkipFields(t *testing.T) {
	tests := []struct {
		name   string
		schema map[string]*schema.Schema
	}{
		{"wifi", ResourceWifiAccessList().Schema},
		{"capsman", ResourceCapsManAccessList().Schema},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.schema[MetaSkipFields] == nil {
				t.Fatalf("%s access-list schema missing MetaSkipFields; Update would nil-deref", tt.name)
			}
		})
	}
}
