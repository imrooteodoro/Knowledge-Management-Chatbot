from services.agent import agent
from models.message import TextMessage

def routes(app):
    @app.post("/message")
    async def message(user_message: TextMessage):
        return agent(user_message)
