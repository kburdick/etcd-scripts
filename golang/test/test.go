package main

import "time"
import "context"
import "os"
import "log"
import "go.etcd.io/etcd/client/v3"

func client() clientv3.KV {

   client, err := clientv3.New(clientv3.Config{
      DialTimeout:          2 * time.Second,
      Endpoints:            []string{"localhost:2379"},
      Username:             "test",
      Password:             os.Getenv("AUTH"),
      DialKeepAliveTime:    2 * time.Second,
      DialKeepAliveTimeout: 2 * time.Second,
   })
   
   if err != nil {
      log.Fatal("Failed creating client!")
   }

   etcdClient := clientv3.NewKV(client)

   log.Println("Connected")
   resp, err := client.Status(context.Background(), "localhost:2379")
   if err != nil {
      log.Fatal(err)
   }

   log.Println("etcd Version:", resp.Version)

   return etcdClient
}


func watchTest(etcdClient clientv3.KV) {

   for {
      ctxPut, _ := context.WithTimeout(context.Background(), 5 * time.Second)
      timestampString := time.Now().String()
      putRes, err := etcdClient.Put(ctxPut, "foo", timestampString)
      log.Println(timestampString)
      if err != nil {
         log.Fatal("Failed putting value!")
      } else {
         log.Printf("result =%+v", putRes.Header.Revision)
      }

      ctxGet, _ := context.WithTimeout(context.Background(), 5 * time.Second)
      getRes, err := etcdClient.Get(ctxGet, "foo")
      if err != nil {
         log.Fatal("Failed getting value!")
      } else {
         keyString := string(getRes.Kvs[0].Key)
         log.Printf("Key String: %v", keyString)
         //log.Printf("Key found =%+v", getRes.Kvs[0].Key)
         valueString := string(getRes.Kvs[0].Value)
         log.Printf("Value String: %v", valueString)
         //log.Printf("Value found =%+v", getRes.Kvs[0].Value)
      }

      log.Println("sleeping for 5 seconds")
      time.Sleep(5 * time.Second)
   }

}


func putAndGetTest(etcdClient clientv3.KV) {

   timeout := 5 * time.Second
   ctx, cancel := context.WithTimeout(context.Background(), timeout)
   resp, err := etcdClient.Put(ctx, "sample_key", "sample_value")
  
   if err != nil {
      // handle error!
      log.Printf("error running put command! %v", err) 
   }

   log.Printf("PUT successful, key revision: %v", resp.Header.Revision)

   getResp, err := etcdClient.Get(ctx, "sample_key")
   cancel()
   if err != nil {
      // handle error!
      log.Printf("error running get command! %v", err)
   }

   if getResp.Count > 0 {
      for _, kv := range getResp.Kvs {
         log.Printf("Key: %s, Value: %s\n", kv.Key, kv.Value)
      } 
   }

   log.Printf("GET successful, key revision: %v", getResp.Header.Revision)
}

func main() {
   log.Println("Creating client and connecting to etcd")
   etcdClient := client()
   log.Println("starting basic put and get test")
   putAndGetTest(etcdClient)
   log.Println("sleeping for 10 seconds...")
   time.Sleep(10 * time.Second)
   log.Println("starting continuous watch test")
   watchTest(etcdClient)
}
