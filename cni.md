# CNI: Container Network Interface

- CNI is
  - a specification
  - a set of libraries
  for writing plugins to configure network interfaces in Linux containers,
  along with a number of supporting plugins.

CNI is used by `containerd`, `CRI-O`,

CNI concerns itself only with network connectivity of containers and
removing allocated resources when the container is deleted.
Because of this:
- CNI has a wide range of support
- the specification is simple to implement

- https://github.com/containernetworking/cni
- https://github.com/containernetworking/plugins
    - A set of reference plugins
    - Maintained by the container networking team


## The SPEC

- https://github.com/containernetworking/cni/blob/main/SPEC.md
- A generic plugin-based networking solution for application containers on Linux

- Three terms are defined specifically:
  - `container` is a network isolation domain.
    - However, the actual isolation technology is not defined by this specification.
    - It could be a `nework namespace`, or a `virtual machine`, for example
  - `network` refers to a group of endpoints that
    - are uniquely addressable
    - and can communcate amongst each other
    - this could be
      - an individual container (per above `container` specified)
      - a machine
      - some other network device (such as a router)
    - containers can be conceptually added to or remove from one or more networks
  - `runtime` is the program responsible for executing CNI plugins
    - plugins are `binaries`
    - they are not `daemons`
  - `plugin` is a program that applies a specified network configuration
- This SPEC aims to specify:
  - the interface between `runtimes` an `plugins`
  - Keywords are used as specified in https://www.ietf.org/rfc/rfc2119.txt

## The CNI Specification Defines

1. A `format` for administrators to define network configurations
2. A `protocol` for container runtimes to make requests to network plugins
3. A `procedure` for executing plugins based on supplied configuration
4. A `procedure` for plugins to delegate functionality to other plugins
5. Data types for plugins to return their results to the runtimes


## Section 1: Network Configuration Format

- CNI defines a network configuration format for administrators
- It contains directives for
  - the container runtime
  - the plugins to consume
- At plugin execution time
  - this configuration format is interpreted
    - by the runtime
    - and transformed into a form
      - to be passed to the plugins
- In general, network configuration is intended to be static
  - Conceptually, it is "on disk"
    - though the specification does not require this

Configuration Format
- The network configuration is a JSON object:
```json
{
    "cniVersion": "string",         // currently 1.1.0
    "cniVersions": "string list",   // if support multiple versions
    "name": "string",               // network name, unique across all network configurations on a host
    "disableCheck": "boolean",      // do not all CHECK for this network config
    "disableGC": "boolean",         // do not invoke GC (garbage collection) for this network config
    "plugins": "list"               // a list of CNI plugins & config
}
```

Plugin configuration
- objects (presumably JSON) that contain additional fields
- runtimes MUST pass through these fields, unchanged, to the plugins
- there are examples of keys in the SPEC

## Section 2: Execution Protocol

- The CNI protocol is baed on
  - execution of binaries
  - by the cotnainer runtime
- CNI defines the protocol between the plugin binary and the runtime
- A CNI plugin is responsible for configuring a container's network interface

Plugins fall into two categories:
- interface plugins, which
  - create a network interface inside the container
  - ensure the contianer has connectivity
- chained plugins, which
  - adjust the configuration of an already-created interface
  - but may also need to create more interfaced to do so

The runtime will
- pass parameters to the plugin
  - via environment variables and configuration
- configuration is supplied via stdin

The plugin will
- return a `result` to `stdout` on success
- return an `error` to `stderr` if the operation fails

Configuration and results are encoded in JSON

Parameters define
- invocation-specific settings

Configuration is
- the same for any given network (with some exceptions)

The runtime must execute the plugin in the runtime's network domain
- For most cases, this means the root network namespace `dom0`

Protocol Parameters
- passed to the plugins via OS environment variables

- `CNI_COMMMAND`: ADD, DEL, CHECK, GC, VErsioN
- `CNI_CONTAINERID`: container id
- `CNI_NETNS`: the container's isolation doamin
- `CNI_IFNAME`: name of the interface created inside the container
- `CNI_ARGS`: extra args passed in by the user at invocation time
- `CNI_PATH`: list of paths to search for CNI plugin executables

The CNI Operations

- `ADD` adds container to network, or applies modifications
- `DEL` remove the container from the network, or un-apply modifications
- `CHECK` check if the container's networking is as expected
- `GC` clean up any stale resources
- `STATUS` check plugin status
- `VERSION` probe plugin version support

## Section 3: Execution of Network Configurations

This section is
- a description of how container runtime
  - interprets a network configuration
  - executes a plugin accordingly

Lifecycle is as follows
- container runtime creates a network namespace for container
  - before invoking any plugins
- container runtime must not invoke parallel operators for the same container
  - but can invoke parallel operations for different containers
- plugins must handle being executed concurrently across different containers
  - if necessary, impelement locking on shared resources, such as databases
- container runtime must ensure `add` is eventually followed by corresponding `delete`
  - with exception of catastrophic failure
  - delete must be executed even if the add fails
- delete may be followed by additional deletes
- network config should not change between add and delete
- network config should not change between attachments
- the container runtime is responsible for cleanup of the container's network namespace

Additional sections on
- adding attachments
- deleting attachments
- checking attachment
- garbage-collecting a network
- deriving request configuration from plugin configuration

## Section 4: Plugin Delegation

Some operations cannot be implemented by a single plugin.
Therefore, some functionality can be delegated to other plugins

## Section 5: Result Types

For certain operations, plugins must output result information
The expected output is a JSON object with a set of expected keys
- this JSON object is not small, but not excessive

```yaml
# this is a JSON object in practice, but YAML is easier
# to write here in my notes
cniVersion:         # version string
interfaces:         # array of all interfaces created by the attachment
    name:           # name of interface
    mac:            # the harware address of the interface, if applicable
    mtu:            # the MTU of the interface, if applicable
    sandbox:        # the isolation domain reference (path to network namespace)
    socketPath:     # absolute path to a socket file corresponding to this interface
    pciID:          # platform-specific identifier of the PCI device for this interface, if applicable
ips:                # IPs assigned to this interface
    address:        # IP address in CIDR notation 192.168.1.3/24
    gateway:        # the default gateway for this subnet, if exists
    interface:      # the index into the interfaces list
routes:             # routes created by this attachment
    dst:            # destination of this route, in CIDR notation
    gw:             # the next hop address
    mtu:            # maximum transmission unit
    advmss:         # maximum segment size
    priority:       # prority of route, lower is higher
    table:          # table to add route to
    scope:          # scope of the destination
dns:                # dictionary consisting of DNS config information
    nameservers:    # list or priority-ordered DNS nameservers
    domain:         # local domain for short hostname lookups
    search:         # list of priority ordered search domains
    options:        # list of options to pass to resolver
```

## Appendix: Examples
