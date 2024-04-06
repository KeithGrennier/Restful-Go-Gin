import requests

url="http://localhost:8080/albums"
header={"Content-Type": "applicaiton/json"}
payload='{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'
r1=requests.get(url)

r2=requests.post(url,data=payload,headers=header)

print(r1.text,"\n\n",r2.text)