import requests
import json
import time
import random

class GameConnection:
    def __init__(self, address):
        self.token = "asd123"
        self.id = "1"
        self.lobbyid = ""
        self.address = address
        self.x = 0
        self.currentField = [[0,0,0], [0,0,0], [0,0,0]]

    def sendActivityRequest(self):
        url = f"{self.address}ready/{self.id}/{self.token}"
        r = requests.get(url)
        if r.status_code != 200:
            print(r.text)
            raise RuntimeError("Something went wrong during the request") 
            

        jsonResponse = json.loads(r.text)
        self.lobbyid = jsonResponse["lobbyId"]
        
        return jsonResponse
    
    def getGrid(self):
        url = f"{self.address}/getgrid/{self.lobbyid}"
        r = requests.get(url)
        if r.status_code != 200:
            print(r.text)
            raise RuntimeError("Something went wrong during the request") 
            

        jsonResponse = json.loads(r.text)
        self.currentField = jsonResponse["Field"]

        return jsonResponse
    

    def sendRandomPlacement(self):
        for i in range(8):
            randomX = random.randint(0,2)
            randomY = random.randint(0,2)

            if self.currentField[randomY][randomX] != 0:
                continue

            if i == 8:
                print("All space are occupied.")

        url = f"{self.address}/place/{self.lobbyid}/"
        payload = {"token": self.token, "cor_x": randomX, "cor_y": randomY}

        r = requests.post(url, data=payload)
        print(r.status_code)
            