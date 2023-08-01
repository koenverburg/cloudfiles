## Still to over
- [ ] Helm
- [ ] Split traffic over two version of a release (40% on A, 60% on B)
- [ ] Create an canary release
- [ ] Setup bookmarks.html

## Tools
- yq - yaml query
- nvim
- rg - ripgrep
- fzf

## Quick Note's
Helm
- adding a repo `helm repo add <name> <link>`
- search hub `helm search hub <name>`
- repo hub `helm search repo <name>`
- install chart `helm install bravo bitnami/drupal` `helm install <release-name> <chart name>`

## Repository's
- add-apt-repository https://launchpad.net/~neovim-ppa/+archive/ubuntu/stable

## prepare env

Use the Kodekloud and killer.sh shells to find a way to have neovim and tmux up and running so the environment feels familiar

## Prep exames
- KodeKloud
- (killer.sh/ckad)[https://killer.sh/ckad]


## tools
**nvim**
```
curl -LO https://github.com/neovim/neovim/releases/latest/download/nvim.appimage
chmod u+x nvim.appimage
./nvim.appimage
```

If the `./nvim.appimage` command fails, try:

```
./nvim.appimage --appimage-extract
./squashfs-root/AppRun --version

# Optional: exposing nvim globally.
sudo mv squashfs-root /
sudo ln -s /squashfs-root/AppRun /usr/bin/nvim
nvim
```

- k9s `curl -sS https://webinstall.dev/k9s | bash`
- kubeconform `go install github.com/yannh/kubeconform/cmd/kubeconform@latest`
