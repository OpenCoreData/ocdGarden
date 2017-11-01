import datapackage
import pandas

url = 'http://7af2d071.ngrok.io/datapackage.json'
# url = 'https://raw.githubusercontent.com/OpenCoreData/ocdGarden/master/frictionlessdata/fdpDemo/datapackage.json'
storage = datapackage.push_datapackage(descriptor=url,backend='pandas')
storage.tables
storage['154_925C_JanusChemCarb_FIUtooys'].head()


