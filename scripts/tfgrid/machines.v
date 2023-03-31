module tfgrid

import freeflowuniverse.crystallib.rpcwebsocket { RpcWsClient }


// Deploy machines workload
pub fn machines_deploy(mut client RpcWsClient, params MachinesModel) !MachinesResult {
	return client.send_json_rpc[MachinesModel, MachinesResult]('tfgrid.machines.deploy', params, default_timeout)!
}

// Get machines deployment info using deployment name
pub fn machines_get(mut client RpcWsClient, params string) !MachinesResult {
	return client.send_json_rpc[string, MachinesResult]('tfgrid.machines.get', params, default_timeout)!
}

// Delete a deployed machines using project name
pub fn machines_delete(mut client RpcWsClient, params string) ! {
	_ := client.send_json_rpc[string, string]('tfgrid.machines.delete', params, default_timeout)!
}

// Add new machine to a machines deployment
pub fn machines_add_machine(mut client RpcWsClient, params AddMachine) !MachinesResult {
	return client.send_json_rpc[AddMachine, MachinesResult]('tfgrid.machines.add', params, default_timeout)!
}

// // Delete machine from a machines deployment
pub fn machines_delete_machine(mut client RpcWsClient, params RemoveMachine) !MachinesResult {
	return client.send_json_rpc[RemoveMachine, MachinesResult]('tfgrid.machines.remove', params, default_timeout)!
}
