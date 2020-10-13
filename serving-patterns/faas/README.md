# Function as a Service

## Deploy

```sh
# Build and load docker images into kind cluster
make build-docker
make load-image

# Deploy faas
make deploy
```

## Dependencies

This deployment requires [kind](https://kind.sigs.k8s.io/docs/user/quick-start/) and [kubeless](https://kubeless.io/).

If you dont want to configure k8s cluster yourself you can use the pre-configured [multinode-k8s](https://github.com/weak-head/multinode-k8s).
