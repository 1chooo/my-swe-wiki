"""
@version: 1.0.0
@author: @fefei69, @1chooo
@date: 2023/05/18
@brief Python code to implement Conway's Game Of Life

`main.py`
"""

import argparse
import numpy as np
import matplotlib.pyplot as plt
import matplotlib.animation as animation
import Conway as conway


def main() -> None:

    """
    Runs Conway's Game of Life simulation.
    
    Returns:
        None
    """

    # Command line args are in sys.argv[1], sys.argv[2] ..
    # sys.argv[0] is the script name itself and can be ignored
    # parse arguments
    parser = argparse.ArgumentParser(
        description="Runs Conway's Game of Life simulation.",
    )

    # add arguments
    parser.add_argument('--grid-size', dest='N', required=False)
    parser.add_argument('--mov-file', dest='movfile', required=False)
    parser.add_argument('--interval', dest='interval', required=False)
    parser.add_argument('--glider', action='store_true', required=False)
    parser.add_argument('--gosper', action='store_true', required=False)
    args = parser.parse_args()
    
    # set grid size
    N = 100

    if args.N and int(args.N) > 8:
        N = int(args.N)
        
    # set animation update interval
    updateInterval = 50
    if args.interval:
        updateInterval = int(args.interval)

    # declare grid
    grid = np.array([])

    # check if "glider" demo flag is specified
    if args.glider:
        grid = np.zeros(N*N).reshape(N, N)
        conway.addGlider(1, 1, grid)
    elif args.gosper:
        grid = np.zeros(N*N).reshape(N, N)
        conway.addGosperGliderGun(0, 0, grid)
    else:
        # populate grid with random on/off -
        # more off than on
        grid = conway.randomGrid(N)

    # set up animation
    fig, ax = plt.subplots()
    img = ax.imshow(grid, interpolation='nearest')
    ani = animation.FuncAnimation(
        fig, 
        conway.update, 
        fargs=(img, grid, N),
        frames=10,
        interval=updateInterval,
        save_count=50,
    )

    # # of frames?
    # set output file
    if args.movfile:
        ani.save(
            args.movfile, 
            fps=30, 
            extra_args=['-vcodec', 'libx264'],
        )
        
    plt.title('Conway\'s Game Of Life')

    plt.show()

if __name__ == '__main__':
	main()