## kind

- `kind` is a tool for running local `k`ubernetes clusters `in` `d`ocker.
[Quick Start](https://kind.sigs.k8s.io/docs/user/quick-start)
- [Quick Start](https://kind.sigs.k8s.io/docs/user/quick-start#creating-a-cluster)

## Using Kind

In order for `kind` to run clusters, `Docker` must be running on your machine.

```bash
# list existing clusters
kind get clusters
# create a cluster
time kind create cluster --name my-clusta
# Thanks for using kind! 😊
# kind create cluster  5.00s user 3.09s system 15% cpu 51.332 total
# named
time kind create cluster --name 2021-june-4
# user the cluster
# specify the name of the cluster as the --context flag
# to the kubectl cli
kubectl cluster-info --context kind-kind
# delete the cluster
kind delete cluster --name 2021-june-4
```
