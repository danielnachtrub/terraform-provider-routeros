# routeros_wifi_cap (Resource)
*<span style="color:red">This resource requires a minimum version of RouterOS 7.13.</span>*

## Example Usage
```terraform
resource "routeros_wifi_cap" "settings" {
  enabled              = true
  discovery_interfaces = ["bridge1"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `caps_man_addresses` (List of String) List of Manager IP addresses that CAP will attempt to contact during discovery.
- `caps_man_certificate_common_names` (List of String) List of manager certificate common names that CAP will connect to.
- `caps_man_names` (List of String) An ordered list of CAPs Manager names that the CAP will connect to.
- `certificate` (String) Certificate to use for authentication.
- `discovery_interfaces` (Set of String) List of interfaces over which CAP should attempt to discover CAPs Manager.
- `enabled` (Boolean) Disable or enable the CAP functionality.
- `lock_to_caps_man` (Boolean) Lock CAP to the first CAPsMAN it connects to.
- `slaves_datapath` (String) Name of the bridge interface the CAP will be added to.
- `slaves_static` (Boolean) An option that creates static virtual interfaces.

### Read-Only

- `current_caps_man_address` (String) Currently used CAPsMAN address.
- `current_caps_man_identity` (String) Currently used CAPsMAN identity.
- `id` (String) The ID of this resource.
- `locked_caps_man_common_name` (String) Common name of the CAPsMAN that the CAP is locked to.
- `requested_certificate` (String) Requested certificate.

## Import
Import is supported using the following syntax:
```shell
terraform import routeros_wifi_cap.settings .
```
