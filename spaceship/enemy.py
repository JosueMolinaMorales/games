import pygame
from utils import WINDOW_HEIGHT
class EnemyShip(pygame.sprite.Sprite):
    def __init__(self, surf, pos, groups):
        pygame.sprite.Sprite.__init__(self, groups)
        self.image = surf
        self.rect = self.image.get_frect(midbottom=pos)
        self.speed = 100
        self.direction = pygame.math.Vector2(0, 1)

        # Where the ship will stop moving
        self._y_stop = 100

    def update(self, dt):
        self.rect.centery += self.speed * dt
        # If the enemy flies off the screen, kill it
        if self.rect.top >= WINDOW_HEIGHT:
            self.kill()