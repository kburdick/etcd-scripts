import etcd3
import time
 
etcd = etcd3.client()
 
def etcd_watch_callback(watchResp):
  #print(watchResp)
  #print(type(watchResp))
  #print(watchResp.events)
  events = watchResp.events
  
  for event in events:
    #print(event)  
    key = event.key.decode('utf-8')
    value = event.value.decode('utf-8')

    if type(event) == etcd3.events.PutEvent:
      print("You created the key " + key + " I was looking for with value " + value)
    elif type(event) == etcd3.events.DeleteEvent:
      print("You deleted the key " + key + " I was looking for with value " + value)

watch_id = etcd.add_watch_callback('/list', etcd_watch_callback, range_end='/list0') 
# BROKEN etcd.add_watch_callback(range_start="/list/", range_end="/list0", callback=etcd_watch_callback)
#etcd.add_watch_callback("/list/kurt", etcd_watch_callback)
 
while True:
  time.sleep(1)
