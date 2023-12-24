# Computer Network FINAL EXAM

**TOCs:**
- [Computer Network FINAL EXAM](#computer-network-final-exam)
  - [3-1 Introduction and Transport-Layer Services](#3-1-introduction-and-transport-layer-services)
  - [3-2 Multiplexing and Demultiplexing](#3-2-multiplexing-and-demultiplexing)
  - [3-3 Connectionless Transport: UDP](#3-3-connectionless-transport-udp)


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