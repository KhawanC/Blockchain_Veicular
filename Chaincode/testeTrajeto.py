# import requests

# key = 'AIzaSyBhKwgrSEqNJPX98SH1lVQM_kL-FOAmGpo'
# d1 = 'Bataillard, Mosela, Petróplis'
# d2 = 'Av. Piabanha, Centro, Petrópolis'
# url = "https://maps.googleapis.com/maps/api/directions/json?origin={a}&destination={b}&key={c}".format(a=d1 , b= d2, c=key)

# payload={}
# headers = {}

# response = requests.request("GET", url, headers=headers, data=payload)

# print(response.text)


import requests, random, json


key = 'AIzaSyBhKwgrSEqNJPX98SH1lVQM_kL-FOAmGpo'
geolocate_url = 'https://www.googleapis.com/geolocation/v1/geolocate?key={c}'.format(c=key)
# nearby_url = "https://maps.googleapis.com/maps/api/place/nearbysearch/json?location={a}%&radius=50000&type=lodging&key={c}".format(a=lat, b=long, c=key)

response1 = requests.post(geolocate_url, {
    'location',
    'accuracy'
})

print(response1.text)

# payload={}
# headers = {}

# response2 = requests.request("GET", nearby_url, headers=headers, data=payload)

# info = json.loads(response2.text)
# aleat_num = random.randint(0, len(info['results']))
# print(info['results'][aleat_num]['vicinity'])



    