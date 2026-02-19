

#### Dagger commands

```bash
# Test shared setup
dagger call shared
# Test control plane
dagger call control-plane
# Test worker node
dagger call worker-node
# Test worker node with custom consul server
dagger call worker-node --consul-server="192.168.1.10"
# Test AI node
dagger call ai-node
# Test AI node with custom consul server
dagger call ai-node --consul-server="192.168.1.10"
```
