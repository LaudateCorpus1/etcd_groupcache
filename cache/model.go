package etcd_groupcache

import "github.com/golang/groupcache"

type PeerManager interface {
	GetPeers() []string
	PeersChanged() chan struct{}
}

type Cache struct {
	pm PeerManager
	*groupcache.HTTPPool
}

func New(listeningAddr string, pm PeerManager, c chan struct{}) *Cache {
	pool := groupcache.NewHTTPPoolOpts(listeningAddr, nil)
	if pm != nil {
		pool.Set(pm.GetPeers()...)
		go func() {
			changeChannel := pm.PeersChanged()
			for {
				select {
				case <-c:
				case <-changeChannel:
					pool.Set(pm.GetPeers()...)
				}
			}
		}()
	}
	return &Cache{HTTPPool: pool, pm: pm}
}
