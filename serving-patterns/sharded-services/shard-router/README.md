```sh
# Deploy shard router, sharded cache and memcache client
make deploy

# Port forward so we can access it
kubectl port-forward service/gomem-service 8989:8989

# Set and Get the key from the cache
curl localhost:8989/set/key1/val1
curl localhost:8989/get/key
```
