'''
Player module contains the Player class which is responsible for handling the player's movement.
'''
from settings import WINDOW_HEIGHT, WINDOW_WIDTH, join
import pygame


class Player(pygame.sprite.Sprite):
    '''
        Player Sprite
    '''

    def __init__(self, group):
        super().__init__(group)
        path = join("assets", "images", "player", "down", "0.png")
        self.image = pygame.image.load(path).convert_alpha()
        self.rect = self.image.get_frect(
            center=(WINDOW_WIDTH // 2, WINDOW_HEIGHT // 2))
        self.direction = pygame.Vector2(0, 0)
        self.speed = 200

    def input(self):
        '''
        Handles the player's input
        '''
        keys = pygame.key.get_pressed()

        self.direction.x = int(keys[pygame.K_d]) - int(keys[pygame.K_a])
        self.direction.y = int(keys[pygame.K_s]) - int(keys[pygame.K_w])
        # normalize the direction vector so that moving diagonally isn't faster
        self.direction = self.direction.normalize() if self.direction else self.direction

    def move(self, dt):
        '''
        Handles the player's movement
        '''
        self.rect.center += self.direction * self.speed * dt

    def update(self, dt):
        '''
        Updates the player's position
        '''
        self.input()
        self.move(dt)
