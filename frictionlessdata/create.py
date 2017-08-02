import datapackage
import io
import csv
from jsontableschema import infer

dp = datapackage.DataPackage()
dp.descriptor['name'] = '154_925C_JanusChemCarb_FIUtooys'
dp.descriptor['title'] = '154 925 C JanusChemCarb'


filepath = './testdata/data.csv'

with io.open(filepath) as stream:
    headers = stream.readline().rstrip('\n').split('\t')
    values = csv.reader(stream)
    schema = infer(headers, values)
    dp.descriptor['resources'] = [
        {
            'name': 'data',
            'path': filepath,
            'schema': schema
        }
    ]

with open('datapackage.json', 'w') as f:
  f.write(dp.to_json())

# dp.descriptor['resources'] = [
#     {'name': 'data'}
# ]

# resource = dp.resources[0]
# resource.descriptor['data'] = [
#     7, 8, 5, 6, 9, 7, 8
# ]

# with open('datapackage.json', 'w') as f:
#   f.write(dp.to_json())
# # {"name": "my_sleep_duration", "resources": [{"data": [7, 8, 5, 6, 9, 7, 8], "name": "data"}]}
