# A Docker Reference Sheet

<!--
  NOTE: all the flags in this doc are long form.  Long form is much more intuitive, once the
  available flags are known, it seems much less of a reach to swap to short form.
-->

Resources used:

- [Docker Up & running](http://shop.oreilly.com/product/0636920036142.do)


## Some basic docker commands

### docker version

what version? reports on both the client & the server.

```bash
$ docker version
# Client:
#  Version:      1.12.1
#  API version:  1.24
#  Go version:   go1.7.1
#  Git commit:   6f9534c
#  Built:        Thu Sep  8 10:31:18 2016
#  OS/Arch:      darwin/amd64

# Server:
#  Version:      1.12.1
#  API version:  1.24
#  Go version:   go1.6.3
#  Git commit:   23cf638
#  Built:        Thu Aug 18 17:52:38 2016
#  OS/Arch:      linux/amd64

```

### docker info

display system wide information.  will show warnings if your system lacks support for a docker feature

### docker exec

Run a command in a running container

```bash
# example:
# docker exec -it <container-name> <command>
docker exec -nteractive -tty origin bash
```

### docker build

Build a docker images with the `Dockerfile` in the current directory. Name the image & tag it.

```bash
  # example:
  # docker build --tag example/docker-node-hello:latest .
  # you can use multiple tags:
  # docker build --tag example/docker-node-hello:v0.0.1 example/docker-node-hello:latest .
  docker build --tag <name-of-image-to-build:with-optional-tag> .
```

### docker run

A wrapper around `docker create` and `docker start`. Creates a container from an underlying image & executes it.

Use `docker images` if you can't remember your container name (likely just created it with the `docker build` command),
and then create a container & run it based on that image:

```bash
  # run is really a convenience command,
  # it wraps `docker create` and `docker start`
  #
  # run a container
  docker run --detach --publish 8080:8080 <name-of-image-to-run:with-tag>
  # never run as `root`!
  # by default, docker runs containers as root
  # this is... well, not good
  docker run --user not-root
  # add env vars
  docker run --user batman --env WHO="Batman" --env WHAT="Vigilante" --publish 8080:8080 <name-of-image-to-run>
  # add labels
  # note: labels are automatically inherited from a parent image
  # note: you can use labels to apply metadata that may be specific to a single container
  docker run --detach --user wolverine --name="Wolverine" --label super.power=claws --label regenerates --label deployer=logan --label author=xmen
  # name it, otherwise docker will assign a random name
  # build from an adjective+famous-person-name
  docker run --user chuck --name="chuck-norris"
  # set a hostname
  docker run --hostname=mycontainer.example.com
  # moount a volume from the docker host
  # note that this is general not advised, as it
  # ties your container to a specific docker host
  # for persistent state
  # --volume <host-volume>:<container-volume-name>
  docker run --volume /mnt/session_data:/data ubuntu:latest /bin/bash
  docker run --rm --interactive -t --volume /mnt/session_data:/data ubuntu:latest /bin/bash
  # an OSX volume mount
  docker run --rm --interactive -t --volume ~/Desktop:/Desktop r-base /bin/bash
  # use readonly to ensure that the container's own root filesystem
  # cannot be written.  this is handy for avoiding unexpected things
  # like logfiles from filling up your container's filesystem
  docker run --rm --interactive -t --read-only=true --volume ~/Desktop:/Desktop r-base /bin/bash
  # control cpu usage
  docker run --rm -it progrium/stress --cpu 2 --io 1 --vm 2 --vm-bytes 128M --timeout 10s

```

### docker images

See a list of docker images

```bash
docker images
# show only the ids
docker images --quiet
```

### docker rmi

remove a container image

```bash
# list images
docker images
# remove an image
docker rmi <image-id>
# you cannot remove the image associated with a running container,
# the container must be stopped first
#
# remove all images:
docker rmi $(docker images --quiet -)
# remove all untagged images
docker rmi $(docker images --quiet --filter "dangling=true")
```

### docker ps

See a list of running containers (docker processes)

```bash
# shows only running containers
docker ps
# shows all containers, rather than just the running containers
# ps defaults to show just running containers
docker ps --all
# filter
docker ps --all --filter label=deployer=conroy.cage
# follow up with a `docker inspect <container-id>` to see all its labels & other goodies in a JSON object
```

### docker start

Start a container

```bash
# see a list of all containers on system
docker ps -a
# pick out the CONTAINER ID of the container you want to run
docker start <container-id>
# ensure its running
docker ps
# ssh into the container and mess around
docker exec --interactive --tty 8f747dcb24bf /bin/bash
```

### docker stop

Stop a running container

```bash
# get the id from docker ps
docker stop <container-id>
# set a time limit
docker stop --time 25 <container-id>
```

### docker kill

forcefully stop a container

```bash
# end it
docker kill <container-id>
# or... just kidding
# docker kill can send any signal, just like linux kill commands
# (which, seems totally messed up, imho.)
docker kill --signal=USR1 <container-id>
# basically ...
docker kill --signal=NOPEJK <container-id>
# and then kill will not kill. if the running container knows what to do with
# the signal, it will respond, else it will do nothing
```

### docker rm

remove one or more containers.  Pass `-v` to remove volumes as well, else containers can waste disk space!

```bash
docker rm -v <container-id>
# remove containers that exited with nonzero state
docker rm $(docker ps --all --quiet --filter 'exited!=0')
```

### docker login

Login to Docker Hub to push your images

```bash
# docker hub by default
docker login
# or another registry
docker login <registry.somewhereelse.com>
# view a list of registries you *may* be logged into
# note that creds are stored elsewhere, this will just dump urls
less ~/.docker/config.json | jq '.auths'
# {
#   "harbor-repo.vmware.com": {},
#   "https://index.docker.io/v1/": {},
#   "quay.io": {}
# }
```

### docker logout

Grabs you coffee & a bagel.  Or does the obvious.

### docker pause

Pauses a container. leverages the `cgroups` freezer. containers are not aware of paused state (process cannot be scheduled),
unlike `stop` which sends the SIGSTOP signal.

### docker unpause

## Env vars

```bash
# Get the Docker host IP address
echo $DOCKER_HOST

```

### docker inspect

Gets a JSON representation of a container. Useful for looking at the container's labels, network settings, config, etc.

```bash
# find a running container
docker ps
# now inspect it
docker inspect <container-id>
# output will be something like the following:
# NOTE: Args: [] is especially useful.
#       containers are typically configured via cli args
#       seeing what args were passed when created is essential for debugging
# [{
#   Args: [],
#   Config: {},
#   Created: '2017-05-26T02:54:35.701369716Z',
#   HostConfig: {},
#   HostnamePath: '/var/lib/docker/containers/bab5433..../hostname',
#   HostPath: '/var/lib/docker/containers/bab5433..../hosts',
#   Id: 'bab5433....',
#   Image: 'sha256:ebcd9e...',
#   LogPath: '/var/lib/docker/containers/bab5433....-json.log',
#   MountLabel: '',
#   Mounts: [],
#   Name: '/zen_mcnulty',
#   NetworkSettings: {
#     Bridge: '',
#     IPAddress: '172.17.0.2',
#     MacAddress: '02:42:ac:11:00:02',
#     Networks: {
#       Gateway: '172.17.0.1',
#       MacAddress: '02:42:ac:11:00:02',
#       NetworkID: '6ea9cb5495c34165675b1e0a6183f26aa9657f2d350579c962de66e2deab6455'
#     },
#   },
#   Path: '/bin/bash',
#   ProcessLabel: '',
#   ResolvConfPath: '/var/lib/docker/containers/bab5433..../resolv.conf',
#   RestartCount: 0,
#   State: {
#     Dead: false
#     Error: ''
#     00MKilled: false
#     Pid: 12534,
#     Restarting: false,
#     Running: true,
#     StartedAt: '2017-05-26T02:54:36.411379089Z'
#     Status: 'running'
#   }
# }]

```

### docker tag

Tag a built docker image.  If the image is not yours, you need to namespace it if you want to push it to a registry like
docker hub:

```bash
# list your images
$ docker images
REPOSITORY                                TAG                 IMAGE ID            CREATED             SIZE
example/docker-node-hello                 latest              b9c6ffc36d66        44 minutes ago      648 MB
# tag the image by id:
# docker tag b9c6ffc36d66 benjaminapetersen/docker-node-hello
$ docker tag b9c6ffc36d66 <your-username>/docker-node-hello
# be sure you login!
$ docker login
# push it
$ docker push <your-username>/docker-node-hello
```


### docker push

Push a tagged image to a registry.  See `docker tag` above for a good workflow.

```bash

```

### docker pull

Pull down an image you have pushed to a registry (or any image, you don't have to own it).
This command only pulls down the layers that have changed since the last pull.
Be sure to remove the local copy, otherwise pull will do nothing:

```bash
# what we got locally?
$ docker images
REPOSITORY                                TAG                 IMAGE ID            CREATED             SIZE
benjaminapetersen/docker-node-hello       latest              b9c6ffc36d66        50 minutes ago      648 MB
# eliminate the local
docker rmi -f b9c6ffc36d66
# pull it!
docker pull benjaminapetersen/docker-node-hello
# NOTE: you can also just do `docker run`, which will automatically pull for you:
docker run benjaminapetersen/docker-node-hello
```

## Testing your containers

### curl

Curl is handy for making a quick cli request & verifying your app within the
container is responsive:

```bash
# curl -i <endpoint>
# example:
# be sure to hit the public PORT that is being mapped to the internal!
# IE, when you did docker run --publish 49160:8080,
# curl (and external traffic) needs to hit the first port
curl -i localhost:49160
```

## Commands to run within containers

For debugging, etc, the following commands can prove to be extremely helpful:

### ps

Check for running processes.  This is enlightening as containers are light weight.  There will not be much running!

```bash
ps -ef
```


### install

Within a container you will not have much.  Fortunately, you can install packages (that will only survive as long as
the container itself).  `apt-get` will work in ubuntu based containers.

```bash
apt-get install cowsay
yum install cowsay
rpm -ivh <package>.rpm
dpkg -i xargs.deb

```


### mount

Running `mount` from within a container will result in many ~bind mounts~ for the container.

One notable mount is the mount point for `/etc/hostname`.  Docker prepares a hostname file for each container, but it typically
contains the container's ID only, not a fully qualified domain name.  A container contains an `/etc/hostname` file, running
`mount` will connect it to Docker's hostname file for the container.  With an additional command we can hook up a
domain name.

``` bash
# first run a container
$ docker run -rm -rq ubuntu:latest /bin/#!/usr/bin/env bash
# then run the `mount` command:
$ mount
# should be quite an output:
 none on / type aufs (rw,relatime,si=b220f0ecb9dbc932,dio,dirperm1)
 proc on /proc type proc (rw,nosuid,nodev,noexec,relatime)
# ...
 /dev/sda1 on /etc/resolv.conf type ext4 (rw,relatime,data=ordered)
 /dev/sda1 on /etc/hostname type ext4 (rw,relatime,data=ordered)  # <-- aha! this is especially interesting
# look for something specific, such as hostname:
mount | grep hostname
# /dev/vda1 on /etc/hostname type ext4 (rw,relatime,data=ordered)
#
# so we can set a specific hostname when we run a container like this:
$ docker run --hostname="mycontainer.example.com" ubuntu:latest /bin/bash
```

## hostname

Get or set the hostname for the container. Running `hostname -f` should retrieve the hostname.

`hostname --file`

## nsenter

From the `util-linux` package, nsenter (Namespace enter) allows you to enter a linux namespace. This function can be used to
get into a docker container as root on the server.   Use `docker exec` first, but if really in a bind, `nsenter` is the big gun.

the following will pull an image that will run a container that installs nsenter in your docker server:

```bash
# cleva
docker run --rm -v /usr/local/bin:/target jpetazzo/nsenter
```


## Recipes

A set of useful commands gathered in one place.  Some of these borrowed from [here](https://www.digitalocean.com/community/tutorials/how-to-remove-docker-images-containers-and-volumes).


### containers
```bash
# list containers using multiple filters
docker ps --all --filter status=exited --filter status=created
# remove containers using multiple filters
docker rm $(docker ps --all --filter status=exited --filter status=created --quiet)

# list all exited containers (only id)
docker ps --all --quiet --filter status=exited
# remove exited containers (and their volumes)
docker rm $(docker ps --all --filter status=exited --quiet) --volumes

# list containers via a pattern
docker ps -a |  grep "pattern"
# remove containers via a pattern
docker ps -a | grep "pattern" | awk '{print $3}' | xargs docker rmi

# stop and remove all containers
docker stop $(docker ps --all --quiet)
docker rm $(docker ps --all --quiet)

# count running containers
docker ps --quiet | wc -l  # 13
```

### images

```bash
# list all images
docker images --all
# remove all images
docker rmi $(docker images --all --quiet)

# list dangling images
docker images --filter dangling=true
# remove dangling images
docker rmi $(docker images --filter dangling=true --quiet)
```

### volumes

```bash
# list volumes
docker volume ls
# remove volume
docker rm <volume_name>
# remove a container and its volume
docker rm -v <container_name>
# list dangling volumes
docker volume ls -f dangling=true
# remove dangling volumes (this is a good cleanup step, if forgot previously)
docker volume rm $(docker volume ls -f dangling=true -q)

```

### docker daemon options

```bash
# Mo's args to the docker daemon
# docker version 1.13.1
ExecStart=/usr/bin/dockerd --exec-opt native.cgroupdriver=systemd --selinux-enabled --log-driver=json-file --insecure-registry 172.30.0.0/16 --insecure-registry=ci.dev.openshift.redhat.com:5000

```

### systemctl

```bash
# restart docker in rhel
sudo systemctl restart docker
```