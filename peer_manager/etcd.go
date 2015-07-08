package peer_manager

import (
	"time"

	"github.com/bountylabs/log"
	"github.com/coreos/go-etcd/etcd"
)

func (this *EtcdPeerManager) GetPeers() []string {
	out, err := GetEntries(this.client, this.Directory, this.Generator)
	if err != nil {
		log.Errorln(err)
	}
	return out
}

func (this *EtcdPeerManager) PeersChanged(stop chan bool) chan struct{} {
	receiver := make(chan *etcd.Response, 10)
	out := make(chan struct{}, 1)
	go func() {

		for {
			_, err := this.client.Watch(this.Directory, 0, true, receiver, stop)
			if err != nil {
				log.Errorln(err)
			}
			time.Sleep(20 * time.Second)
		}

		for {
			_, ok := <-receiver
			if !ok {
				return
			}
			out <- struct{}{}
		}

	}()
	return out
}

func GetEntries(client *etcd.Client, directory string, generator func(node string) string) ([]string, error) {

	resp, err := client.Get(directory, false, true)
	if err != nil {
		if casted, ok := err.(*etcd.EtcdError); ok {
			//100 = Not Found
			//if not found continue, else err
			if casted.ErrorCode == 100 {
				return nil, nil
			}
		}
		return nil, err
	}

	out := []string{}
	for _, host := range resp.Node.Nodes {
		for _, node := range host.Nodes {
			if record := generator(node.Value); record != "" {
				out = append(out, record)
			}
		}
	}
	return out, nil
}
