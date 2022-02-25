# Настройка вручную

Создайте директорию `~/.platformctl`:

```shell
mkdir -p ~/.platformctl
```

Создайте конфигурационный файл `~/.platformctl/platformctl.yaml`:

```shell
cat <<EOF > ~/.platformctl/platformctl.yaml
platform:
  flavor:
    container-runtime: docker
    container-runtime-ctl: docker
    container-runtime-vm: minikube

  minikube:
    memory: "8g"
    cpus: "8"
    disk-size: "50g"

  go_env_vars:
    - "GOSUMDB=on"
    - "GONOPROXY=none"
    - "GOPROXY=https://your.proxy.address"
    - "GONOSUMDB=your.go.sum.db/*"
    - "GOPRIVATE=*.your.go.private"
EOF
```
