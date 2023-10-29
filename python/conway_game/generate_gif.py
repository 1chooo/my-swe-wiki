"""
@version: 1.0.0
@author: @fefei69, @1chooo
@date: 2023/05/18
@brief create animated GIF

`generate_gif.py`
"""

from PIL import Image

image_frames = [Image.open('./imgs/frame_{}.png'.format(frame+1)) for frame in range(10)]
image_frames[0].save(
    './imgs/conway.gif', 
    format='gif', 
    append_images=image_frames[1:], 
    save_all=True, 
    duration=150, 
    loop=0
)