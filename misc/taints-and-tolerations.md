# Taints and Tolerations Scratch

 # taints and tolerations vs node affinity

 - `taints and tolerations`
    - no magnetic affect
    - do not guarantee where a workload will deploy
    - tainted nodes only evict workloads that do not tolerate the taint
    - workloads may still end up on completely different nodes
      - workload tolerates taint "foo", but may still deploy on "bar"
      - taints and tolerations are not like "magnets"
- `node affinity`
  - has a magnetic affect
    - workloads are drawn to certain nodes
  - TODO: go learn some more about this!


# a subset of taints on a node

```bash
kubectl get nodes -o yaml
```

```yaml
apiVersion: v1
items:
- apiVersion: v1
  kind: Node
  metadata:
    annotations:
      cluster.x-k8s.io/cluster-name: tkg-mgmt-vc
      cluster.x-k8s.io/cluster-namespace: tkg-system
      cluster.x-k8s.io/machine: tkg-mgmt-vc-control-plane-6qbt4
      cluster.x-k8s.io/owner-kind: KubeadmControlPlane
      cluster.x-k8s.io/owner-name: tkg-mgmt-vc-control-plane
      csi.volume.kubernetes.io/nodeid: '{"csi.vsphere.vmware.com":"tkg-mgmt-vc-control-plane-6qbt4"}'
      kubeadm.alpha.kubernetes.io/cri-socket: /var/run/containerd/containerd.sock
      node.alpha.kubernetes.io/ttl: "0"
      volumes.kubernetes.io/controller-managed-attach-detach: "true"
    creationTimestamp: "2022-01-03T20:09:53Z"
    labels:
      beta.kubernetes.io/arch: amd64
      beta.kubernetes.io/instance-type: vsphere-vm.cpu-2.mem-8gb.os-photon
      beta.kubernetes.io/os: linux
      kubernetes.io/arch: amd64
      kubernetes.io/hostname: tkg-mgmt-vc-control-plane-6qbt4
      kubernetes.io/os: linux
      node-role.kubernetes.io/control-plane: ""
      node-role.kubernetes.io/master: ""
      node.kubernetes.io/exclude-from-external-load-balancers: ""
      node.kubernetes.io/instance-type: vsphere-vm.cpu-2.mem-8gb.os-photon
    name: tkg-mgmt-vc-control-plane-6qbt4
    resourceVersion: "540689"
    uid: d346eaa3-2c0d-4fc8-ae9a-4d92e2e3ac2e
  spec:
    podCIDR: 100.96.0.0/24
    podCIDRs:
    - 100.96.0.0/24
    providerID: vsphere://4224ce84-5851-6610-8d7e-eeb61eec681c
    # TAINTS
    taints:
    # NoSchedule, PreferNoSchedule, NoExecute
    - effect: NoSchedule
      key: node-role.kubernetes.io/master
```

```bash
# example tainting a node
# this is the breakdown:
#  key=app
#  operator="="
#  value=blue
#  effect=NoSchedule
kubectl taint nodes node1 app=blue:NoSchedule
```

```yaml
# tolerations on a pod
apiVersion: v1
kind: Pod
metadata:
  name: myapp-pod
spec:
  containers:
  - name: nginx-container
    image: nginx
  # this pod will tolerate this taint on the node
  # NOTE: all values need to be encoded in double quotes
  tolerations:
  - key: "app"
    operator: "Equal"
    value: "blue"
    effect: "NoSchedule"
```


