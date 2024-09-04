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
        self._dt = 0
        self._all_sprites = pygame.sprite.Group()
        self.player = Ship(self._all_sprites)

    def run(self):
        # Run the game's main loop
        while self._running:
            self.dt = self._clock.tick() / 1000.0

            # Update
            for event in pygame.event.get():
                if event.type == pygame.QUIT:
                    self._running = False
        
            self._all_sprites.update(self.dt)

            # Draw   
            self._display_surface.blit(self._background, (0, 0))
            self._all_sprites.draw(self._display_surface) 

            pygame.display.update()

        pygame.quit()