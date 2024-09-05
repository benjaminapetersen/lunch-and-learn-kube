# The Flow of Kubernetes Components

The User
- uses `kubectl` to do a CRUD action
- `kubectl` sends an API request to the `api-server`
- the `api-server` sends requests to the `kube-controller-manager`
- the `kube-scheduler`
  - maintains a scheduled `podQueue`
  - listens to the API server

When the user creates a pod
- Pod metadata is written to the `api server`
  - Which writes the metadata to `etcd`
- The scheduler listens to pod status through an `informer`
  - When the pod is added
  - The scheduler adds it to the podQueue
- The main process
  - continually extracts Pods from podQueue
    - assigns nodes to Pods that do not have a node

Scheduling consists of
- Filter matching nodes and prioritize nodes
  - based on pod configuration
  - to score nodes
  - and select nodes with the highest score for the pod

Once a node is assigned
- invoke the binding pod interface
  - of the api-server
    - set pod.spec.nodeName to the assigned node

The `kubelet` on the node
- listens to the api-server
- if it finds a new pod scheduled to its node
- it invokes the local container runtime to run the container
  - docker
  - containerd
  - CRI-O

If the scheduler fails to schedule a pod
- there are mechanisms that can be configured
  - such as priority and preemption
    - which may cause low prority pods to be deleted in favor of high priority pods
