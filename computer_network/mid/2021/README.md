# 2021

### Q1
(3 marks) Naming in the Internet uses a hierarchical scheme called the ___?

**Answer:**
Domain Name System (DNS)

### Q2
- The UDP server described needed only one socket, whereas the TCP server needed two sockets. Why?
- If the TCP server were to support $n$ simultaneous connections, each from a different client host, how many sockets would the TCP server need?


**Answer:**

With the UDP server, there is no welcoming socket, and all data from different clients enters the server through this one socket.

With the TCP server, there is a welcoming socket, and each time a client initiates a connection to the server, a new socket is created.

Thus, to support n simultaneous connections, the server would need n + 1 sockets.

### Q3
(2 marks) What encryption services are provided by HTTP?

**Answer:**
None

### Q4
(3 marks) Which system call attaches a local address to a socket?

**Answer:**
bind

### Q5
(2 marks) Which system call is to specify queu size for a server socket?

**Answer:**

### Q6
(10 marks) How long does it take a packet of length $1,000 bytes$ to propagate over a link of distance $2,500 km$, propagation speed $2.5 × 10^8 m/s$, and transmission rate $2 Mbps$? 

**Answer:**

- Transmission Delay: $L/R = 7 (bits/ byte) * 1000 (bytes) / 2000000 bps = 0.004$
- Propagation Delay: $d/s = 2500 (km) / 2.5 * 10^8 (m/s) = 0.01$

Therefore, the total time = $0.004 + 0.01 = 0.014$

### Q7
(10 marks) DNS resource record format is (name, value, type, ttl). A diferent type has the corresponding name and value which has the different meaning. What are the meanings of name and value when the types is "A" (5 marks) and "NS" (5 marks) respectively?

**Answer:**

- A: name = hostname, value = IP address
- NS: name = domain, value = hostname of authoritative name server for this domain
- CNAME: name = alias name for some "canonical" (the real) name, value = canonical name
- MX: value = name of SMTP mail server associated with name

### Q8
(2 marks) If a web server wants to save the user login name in the client side, then the client can send back the user login name at later access, which method can it use?

**Answer:**

Cookie

### Q9
(3 marks) In the TCP/IP model, which layer deals with reliability, flowm control, and error congestion?

**Answer:**

Transport Layer

### Q10
(10 marks) Following is a DNS database for `ncu.edu.tw`:

```shell
ncu.edu.tw 86400 IN MX mail.ncu.edu.tw
ncu.edu.tw 86400 IN MX smtp.ncu.edu.tw
mail.ncu.edu.tw 86400 IN A 140.115.54.87
smpt.ncu.edu.tw 86400 IN A 140.115.12.21
theworld.ncu.edu.tw 86400 IN A 140.4.87.63
cs.ncu.edu.tw 86400 IN A 140.115.13.24
cs.ncu.edu.tw 86400 IN MX dio.ncu.edu.tw
dio.ncu.edu.tw 86400 IN CAME theworld.ncu.edu.tw
```

What is the mail server address for `jojo@cs.ncu.edu.tw`?


**Answer:**

`140.4.87.63`

```
dio.ncu.edu.tw -> theworld.ncu.edu.tw -> 140.4.87.63
```

### Q11
(10 marks) Suppose that UDP reciever computes the Internet checksum for the received UDP segment and finds that it matches the value carried in the checksum field. Can the receiver be absolutely certain that no bit errors have occurred? Explain.

**Answer:**

Certainty if no bit errors at receiver end using checksum using UDP:

The receiver in the UDP(User Datagram Protocol) verifies the received segment by calculating internet checksum and comparing it with the value in the checksum field.

The 1's compliment of the sum is considered as checksum. So, when this check sum is used to detect the errors in the packet, the errors remain in under cover.

In case if two 16-bit words are added, then there is a scope for flipping 0's and 1's. If the bits are flipped, the sum will be same and error can't detected.

Therefore, it is not possible to the receiver in UDP to be sure that there are no bit errors have occurred.

### Q12
(10 marks) Is it possible for an applicstion to enjoy reliable data transfer even when the application runs over UDP? If so, how?

**Answer:**


### Q13 
Q: Suppose Host A wants to send a larfe file to Host B. The path from Host A to Host B has three links, of rates $R1 = 500 kbps$, $R2 = 2 Mbps$, and $R3 = 1 Mbps$.

- (5 marks) Assuming no other traffic in the network, what is the thoughput for the file transfer?
- (5 marks) Suppose the file is $4 × 10^6 bytes$ Dividing the file size by the throughput, roughly how long will it take to transfer the file to Host B?


**Answer:**

- $R = min(R_1, R_2, R_3) = 500 kbps$ The throughput for the file transfer is $min(500kbps, 2 Mbps, 1 Mbps)$. So, the throughput is $500 kbps$.
- Consider given data: $File size = 4 * 10^6 bytes$ and $Throughput = 500 kbps$. Dividing the file size by the throughput, roughly how long will it take to transfer the file to Host B: file size / throughput for the transfer $= 32000000 bytes / 500 kbps = 64 sec$

### Q14
Q: The picture below is part of rdt 3.0 finite state machine, fill in th blank below, please write down your answer on the answer sheet. (3 marks for each a, b, c, d, e.)

![](./q14_2.png)

**Hint:**
- The main change from rdt 1.0 to 2.9: ACK, checksum.
- The main change from rdt 2.0 to 3.0: Timer.

**Answer:**


