import etcd3
import os


etcd = etcd3.client()
 
etcd.put('/names/jon', 'smith')
etcd.put('/names/bob', 'younger')
etcd.put('/names/mary', 'heston')
 
print("Jon:  " + str(list('/names/jon'.encode(('utf-8')))))
print("Bob:  " + str(list('/names/bob'.encode(('utf-8')))))
print("Mary: " + str(list('/names/mary'.encode(('utf-8')))))
 
etcd.put('/addresses/jon', '101 Swiss Hill')
etcd.put('/addresses/bob', '510 Cherry Tree Ln')
etcd.put('/addresses/mary', '123 Mark Drive')

for value_names, metadata_names in etcd.get_range(range_start="/names/", range_end="/names0"): 
    f_name = metadata_names.key.decode('utf-8').split('/')[2]
    l_name = value_names.decode('utf-8')
    for value_addresses, metadata_addresses in etcd.get_prefix('/addresses/'):
        if metadata_addresses.key.decode('utf-8').split('/')[2] == f_name:
            address = value_addresses.decode('utf-8')
    print(f_name + " " + l_name + " lives at " + address)
