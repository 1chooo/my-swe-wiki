# Priority Queue

## 描述

小明最近期中考週好多事要做

他列出了一個代辦事項清單

請幫小明按照事情的輕重緩急找出他最先該做的三件事

## 輸入

首先輸入事項的總數

接著依序輸入事情名稱與其優先度(越大代表這件事越急)

## 輸出

按照優先度依序輸出最先該做的三件事

## Sample Input 1

5<br>
write_DS_homework 100<br>
watch_video 20<br>
dump_the_garbage 80<br>
play_video_game 15<br>
prepare_the_final_exam 50

## Sample Output 1

First three things to do:<br>
write_DS_homework<br>
dump_the_garbage<br>
prepare_the_final_exam<br>

## 提示

1. 測資很大請不要寫死，建議使用動態配置
2. 請實作出 Max Heap 結構，可以用 array 或 linked list 來實作
3. 請用 Max Heap 幫你找到最急的三件事
