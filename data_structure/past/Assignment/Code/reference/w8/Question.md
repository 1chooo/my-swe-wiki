# Binary Search Tree

## 描述

小明放學途中，走到一個迷宮，這個迷宮的形狀是一顆樹的結構

請幫幫小明找到正確的路徑


## 輸入

建立一棵binary search tree

第一行:節點個數，節點個數 > 0

第二行:節點值，節點值 > 0

第三行:目標值，目標值 > 0

ex:

5

3 1 2 6 5

建樹是根據輸入順序，去建立binary search tree

結果如下圖所示:

![image](https://github.com/Mes0903/NCU_DS/blob/main/w8/tree1.png?raw=true)

## 輸出

找到是否存在至少一條路徑，其節點值的總和等於目標值，且路徑必須從root走到leaf

若有，印出:There exit at least one path in binary"空白"search"空白"tree.

否則，印出:There have no path in binary search tree.

以範例輸入當作例子:

![image](https://github.com/Mes0903/NCU_DS/blob/main/w8/tree2.png?raw=true)

沿著紅色節點走，其總和等於目標值

## Sample Input 1

5
3 1 2 6 5
14

## Sample Output 1

There exit at least one path in binary search tree.

## 提示

輸出以Output Samples為主
