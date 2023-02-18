import docker
import requests


def getData():
    url = "https://hackattic.com/challenges/dockerized_solutions/problem?access_token=8e80fec0cbe25049"
    response = requests.get(url)
    data = response.json()
    print(data)
    return data


data = getData()

dockerClient = docker.from_env()
dockerRes = dockerClient.login(username=data["credentials"]["user"],
                               password=data["credentials"]["password"], registry="localhost:5000")


postUrl = "https://hackattic.com/_/push/"+data["trigger_token"]
response = requests.post(postUrl, json={
    "registry_host": "0ac0-2405-201-d033-a065-d93b-7a1c-4baa-2a96.ngrok.io"})


dockerRes = requests.get("http://localhost:5000/v2/hack/tags/list")
tags = dockerRes.json()["tags"]

secret = ""
for i in range(len(tags)):
    response = dockerClient.containers.run(
        "localhost:5000/hack:"+tags[0], environment=["IGNITION_KEY="+data["ignition_key"]])
    print(str(response.decode('UTF-8')))
    if str(response).__contains__("wrong"):
        continue
    else:
        secret = response.decode('UTF-8')
        break


postUrl = "https://hackattic.com/challenges/dockerized_solutions/solve?access_token=8e80fec0cbe25049"
response = requests.post(postUrl, json={
    "secret": secret[:-1]})

print(response.json())
