## Open Core Data Example Data Set

### About
This is an example data set from Open Core Data.  It is being used to show an example
of the Frictionless Data Package use.

### URL
You can view this package with the following URL

http://data.okfn.org/tools/view?url=https%3A%2F%2Fraw.githubusercontent.com%2FOpenCoreData%2FocdGarden%2Fmaster%2Ffrictionlessdata%2FfdpDemo%2Fdatapackage.json 


### iPython

```
$ pip install datapackage
```



```
import datapackage

dp = datapackage.DataPackage('https://pkgstore.datahub.io/core/co2-ppm/latest/datapackage.json')

# see metadata
print(dp.descriptor)

# get list of csv files
csvList = [dp.resources[x].descriptor['name'] for x in range(0,len(dp.resources))]
print(csvList) # ["resource name", ...]

# access csv file by the index starting 0
print(dp.resources[0].data)
```