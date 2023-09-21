module tfgrid

import os
import json
import time
import log
import models

pub struct Deployer {
pub:
	mnemonics     string
	substrate_url string
	twin_id       u32
	relay_url     string
pub mut:
	logger log.Log
}

pub enum ChainNetwork {
	dev
	qa
	test
	main
}

const substrate_url = {
	ChainNetwork.dev:  'wss://tfchain.dev.grid.tf/ws'
	ChainNetwork.qa:   'wss://tfchain.qa.grid.tf/ws'
	ChainNetwork.test: 'wss://tfchain.test.grid.tf/ws'
	ChainNetwork.main: 'wss://tfchain.grid.tf/ws'
}

const relay_url = {
	ChainNetwork.dev:  'wss://relay.dev.grid.tf'
	ChainNetwork.qa:   'wss://relay.qa.grid.tf'
	ChainNetwork.test: 'wss://relay.test.grid.tf'
	ChainNetwork.main: 'wss://relay.grid.tf'
}

pub fn get_mnemonics() !string {
	mnemonics := os.getenv('MNEMONICS')
	if mnemonics == '' {
		return error('failed to get mnemonics, run `export MNEMONICS=....`')
	}
	return mnemonics
}

pub fn new_deployer(mnemonics string, chain_network ChainNetwork, mut logger log.Log) !Deployer {
	twin_id := get_user_twin(mnemonics, tfgrid.substrate_url[chain_network])!

	return Deployer{
		mnemonics: mnemonics
		substrate_url: tfgrid.substrate_url[chain_network]
		twin_id: twin_id
		relay_url: tfgrid.relay_url[chain_network]
		logger: logger
	}
}

fn (mut d Deployer) handel_deploy(node_id u32, mut dl models.Deployment, body string, solution_provider u64, hash_hex string) !u64 {
	signature := d.sign_deployment(hash_hex)!
	dl.add_signature(d.twin_id, signature)
	payload := dl.json_encode()

	node_twin_id := get_node_twin(node_id, d.substrate_url)!
	d.rmb_deployment_deploy(node_twin_id, payload)!
	workload_versions := d.assign_versions(dl)
	d.wait_deployment(node_id, dl.contract_id, workload_versions)!
	return dl.contract_id
}

pub fn (mut d Deployer) deploy(node_id u32, mut dl models.Deployment, body string, solution_provider u64) !u64 {
	public_ips := dl.count_public_ips()
	hash_hex := dl.challenge_hash().hex()

	contract_id := d.create_node_contract(node_id, body, hash_hex, public_ips, solution_provider)!
	d.logger.info('ContractID: ${contract_id}')
	dl.contract_id = contract_id

	return d.handel_deploy(node_id, mut dl, body, solution_provider, hash_hex) or {
		d.logger.info('Rolling back...')
		d.logger.info('deleting contract id: ${contract_id}')
		d.cancel_contract(contract_id)!
		return err
	}
}

pub fn (mut d Deployer) assign_versions(dl models.Deployment) map[string]u32 {
	mut workload_versions := map[string]u32{}
	for wl in dl.workloads {
		workload_versions[wl.name] = wl.version
	}
	return workload_versions
}

pub fn (mut d Deployer) wait_deployment(node_id u32, contract_id u64, workload_versions map[string]u32) ! {
	start := time.now()
	num_workloads := workload_versions.len
	for {
		mut state_ok := 0
		changes := d.deployment_changes(node_id, contract_id)!
		for wl in changes {
			if wl.version == workload_versions[wl.name]
				&& wl.result.state == models.result_states.ok {
				state_ok++
			} else if wl.version == workload_versions[wl.name]
				&& wl.result.state == models.result_states.error {
				return error('failed to deploy deployment due error: ${wl.result.message}')
			}
		}
		if state_ok == num_workloads {
			return
		}
		if (time.now() - start).minutes() > 5 {
			return error('failed to deploy deployment: contractID: ${contract_id}, some workloads are not ready after wating 5 minutes')
		} else {
			d.logger.info('Waiting for deployment to become ready')
			time.sleep(1 * time.second)
		}
	}
}

pub fn (mut d Deployer) get_deployment(contract_id u64, node_id u32) !models.Deployment {
	twin_id := get_node_twin(node_id, d.substrate_url)!
	payload := {
		'contract_id': contract_id
	}
	res := d.rmb_deployment_get(twin_id, json.encode(payload))!
	return json.decode(models.Deployment, res)
}

pub fn (mut d Deployer) deployment_changes(node_id u32, contract_id u64) ![]models.Workload {
	twin_id := get_node_twin(node_id, d.substrate_url)!

	res := d.rmb_deployment_changes(twin_id, contract_id)!
	return json.decode([]models.Workload, res)
}

pub fn (mut d Deployer) sign_deployment(hash string) !string {
	res := os.execute("grid-cli sign  --substrate \"${d.substrate_url}\" --mnemonics \"${d.mnemonics}\"  --hash \"${hash}\"")
	if res.exit_code != 0 {
		return error(res.output)
	}
	return res.output
}

pub fn (mut d Deployer) deploy_single_vm() !string {
	network_name := 'nettest2'
	name := 'vm1'
	flist := 'https://hub.grid.tf/tf-official-apps/base:latest.flist'
	cpu := 2
	memory := 10
	rootfs := 4
	entrypoint := '/sbin/zinit init'
	node := 11
	sshkey := 'ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDKkaopr/tRilmLprntgqAO6JxnhzhrrB02KeUsNys9qj/G4RtXK4hKZo3yj42Kuoub53TxoW/BfZSRUbY0VNUgZSsifCyDE4g1UXi83ic+uepPb1VIfzlFdtZPUo/dEtjfS5FM7GVAPZCDik08w2+uXeZiAKavLTDQFh2cuIqE5QnQ44enODTMkSDPVJSKJ6NSoNcrY2I++AtcxgNFzy/7YWoT/bA19CliGqxBMSQ/GjOEAF3iQjUPc5LYcqCYAc0Y6WPt2l8uEVvhJAtsLelGApt8v/Nq/OBEKJpUQB2cfkyKlwLLphKFtQ8gcKjrbpE47lVsfbW68uqCw/5avA71HlS6AXxMeda8GZ99UqDTOLoUbd+EEo17hiHs6Nvle8DHNfWdGT2wx+DhzZeCO719UadmQxFLYDd75dDH5gLkMxr9JXPWDqdxvsyMZilxZPk3UQbK811obYSMrc9L+u8vwLs0weUBxnygutnU0eF9cRQxFgx3zOOfwI1ugI9SgSU= omarkassem099@gmail.com'

	res := os.execute("grid-cli deploy-single --mnemonics \"${d.mnemonics}\" --env dev  --network_name \"${network_name}\" --vm_name \"${name}\" --flist \"${flist}\" --cpu \"${cpu}\" --memory \"${memory}\" --node \"${node}\" --entrypoint \"${entrypoint}\" --sshkey \"${sshkey}\" --rootfs \"${rootfs}\"")
	return res.output
}
