# pygame-snake

This is the project that build a little python game -> "Snake Game" through the package "pygame". 

However, when we want to do this project, we will face a lot of problems when we build the python enviroment. Therefore, we have tested a lot of times, then we choose "anaconda" to build. So, here is the steps.

## Create enviroment

Build the enviroment through anaconda (Actually, I use miniconda in my mac, but it's still the same steps.)

my Conda --version: conda 4.12.0

![plot](https://miro.medium.com/max/1400/1*1uDdGsxFlihzmPfLP6Mo4A.png)

* ### First steps: 

    Build a virtual enviroment which called pygame and install python 3.6.13 version. Then activate it...

    ``` vim
    $ conda create --name pygame python=3.6.13
    $ conda activate pygame
    ```

* ### Second steps: Install pygame

    ```vim
    $ python3 -m pip install -U pygame --user
    ```

    If omitting `--user`. The package gets installed into the conda environment.

* ### Third steps: Test pygame
    ```vim
    $ python3 -m pygame.examples.aliens
    ```
    ![plot](https://miro.medium.com/max/1400/1*xRYWm3kCCimYkrsHfh88KA.png)

* ### Well done.


## Run Snake Game

```vim
$ python main.py
```

![plot](https://miro.medium.com/max/1400/1*okV2P3qTibFMO1NWbaRBJQ.png)