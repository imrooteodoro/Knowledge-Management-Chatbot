from pydantic import BaseModel

class TextMessage(BaseModel):
    user_message: str