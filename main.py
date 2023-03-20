from fastapi import FastAPI
from gamemap import TicTacToe

app = FastAPI()
arena = TicTacToe()

@app.get("/")
async def root():
    return {"message": "Hello World"}

@app.get("/game/{move}")
async def player_move(move):
    return {"message": f"{move}"}