# CSP(Communicating Sequential Process)

並行機制有很多像是多線程,CSP,Actor等等.拿多線程來說,就有諸多問題,譬如:死鎖,可擴展性差,共享狀態.就像共享狀態會產生很多問題，它涉及到內存的變化。只有一個進程發生變化沒問題。但如果有多個進程共享和修改相同的數據，這將非常糟糕.為解決這些問題,提出了很多解決方法呢.比如,CSP和ActorModel.

如下是一些編程語言，以及它們相應的並發機制:

Actors Model — Erlang, Scala, Rust
CSP——Go 語言
多線程— Java, C#, C++

## Go 的設計者Rob Pike有一句經典的名言
:::success
Do not communicate by sharing memory; instead, share memory by communicating.
:::


這句話是說“不要使用共享內存通信，而是應該使用通信去共享內存”，Go 語言推薦我們使用通信來進行進程間同步消息。這樣做有三點好處：

* 首先，使用發送消息來同步信息相比於直接使用共享內存和互斥鎖是一種更高級的抽象，使用更高級的抽象能夠為我們在程序設計上提供更好的封裝，讓程序的邏輯更加清晰。
* 其次，消息發送在解耦方面與共享內存相比也有一定優勢，我們可以將線程的職責分成生產者和消費者，並通過消息傳遞的方式將它們解耦，不需要再依賴共享內存。
* 最後，Go 語言選擇消息發送的方式，通過保證同一時間只有一個活躍的線程能夠訪問數據，能夠從設計上天然地避免線程競爭和數據衝突的問題。

## 傳統並發和基於消息傳遞的並發
在多核CPU機器下，為了充分利用計算機的資源，我們需要進行並發編程。

### 1.1 傳統並發模型
多線程編程，就是傳統的並發編程模式。

傳統的多線程編程，使用的是ShreadMemory(共享內存)的方式，來實現的。

有並發的地方就有競爭，傳統多線程的並發模式使用lock(鎖)，condition (條件變量)等同步原語來強制規定了進程的執行順序，而這些同步原語本質上都是在各個線程中使用了鎖來來實現。

除了基於shared memory(共享內存)以外，還有什麼其他的並發模型嗎？

答案是有的，就是基於消息的並發模型

### 1.2 基於消息的並發模型
基於消息傳遞(Message Passing)的並發模型有CSP和Actor

這兩種模型很像，但還是有一些不同的地方

#### Actor模型：
在Actor模型中，有一點類似物件導向模型，世界上所有的東西都被命名為Actor。
單個Actor會擁有一些狀態，比如為名字是book的Actor可能被描述為：

> name : 西遊記
> author : 吳承恩
> price : 99，Actor彼此之間直接發送消息，不需要經過什麼中介，消息是異步發送和處理的

到此為止好像和對象object沒有什麼不同，但是Actor不會給外界提供任何的行為接口.
比如book.getPrice()這是很自然的面向對象的寫法，在Actor是不被允許的。
每一個Actor的屬性絕不對外暴露，想和外界進行通信必鬚髮送message，所以每個Actor自身都有一個郵箱( MailBox)

Actor模型描述了一組為了避免並發編程的常見問題的公理:

* 所有Actor狀態是Actor本地的，外部無法訪問。
* Actor必須只有通過消息傳遞進行通信。
* 一個Actor可以響應消息:推出新Actor,改變其內部狀態,或將消息發送到一個或多個其他參與者。
* Actor可能會堵塞自己,但Actor不應該堵塞它運行的線程


![](https://i.imgur.com/J7yJ502.png)

#### CSP模型：
CSP的是Communicating Sequential Processes( CSP)的縮寫，翻譯成中文是順序通信進程，簡稱CSP

CSP的核心思想是多個線程之間通過Channel來通信，對應到**golang中的chan結構**，對應到Python中是Queue

##### Go語言的CSP模型是由協程Goroutine與通道Channel實現：

* Go協程goroutine: 是一種輕量線程，它不是操作系統的線程，而是將一個操作系統線程分段使用，通過調度器實現協作式調度。是一種綠色線程，微線程，它與Coroutine協程也有區別，能夠在發現堵塞後啟動新的微線程。
* 通道channel: 類似Unix的Pipe，用於協程之間通訊和同步。協程之間雖然解耦，但是它們和Channel有著耦合

![](https://i.imgur.com/dDDnfKR.png)




##### go語言中的unbuffered Channel :

當在程序代碼內丟了一個值到channel，這時候main 函數就需要等到一個channel 值被讀出來才會結束.
![](https://i.imgur.com/qyQi4TI.png)
##### BufferedChannel:

buffered channel 就是只要有容量，你都可以塞值進去，但是不用讀出來沒關係
![](https://i.imgur.com/2F5bzus.png)








## Actor vs CSP模型

### Actor
* 優點

    * 消息傳輸和封裝，多個Actor 可以同時運行，但不共享狀態，而且單個actor 中的事件是串行執行（這歸功於隊列）
    * Actor 模型支持共享內存模型，也支持分佈式內存模型

* 缺點

   * 儘管Actor 模型比使用線程和鎖模型的程序更易debug，但是也會存在死鎖的問題，而且還需要擔心綁定進程的隊列溢出的問題
    * 沒有對並行提供直接支持，需要通過並發的技術來構造並行方案


### CSP

* 優點

    * 與Actor 相比，CSP 最大的優點是靈活性。Actor 模型，負責通信的媒介和執行單元是耦合的。而CSP 中，channel 是第一類對象，可以被獨立創造、寫入、讀出數據，也可以在不同執行單元中傳遞。


* 缺點

    * CSP 模型也易受死鎖影響，且沒有提供直接的並行支持。並行需要建立在並發基礎上，引入了不確定性。


區別Actor 模型重在參與交流的實體(即進程)，而CSP 重在交流的通道，如Go 中的channel
CSP 模型不關注發送消息的進程，而是關注發送消息時使用的channel，而channel 不像Actor 模型那樣進程與隊列緊耦合。而是可以單獨創建和讀寫，並在進程(goroutine) 之間傳遞。



## 主要的不同點
### 關於消息發送方和接收方
* Actor：注重的處理單元，也就是Actor，而不是消息傳送方式。發送消息時，都需要知道對方是誰。

    這裡的“都需要知道對方是誰”的意思，當ActorX要給ActorY發消息時，必須明確知道ActorY的地址。ActorY接收到消息時，就能夠知道消息發送者（ActorX）的地址。返回消息給發送者時，只需要按發送者的地址往回傳消息就行。

* CSP：注重的是消息傳送方式（channel），不關心發送的人和接收的人是誰。

    向channel寫消息的人，不知道消息的接收者是誰；讀消息的人，也不知道消息的寫入者是誰。


兩者比較看來，CSP把發送方和接收方給解耦了，但這種解耦帶的好處是什麼呢？

### 消息傳輸方式
* Actor：每一對Actor之間，都有一個“ MailBox”來進行收發消息。消息的收發是異步的。
* CSP：使用定義的 channel 進行收發消息。消息的收發是同步的（也可以做成異步的，但是一個有限異步）

Actor 模式消息傳輸，只有一個通道（MailBox），所以無論什麼“類型”的消息都可能發過來，所以要做好模式配置。而 CSP 中的通道（channel）類型是定好的，而且兩個對象可以使用多個通道傳輸消息。（CSP 把通信給細化了，讓你在通信時有多種選擇，例如：用一個 channel 傳一類數據，用另一個 channel 傳另一類數據）

這就和MQ的機制有點像了。在通過MQ傳輸消息時有兩種選擇：

* 選擇把這個消息發送到哪個Exchange（類似channel）裡，對於不同的 Exchange 可以有不同的處理程序。
* 還可以把數據發送到一個 Exchange 裡，然後設置分發規則，選擇不同的處理程序。


# 並發設計模式

上面介紹了Go 中使用的並發模型，而在這種並發模型下面channel 是一個重要的概念，而下面每一種模式的設計都依賴於channel，所以有必要了解一下(詳細可以看範例的testing code)。

## 1. Barrier 模式
barrier 屏障模式故名思義就是一種屏障，**用來阻塞直到聚合所有goroutine 返回結果**。可以使用channel 來實現。

### 使用場景
 * 多個網絡請求並發，聚合結果
 * 粗粒度任務拆分並發執行，聚合結果

![](https://i.imgur.com/6oW1OtC.png)


## 2. Future 模式
這個模式常用在異步處理也稱為Promise 模式，採用一種fire-and-forget 的方式，是指主goroutine 不等子goroutine 執行完就直接返回了，然後等到未來執行完的時候再去取結果。在Go 中由於goroutine 的存在，實現這種模式是挺簡單的。

### 使用場景
* 異步
![](https://i.imgur.com/jJKoe0F.png)


## 3. Pipeline 模式
### 使用場景
* 可以利用多核的優勢把一段粗粒度邏輯分解成多個goroutine 執行Pipeline 本身翻譯過來就是管道的意思，注意和Barrire 模式不同的是，它是按順序的，類似於流水線。
![](https://i.imgur.com/CCA6InE.png)
這個圖不是很能表達並行的概念，其實三個goroutine 是同時執行的，通過buffer channel 將三者串起來，只要前序goroutine 處理完一部分數據，就往下傳遞，達到並行的目的。


## 4. Workers Pool 模式
### 使用場景
* 高並發任務

在Go 中goroutine 已經足夠輕量，甚至net/http server 的處理方式也是goroutine-per-connection 的，所以比起其他語言來說可能場景稍微少一些。每個goroutine 的初始內存消耗在2~8kb，當我們有大批量任務的時候，需要起很多goroutine 來處理，這會給系統代理很大的內存開銷和GC 壓力，這個時候就可以考慮一下協程池。

## 5. Pub/Sub 模式
發布訂閱模式是一種消息通知模式，發布者發送消息，訂閱者接收消息。

### 使用場景
* 消息隊列
![](https://i.imgur.com/qhnINvi.png)

## 注意事項
* 同步問題，尤其同步和channel 一起用時，容易出現死鎖
* goroutine 崩潰問題，如果子goroutine panic 沒有recover 會引起主goroutine 異常退出
* goroutine 洩漏問題，確保goroutine 能正常關閉
