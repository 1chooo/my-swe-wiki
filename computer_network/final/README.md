# Computer Network FINAL EXAM

**TOCs:**
- [Computer Network FINAL EXAM](#computer-network-final-exam)
  - [3-1 Introduction and Transport-Layer Services](#3-1-introduction-and-transport-layer-services)
  - [3-2 Multiplexing and Demultiplexing](#3-2-multiplexing-and-demultiplexing)
  - [3-3 Connectionless Transport: UDP](#3-3-connectionless-transport-udp)
    - [Actions](#actions)
    - [Checksum](#checksum)
  - [3-4 Principles of Reliable Data Transfer](#3-4-principles-of-reliable-data-transfer)
  - [4-1 Network Layer: Overview](#4-1-network-layer-overview)
  - [4-2 What's Inside a Router?](#4-2-whats-inside-a-router)
    - [Longest-Prefix Matching](#longest-prefix-matching)
    - [計算 Buffering](#計算-buffering)
  - [4-3 IP: the Internet Protocol](#4-3-ip-the-internet-protocol)


## 3-1 Introduction and Transport-Layer Services

- Sender: 把訊息碎片化，傳到網路層，決定 header fields value，建立 segment 並傳導 IP
- Receiver: 把碎片組合起來，傳到應用層，檢查 header value，extract application layer message，demultiplexing 到應用程式藉由 socket
- 兩個傳輸協定: TCP, UDP
- Transport Layer processes 根據網路層的服務
- 網路層通訊在 host 之間
- 傳輸層協定執行在 end systems 而不是 routers (Transport-layer protocols are implemented in the end systems but not in network routers.)

TCP V.S. UDP
- TCP: reliable, in-order delivery, congestion control, flow control, connection setup
- UDP: unreliable, unordered delivery, no-frills extension of "best-effort" IP

UDP only provides these two services:

- process-to-process data delivery
- error checking

Internet Protocol (in Network Layer) does not guarantee:

- segment delivery
- orderly delivery of segments
- the integrity of the data in the segments

**Analogy:**

- application messages = letters in envelopes
- processes = cousins
- hosts (also called end systems) = houses
- transport-layer protocol = Ann and Bill
- network-layer protocol = postal service (including mail carriers)
- Ann and Bill do all their work within their respective homes, they are not involved in sorting mail in any intermediate mail center.

**Transport protocol are often constrained by network-layer protocol.**

For instance:

- delay
- bandwidth

However, some services can be offered by a transport protocol even when network-layer protocol doesn’t offer.

For instance:

- reliable data transfer
- encryption


## 3-2 Multiplexing and Demultiplexing

- Data passes from the network to the process, and from the process to the network through sockets.
- A process (as part of a network application) can have one or more sockets.

- Demultiplexing: Delivering the data in a transport-layer segment to the correct socket. 發散出去
- Multiplexing: Gathering data chunks at the source host from different sockets, encapsulating each data chunk with header information (that will later be used in demultiplexing) to create segments, and passing the segments to the network layer. 聚合起來，有 source IP, destination IP, each datagram carries one transport layer segment, each segment has source, destination port number

IP/UDP datagrams with same destination port number but different source IP addresses and/or source port numbers will be directed to the same socket at the destination host. 目標一樣的話，就會被送到同一個 socket

```cpp
mySocket = socket(AF_INET, SOCK_STREAM);

mySocket.bind(myAddress, 9157);
```

```cpp
mySocket = socket(AF_INET, SOCK_DGRAM);

mySocket.bind(myAddress, 6428);
```

```cpp
mySocket = socket(AF_INET, SOCK_STREAM);

mySocket.bind(myAddress, 5775);
```

TCP 由 source port number, destination port number, source IP address, destination IP address 來 demultiplexing，receiver 會用 4-tuple 來 direct segment 到 correct socket

每個 socket 被自己的 4-tuple 所識別，每個 socket associated with different connecting client

- Multiplexing, demultiplexing 基於 segment, datagram header fields value，且會發生在所有 layer
- UDP 用 destination port number, IP 來 demultiplexing，TCP 會用 4-tuple 來 demultiplexing

## 3-3 Connectionless Transport: UDP

- 最佳化的 UDP 可能會 lost, duplicate, deliver out of order，但是 TCP 會保證 in-order, no loss，但是 UDP 比 TCP 快
- connectionless: 不用在 UDP sender, receiver 之間建立 connection
- 每個 UDP segment 都是獨立的，不會被其他 segment 影響


- UDP 用在 streaming multimedia, telephony, DNS, Internet telephony, Internet TV, DNS, SNMP, HTTP/3 (DNS: Domain Name System, SMTP: Simple Mail Transfer Protocol, HTTP/3: Hypertext Transfer Protocol)
- 如果要穩定傳輸 over UDP，就要在 application layer 做 needed reliability, congestion control, flow control

```
Format

------
0      7 8     15 16    23 24    31
+--------+--------+--------+--------+
|     Source      |   Destination   |
|      Port       |      Port       |
+--------+--------+--------+--------+
|                 |                 |
|     Length      |    Checksum     |
+--------+--------+--------+--------+
|
|          data octets ...
+---------------- ...

    User Datagram Header Format

Fields
------
```

- length: in bytes of UDP segment, including header and data
- data: to/form application layer

### Actions

- Sender: passed an application layer message, 決定 UDP segment header fields value, 建立 UDP segment，傳導 IP
- Receiver: 接收 segment 來自 IP, 檢查 UDP checksum header value, extract application layer message，demultiplexing 到應用程式藉由 socket

### Checksum

Detect "errors" in transmitted segment

![](https://i.imgur.com/T7mCUiI.png)

- Sender: 把 UDP segment 當成 16-bit word sequence, addition, 1's complement sum of segment, checksum field value = 1's complement of sum 放在 checksum field
- Receiver: 計算 received segment 的 segment, 相等的話就沒有 error，不相等的話就有 error，但是 checksum 不能偵測所有 error，只能偵測一些 error

example: add two 16-bit integers

[HW2-3](https://github.com/1chooo/my-swe-wiki/tree/main/computer_network/hw02#q3-2-3)

```
  1101011010110101
+ 0110011101101101
------------------
  0011111000010010
+ 0000000000000001  // 1's complement (wrap around)
------------------
  0011111000010011
```

- 1's complement sum: 0011111000010011

如果兩個結尾分別是 01, 10 又或是 10, 01 儘管有改變，但是 1's complement sum 會是一樣的，所以 checksum 不能偵測所有 error

TCP V.S. UDP
- TCP: 
  - reliable
  - in-order delivery
  - congestion control
  - flow control
  - connection setup
- UDP: 
  - unreliable
  - unordered delivery
  - no-frills extension of "best-effort" IP

UDP only provides these two services:

- process-to-process data delivery
- error checking

Internet Protocol (in Network Layer) does not guarantee:

- segment delivery
- orderly delivery of segments
- the integrity of the data in the segments


## 3-4 Principles of Reliable Data Transfer


## 4-1 Network Layer: Overview

```
Application Layer
-----------------
Transport Layer
-----------------
Network Layer
-----------------
Data Link Layer
-----------------
Physical Layer
```

transport segment from sending to receiving host
- sender: encapsulates segments into datagrams, pass to link layer
- receiver: delivers segments to transport layer protocol
- network layer protocols in every host, router
- router examines header fields in all IP datagrams passing through it
- network-layer functions:
  - forwarding: move packets from router's input to appropriate router output
  - routing: determine route taken by packets from source to dest.
  - path determination: route taken by packets from source to dest.
  - call setup: some network architectures require router call setup along path before data flows
  - congestion control: routers may throttle senders when network congested


## 4-2 What's Inside a Router?

- Input ports: 負責 forwarding datagrams from input to output ports
- input port queuing: 如果 datagrams 抵達速度快過魚 forward rate，就會發生 queueing delay，如果 queueing delay 太長，就會發生 packet loss

### Longest-Prefix Matching 

[HW2-5](https://github.com/1chooo/my-swe-wiki/tree/main/computer_network/hw02#q5-4)

找最長的去配對

```
| Destination Address Range  |  Link Interface  |
|----------------------------|------------------|
| 11001000 00010111 00010*** ******** |        0         |
| 11001000 00010111 00011000 ******** |        1         |
| 11001000 00010111 00011*** ******** |        2         |
| otherwise |        3         |
```

- 11001000 00010111 00010110 10100001 which interface? Ans: 0
- 11001000 00010111 00011000 10101010 which interface? Ans: 2

- Often done using TCAMs (ternary content addressable memories)
- content addressable: present address to TCAM, retrieve address in one clock cycle
- ternary: each TCAM entry has 3 parts
  - value
  - mask
  - output port number


- Switching fabrics: 傳送 packet 曾 input link 到合適的 output link
- switching rate: rate at which packets can be transfered from inputs to outputs
- Oftem measured as multiple of input/output line rate
- 三個主要的 switching fabrics:
  - memory: 第一代 routers, switch 在 CPU, packet 要複製到 system memory, 速度慢限制在記憶體 bandwidth (2 bus crossings per datagram)
  - bus: datagram 從 input port 的 memory 到 output port 的 memory 藉由分享的 bus, 32 Gbps bus, Cisco 5600: sufficient
  - interconnection network: crossbar, clos network, other interconnection nets developed to connect processors in multiprocessor, multistage switch nxn switch frommultiple stages of smaller switches, exploiting parallelism: fragment datagram into fixed length cells on entry
    - scaling, using multiple switching planes in parallel, speedup, scaleup via parallelism
    - Cisco CRS router: 8 switching planes in basic unit, 3-stage interconnection network, up to 100's Tb/s


input port queuing:
- fabric slower than input ports combined -> queueing may occur at input queues
- queueing delay and loss due to input buffer overflow!
- head-of-the-line (HOL) blocking: queued datagram at front of queue prevents others in queue from moving forward

輸入端口排隊：

- 當網絡連接速度比輸入端口速度慢時，可能導致輸入隊列的排隊現象。
- 排隊延遲和由於輸入緩衝區溢出而導致的丟失！
- 隊列中排隊的數據報在隊列前端阻止其他數據報前進，稱為頭部阻塞（HOL）。


Output port queuing:
- buffering required when datagrams arrive from fabric faster than the transmission rate
- Datagrams can be lost due to congestion, lack of buffers
- scheduling discipline chooses among queued datagrams for transmission
- priority scheduling - who gets best performance, network neutrality?
- buffering when arrival rate via switch exceeds output line speed
- queueing delay and loss due to output buffer overflow!

輸出端口排隊：

- 當數據報從傳輸織理到達速度快於傳輸速率時，需要緩衝。
- 可能由於擁塞或緩衝區不足導致數據報丟失。
- 排程紀律從排隊的數據報中選擇要傳輸的數據報。
- 優先級排程 - 誰獲得最佳性能，網絡中立性？
- 當網絡交換機的到達速率超過輸出線速度時，需要進行緩衝。
- 可能由於輸出緩衝區溢出導致的排隊延遲和丟失！


### 計算 Buffering 

[HW2-4](https://github.com/1chooo/my-swe-wiki/tree/main/computer_network/hw02#q4-3)

$\frac{RTT \cdot C}{\sqrt{N}}$

太多的 buffer 會造成 delay 增加，特別在家裡的 router
- long RTTs: 在即時的 app 有很差的效能，sluggish TCP response
- recall delay-based congestion control: keep bottleneck link just full enough but no fuller

管理 buffer 的方法:
- drop: which packet to add, drop when buffers are full, tail drop, priority drop
- marking: which packets to mark to signal congestion, ECN (Explicit Congestion Notification), RED (Random Early Detection)

FCFS (First Come First Serve) scheduling:
- packet scheduling: 
  - deciding which packet to send next on link:
  - FIFO (first in first out)
  - priority: by classification (any header fields), send packet from highest priority queue that has buffed packets
  - round robin: by classification (any header fields), sever cyclically, repeatedly scaning, sending one complete packet from each class (if available)
  - weighted fair queuing (WFQ): generalized Round Robin, each class gets weighted amount of service in each cycle, minimum bandwidth guarantee per class
- FCFS: packets transmitted in order of arrival to output port
  - FIFO (first in first out)


Network Neutrality:
- How an ISP (Internet Service Provider) should share/allocation its resources among its customers?
- social, economic, political, regulatory issue


2015年美國聯邦通信委員會（FCC）有關保護和促進開放互聯網的法令包括三條“明確、明亮的界線”規則：

1. **不得封鎖（No Blocking）：** 互聯網服務提供商不得阻止合法內容、應用、服務或無害設備，但允許進行合理的網絡管理。

2. **不得限速（No Throttling）：** 互聯網服務提供商不得基於互聯網內容、應用或服務的性質，或使用無害設備，而對合法互聯網流量進行損害或降級，同樣也允許進行合理的網絡管理。

3. **禁止付費優先（No Paid Prioritization）：** 互聯網服務提供商不得進行付費優先處理，即不得對某些流量提供特殊的優先處理，同樣也不能進行付費的特殊流量處理。

這些規則旨在確保互聯網的開放性和公平性，防止網絡提供商對互聯網內容、服務或使用者的不公平差別對待，並允許他們在合理範圍內進行網絡管理以確保網絡安全和效能。

## 4-3 IP: the Internet Protocol