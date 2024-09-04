import pygame
from utils import WINDOW_HEIGHT, WINDOW_WIDTH
class Ship(pygame.sprite.Sprite):
    def __init__(self, groups):
        pygame.sprite.Sprite.__init__(self, groups)
        self.image = pygame.image.load('assets/Foozle_2DS0011_Void_MainShip/Main Ship/Main Ship - Bases/PNGs/Main Ship - Base - Full health.png').convert_alpha()
        # Fetch the rectangle object that has the dimensions of the image
        # Can update the position of this object by setting the values of rect.x and rect.y
        self.rect = self.image.get_frect(center=(WINDOW_WIDTH // 2, WINDOW_HEIGHT // 2))
    
        self.speed = 100
        self.direction = pygame.math.Vector2(0, 0)

    def update(self, dt):
        keys = pygame.key.get_pressed()

        self.direction.x = int(keys[pygame.K_RIGHT]) - int(keys[pygame.K_LEFT])
        self.direction.y = int(keys[pygame.K_DOWN]) - int(keys[pygame.K_UP])
        # normalize the direction vector so that moving diagonally isn't faster
        self.direction = self.direction.normalize() if self.direction else self.direction
        self.rect.center += self.direction * self.speed * dt

        recent = pygame.key.get_just_pressed()
        if recent[pygame.K_SPACE]:
            print('Pew pew!')


    def move_left(self):
        self.rect.x -= self.speed

    def move_right(self):
        self.rect.x += self.speed
        
    def draw(self, screen, dt):
        screen.blit(self.image, self.rect)