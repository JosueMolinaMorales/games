import pygame
from settings import *
from player import Player


class Game():
    def __init__(self):
        # Setup
        pygame.init()
        self._running = True
        self._display_surf = pygame.display.set_mode(
            (WINDOW_WIDTH, WINDOW_HEIGHT))
        self._clock = pygame.time.Clock()
        self._dt = 0
        pygame.display.set_caption("Vampire Survivor")

        # Groups
        self._all_sprites = pygame.sprite.Group()

        # Sprites
        self._player = Player(self._all_sprites)

    def run(self):
        while self._running:
            self._dt = self._clock.tick() / 1000.0
            # Event handling
            for event in pygame.event.get():
                if event.type == pygame.QUIT:
                    self._running = False
            # Update
            self._all_sprites.update(self._dt)

            # Draw
            self._display_surf.fill((0, 0, 0))

            self._all_sprites.draw(self._display_surf)

            pygame.display.flip()

        pygame.quit()


if __name__ == "__main__":
    Game().run()
