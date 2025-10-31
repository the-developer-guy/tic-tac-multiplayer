import network
import random
import time

network = network.Network("http://localhost:8080/")

network.sendActivityRequest()
while True:
    print("Sent random mark")
    network.sendRandomPlacement()
    time.sleep(1)
    network.getGrid()
    print(network.currentField)
    time.sleep(2)