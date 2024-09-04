import pygame
class Sprite(pygame.sprite.Sprite):
    def __init__(self, image, groups=None):
        pygame.sprite.Sprite.__init__(self, groups)

        self.image = pygame.image.load(image).convert_alpha()

    def get_image(self, row: int, frame: int, width: int, height: int, scale: float):
        image = pygame.Surface((width, height)).convert_alpha()
        image.blit(self.image, (0, 0), ((frame * width), (row * height), width, height))
        image = pygame.transform.scale(image, (width * scale, height * scale))
        image.set_colorkey((0, 0, 0))
        return image