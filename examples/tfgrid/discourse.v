module main

import threefoldtech.threebot.tfgrid
import threefoldtech.threebot.explorer

pub struct Discourse {
pub:
	name            string
	farm_id         u64
	cpu             u32
	memory          u32 // in mega bytes
	rootfs_size     u32 // in mega bytes
	disk_size       u32 // in giga bytes
	ssh_key         string
	developer_email string
	smtp_username   string
	smtp_password   string
	smtp_address    string = 'smtp.gmail.com'
	smtp_enable_tls bool   = true
	smtp_port       u32    = 587

	threebot_private_key string
	flask_secret_key     string
}

pub struct DiscourseResult {
pub:
	name           string
	machine_ygg_ip string
	gateway_name   string
}

fn deploy_discourse(mut client tfgrid.TFGridClient, mut explorer_client explorer.ExplorerClient, discourse Discourse) !DiscourseResult {
	gateway_nodes := explorer_client.nodes(explorer.NodesRequestParams{
		filters: explorer.NodeFilter{
			status: 'up'
			dedicated: false
			farm_ids: [discourse.farm_id]
			domain: true
		}
		pagination: explorer.Limit{
			size: 1
		}
	})!

	if gateway_nodes.nodes.len == 0 {
		return error('failed to find an eligible node for gateway')
	}

	gateway_node_id := gateway_nodes.nodes[0].node_id
	domain := gateway_nodes.nodes[0].public_config.domain
	smtp_enable_tls := if discourse.smtp_enable_tls { 'true' } else { 'false' }

	machine := client.machines_deploy(tfgrid.MachinesModel{
		name: generate_discourse_machine_name(discourse.name)
		network: tfgrid.Network{
			add_wireguard_access: false
		}
		machines: [
			tfgrid.Machine{
				name: 'discourse_vm'
				farm_id: u32(discourse.farm_id)
				cpu: discourse.cpu
				memory: discourse.memory
				rootfs_size: discourse.rootfs_size
				flist: 'https://hub.grid.tf/tf-official-apps/forum-docker-v3.1.2.flist'
				disks: [
					tfgrid.Disk{
						size: discourse.disk_size
						mountpoint: '/var/lib/docker'
					},
				]
				env_vars: {
					'SSH_KEY':                         discourse.ssh_key
					'DISCOURSE_HOSTNAME':              domain
					'DISCOURSE_DEVELOPER_EMAILS':      discourse.developer_email
					'DISCOURSE_SMTP_ADDRESS':          discourse.smtp_address
					'DISCOURSE_SMTP_PORT':             '${discourse.smtp_port}'
					'DISCOURSE_SMTP_ENABLE_START_TLS': smtp_enable_tls
					'DISCOURSE_SMTP_USER_NAME':        discourse.smtp_username
					'DISCOURSE_SMTP_PASSWORD':         discourse.smtp_password
					'THREEBOT_PRIVATE_KEY':            discourse.threebot_private_key
					'FLASK_SECRET_KEY':                discourse.flask_secret_key
				}
				planetary: true
			},
		]
	}) or {
		client.machines_delete(generate_discourse_machine_name(discourse.name))!
		return error('failed to deploy discourse instance: ${err}')
	}

	gateway := client.gateways_deploy_name(tfgrid.GatewayName{
		name: discourse.name
		backends: ['http://${machine.machines[0].ygg_ip}:88']
		node_id: u32(gateway_node_id)
	}) or {
		// if either deployment failed, delete all created contracts
		client.machines_delete(generate_discourse_machine_name(discourse.name))!
		client.gateways_delete_name(discourse.name)!
		return error('failed to deploy discourse instance: ${err}')
	}

	return DiscourseResult{
		name: discourse.name
		machine_ygg_ip: machine.machines[0].ygg_ip
		gateway_name: gateway.fqdn
	}
}

fn delete_discourse(mut client tfgrid.TFGridClient, discourse_name string) ! {
	client.gateways_delete_name(discourse_name)!
	client.machines_delete(generate_discourse_machine_name(discourse_name))!
}

fn get_discourse(mut client tfgrid.TFGridClient, discourse_name string) !DiscourseResult {
	machine := client.machines_get(generate_discourse_machine_name(discourse_name))!
	gateway := client.gateways_get_name(discourse_name)!

	return DiscourseResult{
		name: discourse_name
		machine_ygg_ip: machine.machines[0].ygg_ip
		gateway_name: gateway.fqdn
	}
}

fn generate_discourse_machine_name(discourse_name string) string {
	return '${discourse_name}_discourse_machine'
}
