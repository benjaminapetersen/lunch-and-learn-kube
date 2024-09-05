# Kubernetes Controller Manager

- `https://kubernetes.io/docs/reference/command-line-tools-reference/`kube-controller-manager/

- `Controller manager is a `daemon``
- `embeds the core control loops shipped with kubernetes`
- `kube controllers`
  - `watch the shared state of the cluster`
    - `through the apiserver`
  - `make changes to`
    - `move the current state`
    - `toward the desired state`


Example controllers

- `Replication controller`
- `Endpoints controller`
- `Namespace controller`
- `Serviceaccounts controller`

The controller directory in kubernetes repo:

- `https://github.com/kubernetes/kubernetes/tree/master/pkg/controller`

Many other packages listed in `pkg/controller`:

- `apis`
- `bootstrap`
- `certificates`
- `clusterroleaggregation`
- `controller_ref_manager.go`
- `controller_ref_manager_test.go`
- `controller_utils.go`
- `controller_utils_test.go`
- `cronjob`
- `daemon`
- `deployment`
- `disruption`
- `doc.go`
- `endpoint`
- `endpointslice`
- `endpointslicemirroring`
- `garbagecollector`
- `history`
- `job`
- `lookup_cache.go`
- `namespace`
- `nodeipam`
- `nodelifecycle`
- `podautoscaler`
- `podgc`
- `replicaset`
- `replication`
- `resourcequota`
- `serviceaccount`
- `statefulset`
- `storageversiongc`
- `testutil`
- `ttl`
- `ttlafterfinished`
- `util`
- `volume`

## A description of each controller listed above

Initially reading the `kubernetes/pkg/controller/<ctrl>/doc.go`, then
doing further investigation:

- `apis`
- `bootstrap`
  - Provides automatic processes necessary for bootstrapping
  - Includes managing and expiring tokens
  - Includes signing well known configmaps with tokens
  - Small
    - 8 files
- `certificates`
  - Manage certificates in the cluster
  - Medium
    - Contains several subdirectories
- `clusterroleaggregation`
  - Combines cluster roles
  - Small
    - 2 files
- `cronjob`
  - Manages cronjobs
  - Medium
    - several files
    - several subdirectories
- `daemon`
  - Watch & syncronize daemons
  - Medium
    - several files
    - several subdirectories
- `deployment`
  - Watch & sync deployments
  - Medium
    - several files
    - several subdirectories
- `disruption`
  - seems like disruption budgets
  - small
    - 2 files
- `endpoint`
- `endpointslice`
- `endpointslicemirroring`
- `garbagecollector`
- `history`
- `job`
- `namespace`
- `nodeipam`
- `nodelifecycle`
- `podautoscaler`
- `podgc`
- `replicaset`
- `replication`
- `resourcequota`
- `serviceaccount`
- `statefulset`
- `storageversiongc`
- `testutil`
- `ttl`
- `ttlafterfinished`
- `util`
- `volume`
- `doc.go`
- `controller_ref_manager.go`
- `controller_ref_manager_test.go`
- `controller_utils.go`
- `controller_utils_test.go`
- `lookup_cache.go`
