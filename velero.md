# Backup & Restore with Velero

```bash
# create a kind cluster to use kubernetes 1.17 (or some version)
kind create cluster --name velero --image kindest/node:v1.17.11

# there will be a single node cluster
kubectl get nodes
# {clustername}-control-plane Read master 15m v1.17.11
# velero-control-plane Read master 15m v1.17.11
```

Create some workloads on the cluster

```bash
# have a deployment, services, etc.
# do a hello world or something.
kubectl create -f ./directory/of/stuff
```

## Use Case 1: Setup an Automatic Daily Backup

Run an `alpine` container inside `docker` but with some local filesystems `mounted`
as this makes for a clean environment to work in.  We will stilll talk to the `kind` cluster,
which is why we are mounting `${HOME}/root` as (in this scenerio) this is where the `.kubeconfig`
file is located.

Refresh on [docker run](https://docs.docker.com/engine/reference/commandline/run/) flags.

```bash
# long form
docker run --interactive --tty --rm --volume ${HOME}:/root/ --volume ${PWD}:/work --workdir /work --network host alpine sh
# short form
docker run -it --rm -v ${HOME}:/root/ -v ${PWD}:/work -w /work --net host alpine sh
# now inside, lets install a few tools
# curl will be used to download velero and other needed tools
apk add --no-cache curl nano vim
# download and install kubectl with curl
curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
chmod +x ./kubectl
mv ./kubectl /usr/local/bin/kubectl
# worky?
kubectl get nodes
# then setup the kube editor
export KUBE_EDITOR="nano" # or vim, or whatevs
```

Now, to install `velero` itself

Go to the [github velero releases page](https://github.com/vmware-tanzu/velero/releases)
Pick a binary to download, such as [linux amd64](https://github.com/vmware-tanzu/velero/releases/download/v1.8.1/velero-v1.8.1-linux-amd64.tar.gz)
And install it, similar to above `kubectl`:

```bash
# curl it and drop the tar in tmp
curl --location --output /tmp/velero.tar.gz https://github.com/vmware-tanzu/velero/releases/download/v1.8.1/velero-v1.8.1-linux-amd64.tar.gz
curl -L -o /tmp/velero.tar.gz https://github.com/vmware-tanzu/velero/releases/download/v1.8.1/velero-v1.8.1-linux-amd64.tar.gz
# extract it
tar --directory /tmp --extract --verbose --file /tmp/velero.tar.gz
tar -C /tmp -xvf /tmp/velero.tar.gz
# move it
mv /tmp/velero-v1.8.1-linux-amd64 /usr/local/bin/velero
chmod +x /usr/local/bin/velero
# worky?
velero --help
```

SKIPPING AHEAD....

## Velero with Minio for Quick eval

Velero with Minio (aws compat style) w/o PVs (just uses space on the Node for storage)

- [minio quick eval](https://velero.io/docs/main/contributions/minio/)
- [minio deployment example](https://github.com/vmware-tanzu/velero/blob/main/examples/minio/00-minio-deployment.yaml)
    - note that i needed to update to use a proxy to avoid rateliminting

```yaml
# for both the job and the deployment, this is necessary
# Copyright 2017 the Velero contributors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

---
apiVersion: v1
kind: Namespace
metadata:
  name: velero

---
apiVersion: apps/v1
kind: Deployment
metadata:
  namespace: velero
  name: minio
  labels:
    component: minio
spec:
  strategy:
    type: Recreate
  selector:
    matchLabels:
      component: minio
  template:
    metadata:
      labels:
        component: minio
    spec:
      volumes:
      - name: storage
        emptyDir: {}
      - name: config
        emptyDir: {}
      containers:
      - name: minio
        # github rate limiting
        # image: minio/minio:latest
        image: harbor-repo.vmware.com/dockerhub-proxy-cache/minio/minio:latest
        imagePullPolicy: IfNotPresent
        args:
        - server
        - /storage
        - --config-dir=/config
        env:
        - name: MINIO_ACCESS_KEY
          value: "minio"
        - name: MINIO_SECRET_KEY
          value: "minio123"
        ports:
        - containerPort: 9000
        volumeMounts:
        - name: storage
          mountPath: "/storage"
        - name: config
          mountPath: "/config"

---
apiVersion: v1
kind: Service
metadata:
  namespace: velero
  name: minio
  labels:
    component: minio
spec:
  # ClusterIP is recommended for production environments.
  # Change to NodePort if needed per documentation,
  # but only if you run Minio in a test/trial environment, for example with Minikube.
  type: ClusterIP
  ports:
    - port: 9000
      targetPort: 9000
      protocol: TCP
  selector:
    component: minio

---
apiVersion: batch/v1
kind: Job
metadata:
  namespace: velero
  name: minio-setup
  labels:
    component: minio
spec:
  template:
    metadata:
      name: minio-setup
    spec:
      restartPolicy: OnFailure
      volumes:
      - name: config
        emptyDir: {}
      containers:
      - name: mc
        # github rate limiting
        # image: minio/mc:latest
        image: harbor-repo.vmware.com/dockerhub-proxy-cache/minio/mc:latest
        imagePullPolicy: IfNotPresent
        command:
        - /bin/sh
        - -c
        - "mc --config-dir=/config config host add velero http://minio:9000 minio minio123 && mc --config-dir=/config mb -p velero/velero"
        volumeMounts:
        - name: config
          mountPath: "/config"
```

```bash
# ./velero-stuff/install.sh
# ./velero-stuff is where i put my cred file
#   which error'd, as velero expected:
#   An error occurred: open /home/kubo/credentials-velero: no such file or directory

#!/bin/bash
velero install \
    --provider aws \
    --plugins velero/velero-plugin-for-aws:v1.2.1 \
    --bucket velero \
    --secret-file ./credentials-velero \
    --use-volume-snapshots=false \
    --backup-location-config region=minio,s3ForcePathStyle="true",s3Url=http://minio.velero.svc:9000 \
    --wait

velero install --provider aws --plugins velero/velero-plugin-for-aws:v1.2.1 --bucket velero --secret-file ./credentials-velero --use-volume-snapshots=false --backup-location-config region=minio,s3ForcePathStyle="true",s3Url=http://minio.velero.svc:9000 --wait
```

This will create the velero deployment, which may have some proxy issues as well:

```yaml
# take a look @ the deployment...
# kubectl get deployment velero -n velero -o yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  annotations:
    deployment.kubernetes.io/revision: "1"
  labels:
    component: velero
  name: velero
  namespace: velero
spec:
  progressDeadlineSeconds: 600
  replicas: 1
  revisionHistoryLimit: 10
  selector:
    matchLabels:
      deploy: velero
  strategy:
    rollingUpdate:
      maxSurge: 25%
      maxUnavailable: 25%
    type: RollingUpdate
  template:
    metadata:
      annotations:
        prometheus.io/path: /metrics
        prometheus.io/port: "8085"
        prometheus.io/scrape: "true"
      creationTimestamp: null
      labels:
        component: velero
        deploy: velero
    spec:
      containers:
      - args:
        - server
        - --features=
        command:
        - /velero
        env:
        - name: VELERO_SCRATCH_DIR
          value: /scratch
        - name: VELERO_NAMESPACE
          valueFrom:
            fieldRef:
              apiVersion: v1
              fieldPath: metadata.namespace
        - name: LD_LIBRARY_PATH
          value: /plugins
        - name: GOOGLE_APPLICATION_CREDENTIALS
          value: /credentials/cloud
        - name: AWS_SHARED_CREDENTIALS_FILE
          value: /credentials/cloud
        - name: AZURE_CREDENTIALS_FILE
          value: /credentials/cloud
        - name: ALIBABA_CLOUD_CREDENTIALS_FILE
          value: /credentials/cloud
        # this image should be fine
        image: projects.registry.vmware.com/tkg/velero/velero:v1.7.0_vmware.1
        imagePullPolicy: IfNotPresent
        name: velero
        ports:
        - containerPort: 8085
          name: metrics
          protocol: TCP
        resources:
          limits:
            cpu: "1"
            memory: 512Mi
          requests:
            cpu: 500m
            memory: 128Mi
        terminationMessagePath: /dev/termination-log
        terminationMessagePolicy: File
        volumeMounts:
        - mountPath: /plugins
          name: plugins
        - mountPath: /scratch
          name: scratch
        - mountPath: /credentials
          name: cloud-credentials
      dnsPolicy: ClusterFirst
      initContainers:
      - name: velero-velero-plugin-for-aws
        # need to update the init container image with the proxy:
        # image: velero/velero-plugin-for-aws:v1.2.1
        image: harbor-repo.vmware.com/dockerhub-proxy-cache/velero/velero-plugin-for-aws:v1.2.1
        # harbor-repo.vmware.com/dockerhub-proxy-cache/minio/mc:latest
        imagePullPolicy: IfNotPresent
        resources: {}
        terminationMessagePath: /dev/termination-log
        terminationMessagePolicy: File
        volumeMounts:
        - mountPath: /target
          name: plugins
      restartPolicy: Always
      schedulerName: default-scheduler
      securityContext: {}
      serviceAccount: velero
      serviceAccountName: velero
      terminationGracePeriodSeconds: 30
      volumes:
      - emptyDir: {}
        name: plugins
      - emptyDir: {}
        name: scratch
      - name: cloud-credentials
        secret:
          defaultMode: 420
          secretName: cloud-credentials
status:
  conditions:
  - lastTransitionTime: "2022-06-15T21:33:35Z"
    lastUpdateTime: "2022-06-15T21:33:35Z"
    message: Deployment does not have minimum availability.
    reason: MinimumReplicasUnavailable
    status: "False"
    type: Available
  - lastTransitionTime: "2022-06-15T21:33:35Z"
    lastUpdateTime: "2022-06-15T21:33:35Z"
    message: ReplicaSet "velero-58cf45c794" is progressing.
    reason: ReplicaSetUpdated
    status: "True"
    type: Progressing
  observedGeneration: 1
  replicas: 1
  unavailableReplicas: 1
  updatedReplicas: 1
```
