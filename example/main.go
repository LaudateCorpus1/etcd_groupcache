package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/bountylabs/etcd_groupcache/cache"
	"github.com/bountylabs/log"
	"github.com/golang/groupcache"
)

var port = flag.String("port", "8080", "--port 8080")

func main() {

	if !flag.Parsed() {
		flag.Parse()
	}

	// 	etcd := os.Getenv("ETCD_ADDR")
	ip := os.Getenv("PUBLIC_HOSTIP")
	pport := os.Getenv("PUBLIC_PORT")

	cache := etcd_groupcache.New(fmt.Sprintf("http://%s:%s", ip, pport), nil, nil)
	peers := []string{fmt.Sprintf("http://%s:8080", ip), fmt.Sprintf("http://%s:8081", ip)}
	log.Infoln(peers)
	cache.Set(peers...)
	http.Handle("/_groupcache/", cache)

	var stringcache = groupcache.NewGroup("SlowDBCache", 64<<20, groupcache.GetterFunc(
		func(ctx groupcache.Context, key string, dest groupcache.Sink) error {
			log.Infoln(pport)
			dest.SetString(key + "heyo")
			return nil
		}))

	go func() {
		var data []byte
		key := "some key"
		err := stringcache.Get(nil, key, groupcache.AllocatingByteSliceSink(&data))
		if err != nil {
			log.Errorln(err)
		}
		log.Infoln(string(data))
	}()

	log.Infoln("Groupcache Listening 8080")
	if err := http.ListenAndServe(":"+*port, nil); err != nil {
		panic(err)
	}
}
