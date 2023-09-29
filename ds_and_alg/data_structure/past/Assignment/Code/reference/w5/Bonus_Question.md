# Singly Linked List(Booooooonus)

## Description

加分題

實作Singly Linked List的功能包含:addBack、addFront、addIndex、deleteIndex

addBack: 從linked list 尾端加入元素

addFront:從linked list 頭端加入元素

addIndex: 指定linked list 的索引值做插入元素，起始位置從0開始

deleteIndex:指定linked list 的索引值做刪除元素


## Input

addBack: 從linked list 尾端加入元素

addFront:從linked list 頭端加入元素

addIndex: 指定linked list 的索引值做插入元素，起始位置從0開始，先輸入索引值再輸入元素值

deleteIndex:指定linked list 的索引值做刪除元素



範例:

addBack 0addBack 1 addBack 2

結果:

0 --> 1 --> 2->null



範例:

addBack 0addBack 1addIndex1 2

結果:

0-->2-->1-->null



範例:

addBack 1deleteIndex 0

結果:

null



範例:

addBack 0 addIndex 1 1 exit

結果:

0-->1-->null



最後輸入exit來結束程式執行


## Output

1.印出操作的結果，linked list 最後會指向null。

2.若linked list為空，則會印出null。

## Sample Input 1

addBack 0 addBack 1 addBack 2 exit

## Sample Output 1

0-->1-->2-->null

## Sample Input 2

addBack 1 addBack 2 deleteIndex 1 exit

## Sample Output 2

1-->null

## 我做的測資 Input

deleteIndex 1 addIndex 1 1 addBack 0 deleteIndex 1 deleteIndex 0 addFront 0 deleteIndex 0 addIndex 0 0 addFront 1 addIndex 2 1 exit

## 我做的測資 Output

1-->0-->1-->null