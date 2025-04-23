package main

import "time"
import "log"
import "os"
import "context"
import "go.etcd.io/etcd/client/v3"

func watcher() clientv3.Watcher {
   
   client, err := clientv3.New(clientv3.Config{
      DialTimeout:          2 * time.Second,
      Endpoints:            []string{"localhost:2379"},
      Username:             "test",
      Password:             os.Getenv("AUTH"),
      DialKeepAliveTime:    2 * time.Second,
      DialKeepAliveTimeout: 2 * time.Second,
   })

   etcdWatcher := clientv3.NewWatcher(client)

   log.Println("Connected")
   resp, err := client.Status(context.Background(), "localhost:2379")
   if err != nil {
      log.Fatal(err)
   }

   log.Println("etcd Version:", resp.Version)

   return etcdWatcher

}


func watch(etcdWatcher clientv3.Watcher) {

   watchChannel := etcdWatcher.Watch(context.Background(), "foo")

   for watchResp := range watchChannel {
      if watchResp.Err() != nil {
         log.Fatal(watchResp.Err())
      }

      for _, ev := range watchResp.Events {
         log.Printf("Type: %s, Key: %s, Value: %s\n", ev.Type, string(ev.Kv.Key), string(ev.Kv.Value))
      }
   }
}

func main() {

   //watcher()
   etcdWatcher := watcher()
   log.Println("starting watcher to wait for key updates...")
   watch(etcdWatcher)

}
