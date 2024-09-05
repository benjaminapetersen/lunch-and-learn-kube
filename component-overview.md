# Kubernetes Components

- https://kubernetes.io/docs/concepts/overview/components/

A kubernetes cluster is
- Worker machines, called nodes, that run containerized applications
  - Every cluster has at least 1 worker node
  - Worker nodes host Pods
- Controle plane
  - Manages worker nodes
  - Manages pods in the cluster

In a production environment
- Control plane usually runs across multiple machines
- CLuster usually runs multiple nodes

Control Plane
- Makes global decisions
  - such as scheduling
- Detects & responds to cluster events
  - example: start a new pod
- Can be run on any machine in the cluster
  - Typically run all components on same machine
    - With replicas on other machines


## Control Plane Components

- kube-apiserver
  - exposes the kubernetes API
  - the front end for the kubernetes control plane
  - designed to scale horizontally
    - ie, by deploying more instances
    - with a load balancer
- etcd
  - consistent & highly-available key value store
  - backs all cluster data
    - be sure to back this up!
- kube-scheduler
  - watches for new pods with no assigned node
    - selects a node for them to run on
  - factors into account:
    - individual and collective resource requirements
    - hardware/software/policy constraints
    - affinity and anti-affinity
    - data locality
    - inter-workload interference
    - deadlines
- kube-controller-manager
  - runs controller processes
    - logically, each controller is a separate process
    - however, for simplicity, all are compiled into a single binary
      - run as a single process
  - many different types of controllers such as
    - node controller
    - job controller
    - endpoints slice controller
    - service account controller
    - namespace controller
    - etc
- cloud-controller-manager
  - runs controller processes
    - logically, each controller is a separate process
    - however, for simplicity, all are compiled into a single binary
      - run as a single process
  - embeds cloud-specific control logic
  - links your cluster into a cloud provider's API
  - only runs controllers specific to a particular cloud provider
  - on-prem or local (kind) will not run this component

## Node Components

- kubelet
  - agent that runs on each node
  - ensure containers are running in a pod
  - manages PodSpecs provided
    - either locally by watching a dir
    - or via the kube api server
  - only manages containers created by kubernetes
- kube-proxy
  - network proxy
  - runs on each node in a cluster
  - implements part of the kubernetes service concept
  - maintains network rules on the node
    - allows network communication to pods
      - from network sessions inside or outside cluster
  - uses either
    - the operating system packet filter layer
    - or forwards traffic itself
- container runtime
  - fundamental component
  - manages the execution and lifecycle of containers
  - examples
    - containerd
    - CRI-O
    - etc
    - any impl of the Kubernetes CRI (container runtime interface)

## Addons

- DNS
  - oddly an addon, because cluster DNS is required
  - is a DNS server
  - serves DNS records for Kubernetes services
- Web UI
  - a dashboard
- Container Resource Monitoring
  - records generic time-series metrics
  - records to a central database
  - provides UI for browsing data
- Cluster-level logging
  - saves container logs to central store
  - usually a search/browsing interface
- Network Plugins
  - implements container network interface (CNI)
  - allocates IP addresses to pods
  - enalbes pods to communicate within a cluster
