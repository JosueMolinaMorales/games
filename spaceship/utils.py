import pygame

WINDOW_WIDTH = 1024
WINDOW_HEIGHT = 1024


def strip_from_sheet(spritesheet: str, rows: int, cols: int, scale: float):
    sheet = pygame.image.load(spritesheet)
    rect = sheet.get_rect()
    img_width, img_height = rect.width / cols, rect.height / rows
    print(img_width, img_height) 
    frames = []
    for row in range(rows):
        for col in range(cols):
            rect = pygame.Rect(col*img_width, row*img_height, img_width, img_height)
            frames.append(sheet.subsurface(rect))
            
    return frames

EXPLOSION_FRAMES = strip_from_sheet('assets/Foozle_2DS0014_Void_EnemyFleet_3/Nautolan/Destruction/PNGs/Nautolan Ship - Fighter.png', 1, 9, 1)