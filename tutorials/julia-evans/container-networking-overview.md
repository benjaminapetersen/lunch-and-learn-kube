# A container networking overview

- 2016. This is old.  Conceptually relevant, but not current info.

## What is Container Networking?

When running programs in containers, there are 2 options:

- Run in the host network namespace.
  - Nothing unusual here.
  - Run a program on `:8282`, then it runs on `:8282` on the `host.`
- Run the program in its own network namespace.
  - If you run the program in its own network namespace, other
    programs on other computers need to be able to make network
    connections to this program.

How complicated can this be?
Apparently quite a bit complicated.

## Kubernetes: Every Container Gets an IP

- Kubernetes schedules containers to run on worker nodes.
- Every container on any worker node gets its own IP address.
  - How?
  - By default, containers run on a host.
    - That host machine gets one IP address.
- Other programs in the cluster can talk to the container via this IP address.

## On AWS: Every Container gets an IP

There are several ways to give every container an IP.
This article uses AWS as the context.

Context:
- You have an AWS Instance (computer)
- The computer has an IP address
  - such as 172.9.9.9
- You want your container on this instance to have an IP address
  - such as 10.4.4.4

Let's learn how we can get packets sent to 10.4.4.4 on 172.9.9.9.

On AWS:
- VPC Routing Tables are the technology
  - Effectively can define `send packets for 10.4.4.*` to `172.9.9.9`
  - Limit to 50 entries
    - If you want a cluster of more than 50 instances, need to find another solution

## Networking Basics

- IP addresses
- subnets
- MAC Addresses
- local networking

In computer networking:
- Programs send packets to each other
- Every packet (mostly) has an IP address (or two, source and destination, or more)
- In Linux, the Kernel is responsible for implementing most networking protocols
- Subnets are a division of an IP network
  - IE, a collection of IP addresses
  - IE, 10.4.4.0/24 = 10.4.4.0 through 10.4.4.255
  - Can write as 10.4.4.* (not sure this is standardized, but is blog-common at least)

## Thing 0: Parts of a Network Packet

- A network packet has multiple `layers`
  - frames are data units for the data layer
  - packets are transmission units for the transmission layer
- There are many kinds of network packets
- We will focus on HTTP request packets
  1. the MAC address (MAC address indicating where packet should go)
     - Layer 2 of OCI
  2. The source IP and destination IP
    - Layer 3 of OCI
  3. The port an other TCP/UDP Information
    - Layer 4 of OCI
  4. The contents of your HTTP packet like `GET /`
    - Layer 7 of OCI

## Thing 1: local networking vs far-away networking

- When sending packets directly to a computer on a local network:
  - Packets are addressed by `MAC address`
    - `MAC address` can be found via `ifconfig`
    - `MAC address` is something like `3c:97:ae:44:b3:7f`
- In AWS, local network is `Availability Zone`
  - Basically, anyway
  - If 2 instances are on the same AWS Availabilty Zone
    - MAC address can be used for destination
    - Packets will route to the correct machine
    - IP address is irrelevant
      - Within the same Availability Zone MAC is used
      - This is generally true of local networking
      - You can send to incorrect IPs
        - Since the MAC is correct, it will arrive anyway
- When not in the same `Availabilty Zone`
  - Enter routers!
    - Routers need to look at the IP address
      - And forward packets between networks
- In AWS, you do not configure the routers
  - This is a good thing
  - Don't need to know how they work
  - They do their job (you hope)
- If you manage your own data center
  - Enjoy.
  - Lots you can do to configure routers
  - Enjoy.

## The Route Table

- How to control the MAC address a packet is sent to?
- When a packet is sent to an IP, such as `172.23.2.1`
  - The OS
    - Linux, probably
  - Looks up the MAC address for the IP address
    - In a table it maintains for itself
      - Called the `ARP table`
    - Sets the MAC address
    - Sends the packet

So what if you want to manually manipulate this?
You edit the IP table

```bash
# now 10.4.4.* packets will go to 172.23.1.1
# however, note that this is IP to IP, not MAC
# I believe we are still trusting the OS to figure out MAC addresses
sudo ip route add 10.4.4.0/24 via 172.23.1.1 dev eth0
```

The `ip route` command addres an entry to the `route table ` on your computer.
This route entry table instructs Linux for mapping IPs.

## We can give containers IPs!

We now know all the basic tools for ONE of the ways to route container IP addresses.

What to do:
- Pick a differnet `subnet` for every `computer` (`node` in Kubernetes lingo) on
  your `network` (`cluster` in kubernetes lingo).  This subnet is something like
  `10.4.4.0/24` and indicates `10.4.4.*` or `256` addresses on this computer which
  can each be assigned to a container.
- On every computer (`node`) in the network, add `routes` for every other computer.
  This means routes for:
  - 10.4.1.0/24 (256 IP addresses)
  - 10.4.2.0/24 (256 IP addresses)
  - 10.4.3.0/24 (256 IP addresses)
- Packets on the network can now be routed to many individual containers.

The `route table` is critical for networking containers.

## What if two computers are in different availability zones?

- Route tables only work if computers are directly connected (one network)
- For computers on other networks, need more complexity
Example:
- Send packet to container IP 10.4.4.4
  - On coputer 172.9.9.9
    - This is on another network
    - We have to address the packet to 172.9.9.9
      - Where does IP 10.4.4.4 go?

Encapsulation
- Network packets can be wrapped in network packets

We have inner packet:

```
IP: 10.4.4.4
TCP stuff
HTTP stuff
```
And we will wrap it inside another packet:

```
IP: 172.9.9.9
(extra wrapper stuff)
IP: 10.4.4.4
TCP stuff
HTTP stuff
```

There are AT LEAST 2 different ways of doing `encapsulation`.
- ip-in-ip
- vxlan

In `vxlan`, the whole packet (including the MAC address) is wrapped inside a UDP
packet.  It looks like this:

```
MAC address: 11:11:11:11:11:11
IP: 172.9.9.9
UDP port 8472 (the "vxlan port")
MAC address: ab:cd:ef:12:34:56
IP: 10.4.4.4
TCP port 80
HTTP stuff
```

In `ip-in-ip` encapsulation, an extra IP header is entered over the old IP header.
The MAC address is lost (but, you likely don't need this anyway).

```
MAC:  11:11:11:11:11:11
IP: 172.9.9.9
IP: 10.4.4.4
TCP stuff
HTTP stuff
```

Setting up encapsulation

There are tools for doing encapsulation just like there are tools for routing.

```bash
# You setup a new `network interface` with encapsulation configured.
sudo ip tunnel add mytun mode ipip remote 172.9.9.9 local 10.4.4.4 ttl 255
sudo ifconfig mytun 10.42.1.1

# then you also setup a route table
# tell it to route packets with the encapsulation network interface
sudo route add -net 10.42.2.0/24 dev mytun
sudo route list
```

The above is examples taken from the article.

## How do Routes get Distributed?

- Routes such as `10.4.4.4 should go via 172.9.9.9` can be added
- But how do they get in the routing table?
- Can they be configured automatically?

Each container networking thing runs some kind of `daemon` on every box
in the network which is responsible for adding routes to the route table.

This happens one of two ways:

- The routes are in an `etcd` cluster.  The program asks the etcd cluster
  to figure out which routes to set
- Otherwise, use teh BGP protocol to talk to eachother about routes.
  A daemon (`BIRD`) listens for BGP messages on every box

## What happens when packets get to your box?

When:
- You are running `docker` or another container runtime
- A packet comes in on IP 10.4.4.4
  - Via the box IP of 172.9.9.9
How does the container get the packet?

`bridge networking` solves this problem.

Roughly (author states this is fuzzy and needs work):
- Every packet on the computer goes out through a real interface
  - Such as `eth0` for ethernet
- Docker will create a `fake` (virtual) netwrok interface
  - For every container
  - These have IP addresses such as 10.4.4.4
- The virtual network interfaces are `bridged` to your real network interface
  - This means that packets get copied(?)
    - To the network interface that corresponds to the network card
    - Then sent out ot the internet
  - NOTE: this doesn't explain receiving traffic, however.

## Finale: how all these container networking things work

Several specific Kubernetes networking tools:
- `Flannel`: supports networking via
  - `vxlan`: encapsulate all packets
  - `host-gw`: just set route table entries, no encapsulation
  - The daemon that sets the routes gets them from an etcd cluster
- `Calico`: supports networking via
  - `ip-in-ip`: encapsulation
  - `regular mode`, just set route table entries, no encapsulation
  - the daemon that sets the routes gets them using BGP messages from other hosts
    - There is still an etcd cluster for calico
      - It is not used for distributing routes, however
  - calico has the option to not use encapsulation
    - this is good?
    - flannel can also not use encapsulation
      - i guess good also?

These container networking tools will setup the routes for you.s

## Is this software defined networking?

Maybe?
Need to define networking.
There is software running.
There is networking.
Not sure if it is software defined networking.

## That's all

get used to:

```bash
ip
ifconfig
tcpdump
```

and it will help you understand the basics of networkign in a Kubernetes install.
