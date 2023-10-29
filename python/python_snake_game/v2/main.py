# -*- coding: utf-8 -*-

import pygame
from SnakeGame import SnakeGame

if __name__ == '__main__':
    game = SnakeGame()

    # game loop
    while True:
        gameOver, score = game.play_step()

        if gameOver == True:
            break

    print('Final Score', score)

    pygame.quit()