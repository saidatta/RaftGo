package conf

import "strings"

type RaftConfig struct {
	NodeId       string
	ClusterNodes string
}

func (c *RaftConfig) ParseClusterNodes() []string {
	if c.ClusterNodes == "" {
		panic("Cluster node config cannot be empty")
	}

	if strings.Contains(c.ClusterNodes, ",") {
		return strings.Split(c.ClusterNodes, ",")
	} else if strings.Contains(c.ClusterNodes, ";") {
		return strings.Split(c.ClusterNodes, ";")
	} else {
		return []string{c.ClusterNodes}
	}
}
