WINDOW_WIDTH = 1024
WINDOW_HEIGHT = 1024


def load_png(name):
    """Load image and return image object"""
    fullname = os.path.join('data', name)
    try:
        image = pygame.image.load(fullname)
    except pygame.error as message:
        print('Cannot load image:', name)
        raise SystemExit(message)
    return image