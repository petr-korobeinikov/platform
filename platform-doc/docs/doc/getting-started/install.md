# Установка

## Зависимости

### docker

```shell
brew install --cask docker
```

После завершения установки `docker` (`Docker Desktop`), скорее всего,
потребуется перезагрузить компьютер.

### docker-compose

В подавляющем большинстве случаев подойдёт установка `docker-compose`
через `Homebrew`.

```shell
brew install docker-compose
```

Наиболее рекомендуемым способом является установка через `pipx`.
Этот способ подойдёт для тех, кто хорошо понимает его преимущества.

```shell
python -m pipx install docker-compose
```

### minikube

```shell
asdf install minikube 1.23.2
```

`asdf-vm` позволяет устанавливать несколько разных версий программного
обеспечения независимо друг от друга.

Это своеобразный `rbenv`, `pyenv` или `nvm` для "всего".

Подробнее узнать про `asdf` можно в официальной документации[^1].

Возможно, в будущем установка `asdf-vm` и `minikube` будут автоматизированы.

[^1]: [asdf-vm.com](https://asdf-vm.com/)
