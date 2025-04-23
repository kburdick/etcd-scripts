import etcd3
 
#etcd = etcd3.client('etcd') # in k8s just use the service name to connect(if running in the same cluster) 
#etcd = etcd3.client(host='172.31.2.47', port=2379)
etcd = etcd3.client()
 
etcd.put('my_key', 'my_value')
etcd.put('my_key2', 'my_value2')
etcd.put('my_key3', 'my_value3')
 
for value, metadata in etcd.get_all():
    print("Key: " + metadata.key.decode('utf-8'))
    print("Value: " + value.decode('utf-8'))

 
