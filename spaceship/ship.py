import pygame
from utils import WINDOW_HEIGHT, WINDOW_WIDTH
from rocket import Rocket
class Ship(pygame.sprite.Sprite):
    def __init__(self, groups):
        pygame.sprite.Sprite.__init__(self, groups)
        self.image = pygame.image.load('assets/Foozle_2DS0011_Void_MainShip/Main Ship/Main Ship - Bases/PNGs/Main Ship - Base - Full health.png').convert_alpha()
        # Fetch the rectangle object that has the dimensions of the image
        # Can update the position of this object by setting the values of rect.x and rect.y
        self.rect = self.image.get_frect(center=(WINDOW_WIDTH // 2, WINDOW_HEIGHT // 2))
        self.speed = 100
        self.direction = pygame.math.Vector2(0, 0)
        self._all_sprites = groups

        # Rockets
        self._rocket_surf = pygame.image.load('assets/Foozle_2DS0011_Void_MainShip/Main ship weapons/PNGs/Main ship weapon - Projectile - Rocket.png').convert_alpha()
    
        # Cooldown for shooting
        self._can_shoot = True
        self._laser_shoot_time = 0
        self._cooldown_duration = 400 # in milliseconds

    def update(self, dt):
        keys = pygame.key.get_pressed()

        self.direction.x = int(keys[pygame.K_RIGHT]) - int(keys[pygame.K_LEFT])
        self.direction.y = int(keys[pygame.K_DOWN]) - int(keys[pygame.K_UP])
        # normalize the direction vector so that moving diagonally isn't faster
        self.direction = self.direction.normalize() if self.direction else self.direction
        self.rect.center += self.direction * self.speed * dt

        recent = pygame.key.get_just_pressed()
        if recent[pygame.K_SPACE] and self._can_shoot:
            Rocket(self._rocket_surf, self.rect.midtop, self._all_sprites)
            self._can_shoot = False
            self._laser_shoot_time = pygame.time.get_ticks()
            
        self.laser_time()

    def laser_time(self):
        if not self._can_shoot:
            current_time = pygame.time.get_ticks()
            if current_time - self._laser_shoot_time >= self._cooldown_duration:
                self._can_shoot = True