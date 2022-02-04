build:
	docker build --file dockerfile --tag ansible-run .

deploy-pihole:
	kubectl apply -f .\deployments\pihole.yml

remove-pihole:
	kubectl delete -n default deployment pihole
	kubectl delete -n default service pihole
	kubectl delete pvc pihole-local-containers-claim
	kubectl delete pvc pihole-local-dnsmasq-claim
	kubectl delete pv pihole-local-containers-volume
	kubectl delete pv pihole-local-dnsmasq-volume

deploy-unbound:
	kubectl apply -f .\deployments\unbound.yml

remove-unbound:
	kubectl delete -n default deployment unbound
	kubectl delete -n default service unbound

deploy-assistant-relay:
	kubectl apply -f .\deployments\assistant-relay.yml

remove-assistant-relay:
	kubectl delete -n default deployment assistant-relay
	kubectl delete -n default service assistant-relay
	kubectl delete pvc assistant-relay-claim
	kubectl delete pv assistant-relay-pv