package tfgrid

import (
	"context"
	"net"

	"github.com/pkg/errors"
	client "github.com/threefoldtech/tfgrid-sdk-go/grid-client/node"
	"github.com/threefoldtech/zos/pkg/capacity/dmi"
	"github.com/threefoldtech/zos/pkg/gridtypes"
	"github.com/threefoldtech/zos/pkg/gridtypes/zos"
)

type Statistics struct {
	Total gridtypes.Capacity `json:"total"`
	Used  gridtypes.Capacity `json:"used"`
}

// PoolMetrics stores storage pool metrics
type PoolMetrics struct {
	Name string         `json:"name"`
	Type zos.DeviceType `json:"type"`
	Size gridtypes.Unit `json:"size"`
	Used gridtypes.Unit `json:"used"`
}

func (r *Client) ZOSDeploymentDeploy(ctx context.Context, nodeID uint32, dl gridtypes.Deployment) error {

	nodeClient, err := r.GridClient.GetNodeClient(nodeID)
	if err != nil {
		return errors.Wrapf(err, "failed to get node %d client", nodeID)
	}

	if err := dl.Sign(r.TwinID, r.Identity); err != nil {
		return errors.Wrap(err, "failed to sign deployment")
	}

	return nodeClient.DeploymentDeploy(ctx, dl)
}

func (r *Client) ZOSDeploymentGet(ctx context.Context, nodeID uint32, contractID uint64) (gridtypes.Deployment, error) {

	nodeClient, err := r.GridClient.GetNodeClient(nodeID)
	if err != nil {
		return gridtypes.Deployment{}, errors.Wrapf(err, "failed to get node %d client", nodeID)
	}

	deployment, err := nodeClient.DeploymentGet(ctx, contractID)
	if err != nil {
		return gridtypes.Deployment{}, errors.Wrapf(err, "failed to get deployment with contract id %d", contractID)
	}

	return deployment, nil
}

func (r *Client) ZOSDeploymentDelete(ctx context.Context, nodeID uint32, contractID uint64) error {

	nodeClient, err := r.GridClient.GetNodeClient(nodeID)
	if err != nil {
		return errors.Wrapf(err, "failed to get node %d client", nodeID)
	}

	err = nodeClient.DeploymentDelete(ctx, contractID)
	if err != nil {
		return errors.Wrapf(err, "failed to delete deployment with contract id %d", contractID)
	}

	return nil
}

func (r *Client) ZOSDeploymentUpdate(ctx context.Context, nodeID uint32, dl gridtypes.Deployment) error {

	nodeClient, err := r.GridClient.GetNodeClient(nodeID)
	if err != nil {
		return errors.Wrapf(err, "failed to get node %d client", nodeID)
	}

	if err := dl.Sign(r.TwinID, r.Identity); err != nil {
		return errors.Wrap(err, "failed to sign deployment")
	}

	err = nodeClient.DeploymentUpdate(ctx, dl)
	if err != nil {
		return errors.Wrapf(err, "failed to update deployment with contract id %d", dl.ContractID)
	}

	return nil
}

func (r *Client) ZOSDeploymentChanges(ctx context.Context, nodeID uint32, contractID uint64) ([]gridtypes.Workload, error) {
	nodeClient, err := r.GridClient.GetNodeClient(nodeID)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get node %d client", nodeID)
	}

	workloads, err := nodeClient.DeploymentChanges(ctx, contractID)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get changes for deployment with contract id %d", contractID)
	}

	return workloads, nil
}

func (r *Client) ZOSStatisticsGet(ctx context.Context, nodeID uint32) (Statistics, error) {
	nodeClient, err := r.GridClient.GetNodeClient(nodeID)
	if err != nil {
		return Statistics{}, errors.Wrapf(err, "failed to get node %d client", nodeID)
	}

	total, used, err := nodeClient.Statistics(ctx)
	if err != nil {
		return Statistics{}, errors.Wrapf(err, "failed to get statistics for node with id %d", nodeID)
	}

	return Statistics{
		Total: total,
		Used:  used,
	}, nil
}

func (r *Client) ZOSNetworkListWGPorts(ctx context.Context, nodeID uint32) ([]uint16, error) {
	nodeClient, err := r.GridClient.GetNodeClient(nodeID)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get node %d client", nodeID)
	}

	ports, err := nodeClient.NetworkListWGPorts(ctx)
	if err != nil {
		return nil, err
	}

	return ports, nil
}

func (r *Client) ZOSNetworkInterfaces(ctx context.Context, nodeID uint32) (map[string][]net.IP, error) {
	nodeClient, err := r.GridClient.GetNodeClient(nodeID)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get node %d client", nodeID)
	}

	ips, err := nodeClient.NetworkListInterfaces(ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get statistics for node with id %d", nodeID)
	}

	return ips, nil
}

func (r *Client) ZOSNetworkPublicConfigGet(ctx context.Context, nodeID uint32) (client.PublicConfig, error) {
	nodeClient, err := r.GridClient.GetNodeClient(nodeID)
	if err != nil {
		return client.PublicConfig{}, errors.Wrapf(err, "failed to get node %d client", nodeID)
	}

	cfg, err := nodeClient.NetworkGetPublicConfig(ctx)
	if err != nil {
		return client.PublicConfig{}, errors.Wrapf(err, "failed to get statistics for node with id %d", nodeID)
	}

	return cfg, nil
}

func (r *Client) ZOSSystemDMI(ctx context.Context, nodeID uint32) (dmi.DMI, error) {
	nodeClient, err := r.GridClient.GetNodeClient(nodeID)
	if err != nil {
		return dmi.DMI{}, errors.Wrapf(err, "failed to get node %d client", nodeID)
	}

	resDMI, err := nodeClient.SystemDMI(ctx)
	if err != nil {
		return dmi.DMI{}, errors.Wrapf(err, "failed to get statistics for node with id %d", nodeID)
	}

	return resDMI, nil
}

func (r *Client) ZOSSystemHypervisor(ctx context.Context, nodeID uint32) (string, error) {
	nodeClient, err := r.GridClient.GetNodeClient(nodeID)
	if err != nil {
		return "", errors.Wrapf(err, "failed to get node %d client", nodeID)
	}

	res, err := nodeClient.SystemHypervisor(ctx)
	if err != nil {
		return "", errors.Wrapf(err, "failed to get statistics for node with id %d", nodeID)
	}

	return res, nil
}

func (r *Client) ZOSVersion(ctx context.Context, nodeID uint32) (client.Version, error) {
	nodeClient, err := r.GridClient.GetNodeClient(nodeID)
	if err != nil {
		return client.Version{}, errors.Wrapf(err, "failed to get node %d client", nodeID)
	}

	version, err := nodeClient.SystemVersion(ctx)
	if err != nil {
		return client.Version{}, errors.Wrapf(err, "failed to get statistics for node with id %d", nodeID)
	}

	return version, nil
}

func (r *Client) ZOSStoragePools(ctx context.Context, nodeID uint32) (pools []client.PoolMetrics, err error) {
	nodeClient, err := r.GridClient.GetNodeClient(nodeID)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get node %d client", nodeID)
	}

	pools, err = nodeClient.Pools(ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get node %d storage pools", nodeID)
	}

	return pools, nil
}

func (r *Client) ZosHasPublicIPv6(ctx context.Context, nodeID uint32) (bool, error) {
	nodeClient, err := r.GridClient.GetNodeClient(nodeID)
	if err != nil {
		return false, errors.Wrapf(err, "failed to get node %d client", nodeID)
	}

	hasPublicIPv6, err := nodeClient.HasPublicIPv6(ctx)
	if err != nil {
		return false, errors.Wrapf(err, "failed to get node %d storage pools", nodeID)
	}

	return hasPublicIPv6, nil
}

func (r *Client) ZOSNetworkListAllInterfaces(ctx context.Context, nodeID uint32) (result map[string]client.Interface, err error) {
	nodeClient, err := r.GridClient.GetNodeClient(nodeID)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get node %d client", nodeID)
	}

	interfaces, err := nodeClient.NetworkListAllInterfaces(ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get node %d storage pools", nodeID)
	}

	return interfaces, nil
}

func (r *Client) ZOSNetworkSetPublicExitDevice(ctx context.Context, nodeID uint32, iface string) error {
	nodeClient, err := r.GridClient.GetNodeClient(nodeID)
	if err != nil {
		return errors.Wrapf(err, "failed to get node %d client", nodeID)
	}

	err = nodeClient.NetworkSetPublicExitDevice(ctx, iface)
	if err != nil {
		return errors.Wrapf(err, "failed to get node %d storage pools", nodeID)
	}

	return nil
}

func (r *Client) ZOSNetworkGetPublicExitDevice(ctx context.Context, nodeID uint32) (exit client.ExitDevice, err error) {
	nodeClient, err := r.GridClient.GetNodeClient(nodeID)
	if err != nil {
		return client.ExitDevice{}, errors.Wrapf(err, "failed to get node %d client", nodeID)
	}

	exit, err = nodeClient.NetworkGetPublicExitDevice(ctx)
	if err != nil {
		return client.ExitDevice{}, errors.Wrapf(err, "failed to get node %d storage pools", nodeID)
	}

	return exit, nil
}
