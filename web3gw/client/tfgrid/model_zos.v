module tfgrid

/*
// A request that contains zos node calls information
pub struct ZOSNodeRequest {
pub:
	node_id u32                [json: 'node_id'] // node id of the desired node
	data    ZOSNodeRequestData [json: 'data'] // data of the request, could be a Deployment or a contract id
}

pub type ZOSNodeRequestData = Deployment | u64

// Deployment is a ZOS deployment structure
pub struct Deployment {
pub:
	version               u32                  [json: 'version'] // deployment version
	twin_id               u32                  [json: 'twin_id'] // twin id of the user
	contract_id           u64                  [json: 'contract_id'] // contract id of the deployment
	metadata              string               [json: 'metadata'] // metdata of the deployment
	description           string               [json: 'description'] //  description is human readable description of the deployment
	expiration            i64                  [json: 'expiration'] // expiration [deprecated] is not used
	signature_requirement SignatureRequirement [json: 'signature_requirement'] // signature specifications
	workloads             []Workload           [json: 'workloads'] // workloads is a list of workloads associated with this deployment
}

pub struct DeploymentRaw {
	version               u32                  [json: 'version']
	twin_id               u32                  [json: 'twin_id']
	contract_id           u64                  [json: 'contract_id']
	metadata              string               [json: 'metadata']
	description           string               [json: 'description']
	expiration            i64                  [json: 'expiration']
	signature_requirement SignatureRequirement [json: 'signature_requirement']
	workloads             []WorkloadRaw        [json: 'workloads']
}

// SignatureRequirement struct describes the signatures that are needed to be valid
// for the node to accept the deployment
pub struct SignatureRequirement {
pub:
	requests        []SignatureRequest [json: 'requests'] // signature requests
	weight_required u32                [json: 'weight_required'] // basically describes how many signatures are required
	signatures      []Signature        [json: 'signatures'] // list of actual signatures
	signature_style string             [json: 'signature_style']
}

// SignatureRequest describes which twin should sign the deployment and how much this signature weighs
pub struct SignatureRequest {
pub:
	twin_id  u32  [json: 'twin_id'] // twin id of the signer
	required bool [json: 'required'] // true if this twin is required to sign the deployment
	weight   u32  [json: 'weight'] // weight of this signature
}

// Signature contains actual signature information
pub struct Signature {
pub:
	twin_id        u32    [json: 'twin_id'] // twin id of the signer
	signature      string [json: 'signature'] // signature string
	signature_type string [json: 'signature_type'] // signature style
}

// Workload contains workload information
pub struct Workload {
pub:
	version       u32          [json: 'version'] // workload verison
	name          string       [json: 'name'; required] // workload name, must be unique per deployment
	workload_type string       [json: 'type'] // Type of the workload (zmachine, zdb, vm, etc...)
	data          WorkloadData [json: 'data'] // contains the workload type arguments.
	metadata      string       [json: 'metadata'] // metadata is user specific meta attached to deployment
	description   string       [json: 'description'] // human readale description of the workload
	result        Result       [json: 'result'] // result of reservation, set by the node
}

pub struct WorkloadRaw {
	version       u32       [json: 'version']
	name          string    [json: 'name']
	workload_type string    [json: 'type']
	data          string    [json: 'data'; raw]
	metadata      string    [json: 'metadata']
	description   string    [json: 'description']
	result        ResultRaw [json: 'result']
}

pub type WorkloadData = GatewayFQDNProxyWorkload
	| GatewayNameProxyWorkload
	| NetworkWorkload
	| PublicIP
	| ZDBWorkload
	| ZMachine
	| ZMount
	| Zlogs

// Result is the struct filled by the node
// after a reservation object has been processed
pub struct Result {
pub:
	created i64        [json: 'created'] // Time when the result is sent
	state   string     [json: 'state'] // state ate of the deployment (ok,error, ...)
	message string     [json: 'message'] // if State is "error", then this field contains the error, otherwise it's nil
	data    ResultData [json: 'data'] // data is the information generated by the provisioning of the workload, its type depend on the reservation type
}

pub struct ResultRaw {
	created i64    [json: 'created']
	state   string [json: 'state']
	message string [json: 'message']
	data    string [json: 'data'; raw]
}

pub type ResultData = GatewayNameProxyResult
	| PublicIPResult
	| ZDBResult
	| ZDBResultData
	| ZMachineResult
	| string

// SystemVersion contains system version information
pub struct SystemVersion {
pub:
	zos   string [json: 'zos'] // zos version
	zinit string [json: 'zinit'] // zinit version
}

// Contains DMI information of the node
pub struct DMI {
pub:
	tooling  Tooling   [json: 'tooling']
	sections []Section [json: 'sections']
}

pub struct Tooling {
	aggregator string [json: 'aggregator']
	decoder    string [json: 'decoder']
}

pub struct Section {
	handleline  string       [json: 'handleline']
	typestr     string       [json: 'typestr']
	type_       int          [json: 'typenum']
	subsections []Subsection [json: 'subsections']
}

pub struct Subsection {
	title      string                  [json: 'title']
	properties map[string]PropertyData [json: 'properties']
}

pub struct PropertyData {
	value string   [json: 'value']
	items []string [json: 'items']
}

// PublicConfig is the configuration of the interface
// that is connected to the public internet
pub struct PublicConfig {
pub:
	type_  string [json: 'type'] // Type define if we need to use the Vlan field or the MacVlan
	ipv4   string [json: 'ipv4']
	ipv6   string [json: 'ipv6']
	gw4    string [json: 'gw4']
	gw6    string [json: 'gw6']
	domain string [json: 'domain'] // Domain is the node domain name like gent01.devnet.grid.tf or similar
}

// Statistics contains some info about the node like total and used resources
pub struct Statistics {
pub:
	total Capacity [json: 'total'] // total resources of the node
	used  Capacity [json: 'used'] // used resources of the node
}

pub struct Capacity {
	cru   u64 [json: 'cru']
	hru   u64 [json: 'hru']
	sru   u64 [json: 'sru']
	mru   u64 [json: 'mru']
	ipv4u u64 [json: 'ipv4u']
}
*/