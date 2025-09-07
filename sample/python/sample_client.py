import pygame
import network
from gamestates import GameState
import threading


pygame.init()
screen = pygame.display.set_mode((400,400))
clock = pygame.time.Clock()
running = True
BACKGROUND = (255, 255, 255)
pygame.display.set_caption("TicTacToe")

current_state = GameState.LOBBY

network = network.Network("http://localhost:8080/")
thread = threading.Thread(target=network.InQueue, daemon=True)
thread.start()

 

while running:
    screen.fill(BACKGROUND)

    for event in pygame.event.get():
        if event.type == pygame.QUIT:
            running = False

    if network.lobbyid != "":
        print(network.lobbyid)
        print(thread.is_alive())
    
    pygame.display.flip()
    clock.tick(60)

pygame.quit()


