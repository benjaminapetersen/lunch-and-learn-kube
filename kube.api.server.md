# Kube API Server


The Kube API Server can be seen as pods if you `kubectl get pods -A` and look for them, but there is no
Deployment or ReplicationController or anything associated with these pods. So, how does it work?
(This does make sense given the API server is where Deployments and other resources are sent...)

Basically this is how it works:
- `kubelet` has a directory with YAML manifests in it.
- It makes sure there are containers running that matches to manifests.
- The kubelet also makes a mirror pod resource in the kube API, but that is just for viewing purposes.
- The cluster itself is responsible for its API server config
- So, you shouldn't really but you could:
  - `ssh` into a Master Node and change the API Server YAML.
- you dont need master nodes at all
- that wording implies running kubelets
- you can just have a computer and have it run the api server binary
- its 100% opaque from the cluster's perspective
- i.e. its not a master node
- its not a node at all
