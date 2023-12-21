### Q1

1. In our RDT protocols, why did we need to introduce checksums?
   - Solution To detect packet corruption during transmission.
2. In our RDT protocols, why did we need to introduce sequence numbers?
   - Solution Sequence numbers are required for a receiver to determine whether an arriving packet contains new data or is a retransmission, to support re-ordering, and provide some information about potentially dropped packets.
3. In our RDT protocols, why did we need to introduce acknowledgements?
   - Solution Acknowledgements of some form are necessary to provide feedback to the sending host, so the sending host can know whether packets were successfully received.
4. In our RDT protocols, why did we need to introduce timers?
   - Solution Timers were introduced to detect lost packets. If the ACK for a transmittedpacket is not received within the duration of the timer for the packet, the packet (or its ACK or NACK) is assumed to have been lost. Hence, the packet is retransmitted.

### Q2
1. Answer each of the following as True or False:
a) Host A is sending Host B a large file over a TCP connection. Assume Host B has no
data to send to Host A. Host B will not send acknowledgements to Host A because
Host B cannot piggyback the acknowledgements on data.
Solution False
b) The size of the TCP rwnd never changes throughout the duration of the connection.
Solution False
c) Suppose Host A is sending Host B a large file over a TCP connection. The number
of un-ACKed bytes that A sends cannot exceed the size of the receive buffer.
Solution True except for when the receive buffer has an advertised size of 0, in which
case, Host A will send 1 byte to avoid deadlock.
d) Suppose Host A is sending Host B a large file over a TCP connection. If the sequence
number for a segment of this connection is m, then the sequence number for the
subsequent segment must be m + 1.
Solution False
e) The TCP segment has a field in its header for the receive window.
Solution True
f) Suppose the last SampleRTT in a TCP connection is equal to 1 second. The current
value of TimeoutInterval for the connection will necessarily by ≥ 1 sec.
Solution False, as SampleRTT and EstimatedRTT could be set such that EstimatedRTT is much lower than SampleRTT, so 0.875 x EstimatedRTT + 0.125 SampleRTT could be much less than (1-4DevRTT), which would create a TimeoutInterval < 1.
g) Suppose Host A sends one segment with the sequence number 38 and 4 bytes of
data over a TCP connection to Host B. In this same segment, the acknowledgement
number is necessarily 42.
Solution False

### Q4
The short answer: Yep! By "1500 byte segments" we infer MSS = 1500 bytes. Maximum Segment Size (MSS) is the largest amount of bits we are going to write into a single packet (not counting headers). If we want to send more information than this, we will break it up into multiple pieces and send them all. The receiver will be able to tell the order because we attach a sequence number (which is the total number of bits sent so far, not counting headers).

The longer answer:

I usually start these problems by converting everything into standard units:

100ms = .1 seconds

10 Gbps = 10*109 bits / second

1500 byte segement = 1500* 8 = 12000 bit segment

Potentially unnecessary if you can keep a lot in your head at the same time, but slow and steady is the way to go.

Next, we just plug everything in:

10*109 = (1.22 * 1500 * 8) / (.1 * sqrt(L))

683060 = 1 / (.1 * sqrt(L))

1/683060 = .1 * sqrt(L)

10/683060 = sqrt(L)

L = 2.143 * 10-10

Which is about the loss probability given (2 * 10-10). I realize I didn't help with too much more than plugging in, but maybe that helped out.

The second part will be very similar, so I assume you can get there with algebra. Let me know if you need help.

edit: I didn't factor in congestion window size. Take with a grain of salt.


### Q9

**Step1**  
A flow table is a more comprehensive routing table. A flow table allows more variables to determine the outbound interface of a packet.

It also allows dropping of a packet and modifying packet header values.

The match of a flow table is the matching of the packet header values to the table entry values.

The action of a flow table is the dropping, modify or forwarding of the packet.

**Step2:**

The first requirement is for packets from h3 to s2 should be sent clockwise. From s2 clockwise would be to send the packet to s1, meaning that it should be sent to port 2.

Hosts h1 and h2 belong to the same network, so they will both have destination address 10.1.x.x. Similarly hosts 5 and 6 will both have destination address 10.3.x.x.

Any packet sent from h3 to s2 will be received by s2 on port 3, so the INGR port value will be 3.

The first 2 entries are then:

| Match | Action |
| --- | --- |
| INGR port: 3,Dst IP: 10.1.x.x | Forward(2) |
| INGR port: 3,Dst IP: 10.3.x.x | Forward(2) |

**Step3:** 

Now, packet from h4 to h1, h2, h5 and h6 should be sent in a counter clockwise direction, meaning the output port will be 1.

The entries are then:

| Match | Action |
| --- | --- |
| INGR port: 4,Dst IP: 10.1.x.x | Forward(1) |
| INGR port: 4,Dst IP: 10.3.x.x | Forward(1) |​


**Step4:** 

| Match | Action |
| --- | --- |
| INGR port: 3,Dst IP: 10.1.x.x | Forward(2) |
| INGR port: 3,Dst IP: 10.3.x.x | Forward(2) |
| INGR port: 4,Dst IP: 10.1.x.x | Forward(1) |
| INGR port: 4,Dst IP: 10.3.x.x | Forward(1) |​