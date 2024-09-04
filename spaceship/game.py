import pygame
from ship import Ship
from utils import WINDOW_HEIGHT, WINDOW_WIDTH
class Game:
    def __init__(self):
        # Initialize the game
        pygame.init()
        self._display_surface = pygame.display.set_mode((WINDOW_WIDTH, WINDOW_HEIGHT))
        self._background = pygame.image.load('assets/Blue_Nebula_01-1024x1024.png').convert()
        pygame.display.set_caption("Space Invaders")
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
        self._display_surface.blit(self._background, (0, 0))
        self.player.draw(self._display_surface)
        pygame.display.update()

    def run(self):
        # Run the game's main loop
        while self._running:
            self._clock.tick(60)
            self.update()
            self.draw()

        pygame.quit()