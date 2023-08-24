module main

import freeflowuniverse.crystallib.threefold.rmb
import json
import threefoldtech.zos
import log

fn main() {
	mut logger := log.Logger(&log.Log{
		level: .debug
	})

	mnemonics := '<YOUR MNEMONICS>'
	substrate_url := 'wss://tfchain.dev.grid.tf/ws'
	mut client := rmb.new(nettype: rmb.TFNetType.dev, tfchain_mnemonic: mnemonics)!
	mut deployer := zos.Deployer{
		mnemonics: mnemonics
		substrate_url: substrate_url
		twin_id: 49
		rmb_cl: client
	}

	node_id := u32(14)

	mut network := zos.Znet{
		ip_range: '10.1.0.0/16'
		subnet: '10.1.1.0/24'
		wireguard_private_key: 'GDU+cjKrHNJS9fodzjFDzNFl5su3kJXTZ3ipPgUjOUE='
		wireguard_listen_port: 8080
		peers: [
			zos.Peer{
				subnet: '10.1.2.0/24'
				wireguard_public_key: '4KTvZS2KPWYfMr+GbiUUly0ANVg8jBC7xP9Bl79Z8zM='
				allowed_ips: ['10.1.2.0/24', '100.64.1.2/32']
			},
		]
	}

	mut znet_workload := zos.Workload{
		version: 0
		name: 'network'
		type_: zos.workload_types.network
		data: json.encode_pretty(network)
		description: 'test network2'
	}

	zmachine := zos.Zmachine{
		flist: 'https://hub.grid.tf/tf-official-apps/base:latest.flist'
		network: zos.ZmachineNetwork{
			public_ip: ''
			interfaces: [
				zos.ZNetworkInterface{
					network: 'network'
					ip: '10.1.1.3'
				},
			]
			planetary: true
		}
		compute_capacity: zos.ComputeCapacity{
			cpu: 1
			memory: i64(1024) * 1024 * 1024 * 2
		}
		env: {
			'SSH_KEY': 'ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQCs3qtlU13/hHKLE8KUkyt+yAH7z5IKs6PH63dhkeQBBG+VdxlTg/a+6DEXqc5VVL6etKRpKKKpDVqUFKuWIK1x3sE+Q6qZ/FiPN+cAAQZjMyevkr5nmX/ofZbvGUAQGo7erxypB0Ye6PFZZVlkZUQBs31dcbNXc6CqtwunJIgWOjCMLIl/wkKUAiod7r4O2lPvD7M2bl0Y/oYCA/FnY9+3UdxlBIi146GBeAvm3+Lpik9jQPaimriBJvAeb90SYIcrHtSSe86t2/9NXcjjN8O7Fa/FboindB2wt5vG+j4APOSbvbWgpDpSfIDPeBbqreSdsqhjhyE36xWwr1IqktX+B9ZuGRoIlPWfCHPJSw/AisfFGPeVeZVW3woUdbdm6bdhoRmGDIGAqPu5Iy576iYiZJnuRb+z8yDbtsbU2eMjRCXn1jnV2GjQcwtxViqiAtbFbqX0eQ0ZU8Zsf0IcFnH1W5Tra/yp9598KmipKHBa+AtsdVu2RRNRW6S4T3MO5SU= mario@mario-machine'
		}
	}

	mut zmachine_workload := zos.Workload{
		version: 0
		name: 'vm2'
		type_: zos.workload_types.zmachine
		data: json.encode(zmachine)
		description: 'zmachine test'
	}

	mut deployment := zos.Deployment{
		version: 0
		twin_id: deployer.twin_id
		metadata: 'zm dep'
		description: 'zm kjasdf1nafvbeaf1234t21'
		workloads: [znet_workload, zmachine_workload]
		signature_requirement: zos.SignatureRequirement{
			weight_required: 1
			requests: [
				zos.SignatureRequest{
					twin_id: deployer.twin_id
					weight: 1
				},
			]
		}
	}

	contract_id := deployer.deploy(node_id, mut deployment, '', 0) or {
		logger.error('failed to deploy deployment: ${err}')
		exit(1)
	}
	logger.info('deployment contract id: ${contract_id}')

	dl := deployer.get_deployment(contract_id, node_id) or {
		logger.error('failed to get deployment: ${err}')
		exit(1)
	}
	logger.info('deployment: ${dl}')

	machine_res := get_machine_result(dl)!
	logger.info('zmachine result: ${machine_res}')
}

fn get_machine_result(dl zos.Deployment) !zos.ZmachineResult {
	for _, w in dl.workloads {
		if w.type_ == zos.workload_types.zmachine {
			res := json.decode(zos.ZmachineResult, w.result.data)!
			return res
		}
	}

	return error('failed to get zmachine workload')
}
