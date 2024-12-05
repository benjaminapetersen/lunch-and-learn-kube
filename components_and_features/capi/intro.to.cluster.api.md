# Introduction to Kubernetes Cluster API (CAPI)

## What is Cluster API?

- Declarative API to simplify & automate
  - lifecycle management
    - create
    - scale
    - upgrade
    - destroy
  - including
    - configuration
    - infrastructure
      - virtual machines
      - VPCs
      - load balancers
      - etc
- Makes Kubernetes Clusters themselves 1st class citizens
  - CRDs
  - Controllers

## Why Use CAPI?

- Standardizes lifecyle management of kube clusters
  - across infra providers
    - AWS
    - Aure
    - GCP
    - VMware
    - pre-provisioned
    - many more
- CAPI controllers actively reconciple current state
  of the cluster with what is declared
    - Easier to deploy clusters
    - Easier to maintain clusters

## What are the major CAPI components?

- Requires
  - A kube cluster to host CRDs & controllers
    - A management cluster
      - Any lightweight kube cluster
      - Even a kind cluster locally!
  - Workload clusters
    - Deployed from the management clusters
- Infrastructure specs
  - CAPI CRDs
- Controllers
  - Core Controller Manager
    - Responsible for
      - OwnerRefs
      - cleanup up objects
      - resource state change
      - BoostrapConfig
      - Machine resources
    - Controllers
      - Cluster
      - machine
      - machine
      - machinedeployment
      - etc
  - Bootstrap Controller Manager
    - Responsible for
      - Bootstrapping clusters!
    - Default provider is
      - ClusterAPI KubeADM
  - Infrastructure Controller Manager
    - Responsible for
      - Virtual Machines
      - Networking
    - Providers
      - Unique to the infrastructure provider
      - Most popular
        - AWS
        - Azure
        - GCP
        - VMware
  - Control Plane Controller Manager
    - Responsible for
      - instantiating the control plane

## Custom Resource Definitions (CRDs)

- The Inputs to the CAPI orchestration
- Many CRDs with many parameters

- CRDs
  - Cluster
    - `Cluster`
      - Name
      - CIDRs (for Pods,Services)
      - ControlPlaneRef
      - InfrastructureRef
    - <Provider>Cluster
  - Control Plane
    - KubeadmControlPlane
      - KubeadmConfigSpec
    - <Provider>MachineTemplate
      - ie, AWSMachineTemplate, AzureMachineTemplate,etc
  - Worker Node
  - Workload
