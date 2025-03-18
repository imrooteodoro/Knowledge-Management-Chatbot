from mistralai import Mistral
from dotenv import load_dotenv
import os

def agent(user_message):
    load_dotenv()
    api_key = os.getenv("MISTRAL_API_KEY")
    client = Mistral(api_key=api_key)

    chat_response = client.agents.complete(
        agent_id=os.getenv("AGENT_ID"),
        messages=[
            {
                "role": "user",
                "content": f"{user_message}",
            },
        ],
    )
    return {
        "bot_response": chat_response.choices[0].message.content
    }

