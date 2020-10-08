.PHONY: clean
clean:
	# Delete sharded memcache
	kubectl delete statefulset sharded-memcache 
	kubectl delete service memcache 

	# Delete twemproxy config map
	kubectl delete configmap twem-config

	# Delete ambassador pod
	kubectl delete deployment shared-twemproxy
	kubectl delete service shared-twemproxy-service

.PHONY: deploy
deploy:
	# Deploy sharded memchache
	kubectl create -f sharded-memcache.yaml
	kubectl create -f sharded-memcache-service.yaml

	# Create twemproxy config map
	kubectl create configmap --from-file=nutcracker.yaml twem-config

	# Deploy shared twemproxy
	kubectl create -f twemproxy-deployment.yaml
	kubectl create -f twemproxy-service.yaml