import pygame
from ship import Ship

class Game:
    def __init__(self):
        # Initialize the game
        pygame.init()
        self._screen = pygame.display.set_mode((1280, 720))
        self._clock = pygame.time.Clock()
        self._running = True
        self.player = Ship()

    def update(self):
        for event in pygame.event.get():
            if event.type == pygame.QUIT:
                self._running = False
            if event.type == pygame.KEYDOWN:
                if event.key == pygame.K_LEFT:
                    self.player.move_left()
                if event.key == pygame.K_RIGHT:
                    self.player.move_right()            

    def draw(self):
        self._screen.fill((0, 0, 0))
        self.player.draw(self._screen)
        pygame.display.flip()
        self._clock.tick(60)

    def run(self):
        # Run the game's main loop
        while self._running:
            self.update()
            self.draw()

        pygame.quit()