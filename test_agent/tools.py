# инструменты агента

import random


def get_weather(city: str) -> str:
    """
    Инструмент для получения погоды в городе.
    """
    # Имитация реального API погоды (в реальном проекте здесь был бы HTTP-запрос)
    weather_types = ["солнечно", "облачно", "дождливо", "снежно", "ветрено"]
    temperatures = list(range(-10, 35))
    
    weather = random.choice(weather_types)
    temp = random.choice(temperatures)
    
    return f"В городе {city} сейчас {weather}, температура {temp}°C"