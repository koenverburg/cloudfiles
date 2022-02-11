# Lima-vm

Install Lima using `brew install lima`.

Setup the lima mv with `limactl start docker.yml`

Install docker and docker compose using

```bash
brew install docker docker-compose
```

connect docker cli with the lima docker vm

```bash
  docker context create lima --docker "host=unix://{{.Dir}}/sock/docker.sock"
  docker context use lima
  docker run hello-world
```

