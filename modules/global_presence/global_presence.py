# "198.49.68.80:80",
# "51.15.242.202:8888",
# "165.154.243.154:80",
# "169.57.1.85:8123",
# "172.104.115.95:80",
# "139.99.237.62:80",
# "185.76.10.144:8080",
# "8.209.246.6:80",
# "128.69.6.92:8000",
# "103.149.162.195:80",
# "20.206.106.192:8123",
# "80.48.119.28:8080",
# "165.154.243.154:80",
# "8.219.97.248:80",
# "20.111.54.16:80",
# "80.48.119.28:8080",
# "169.57.1.85:8123",
# "103.149.162.195:80"
import requests

data = requests.get(
    "https://hackattic.com/challenges/a_global_presence/problem?access_token=8e80fec0cbe25049")

token = data.json()["presence_token"]

url = f"https://hackattic.com/_/presence/{token}"


httpsIps = [
    "157.100.12.138:999",
    "49.0.2.242:8090",
    "45.42.177.17:3128",
    "49.0.2.242:8090",
    "65.21.141.242:10100",
    "157.100.12.138:999",
    "145.40.121.163:3128",
    "190.26.201.194:8080",
]


for ip in httpsIps:
    proxy = {"https": ip}
    print(ip)
    data = requests.get(url, proxies=proxy)
    print(data.content)

httpIps = [
    "193.122.71.184:3128"
]


for ip in httpIps:
    proxy = {"http": ip}
    print(ip)
    data = requests.get(url, proxies=proxy)
    print(data.content)

print(requests.get(url).content)

ans = requests.post(
    "https://hackattic.com/challenges/a_global_presence/solve?access_token=8e80fec0cbe25049",
    data="{}").json()
print(ans)
