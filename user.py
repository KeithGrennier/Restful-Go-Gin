import requests

url="http://localhost:8080/albums"
header={"Content-Type": "applicaiton/json"}
payload='{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'
r1=requests.get(url)

r2=requests.post(url,data=payload,headers=header)
r3=requests.get(url+"/2")
print(r1.text,"\n\n",r2.text,"\n\n",r3.text)
#r1.content,"\n\n",r2.content