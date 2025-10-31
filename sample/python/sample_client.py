import random
import time
from game_connection import GameConnection

connection = GameConnection("http://localhost:8080/")

connection.sendActivityRequest()
while True:
    print("Sent random mark")
    connection.sendRandomPlacement()
    time.sleep(1)
    connection.getGrid()
    print(connection.currentField)
    time.sleep(2)