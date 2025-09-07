import requests
import json
import time

class Network:
    def __init__(self, address):
        self.token = "MyToken123"
        self.lobbyid = ""
        self.address = address
        self.x = 0

    def sendActivityRequest(self):
        url = f"{self.address}ready/{self.token}/"
        r = requests.get(url)
        if r.status_code != 200:
            raise RuntimeError("Something went wrong during the request") 

        jsonResponse = json.loads(r.text)
        return jsonResponse
    
    def InQueue(self):
        while True:
            r = self.sendActivityRequest()
            if self.x ==3:
                r["lobbyId"] = "vmi"

            if r["lobbyId"] == "":
                print(r)
                self.x+=1
                time.sleep(2)
                continue
            
            self.lobbyid = r["lobbyId"]
            break