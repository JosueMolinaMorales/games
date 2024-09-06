import os

WINDOW_WIDTH, WINDOW_HEIGHT = 1280, 720
TILE_SIZE = 64

def join(*args):
    return os.path.join(os.path.dirname(__file__), *args)