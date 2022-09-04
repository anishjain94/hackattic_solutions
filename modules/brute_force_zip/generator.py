# m = list('abcdefghijklmnopqrstuvwxyz0123456789')


# def to_base_lm(n):
#     result = ''
#     base = len(m)
#     while n > 0:
#         index = n % base
#         result = m[index] + result
#         n = n // base
#     return result


# def from_base_lm(s):
#     s = s[::-1]
#     result = 0
#     n = len(s)
#     base = len(m)
#     for i in range(n):
#         number = m.index(s[i])
#         result += number * (base ** i)
#     return result


# print(from_base_lm('aaaa'))
# print(from_base_lm('999999'))

# for i in range(from_base_lm('aaaa'), from_base_lm('999999')+1):
#     print(to_base_lm(i).rjust(4, 'a'))

from urllib import request
from zipfile import ZipFile
import subprocess
import requests
import os

resp = requests.get("https://hackattic.com/challenges/brute_force_zip/problem?access_token=8e80fec0cbe25049",
                    )

zipUrl = resp.json()["zip_url"]


zipFile = requests.get(zipUrl, stream=True)

with open("hackattic.zip", 'wb') as output_file:
    output_file.write(zipFile.content)

os.system("/opt/homebrew/Cellar/john-jumbo/1.9.0/share/john/zip2john hackattic.zip > hackattic.hashes")
output = subprocess.check_output(["john", "hackattic.hashes"])
# output = subprocess.check_output(["john", "--show", "hackattic.hashes"])
# print(output)

# output = output.split(":")

password = output[1]

# output = subprocess.check_output(["pwd"])
# print(output)
# password = "yepn7"


file_name = "hackattic.zip"
with ZipFile(file_name, "r") as zip:
    zip.extractall(path="", pwd=password.encode("utf-8"))

# read secret.txt
with open("./modules/brute_force_zip/secret.txt", "r") as f:
    secretTxt = f.read()

print(secretTxt)


# send post request to /challenges/brute_force_zip/solve?access_token=8e80fec0cbe25049


requests.post("https://hackattic.com//challenges/brute_force_zip/solve?access_token=8e80fec0cbe25049",
              data={"secret": secretTxt})
