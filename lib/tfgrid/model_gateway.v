module tfgrid

// NOTE: is there any need to translate json here?

[params]
pub struct GatewayFQDN {
pub:
	name            string   [required]
	node_id         u32      [required]
	tls_passthrough bool
	backends        []string [required]
	fqdn            string   [required]
}

pub struct GatewayFQDNResult {
pub:
	name            string
	node_id         u32
	tls_passthrough bool
	backends        []string
	fqdn            string
	// computed
	contract_id u32
}

[params]
pub struct GatewayName {
pub mut:
	name            string   [json: 'name'; required]
	node_id         u32      [json: 'node_id']
	tls_passthrough bool     [json: 'tls_passthrough']
	backends        []string [json: 'backends'; required]
}

pub struct GatewayNameResult {
pub:
	name            string   [json: 'name']
	node_id         u32      [json: 'node_id']
	tls_passthrough bool     [json: 'tls_passthrough']
	backends        []string [json: 'backends']
	// computed
	fqdn             string [json: 'fqdn'] // the full domain name
	name_contract_id u32    [json: 'name_contract_id']
	contract_id      u32    [json: 'contract_id']
}
