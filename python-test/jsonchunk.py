import etcd3
import json
 
etcd = etcd3.client()
 
list_of_people = [{"f_name": "Jon", "l_name": "Smith", "address":
                  [{"street": "101 Swiss Hill Dr", "state": "MN", "zip": "55455"}]},
                  {"f_name": "Bob", "l_name": "Younger", "address":
                  [{"street": "510 Cherry Tree Ln", "state": "WI", "zip": "53099"}]},
                  {"f_name": "Mary", "l_name": "Heston", "address":
                  [{"street": "123 Mark Drive", "state": "CO", "zip": "52345"}]}]
 
for person in list_of_people:
  json_person = json.dumps(person)
  key = "/list/person/" + person['address'][0]['street']
  etcd.put(key, json_person)
 
for value_name, metadata_names in etcd.get_prefix('/list/person/'):
  person = json.loads(value_name.decode('utf-8'))
  key = metadata_names.key.decode('utf-8').split('/')[3]
  f_name = person['f_name']
  l_name = person['l_name']
  address_street = person['address'][0]['street']
  address_state = person['address'][0]['state']
  address_zip = person['address'][0]['zip']
 
  print(f_name + " "  + l_name + " Lives at " + address_street + ", " + address_state + " " + address_zip)

  if key == address_street:
    print("Found match - deleting key!")
    etcd.delete("/list/person/" + key)
    print()


