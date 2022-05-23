# import requests

# key = 'AIzaSyBhKwgrSEqNJPX98SH1lVQM_kL-FOAmGpo'
# d1 = 'Bataillard, Mosela, Petróplis'
# d2 = 'Av. Piabanha, Centro, Petrópolis'
# url = "https://maps.googleapis.com/maps/api/directions/json?origin={a}&destination={b}&key={c}".format(a=d1 , b= d2, c=key)

# payload={}
# headers = {}

# response = requests.request("GET", url, headers=headers, data=payload)

# print(response.text)


# import requests


# key = 'AIzaSyBhKwgrSEqNJPX98SH1lVQM_kL-FOAmGpo'
# lat = '-22.492277969335433'
# long = '-43.19911005771312'

# url = "https://maps.googleapis.com/maps/api/place/nearbysearch/json?location={a}%2C{b}&radius=1500&key={c}".format(a=lat, b=long, c=key)

# payload={}
# headers = {}

# response = requests.request("GET", url, headers=headers, data=payload)

# print(response.text)

import json

with open("trajetp.json") as f:
    arq_json = json.load(f)
    
print(arq_json)
    