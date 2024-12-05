# Julia Evans

- https://jvns.ca/
- https://x.com/b0rk 
- https://wizardzines.com/ 

## Julia Evans Blog by Category

- https://jvns.ca/#computer-networking
- https://jvns.ca/#dns
- https://jvns.ca/#how-a-computer-thing-works 
- https://jvns.ca/#kubernetes---containers 
- https://jvns.ca/#debugging


## Learning About Kubernetes: Post 1

- A few things about Kubernetes: https://jvns.ca/blog/2017/06/04/learning-about-kubernetes/
- Container Networking Overview: https://jvns.ca/blog/2016/12/22/container-networking/

- Recommends Kelsey Hightower talks 
- Basically Kubernetes is a distributed system that runs programs 
  (well, containers) on computers. You tell it what to run, 
  and it schedules it onto your machines.
- some funny comic illustrations:
    - https://drawings.jvns.ca/drawings/scenes-from-kubernetes-page1.svg
- Kubernetes from the ground up
    - Part 1: The Kubelet
    - Part 2: The API Server
    - Part 3: The Scheduler
    - etcd: kubernetes brain 
        - All Kubernetes components are stateless
            - API server
                - manages auth (many ways)
                - gates what goes into the etcd
            - Scheduler
                - decides which nodes to send pods to 
            - Kubelet
                - runs pods on nodes
            - Controller manager
                - manages controllers
                - controllers do things 
            - etc
        - etcd stores state
        - etcd is a key-value store 
        - generally, kubernetes works by components watching etcd for 
          things they are concerned about, and reacting when they have 
          something to do
    - broken?
        - find the controller responsible
            - check that controller's logs 
            - also, check status of the object?

To run programs on Kubernetes, you need:
- the scheduler
- the api server
- etcd
- kubelets
    - on every node to actually execute containers
- controller manager
    - sets up things


## Cluster Networking

- https://kubernetes.io/docs/concepts/cluster-administration/networking/

Supposedly kubernetes networking is not an enigma.
- every container gets an IP
    - This is not trivial

Things to understand about Kubernetes Networking
- Overlay networks
    - This is a start: https://jvns.ca/blog/2016/12/22/container-networking/
- Network namespaces
- DNS since Kubernetes has its own DNS server
- route tables, how to run `ip route list` and `ip link list`
- network interfaces
- encapsulation (vxlan, UDP)
- basics of iptables, how to use and read iptables configuration
- TLS, server serts, client certs, certificate authorities

## Kubernetes source is easy to read

- its golang
- it moves fast

## Kubernetes slack group is great

- slack.kubernetes.io 