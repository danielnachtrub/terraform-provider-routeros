# routeros_wireguard_peer (Resource)




<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `interface` (String) Name of the interface.
- `public_key` (String) The remote peer's calculated public key.

### Optional

- `___id___` (Number) <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
- `___path___` (String) <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
- `allowed_address` (List of String) List of IP (v4 or v6) addresses with CIDR masks from which incoming traffic for this peer is allowed and to which outgoing traffic for this peer is directed. The catch-all 0.0.0.0/0 may be specified for matching all IPv4 addresses, and ::/0 may be specified for matching all IPv6 addresses.
- `comment` (String)
- `disabled` (Boolean)
- `endpoint_address` (String) An endpoint IP or hostname can be left blank to allow remote connection from any address.
- `endpoint_port` (String) An endpoint port can be left blank to allow remote connection from any port.
- `persistent_keepalive` (Number) A seconds interval, between 1 and 65535 inclusive, of how often to send an authenticated empty packet to the peer for the purpose of keeping a stateful firewall or NAT mapping valid persistently. For example, if the interface very rarely sends traffic, but it might at anytime receive traffic from a peer, and it is behind NAT, the interface might benefit from having a persistent keepalive interval of 25 seconds.
- `preshared_key` (String, Sensitive) A **base64** preshared key. Optional, and may be omitted. This option adds an additional layer of symmetric-key cryptography to be mixed into the already existing public-key cryptography, for post-quantum resistance.

### Read-Only

- `current_endpoint_address` (String) The most recent source IP address of correctly authenticated packets from the peer.
- `current_endpoint_port` (Number) The most recent source IP port of correctly authenticated packets from the peer.
- `id` (String) The ID of this resource.
- `last_handshake` (String) Time in seconds after the last successful handshake.
- `rx` (String) The total amount of bytes received from the peer.
- `tx` (String) The total amount of bytes transmitted to the peer.

