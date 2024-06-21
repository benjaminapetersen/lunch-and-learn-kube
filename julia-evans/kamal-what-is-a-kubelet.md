# What even is a Kubelet?

https://kamalmarhubi.com/blog/2015/08/27/what-even-is-a-kubelet/

- `kubelet` is the lowest level component in Kubernetes.
- It is responsible for what is running on an idnvidual machine.
- It is like `supervisord`, a process watcher, but focused on containers
    - Supervisord http://supervisord.org/
        - Supervisor is a client/server system that allows its users to monitor 
          and control a number of processes on UNIX-like operating systems.
- It has one job:
    - given a set of containers to run, make sure they are running

## Kubelets run pods

- `pods` are the unit of execution in Kubernetes.
- Pods are a collection of containers that share some resources
    - A single `IP`
    - Volumes

Example pod:
- A web server pod may have:
    - a container for the server itself 
    - a container that tails the logs & ships them off to your 
      logging or metrics infrastructure

Pods are defined in `YAML`:

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: nginx
spec:
  containers:
  # a single container in the pod
  - name: nginx
    # a Docker image name
    # by default, the entrypoint defined
    # in the image is what will run. In
    # this case, that would be the nginx server.
    image: nginx
    ports:
    # expose the port from the nginx container
    # this allows the nginx server to connect to 
    # the pod's IP
    - containerPort: 80
```
Adding a second container could look something like this:

```yaml
 1apiVersion: v1
 2kind: Pod
 3metadata:
 4  name: nginx
 5spec:
 6  containers:
    # first container, nginx again
 7  - name: nginx
 8    image: nginx
 9    ports:
10    - containerPort: 80
      # mount a shared volume
11    volumeMounts:
      # the same volume can be mounted at one path in this container,
      # but at another path in the other container.
      # nginx writes logs to /var/log/nginx so this makes sense.
12    - mountPath: /var/log/nginx
13      name: nginx-logs
    # second container, some silly thing to truncate logs
14  - name: log-truncator
      # busybox is a tiny image.
      # its a linux command line environment, designed 
      # especially for running little commands like this.
      # it has no package manager (by design)
15    image: busybox
      # this runs a command
16    command:
17    - /bin/sh
18    args: [-c, 'while true; do cat /dev/null > /logdir/access.log; sleep 10; done']
      # mount a shared volume
19    volumeMounts:
20    - mountPath: /logdir
21      name: nginx-logs
    # a volume to share is defined here.
    # it gets a name, and a type.
    # an emptyDir 
    # - begins its life empty
    # - is cleaned up when the pod exists
    # - but is persistent across container restarts
    #   - can be used while the containers live 
22  volumes:
23  - name: nginx-logs
24    emptyDir: {}
```

## Kubelet running pods

Kubelet finds pods via:
- a directory to poll for pod manifests
- a URL to poll and download pod manifests
- a kubernetes API server

By default, the watched directory polls the directory every 20 seconds looking for manifests.
The kubelet will both launch new pods (if it finds new manifests) and kill 
pods for manifests that are removed.

Kubelet supports several container runtimes
- Docker
- rkt
- others.
    - https://kubernetes.io/docs/setup/production-environment/container-runtimes/

