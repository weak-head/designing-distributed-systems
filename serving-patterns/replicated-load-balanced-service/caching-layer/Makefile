clean:
	kind delete cluster

create-cluster: clean
	kind create cluster

configure: create-cluster
	# Nginx TLS certs
	kubectl create secret tls ssl --cert=server.crt --key=server.key

	# Nginx config
	kubectl create configmap nginx-config --from-file=nginx.conf

	# Varnish cache config
	kubectl create configmap varnish-config --from-file=varnish.vcl

deploy: configure
	# dictionary service
	kubectl create -f dictionary-deployment.yaml
	kubectl create -f dictionary-service.yaml

	# varnish cache
	kubectl create -f varnish-deployment.yaml
	kubectl create -f varnish-service.yaml

	# ssl termination
	kubectl create -f nginx-deployment.yaml
	kubectl create -f nginx-service.yaml

	# -- Client to access varnish directly:
	# kubectl port-forward service/varnish-service 8080:80
	# curl localhost:8080/something

	# -- Client to access nginx:
	# sudo kubectl port-forward service/nginx-service 8080:443
	# curl -k https://localhost:8080/ready