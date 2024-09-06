import pygame
from animated_explosion import AnimatedExplosion

class Rocket(pygame.sprite.Sprite):
    def __init__(self, surf, pos, enemy_sprites, groups=None):
        pygame.sprite.Sprite.__init__(self, groups)
        self.surf = surf
        self.image = self.get_image(0, 0, 32, 32, 1)
        self.rect = self.image.get_frect(midbottom=pos)
        self.speed = 500
        self.group = groups
        self.frames = [self.get_image(0, frame, 32, 32, 1) for frame in range(3)]
        self.frame_index = 0

        self.enemy_sprites = enemy_sprites

        # Mask
        self.mask = pygame.mask.from_surface(self.image)

    def get_image(self, row: int, frame: int, width: int, height: int, scale: float):
        image = pygame.Surface((width, height)).convert_alpha()
        image.blit(self.surf, (0, 0), ((frame * width), (row * height), width, height))
        image = pygame.transform.scale(image, (width * scale, height * scale))
        image.set_colorkey((0, 0, 0))
        return image

    def update(self, dt):
        self.rect.centery -= 400 * dt

        # If the rocket flies off the screen, kill it
        if self.rect.bottom <= 0:
            self.kill()

        # Animation
        self.frame_index += 1 * dt
        self.image = self.frames[int(self.frame_index % len(self.frames))]

        collided = pygame.sprite.spritecollide(self, self.enemy_sprites, True, pygame.sprite.collide_mask)
        if collided:
            AnimatedExplosion(self.frames, self.rect.center, self.group) 
            self.kill()