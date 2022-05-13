import requests

def url_ok(url):
    r = requests.head(url)
    return r.status_code == 200

print(url_ok("https://www.youtube.com"))