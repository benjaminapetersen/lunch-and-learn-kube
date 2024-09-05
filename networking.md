# Cluster Networking

- [Networking Design Proposal 2017](https://github.com/kubernetes/design-proposals-archive/blob/main/network/networking.md)
- [Kubernetes Networking Model](https://kubernetes.io/docs/concepts/cluster-administration/networking/#how-to-implement-the-kubernetes-network-model)

Cluster networking is central to Kubernetes

There are 4 networking problems to address:
- highly-coupled container-to-container communications.
    - This is solved by `pods` and `localhost` communication.
- pod-to-pod communications
    - This document covers this type of communication.
- pod-to-service communications
    - This is covered in services documentation.
    - https://kubernetes.io/docs/concepts/services-networking/service/
- external-to-service communications
    - This is also covered in services documentation.
    - https://kubernetes.io/docs/concepts/services-networking/service/

Kubernetes is about sharing machines among applications.
Sharing requires that two application do not try to use the same ports.
Coordinating ports at scale is very difficult.

Dynamic port allocation introduces a lot of complication to the system.
Rather than dealing with dynamic ports, Kubernetes takes a different
approached outlined in the networking model document:
- https://kubernetes.io/docs/concepts/services-networking/


# Networking Model

As noted above, here again is the services networking model documentation:

- https://kubernetes.io/docs/concepts/services-networking/

Briefly:
- every `pod` gets a unique cluster-wide IP address
    - No need to make explicit links between pods
    - No need to deal with mapping container ports to host ports
- Pods can therefore be treated like VMs or physical hosts in
  terms of port allocation, naming, service discovery, load balancing,
  application configuration, and migration.

Kubernetes requires the following from any networking implementation:
- pods can communicate with all other pods on any node without NAT
- agents on a node (system daemons, kubelet, etc) can communicate
  with all pods on that node.

For platforms that support pods running on the host network
- ex: linux
- when pods are attached to the host network of the node,
    - the pods can still communicate with
        - all other pods on all other nodes
        - without NAT

# Services

- https://kubernetes.io/docs/concepts/services-networking/service/

Services covers 2 of the 4 original networking problems listed at the top:

- pod-to-service communications
- external-to-service communications
