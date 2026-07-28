from gigachat import GigaChat

# Используем внутренний балансировщик, доступный из вашего сегмента сети
giga = GigaChat(
    credentials="MDE5ZjNjZGItYmM3Yy03YTRkLWE0ZTYtZDY2M2I1ODk2OTkwOmZiMDMzYmVhLTAxZjktNGFiNy04ZGZmLTc3MTFlOWU4ZDRhNQ==",
    base_url="https://ngw.devices.sberbank.ru:9443/api/v2/oauth", # Внутренний адрес вместо внешнего
    scope="GIGACHAT_API_B2B" # Явное указание скоупа (так как ключ закодирован под B2B)
)

try:
    response = giga.chat("Привет! Как дела?")
    print(response.choices[0].message.content)
except Exception as e:
    print(f"Критическая ошибка: {e}")