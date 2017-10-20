import datapackage
import pandas

url = "https://pkgstore.datahub.io/core/co2-ppm/latest/datapackage.json"
storage = datapackage.push_datapackage(descriptor=url,backend='pandas')

storage.tables

storage['data__cpi'].head()

