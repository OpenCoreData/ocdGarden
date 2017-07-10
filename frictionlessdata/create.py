import datapackage

dp = datapackage.DataPackage()
dp.descriptor['name'] = 'my_sleep_duration'
dp.descriptor['resources'] = [
    {'name': 'data'}
]

resource = dp.resources[0]
resource.descriptor['data'] = [
    7, 8, 5, 6, 9, 7, 8
]

with open('datapackage.json', 'w') as f:
  f.write(dp.to_json())
# {"name": "my_sleep_duration", "resources": [{"data": [7, 8, 5, 6, 9, 7, 8], "name": "data"}]}
