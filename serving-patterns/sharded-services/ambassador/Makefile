.PHONY: clean
clean:
	# Delete sharded memcache
	kubectl delete statefulset sharded-memcache 
	kubectl delete service memcache 

	# Delete twemproxy config map
	kubectl delete configmap twem-config

	# Delete ambassador pod
	kubectl delete pod sharded-memcache-ambassador 

.PHONY: deploy
deploy:
	# Deploy sharded memchache
	kubectl create -f sharded-memcache.yaml
	kubectl create -f sharded-memcache-service.yaml

	# Create twemproxy config map
	kubectl create configmap --from-file=nutcracker.yaml twem-config

	# Deploy ambassador pod with application container
	kubectl create -f memcache-ambassador-pod.yaml

	# -- Test how it works:
	#
	# > kubectl port-forward pod/sharded-memcache-ambassador 8989:8989
	# > curl localhost:8989/set/key1/val1
	# > curl localhost:8989/get/key1