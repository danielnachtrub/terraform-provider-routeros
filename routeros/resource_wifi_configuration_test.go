package routeros

import (
	"testing"
)

// "manager" is an enum on /interface/wifi/configuration. Clearing it must be
// serialized as an unset ("!manager") - RouterOS rejects manager="" - so the
// resource has to register it as a set/unset field (same mechanism the firewall
// resources use for their address-list fields).
func TestResourceWifiConfiguration_managerIsSetUnsetField(t *testing.T) {
	res := ResourceWifiConfiguration()

	meta, ok := res.Schema[MetaSetUnsetFields]
	if !ok {
		t.Fatalf("resource must declare %s", MetaSetUnsetFields)
	}

	fields := loadSkipFields(meta.Default.(string))
	if _, ok := fields["manager"]; !ok {
		t.Fatalf("manager must be a set/unset field, got %#v", fields)
	}
}
