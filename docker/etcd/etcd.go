package main

import (
	"context"
	"fmt"
	"log"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"http://localhost:2379"},
		DialTimeout: 10 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	rch := cli.Watch(context.Background(), "api/v1/foo", clientv3.WithPrefix())
	for wresp := range rch {
		for _, ev := range wresp.Events {
			fmt.Printf("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}

	// ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	// cli.Put(ctx, "language/julia", "julia")
	// cli.Put(ctx, "language/python", "python")

	// cli.KV = namespace.NewKV(cli.KV, "language")
	// cancel()
	resp, err := cli.Get(context.Background(), "julia")

	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp.Kvs)
	delResp, _ := cli.Delete(context.Background(), "bar")
	log.Println(delResp.Deleted, delResp.PrevKvs)

	//lease use

	leaseResp, _ := cli.Lease.Grant(context.Background(), 20)
	log.Printf("lease ttl:%v", leaseResp.TTL)
	log.Println(leaseResp)
	resp, err = cli.KV.Get(context.Background(), "language/python")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp.Kvs)
	// cli.KV.Txn(context.Background()).Then(clientv3.OpPut("foo", "bar")).Commit()

	putResp, err := cli.KV.Put(context.Background(), "foo", "bar", clientv3.WithLease(leaseResp.ID))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(putResp.PrevKv)

	// resp, err = cli.KV.Get(context.Background(), "foo")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println(resp.Kvs)
	// rc, kerr := cli.KeepAlive(context.Background(), leaseResp.ID)
	// if kerr != nil {
	// 	log.Fatal(kerr)
	// }
	// for i := 0; i < 10; i++ {
	// 	if _, ok := <-rc; ok {
	// 		log.Println("alive")
	// 	}
	// }

	// time.Sleep(30 * time.Second)
	// resp, err = cli.KV.Get(context.Background(), "foo")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println(resp.Kvs)

	//range kv
	// opts := clientv3.WithFromKey()
	resp, err = cli.KV.Get(context.Background(), "language", clientv3.WithRange("language/python"), clientv3.WithRev(0), clientv3.WithFromKey(),
		clientv3.WithSort(clientv3.SortByKey, clientv3.SortAscend))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp.Kvs, resp.Count)

	resp, err = cli.KV.Get(context.Background(), "language", clientv3.WithRange("language"), clientv3.WithCountOnly())
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp.Count)

}
