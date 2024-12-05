# Container Runtimes

- https://kubernetes.io/docs/setup/production-environment/container-runtimes/

Each node in a kubernetes cluster must support a container runtime.
The container runtime must conform to the `Container Runtime Interface` CRI

## Common Container Runtimes

- containerd
- CRI-O
- Docker Engine
- Mirantis Container Runtime
- rkt (end of life)

## Installation Prerequisites

- Linux kernel does not allow IPv4 packet routing between interfaces
  by default.  Most kubernetes cluster networking impelementations
  change this setting.  Some require admin intervention.

This is changed manually by doing something like:

```bash
vim /etc/sysctl.d/k8s.conf
# add the following line:
# net.ipv4.ip_forward = 1

# apply sysctl params without reboot
sudo sysctl --system

# verify
sysctl net.ipv4.ip_forward
```

## cgroup drivers

To enforce resource management, both:
- container runtime
- kubelet
need to interface with control groups. both will need a
`cgroup` driver to do so.  Both should be configured
to use the same driver.  There are currently two available:
- cgroupfs
- systemd

## Container Runtime

- https://karampok.me/posts/container-networking-with-cni/

In a OCI/CNI compatible version:
- the Container Runtime [Engine] is:
  - a daemon process that
    - sits between the
      - container scheduler and
      - the actual implementation of binaries
        - that create a container
- The daemon does not (necessarily) need to run as root
- The container runtime listens for requests from the scheduler
  - it does not touch the kernel
    - it uses external binaries
      - through container standards
        - to create or delete a container

In Kubernetes:
- the Container Runtime can be
  - CRI-O
  - containerd
- it listens for requests from the `kubelet`
  - which is the agent from the `scheduler`
    - and is located on each `node`
  - through the CRI interface

The kubelet
- is instructing the container runtime
  - to spinup a container
- the runtime is executing
  - by calling in a standard way the
  -  `runc` (binary that implements the OCI-runtime specs)
  -  and `flannel` (a binary that impelents CNI)
