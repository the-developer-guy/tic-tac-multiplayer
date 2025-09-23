import network
import random

network = network.Network("http://localhost:8080/")
board = [['' for _ in range(3)] for _ in range(3)]

def move(mark, x,y):
    board[y][x] = mark


move("X", 2,0)
print(board)

network.InQueue()
print(network.lobbyid)



