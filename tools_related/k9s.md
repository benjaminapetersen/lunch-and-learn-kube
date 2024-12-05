# k9s

`k9s` is a terminal based UI to interact with your Kubernetes clusters.

# Commands

- [k9s Commands](https://k9scli.io/topics/commands/)

```bash
k9s help
k9s info 
# only one namespace
k9s -n <namespace>
# pod view
k9s -c <pod-name>
# non-default kubeconfig context
k9s --context <context>
# readonly mode
# modification commands are disabled
k9s --readonly
```

## Key Bindings

Is it vim-like?  Dunno! Let's find out.

```yaml
# back up and out of whatever you are down in
# xray, for example, lets you drill down quite a bit
# so this steps up and back out 
esc
# single letter keys for doing things:
d,v,e,l # describe, view, edit, logs,
# show active keybaord mkemonics
# ie, a convenient help menu
?
# quit
:q, ctrl-c
# view a specific kube resource
:pod
:deploy
# view in a namespace
:pod <namespace>
:deploy <namespace>
# with a filter
# basic regex supported
# easily matches a variety of things with the 
# or | operator
:pod /some-filter
:pod /foo|bar|fb3M8
# in a namespace and with a filter
:pod foobar /baz
# in a namespace, with a label
:pod foobar app=demo
# in a ns, with a label, different context
# the @ context switch is a full context switch
# that persists until switching back
:pod foobar app=demo @my-kind-cluster
# dunno what pulses is
:pulses
# nifty view of things
:xray
```
Further, you can:

`<click>` into a pod to view it's `logs`, which 
does `watch` by default.