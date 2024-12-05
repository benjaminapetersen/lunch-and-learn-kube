# Controllers

- Many controllers are built-in to the Kube Controller Manager deployed
  as part of all Kubernetes clusters, each with a specific purpose.

Six examples of controllers:

- `Node Controller`: Responsible for managing Worker Nodes.  Monitors
  new Nodes connecting to hte cluster, validates the Node's health status
  based on the metrics reported by the Node's Kubelet component, and
  update the Node's .status field.  If a Kubelet stops posting health
  checks to the API Server, the Node Controller will be responsible for
  triggering Pod eviction from the missing Node before removing the Node
  from the cluster.
- `Deployment Controller`: Responsible for managing Deployment objects
  and creating/modifying ReplicaSet objects.
- `ReplicaSet Controller`: Responsible for creating/modifying Pods based
  on the ReplicaSet configuration.
- `Service Controller`: Responsible for configuring ClusterIP, NodePort
  and LoadBalancer configuration on Service objects
- `CronJob Controller`: Responsible for creating Job objects based on the
  Cron schedule defined in CronJob objects
- `StatefulSet Controller`: responsible for creating pods in a guaranteed
  order with a sticky identity


## Controller Design Principles
- `Desired State`: delegate impl details to controllers by declaring desired state
  rather than the imperative steps to achieve state.  controllers constantly
  reconcile current state to desired state.
- `Fault Tolerance`: multiple replicas of controllers on master nodes allows
  for redundancy.  When one copy fails another is there to resume operation.
- `Self-healing`: Drift in cluster configuration will be detected and mitigated
  due to the nature of controllers constantly comparing and reconciling current
  state with desired state.
- `Atomicity`: Distributed systems should be designed to handle failures of
  any component.  Controllers handle interruptions during reconciliation
  activities.  Controllers can pick up reconciliation steps from any point
  where an interruption occurred, wether or not an action was completed.

## Controller Overview

- [Controller Overview](https://able8.medium.com/kubernetes-controllers-overview-b6ec086c1fb)
- A controller is a control loop that watches the state of your cluster
  then makes changes to move the current cluster state closer to the
  desired state.

```go
// repeats forever
// with some "watches" and triggers
// built in, as well as exponential
// backoff and other features to
// avoid flurry of chaos when things
// aren't in a good place.
for {
  desired := getDesiredState()
  current := getCurrentState()
  makeChanges(desired, current)
}
```

- Desired state vs current state

```yaml
spec: specify the desired state
status: report the current state
```
