# Kubeadm

https://kubernetes.io/docs/reference/setup-tools/kubeadm/

```bash
# kubeadm is a cli to build kubernetes clusters
# it automates the process of setting up the control plane and worker nodes

# bootstrap a kubernetes control-plane node
kubeadm init

# bootstrap a kubernetes worker node and join it to the cluster
kubeadm join

# upgrade kubernetes cluster to a newer version
kubeadm upgrade

# if cluster was initialized with kubeam v1.7.x or lower,
# kubeadm config will configure your cluster in order
# to then run kubeadm ugrade
kubeadm config

# makage tokens for kubeadm join
kubeadm token

# revert changes made on this host (control plan or worker node)
# by the kubeadm init or kubeadm join commands
kubeadm reset

# manage Kubernetes Certificates
kubeadm certs

# manage kubeconfigs
kubeadm kubeconfig

# print kubeadm version
kubeadm version

# preview future features for community feedback
kubeadmn alpha
```

## Setting up with kubeadm init

- https://kubernetes.io/docs/setup/production-environment/tools/kubeadm/create-cluster-kubeadm/

## kubeadm init gotchas

- https://stackoverflow.com/questions/49112336/container-runtime-network-not-ready-cni-config-uninitialized

- run `kubeadm reset` if you already created a previous cluster
- remove `.kube` from home dir
- stop the `kubelet` with `sytemctl`
- disable `swap` permanently on the machine, especially if rebooting the Linux system
- install your pod network addon per instructions for your container runtime
- follow post initialization steps provided by kubeadm
- tain the nodes to enable the scheduler
  `kubectl taint nodes --all node-role.kubernetes.io/master-`

If installing behind a proxy, read this:
- https://stackoverflow.com/questions/45580788/how-to-install-kubernetes-cluster-behind-proxy-with-kubeadm
