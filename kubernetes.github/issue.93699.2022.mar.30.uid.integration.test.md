# 2022, Jan 28

```bash
# TODO: read through this and understand it better
make help

# run an integration test
# make will do this, these args specify what to run
make test-integration WHAT=./test/integration/examples KUBE_TEST_ARGS="-run Aggregated"

# number of open file descriptors (remember in unix everything is a file)
# needs to be a lot or tests may fail
ulimit -n 60000
```
