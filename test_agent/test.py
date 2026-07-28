import requests

# Version 2
# url = "https://ngw.devices.sberbank.ru:9443/api/v2/oauth"
# value = "af6e3a2d-a228-443f-9fa2-cb13d024b82c"
# payload = 'scope=GIGACHAT_API_B2B'

# headers = {
#   'Content-Type': 'application/x-www-form-urlencoded',
#   'Accept': 'application/json',
#   'RqUID': '{value}',
#   'Authorization': 'Basic MDE5ZjNjZGItYmM3Yy03YTRkLWE0ZTYtZDY2M2I1ODk2OTkwOmZiMDMzYmVhLTAxZjktNGFiNy04ZGZmLTc3MTFlOWU4ZDRhNQ=='
# }

# response = requests.request("POST", url, headers=headers, data=payload)
# print(response.text)


# Version 2
# url = "https://ngw.devices.sberbank.ru:9443/api/v2/oauth"
# url = "https://gigachat-api.sberbank.ru/api/v2/oauth"
url = "https://gigachat.devices.sberbank.ru/api/"

payload={
  'scope': 'GIGACHAT_API_B2B'
}
headers = {
  'Content-Type': 'application/x-www-form-urlencoded',
  'Accept': 'application/json',
  'RqUID': '8642167b-8e00-4a33-884b-9d5f4f33a935',
  'Authorization': 'Basic MDE5ZjNjZGItYmM3Yy03YTRkLWE0ZTYtZDY2M2I1ODk2OTkwOmZiMDMzYmVhLTAxZjktNGFiNy04ZGZmLTc3MTFlOWU4ZDRhNQ=='
}

response = requests.request("POST", url, headers=headers, data=payload)

print(response.text)



# curl -L -X POST 'https://ngw.devices.sberbank.ru:9443/api/v2/oauth' \
# -H 'Content-Type: application/x-www-form-urlencoded' \
# -H 'Accept: application/json' \
# -H 'RqUID: 9d989aae-6189-49d1-ac9f-83eeca0eda47' \
# -H 'Authorization: Basic MDE5ZjNjZGItYmM3Yy03YTRkLWE0ZTYtZDY2M2I1ODk2OTkwOmZiMDMzYmVhLTAxZjktNGFiNy04ZGZmLTc3MTFlOWU4ZDRhNQ==' \
# --data-urlencode 'scope=GIGACHAT_API_B2B'
