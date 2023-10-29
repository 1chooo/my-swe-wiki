"""
@version: 1.0.0
@author: @fefei69, @1chooo
@date: 2023/05/18
@brief Python code to implement Conway's Game Of Life

`Conway.py`
"""

import numpy as np
import matplotlib.pyplot as plt

# setting up the values for the grid
ON = 255
OFF = 0
vals = [ON, OFF]


def randomGrid(N: int) -> np.ndarray:

    """
    Returns a grid of NxN random values.

    Parameters:
        N (int): The size of the grid.

    Returns:
        np.ndarray: The generated grid of random values.
    """

    return np.random.choice(vals, N*N, p=[0.2, 0.8]).reshape(N, N)

def addGlider(i: int, j: int, grid: np.ndarray) -> None:

    """
    Adds a glider with the top left cell at position (i, j) in the grid.

    Parameters:
        i (int): The row index of the top left cell.
        j (int): The column index of the top left cell.
        grid (np.ndarray): The grid to add the glider to.

    Returns:
        None
    """

    glider = np.array(
        [[  0,   0, 255],
         [255,   0, 255],
         [  0, 255, 255],]
    )
    grid[i:i+3, j:j+3] = glider


def addGosperGliderGun(i: int, j: int, grid: np.ndarray) -> None:

    """
    Adds a Gosper Glider Gun with the top left cell at position (i, j) in the grid.

    Parameters:
        i (int): The row index of the top left cell.
        j (int): The column index of the top left cell.
        grid (np.ndarray): The grid to add the Gosper Glider Gun to.

    Returns:
        None
    """
    
    gun = np.zeros((46+35)*38).reshape((46+35), 38)
	#gun = np.zeros(11*38).reshape(11, 38)

    gun[5][1] = gun[5][2] = 255
    gun[6][1] = gun[6][2] = 255

    gun[3][13] = gun[3][14] = 255
    gun[4][12] = gun[4][16] = 255
    gun[5][11] = gun[5][17] = 255
    gun[6][11] = gun[6][15] = gun[6][17] = gun[6][18] = 255
    gun[7][11] = gun[7][17] = 255
    gun[8][12] = gun[8][16] = 255
    gun[9][13] = gun[9][14] = 255

    gun[1][25] = 255
    gun[2][23] = gun[2][25] = 255
    gun[3][21] = gun[3][22] = 255
    gun[4][21] = gun[4][22] = 255
    gun[5][21] = gun[5][22] = 255
    gun[6][23] = gun[6][25] = 255
    gun[7][25] = 255

    gun[3][35] = gun[3][36] = 255
    gun[4][35] = gun[4][36] = 255
    ##############################
    gun[5+35][1] = gun[5+35][2] = 255
    gun[5+35][1] = gun[5+35][2] = 255
    gun[6+35][1] = gun[6+35][2] = 255

    gun[3+35][13] = gun[3+35][14] = 255
    gun[4+35][12] = gun[4+35][16] = 255
    gun[5+35][11] = gun[5+35][17] = 255
    gun[6+35][11] = gun[6+35][15] = gun[6+35][17] = gun[6+35][18] = 255
    gun[7+35][11] = gun[7+35][17] = 255
    gun[8+35][12] = gun[8+35][16] = 255
    gun[9+35][13] = gun[9+35][14] = 255

    gun[1+35][25] = 255
    gun[2+35][23] = gun[2+35][25] = 255
    gun[3+35][21] = gun[3+35][22] = 255
    gun[4+35][21] = gun[4+35][22] = 255
    gun[5+35][21] = gun[5+35][22] = 255
    gun[6+35][23] = gun[6+35][25] = 255
    gun[7+35][25] = 255

    gun[3+35][35] = gun[3+35][36] = 255
    gun[4+35][35] = gun[4+35][36] = 255


    gun[5+70][1] = gun[5+70][2] = 255
    gun[6+70][1] = gun[6+70][2] = 255

    gun[3+70][13] = gun[3+70][14] = 255
    gun[4+70][12] = gun[4+70][16] = 255
    gun[5+70][11] = gun[5+70][17] = 255
    gun[6+70][11] = gun[6+70][15] = gun[6+70][17] = gun[6+70][18] = 255
    gun[7+70][11] = gun[7+70][17] = 255
    gun[8+70][12] = gun[8+70][16] = 255
    gun[9+70][13] = gun[9+70][14] = 255

    gun[1+70][25] = 255
    gun[2+70][23] = gun[2+70][25] = 255
    gun[3+70][21] = gun[3+70][22] = 255
    gun[4+70][21] = gun[4+70][22] = 255
    gun[5+70][21] = gun[5+70][22] = 255
    gun[6+70][23] = gun[6+70][25] = 255
    gun[7+70][25] = 255

    gun[3+70][35] = gun[3+70][36] = 255
    gun[4+70][35] = gun[4+70][36] = 255

    grid[i:i+46+35, j:j+38] = gun
    
def update(frameNum: int, img, grid: np.ndarray, N: int) -> tuple:
    """
    Update the grid based on Conway's Game of Life rules 
    and return the updated image and grid.

    Parameters:
        frameNum (int): The current frame number.
        img (matplotlib.image.AxesImage): The image to be updated.
        grid (np.ndarray): The grid representing the current state of the Game of Life.
        N (int): The size of the grid.

    Returns:
        tuple: A tuple containing the updated image and grid.
    """

    # Copy grid since we require 8 neighbors for calculation and we go line by line
    newGrid = grid.copy()

    for i in range(N):
        for j in range(N):
            # Compute 8-neighbor sum using toroidal boundary conditions
            total = int((
                	grid[i, (j-1)%N] + grid[i, (j+1)%N] +
                 	grid[(i-1)%N, j] + grid[(i+1)%N, j] +
                 	grid[(i-1)%N, (j-1)%N] + grid[(i-1)%N, (j+1)%N] +
                 	grid[(i+1)%N, (j-1)%N] + grid[(i+1)%N, (j+1)%N]
                ) / 255
            )

            # Apply Conway's rules
            # Cell is alive
            if grid[i, j] == ON:
                if total < 2 or total > 3:
                    newGrid[i, j] = OFF
            # Cell is dead
            else:
                # Reproduction
                if total == 3:
                    newGrid[i, j] = ON

    # Update data
    img.set_data(newGrid)
    grid[:] = newGrid[:]

    plt.savefig(f'./imgs/frame_{frameNum}.png')

    return img, grid