# Computer Network FINAL EXAM

**TOCs:**
- [Computer Network FINAL EXAM](#computer-network-final-exam)
  - [3-1 Introduction and Transport-Layer Services](#3-1-introduction-and-transport-layer-services)
  - [3-2 Multiplexing and Demultiplexing](#3-2-multiplexing-and-demultiplexing)
  - [3-3 Connectionless Transport: UDP](#3-3-connectionless-transport-udp)
    - [Actions](#actions)
    - [Checksum](#checksum)
  - [3-4 Principles of Reliable Data Transfer](#3-4-principles-of-reliable-data-transfer)


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