package peer_manager

import "github.com/coreos/go-etcd/etcd"

type EtcdPeerManager struct {
	Directory string
	Generator func(string) string
	client    *etcd.Client
}

func NewEtcdPeerManger(directory string, endpoint string) *EtcdPeerManager {
	machines := []string{endpoint}
	client := etcd.NewClient(machines)
	return &EtcdPeerManager{Directory: directory, client: client}
}
