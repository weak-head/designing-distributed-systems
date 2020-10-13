# Function as a Service

## Deploy

```sh
# Build and load docker images into kind cluster
make build-docker
make load-image

# Deploy faas
make deploy
```

## Access the app

```sh
# Forward requests to the service
kubectl port-forward service/faas-service 8888:80

# Access the frontend
curl http://localhost:8888
```

## Scraper package

```sh
# Install scraper dependencies and package them into zip
make scraper-package
```

## Dependencies

This deployment requires [kind](https://kind.sigs.k8s.io/docs/user/quick-start/) and [kubeless](https://kubeless.io/).

If you dont want to configure k8s cluster yourself you can use the pre-configured [multinode-k8s](https://github.com/weak-head/multinode-k8s).
