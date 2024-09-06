import pygame
from utils import EXPLOSION_FRAMES

class AnimatedExplosion(pygame.sprite.Sprite):
    def __init__(self, frames, pos, groups):
        super().__init__(groups)
        self.frames = EXPLOSION_FRAMES
        self.frames_index = 0

        self.image = self.frames[self.frames_index]
        self.rect = self.image.get_frect(center=pos) 

    def update(self, dt):
        self.frames_index += 10 * dt
        self.image = self.frames[int(self.frames_index % len(self.frames))]

        if self.frames_index >= len(self.frames):
            self.kill()