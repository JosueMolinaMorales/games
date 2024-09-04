import pygame
from utils import WINDOW_HEIGHT, WINDOW_WIDTH
class Ship(pygame.sprite.Sprite):
    def __init__(self):
        pygame.sprite.Sprite.__init__(self)

        
        self.player_surface = pygame.image.load('assets/Foozle_2DS0011_Void_MainShip/Main Ship/Main Ship - Bases/PNGs/Main Ship - Base - Full health.png').convert_alpha()

        # Fetch the rectangle object that has the dimensions of the image
        # Can update the position of this object by setting the values of rect.x and rect.y
        self.rect = self.player_surface.get_frect(center=(WINDOW_WIDTH // 2, WINDOW_HEIGHT // 2))
    
        self.speed = 2
        self.direction = pygame.math.Vector2(5, -1)

    def update(self):
        pass

    def move_left(self):
        self.rect.x -= self.speed

    def move_right(self):
        self.rect.x += self.speed
        
    def draw(self, screen):
        self.rect.center += self.direction * self.speed
        screen.blit(self.player_surface, self.rect)