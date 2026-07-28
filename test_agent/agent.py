# основной код агента

import os
import json
from gigachat import GigaChat
from gigachat.exceptions import ResponseError

from dotenv import load_dotenv
from tools import get_weather

# Загружаем API ключ из .env файла
load_dotenv()


# giga = GigaChat(
#     credentials=os.getenv("GIGACHAT_CREDENTIALS"),
#     base_url="https://ngw.devices.sberbank.ru:9443/api/v2/oauth", # Внутренний адрес вместо внешнего
#     scope="GIGACHAT_API_B2B" # Явное указание скоупа (так как ключ закодирован под B2B)
# )


class SimpleGigaAgent:
    def __init__(self):

        self.client = GigaChat(
            credentials=os.getenv("GIGACHAT_CREDENTIALS"),
            scope="GIGACHAT_API_B2B", # Явное указание скоупа (так как ключ закодирован под B2B)
            model="GigaChat-2",
            ca_bundle_file="/usr/local/share/ca-certificates/russian-trusted/russian_trusted_root_ca_pem.crt"
        )

        self.tools = [
            {
                "type": "function",
                "function": {
                    "name": "get_weather",
                    "description": "Получить текущую погоду в указанном городе",
                    "parameters": {
                        "type": "object",
                        "properties": {
                            "city": {
                                "type": "string",
                                "description": "Название города, например: Москва, Лондон"
                            }
                        },
                        "required": ["city"]
                    }
                }
            }
        ]

        # Соответствие названий функций реальным Python-функциям
        self.available_functions = {
            "get_weather": get_weather
        }


    def run(self, user_message: str) ->str:
        """
        Основной цикл агента: Подумал -> Сделал -> Посмотрел -> Ответил
        """

        print(f"\n👤 Пользователь: {user_message}")


        # Шаг 1: Отправляем запрос в "мозг" (LLM)
        messages = [
            {"role": "system", "content": "Ты полезный помощник. Используй инструменты когда это необходимо."},
            {"role": "user", "content": user_message}
        ]


        payload = {
            "messages": messages,
            "tools": self.tools,
            "tool_choice": "auto"
        }

        try:
            response = self.client.chat(payload=payload)
            print(response)
        
        except ResponseError as e:
            print(f"Ошибка {e.status_code}: {e.content}")
            
            
            response_message = response.choices[0].message


        # Шаг 3: Проверяем, хочет ли модель вызвать инструмент
        if response_message.tool_calls:
            print("🧠 Агент решил: Нужно использовать инструмент!")

            # Шаг 4: Выполняем инструмент
            tool_call = response_message.tool_calls[0]
            function_name = tool_call.function.name
            function_args = json.loads(tool_call.function.arguments)

            print(f"🔧 Вызываю инструмент: {function_name}({function_args})")


            # Получаем реальную функцию и выполняем её
            function_to_call = self.available_functions[function_name]
            function_result = function_to_call(**function_args)

            print(f"📊 Результат инструмента: {function_result}")


            # Шаг 5: Отправляем результат инструмента обратно в модель для финального ответа
            messages.append(response_message.model_dump())  # Добавляем сообщение о вызове инструмента
            messages.append({
                "role": "tool",
                "tool_call_id": tool_call.id,
                "content": function_result
            })


            # Шаг 4: Отправляем второй запрос с результатом инструмента
            second_payload = {
                "messages": messages
                # tools больше не передаем - модель уже приняла решение
            }


            # Шаг 6: Модель генерирует финальный ответ на основе результата инструмента
            second_response = self.client.chat(payload=second_payload)


            final_answer = second_response.choices[0].message.content
            print(f"🤖 Агент: {final_answer}")
            return final_answer
        

        else:
            # Если инструмент не нужен, просто отвечаем сразу
            print("🧠 Агент решил: Отвечу сразу, без инструментов!")
            print(f"🤖 Агент: {response_message.content}")
            return response_message.content



# Шаг 7: Запускаем агента
if __name__ == "__main__":
    agent = SimpleGigaAgent()
    
    print("🤖 Агент запущен!")
    print("-" * 50)
    
    # Тестируем различные запросы
    agent.run("Привет! Как дела?")
    print("-" * 50)
    
    agent.run("Какая погода в Москве?")
    print("-" * 50)
    
    agent.run("А в Лондоне?")