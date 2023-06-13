module main

import threefoldtech.threebot.tfgrid { TFGridClient, Taiga, TaigaResult }
import log { Logger }
import flag { FlagParser }
import os
import freeflowuniverse.crystallib.rpcwebsocket

const (
	default_server_address = 'ws://127.0.0.1:8080'
)

fn deploy_taiga(mut fp FlagParser, mut t TFGridClient) !TaigaResult {
	fp.usage_example('deploy [options]')

	name := fp.string_opt('name', `n`, 'Name of the Taiga instance')!
	farm_id := fp.int('farm_id', `f`, 0, 'Farm ID to deploy on')
	capacity := fp.string('capacity', `c`, 'medium', 'Capacity of the instance')
	ssh_key := fp.string('ssh_key', `s`, '', 'Public SSH Key to access the instance')
	admin_username := fp.string('admin_username', `u`, '', 'Admin username')
	admin_password := fp.string('admin_password', `p`, '', 'Admin password')
	admin_email := fp.string('admin_email', `e`, '', 'Admin email')
	_ := fp.finalize()!

	taiga := Taiga{
		name: name
		farm_id: u64(farm_id)
		capacity: capacity
		ssh_key: ssh_key
		admin_username: admin_username
		admin_password: admin_password
		admin_email: admin_email
	}

	return t.deploy_taiga(taiga)!
}

fn get_taiga(mut fp FlagParser, mut t TFGridClient) !TaigaResult {
	fp.usage_example('get [options]')

	name := fp.string_opt('name', `n`, 'Name of the Taiga instance')!
	_ := fp.finalize()!

	return t.get_taiga(name)!
}

fn delete_taiga(mut fp FlagParser, mut t TFGridClient) ! {
	fp.usage_example('delete [options]')

	name := fp.string_opt('name', `n`, 'Name of the Taiga instance')!
	_ := fp.finalize()!

	return t.delete_taiga(name)
}

fn main() {
	mut fp := flag.new_flag_parser(os.args)
	fp.application('Welcome to the web3_proxy client. The web3_proxy client allows you to execute all remote procedure calls that the web3_proxy server can handle.')
	fp.description('')
	fp.skip_executable()
	fp.allow_unknown_args()

	mnemonic := fp.string_opt('mnemonic', `m`, 'The mnemonic to be used to call any function') or {
		eprintln('${err}')
		exit(1)
	}
	network := fp.string('network', `n`, 'dev', 'TF network to use')
	address := fp.string('address', `a`, '${default_server_address}', 'The address of the web3_proxy server to connect to.')
	debug_log := fp.bool('debug', 0, false, 'By setting this flag the client will print debug logs too.')
	operation := fp.string_opt('operation', `o`, 'Required operation to perform ')!
	remainig_args := fp.finalize() or {
		eprintln('${err}')
		exit(1)
	}

	mut logger := Logger(&log.Log{
		level: if debug_log { .debug } else { .info }
	})

	mut myclient := rpcwebsocket.new_rpcwsclient(address, &logger) or {
		logger.error('Failed creating rpc websocket client: ${err}')
		exit(1)
	}

	_ := spawn myclient.run()

	mut tfgrid_client := tfgrid.new(mut myclient)

	tfgrid_client.load(tfgrid.Credentials{
		mnemonic: mnemonic
		network: network
	})!

	match operation {
		'deploy' {
			mut new_fp := flag.new_flag_parser(remainig_args)
			res := deploy_taiga(mut new_fp, mut tfgrid_client) or {
				logger.error('${err}')
				exit(1)
			}
			logger.info('${res}')
		}
		'get' {
			mut new_fp := flag.new_flag_parser(remainig_args)
			res := get_taiga(mut new_fp, mut tfgrid_client) or {
				logger.error('${err}')
				exit(1)
			}
			logger.info('${res}')
		}
		'delete' {
			mut new_fp := flag.new_flag_parser(remainig_args)
			delete_taiga(mut new_fp, mut tfgrid_client) or {
				logger.error('${err}')
				exit(1)
			}
		}
		else {
			logger.error('operation ${operation} is invalid')
			exit(1)
		}
	}
}
