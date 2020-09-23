clean:
	kind delete cluster

create-cluster: clean
	kind create cluster

configure: create-cluster
	kubectl create configmap varnish-config --from-file=varnish.vcl

deploy: configure
	kubectl create -f dictionary-deployment.yaml
	kubectl create -f dictionary-service.yaml
	kubectl create -f varnish-deployment.yaml
	kubectl create -f varnish-service.yaml

	# kubectl port-forward service/varnish-service 8080:80
	# curl localhost:8080/something