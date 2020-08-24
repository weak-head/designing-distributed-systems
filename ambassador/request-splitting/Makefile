clean:
	kind delete cluster

create-cluster: 
	kind create cluster

configure: clean create-cluster
	kubectl create configmap ambassador-config --from-file=ambassador/nginx.conf
	kubectl create configmap experiment-config --from-file=experiment/nginx.conf
	kubectl create configmap prod-config --from-file=prod/nginx.conf

deploy: 
	kubectl create -f pod.yaml
	kubectl create -f services.yaml