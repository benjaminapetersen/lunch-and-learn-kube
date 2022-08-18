# Kube Builder

- [Quick Start](https://book.kubebuilder.io/quick-start.html)
- [Tutorial: Building the CronJob Controller](https://book.kubebuilder.io/cronjob-tutorial/cronjob-tutorial.html)
  - This tutorial is a rebuild of the CronJob controller with KubeBuilder


## Installation

The official Quick Start above lists out:

```bash
# download kubebuilder and install locally.
curl -L -o kubebuilder https://go.kubebuilder.io/dl/latest/$(go env GOOS)/$(go env GOARCH)
chmod +x kubebuilder && mv kubebuilder /usr/local/bin/
```
But I had to make a slight tweak due to zsh errors:

```bash
# quotes, and for some reason copy/paste inserted an additional slash
#   go env GOOS=darwin
#   go env GOARCH=amd64
curl -L -o kubebuilder  "https://go.kubebuilder.io/dl/latest/$(go env GOOS)/$(go env GOARCH)"
```
