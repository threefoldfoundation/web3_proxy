package tfgrid

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/threefoldtech/tfgrid-sdk-go/grid-client/state"
	"github.com/threefoldtech/tfgrid-sdk-go/grid-client/workloads"
	"github.com/threefoldtech/zos/pkg/gridtypes"
)

/* ***** Types ***** */

// K8sCluster struct for k8s cluster
type K8sCluster struct {
	Name        string    `json:"name"`
	Master      *K8sNode  `json:"master"`
	Workers     []K8sNode `json:"workers"`
	Token       string    `json:"token"`
	NetworkName string    `json:"network_name"`
	SSHKey      string    `json:"ssh_key"`
	AddWGAccess bool      `json:"add_wg_access"`
}

// K8sNode kubernetes data
type K8sNode struct {
	Name      string `json:"name"`
	NodeID    uint32 `json:"node_id"`
	FarmID    uint32 `json:"farm_id"`
	DiskSize  int    `json:"disk_size"`
	PublicIP  bool   `json:"public_ip"`
	PublicIP6 bool   `json:"public_ip6"`
	Planetary bool   `json:"planetary"`
	Flist     string `json:"flist"`
	CPU       int    `json:"cpu"`
	Memory    int    `json:"memory"`

	// computed
	ComputedIP4 string `json:"computed_ip4"`
	ComputedIP6 string `json:"computed_ip6"`
	WGIP        string `json:"wg_ip"`
	YggIP       string `json:"ygg_ip"`
}

type GetClusterParams struct {
	ClusterName string `json:"cluster_name"`
	MasterName  string `json:"master_name"`
}

type AddWorkerParams struct {
	ClusterName string  `json:"cluster_name"`
	MasterName  string  `json:"master_name"`
	Worker      K8sNode `json:"worker"`
}

type RemoveWorkerParams struct {
	ClusterName string `json:"cluster_name"`
	MasterName  string `json:"master_name"`
	WorkerName  string `json:"worker_name"`
}

/* ***** Methods ***** */

// K8sDeploy deploys a kubernetes cluster
func (c *Client) K8sDeploy(ctx context.Context, cluster K8sCluster) (K8sCluster, error) {
	log.Info().Msgf("Deploying k8s cluster with name %s", cluster.Name)

	if err := c.validateProjectName(ctx, cluster.Name); err != nil {
		return K8sCluster{}, err
	}

	if err := assignNodesIDsForCluster(ctx, c, &cluster); err != nil {
		return K8sCluster{}, errors.Wrapf(err, "Couldn't find node for all cluster nodes")
	}

	// deploy network
	nodes := []uint32{cluster.Master.NodeID}
	for _, worker := range cluster.Workers {
		nodes = append(nodes, worker.NodeID)
	}

	znet, err := c.deployNetwork(ctx, cluster.Name, nodes, "10.1.0.0/16", false)
	if err != nil {
		return K8sCluster{}, errors.Wrap(err, "failed to deploy network")
	}

	cluster.NetworkName = znet.Name

	// map to workloads.k8sCluster
	k8s := toGridK8s(cluster)

	// Deploy workload
	if err := c.client.DeployK8sCluster(ctx, &k8s); err != nil {
		return K8sCluster{}, errors.Wrapf(err, "Failed to deploy K8s Cluster")
	}

	// update state for both network and cluster
	nodeContracts := map[uint32]state.ContractIDs{}
	for nodeID, contract := range k8s.NodeDeploymentID {
		nodeContracts[nodeID] = append(nodeContracts[nodeID], contract)
	}

	for nodeId, contract := range znet.NodeDeploymentID {
		nodeContracts[nodeId] = append(nodeContracts[nodeId], contract)
	}

	projectName := generateProjectName(cluster.Name)

	c.Projects[projectName] = ProjectState{
		nodeContracts: nodeContracts,
	}

	return c.K8sGet(ctx, GetClusterParams{
		ClusterName: cluster.Name,
		MasterName:  cluster.Master.Name,
	})
}

// K8sDelete deletes a kubernetes cluster specified by the cluster name
func (c *Client) K8sDelete(ctx context.Context, clusterName string) error {
	log.Info().Msgf("Deleting k8s cluster with name %s", clusterName)

	if err := c.cancelModel(ctx, clusterName); err != nil {
		return errors.Wrapf(err, "failed to cancel project: %s", clusterName)
	}

	return nil
}

// K8sGet retreives a kubernetes cluster specified by the cluster name
func (c *Client) K8sGet(ctx context.Context, params GetClusterParams) (K8sCluster, error) {
	log.Info().Msgf("Getting k8s cluster with name %s", params.ClusterName)

	// load the cluster contracts
	clusterContracts, err := c.loadModelContracts(ctx, params.ClusterName)
	if err != nil {
		return K8sCluster{}, errors.Wrapf(err, "failed to get cluster %s contracts", params.ClusterName)
	}

	if len(clusterContracts.nodeContracts) == 0 {
		return K8sCluster{}, fmt.Errorf("found 0 contracts for cluster %s", params.ClusterName)
	}

	fmt.Println("Loading....", params.MasterName, clusterContracts.nodeContracts)

	// update state from the created contracts & load info from the grid
	cluster, err := c.loadK8s(params.MasterName, clusterContracts.nodeContracts)
	if err != nil {
		return K8sCluster{}, errors.Wrapf(err, "failed to load kubernetes cluster %s", params.MasterName)
	}

	// get farms to construct the cluster node object
	nodeFarms, err := getNodeFarmsIDs(c.client, &cluster)
	if err != nil {
		return K8sCluster{}, errors.Wrapf(err, "failed to get node farms ids")
	}

	// convert the cluster to the local type
	ret := fromGridK8s(cluster, params.MasterName, nodeFarms)
	return ret, nil
}

// AddK8sWorker adds a worker to a deployed kubernetes cluster
func (c *Client) AddK8sWorker(ctx context.Context, params AddWorkerParams) (K8sCluster, error) {
	log.Info().Msgf("Adding worker %s", params.Worker.Name)

	clusterContracts, err := c.loadModelContracts(ctx, params.ClusterName)
	if err != nil {
		return K8sCluster{}, errors.Wrapf(err, "failed to get kubernetes cluster %s contracts", params.ClusterName)
	}

	if len(clusterContracts.nodeContracts) == 0 {
		return K8sCluster{}, fmt.Errorf("found 0 contracts for cluster %s", params.ClusterName)
	}

	znet, err := c.loadNetwork(params.ClusterName)
	if err != nil {
		return K8sCluster{}, errors.Wrapf(err, "failed to load network for cluster %s", params.ClusterName)
	}

	nodeIds := znet.Nodes

	// use slice.Contains
	if !doesNetworkIncludeNode(znet.Nodes, params.Worker.NodeID) {
		znet.Nodes = append(znet.Nodes, params.Worker.NodeID)
		err = c.client.DeployNetwork(ctx, &znet)
		if err != nil {
			return K8sCluster{}, errors.Wrap(err, "failed to deploy network")
		}
	}

	fmt.Println("Loading....", params.MasterName, znet.Nodes, nodeIds)

	cluster, err := c.client.LoadK8s(params.MasterName, nodeIds)
	if err != nil {
		return K8sCluster{}, errors.Wrap(err, "failed to load kubernetes cluster")
	}

	fmt.Println("Loaded....")

	cluster.Workers = append(cluster.Workers, toGridK8sNode(params.Worker))
	if err := c.client.DeployK8sCluster(ctx, &cluster); err != nil {
		return K8sCluster{}, errors.Wrap(err, "failed to update kubernetes cluster")
	}

	// update state for both network and cluster
	nodeContracts := map[uint32]state.ContractIDs{}
	for nodeID, contract := range cluster.NodeDeploymentID {
		nodeContracts[nodeID] = append(nodeContracts[nodeID], contract)
	}

	for nodeId, contract := range znet.NodeDeploymentID {
		nodeContracts[nodeId] = append(nodeContracts[nodeId], contract)
	}

	projectName := generateProjectName(params.ClusterName)

	c.Projects[projectName] = ProjectState{
		nodeContracts: nodeContracts,
	}

	return c.K8sGet(ctx, GetClusterParams{
		ClusterName: params.ClusterName,
		MasterName:  params.MasterName,
	})
}

// RemoveK8sWorker removes a worker from a deployed kubernetes cluster
func (c *Client) RemoveK8sWorker(ctx context.Context, worker RemoveWorkerParams) (K8sCluster, error) {
	log.Info().Msgf("removing worker %s", worker.WorkerName)

	k8sContracts, err := c.loadModelContracts(ctx, worker.ClusterName)
	if err != nil {
		return K8sCluster{}, errors.Wrapf(err, "failed to get kubernetes cluster %s contracts", worker.ClusterName)
	}

	if len(k8sContracts.nodeContracts) == 0 {
		return K8sCluster{}, fmt.Errorf("found 0 contracts for cluster %s", worker.ClusterName)
	}

	znet, err := c.loadNetwork(worker.ClusterName)
	if err != nil {
		return K8sCluster{}, errors.Wrapf(err, "failed to load network for cluster %s", worker.ClusterName)
	}

	cluster, err := c.loadK8s(worker.MasterName, k8sContracts.nodeContracts)
	if err != nil {
		return K8sCluster{}, errors.Wrap(err, "failed to load kubernetes cluster")
	}

	workerIdx, err := getWorkerIndex(&cluster, worker.WorkerName)
	if err != nil {
		return K8sCluster{}, err
	}

	workerNodeID := cluster.Workers[workerIdx].Node

	cluster.Workers = append(cluster.Workers[:workerIdx], cluster.Workers[workerIdx+1:]...)

	// nodeIDs := []uint32{}
	// for _, worker := range cluster.Workers {
	// 	nodeIDs = append(nodeIDs, worker.Node)
	// }
	// nodeIDs = append(nodeIDs, cluster.Master.Node)

	if err := c.client.DeployK8sCluster(ctx, &cluster); err != nil {
		return K8sCluster{}, err
	}

	// TODO: check if there is no other worker on workerNodeID before updating network
	// for idx, nodeID := range znet.Nodes {
	// 	if nodeID == workerNodeID {
	// 		znet.Nodes = append(znet.Nodes[:idx], znet.Nodes[idx+1:]...)
	// 		break
	// 	}
	// }
	// if doesNetworkIncludeNode(znet.Nodes, workerNodeID) {

	// }

	// if the node doesn't have other workers remove
	if _, ok := cluster.NodeDeploymentID[workerNodeID]; !ok {
		if err := c.removeNodeFromNetwork(ctx, &znet, workerNodeID); err != nil {
			return K8sCluster{}, err
		}
	}

	// update state for both network and cluster
	nodeContracts := map[uint32]state.ContractIDs{}
	for nodeID, contract := range cluster.NodeDeploymentID {
		nodeContracts[nodeID] = append(nodeContracts[nodeID], contract)
	}

	for nodeId, contract := range znet.NodeDeploymentID {
		nodeContracts[nodeId] = append(nodeContracts[nodeId], contract)
	}

	projectName := generateProjectName(worker.ClusterName)

	c.Projects[projectName] = ProjectState{
		nodeContracts: nodeContracts,
	}

	return c.K8sGet(ctx, GetClusterParams{
		ClusterName: worker.ClusterName,
		MasterName:  worker.MasterName,
	})
}

/* ***** Helpers ***** */

// Assign chosen NodeIds to cluster node. with both way conversions to/from Reservations array.
func assignNodesIDsForCluster(ctx context.Context, client *Client, cluster *K8sCluster) error {
	// all units unified in bytes

	workloads := []*PlannedReservation{}

	ms := PlannedReservation{
		WorkloadName: cluster.Master.Name,
		FarmID:       cluster.Master.FarmID,
		MRU:          uint64(cluster.Master.Memory * int(gridtypes.Megabyte)),
		SRU:          uint64(cluster.Master.DiskSize * int(gridtypes.Gigabyte)),
		PublicIps:    cluster.Master.PublicIP,
		NodeID:       cluster.Master.NodeID,
	}

	workloads = append(workloads, &ms)

	for idx := range cluster.Workers {

		wr := PlannedReservation{
			WorkloadName: cluster.Workers[idx].Name,
			FarmID:       cluster.Workers[idx].FarmID,
			MRU:          uint64(cluster.Workers[idx].Memory * int(gridtypes.Megabyte)),
			SRU:          uint64(cluster.Workers[idx].DiskSize * int(gridtypes.Gigabyte)),
			PublicIps:    cluster.Workers[idx].PublicIP,
			NodeID:       cluster.Workers[idx].NodeID,
		}

		workloads = append(workloads, &wr)
	}

	err := client.AssignNodes(ctx, workloads)
	if err != nil {
		return err
	}

	if cluster.Master.NodeID == 0 {
		for _, workload := range workloads {
			if workload.WorkloadName == cluster.Master.Name {
				cluster.Master.NodeID = uint32(workload.NodeID)
			}
		}
	}

	for idx := range cluster.Workers {
		if cluster.Workers[idx].NodeID == 0 {
			for _, workload := range workloads {
				if workload.WorkloadName == cluster.Workers[idx].Name {
					cluster.Workers[idx].NodeID = uint32(workload.NodeID)
				}
			}
		}
	}

	return nil
}

// convert k8s cluster from local type to grid type
func toGridK8s(model K8sCluster) workloads.K8sCluster {
	master := toGridK8sNode(*model.Master)
	workers := []workloads.K8sNode{}
	for _, w := range model.Workers {
		workers = append(workers, toGridK8sNode(w))
	}

	return workloads.K8sCluster{
		Master:       &master,
		Workers:      workers,
		Token:        model.Token,
		NetworkName:  model.NetworkName,
		SolutionType: generateProjectName(model.Name),
		SSHKey:       model.SSHKey,
	}
}

// convert k8s node from local type to grid type
func toGridK8sNode(model K8sNode) workloads.K8sNode {
	return workloads.K8sNode{
		Name:        model.Name,
		Node:        model.NodeID,
		DiskSize:    model.DiskSize,
		PublicIP:    model.PublicIP,
		PublicIP6:   model.PublicIP6,
		Planetary:   model.Planetary,
		Flist:       model.Flist,
		ComputedIP:  model.ComputedIP4,
		ComputedIP6: model.ComputedIP6,
		YggIP:       model.YggIP,
		IP:          model.WGIP,
		CPU:         model.CPU,
		Memory:      model.Memory,
	}
}

// convert k8s cluster from grid type to local type
func fromGridK8s(cluster workloads.K8sCluster, clusterName string, nodeFarms map[uint32]uint32) K8sCluster {
	master := fromGridK8sNode(*cluster.Master, nodeFarms)
	workers := []K8sNode{}
	for _, worker := range cluster.Workers {
		workers = append(workers, fromGridK8sNode(worker, nodeFarms))
	}

	return K8sCluster{
		Name:        clusterName,
		Master:      &master,
		Workers:     workers,
		Token:       cluster.Token,
		NetworkName: cluster.NetworkName,
		SSHKey:      cluster.SSHKey,
	}
}

// convert k8s node from grid type to local type
func fromGridK8sNode(node workloads.K8sNode, nodeFarms map[uint32]uint32) K8sNode {
	return K8sNode{
		Name:        node.Name,
		NodeID:      node.Node,
		FarmID:      nodeFarms[node.Node],
		DiskSize:    node.DiskSize,
		PublicIP:    node.PublicIP,
		PublicIP6:   node.PublicIP6,
		Planetary:   node.Planetary,
		Flist:       node.Flist,
		CPU:         node.CPU,
		Memory:      node.Memory,
		ComputedIP4: node.ComputedIP,
		ComputedIP6: node.ComputedIP6,
		WGIP:        node.IP,
		YggIP:       node.YggIP,
	}
}

// get farmsIds for the nodes where the cluster nodes are deployed
func getNodeFarmsIDs(c TFGridClient, cluster *workloads.K8sCluster) (map[uint32]uint32, error) {
	nodeFarms := map[uint32]uint32{}

	farm, err := c.GetNodeFarm(cluster.Master.Node)
	if err != nil {
		return nil, err
	}

	nodeFarms[cluster.Master.Node] = farm

	for _, w := range cluster.Workers {
		farm, err := c.GetNodeFarm(w.Node)
		if err != nil {
			return nil, err
		}
		nodeFarms[w.Node] = farm
	}

	return nodeFarms, nil
}

// get worker index in the cluster.Workers array
func getWorkerIndex(cluster *workloads.K8sCluster, workerName string) (int, error) {
	for idx, worker := range cluster.Workers {
		if worker.Name == workerName {
			return idx, nil
		}
	}

	return 0, fmt.Errorf("failed to find a worker with name %s", workerName)
}
