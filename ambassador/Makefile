delete-cluster:
	kind delete cluster

create-cluster: delete-cluster
	kind create cluster

deploy-redis: create-cluster
	# Redis service (StetefulSet) with fixed DNS endpoints
	kubectl create -f redis-shards.yaml
	kubectl create -f redis-service.yaml

deploy-client:
	# Ambassador load balancer (twemproxy)
	kubectl create configmap twem-config --from-file=twemproxy.config.yaml
	kubectl create -f ambassador-example.yaml
