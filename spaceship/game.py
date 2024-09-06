import pygame
from ship import Ship
from utils import WINDOW_HEIGHT, WINDOW_WIDTH, strip_from_sheet
from enemy import EnemyShip
import random
class Game:
    def __init__(self):
        # Initialize the game
        self._display_surface = pygame.display.set_mode((WINDOW_WIDTH, WINDOW_HEIGHT))
        self._background = pygame.image.load('assets/Blue_Nebula_01-1024x1024.png').convert()
        pygame.display.set_caption("Space Invaders")
        self._clock = pygame.time.Clock()
        self._running = True
        self._dt = 0
        self._all_sprites = pygame.sprite.Group()
        self._enemey_sprites = pygame.sprite.Group()
        self.player = Ship(self._all_sprites, self._enemey_sprites)

        # Custom Events --> Enemy Spawn
        self._enemy_spawn_event = pygame.event.custom_type()
        pygame.time.set_timer(self._enemy_spawn_event, 1000)
        self._enemy_surf = pygame.image.load('assets/Foozle_2DS0014_Void_EnemyFleet_3/Nautolan/Designs - Base/PNGs/Nautolan Ship - Fighter - Base.png').convert_alpha()
        
        self._explosion_surf = pygame.image.load('assets/Foozle_2DS0014_Void_EnemyFleet_3/Nautolan/Destruction/PNGs/Nautolan Ship - Fighter.png').convert_alpha()
        self.frame_idx = 0

    def run(self):
        # Run the game's main loop
        while self._running:
            self.dt = self._clock.tick() / 1000.0

            # Update
            for event in pygame.event.get():
                if event.type == pygame.QUIT:
                    self._running = False
                if event.type == self._enemy_spawn_event:
                    x, y = random.randint(0, WINDOW_WIDTH), -10
                    EnemyShip(self._enemy_surf, (x, y), (self._all_sprites, self._enemey_sprites))
                    
       
            self._all_sprites.update(self.dt)
            

            # Draw
            self._display_surface.blit(self._background, (0, 0))
            self._all_sprites.draw(self._display_surface) 
           
            pygame.display.update()

        pygame.quit()
        
    
