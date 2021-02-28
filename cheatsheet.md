# Cheatsheet
> This file hold all the commands I use to interact with Kubernetes, docker and other devops related tooling.

---
Get all pods
```bash
kubectl get pods
```

Start local proxy
```bash
kubectl proxy --port=8000
```

Get Dashboard token
```bash
kubectl -n kubernetes-dashboard describe secret admin-user-token | grep ^token # unix
kubectl -n kubernetes-dashboard describe secret admin-user-token # windows
```


---
## Docker

Clean up all images, volumes and networks
```bash
docker system prune --volumes
```

Clean up all images, volumes and networks that are 24 hours of older
```bash
docker container prune --filter "until=24h"
```

Get all running docker containers
```bash
docker ps
```

Dropping to a shell in a docker container
```bash
docker exec -it containerid123 bash
```