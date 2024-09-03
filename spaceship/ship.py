import pygame

class Ship(pygame.sprite.Sprite):
    def __init__(self):
        pygame.sprite.Sprite.__init__(self)

        self.image = pygame.image.load('assets/Main Ship/Main Ship - Bases/PNGs/Main Ship - Base - Damaged.png')

        # Fetch the rectangle object that has the dimensions of the image
        # Can update the position of this object by setting the values of rect.x and rect.y
        self.rect = self.image.get_rect()

        self.speed = 10

    def update(self):
        pass

    def move_left(self):
        self.rect.x -= self.speed

    def move_right(self):
        self.rect.x += self.speed
        
    def draw(self, screen):
        screen.blit(self.image, self.rect)