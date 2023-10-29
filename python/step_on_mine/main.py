# 踩地雷 Minesweeper
# 輸入座標: 大寫英文字母+數字 (不用空格)
# 插旗子:  在座標後面直接加上一個* (也不用空格)

size = 10 # 因為座標使用英文，26以上會有bug(・∀・)
mine = 10

from random import randint

def Print():
    l = len(str(size))
    print(" " * (l+1) + ABC[:size])
    for i in range(size):
        print(str(i+1).rjust(l), end = "|")
        for j in range(size):
            print(display[i][j], end = "")
        print()
def Pos(s):
    global flag
    try:
        if s[-1] == "*":
            flag = True
            [x, y] = Pos(s[:-1])
            if display[y][x] == "-" or "*":
                return [x, y]
            else: return Err()
        for x in range(size):
            if ABC[x] == s[0]: break
        else: return Err()
        y = int(s[1:]) - 1
        if y < 0 or y >= size: return Err()
        if display[y][x] != "-" and display[y][x] != "*": return Err()
    except: return Err()
    return [x, y]
def Open(pos, flag):
    [x, y] = pos
    if flag:
        if display[y][x] == "-": display[y][x] = "*"
        else: display[y][x] = "-"
        return False
    if game[y][x] == 1:
        for i, item in enumerate(game):
            for j, jtem in enumerate(item):
                if jtem == 1:
                    if display[i][j] != "*": display[i][j] = "#"
                else: C2D([j, i])
        Print()
        print("GG!")
        return True
    else:
        if count([x, y]) == 0: spread([x, y])
        else: display[y][x] = count([x, y])
        print("SAFE!")
        return False
def spread(pos):
    [x, y] = pos
    if display[y][x] != "-": return
    C2D(pos)
    if count([x, y]) != 0:
        return
    if x >= 1:
        try: spread([x-1, y])
        except: pass
    if y >= 1:
        try: spread([x, y-1])
        except: pass
    try: spread([x+1, y])
    except: pass
    try: spread([x, y+1])
    except: pass
def count(pos):
    [x, y] = pos
    c = 0
    for i in range(-1, 2):
        for j in range(-1, 2):
            if y+j >= 0 and x+i >= 0:
                try:
                    c += game[y+j][x+i]
                except: pass
    return c
def Err():
    global flag
    flag = False
    print("Error!")
    return Pos(input("pos:"))
def C2D(pos):
    [x, y] = pos
    display[y][x] = " " if count([x, y]) == 0 else count([x, y])

ABC = "ABCDEFGHIJKMNOPQRSTUVWXYZ"

game = []
for i in range(size):
    game.append([0] * size)

b = 0
while b < mine:
    x, y = randint(0, size - 1), randint(0, size - 1)
    if game[y][x] == 0:
        b += 1
        game[y][x] = 1
for i in game:
    for j in i:
        if j == 1:
            print("#", end = "")
        else:
            print(" ", end = "")
    print("|")

display = []
for i in range(size):
    display.append(["-"] * size)

while True:
    Print()
    flag = False
    pos = Pos(input("pos:"))
    if Open(pos, flag): break
    for i in display:
        if "-" in i: break
    else:
        print("WIN!!!")
        break