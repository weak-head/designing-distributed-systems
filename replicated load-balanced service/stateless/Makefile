clean:
	kind delete cluster

create-cluster:
	kind create cluster

deploy: clean create-cluster
	kubectl create -f deployment.yaml
	kubectl create -f service.yaml

	# kubectl port-forward service/dictionary-server-service 8080:8080
	# curl localhost:8080/something