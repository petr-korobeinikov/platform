# asdf-vm

## Введение

* что такое `asdf-vm`;
* как его установить;
* как установить разные версии рантаймов;
* как переключаться между версиями;
* приведу ссылки на простую и понятную документацию.

`asdf-vm` — это единый менеджер версий различных рантаймов (Go, Python,
Postgres, node, ruby, ...). Он позволяет иметь несколько одновременно
установленных версий (в домашней папке, unix-way) и легко переключаться между
ними. Одну из версий каждого рантайма можно настроить как версию по умолчанию.

Сайт и документация:

- [https://asdf-vm.com/](https://asdf-vm.com/)

## Установка

Наиболее подходящий способ для macOS я опишу здесь.

```shell
brew install asdf
```

Не забудьте добавить в ваш `.rc`-файл строку:

```shell
. $(brew --prefix asdf)/asdf.sh
```

Все возможные способы установки можно найти в документации:

- [https://asdf-vm.com/guide/getting-started.html#_3-install-asdf](https://asdf-vm.com/guide/getting-started.html#_3-install-asdf)

## Установка плагинов

Плагины отвечают за механизм установки и удаления рантаймов. Рассмотрим
установку плагина на примере Go:

```shell
asdf plugin install golang
```

Подробнее про установку плагинов в документации:

- [https://asdf-vm.com/guide/getting-started.html#install-the-plugin](https://asdf-vm.com/guide/getting-started.html#install-the-plugin)

## Установка рантайма

Теперь давайте установим одну или несколько версий Go:

```shell
asdf install golang 1.16.2
asdf install golang 1.15.6
```

Получить список установленных версий можно так:

```shell
asdf list golang
```

Подробнее про установку можно прочитать в документации:

- [https://asdf-vm.com/guide/getting-started.html#_5-install-a-version](https://asdf-vm.com/guide/getting-started.html#_5-install-a-version)

## Назначение рантайма по умолчанию

Чтобы понимать, какой рантайм назначен по умолчанию, `asdf-vm` использует
файл `~/.tool-versions`. Его можно заполнить вручную или, что более удобно с
помощью следующей команды:

```shell
asdf global golang 1.15.6
```

Список текущих версий рантаймов можно посмотреть так:

```shell
asdf current
```

Подробнее в документации:

- [https://asdf-vm.com/guide/getting-started.html#global](https://asdf-vm.com/guide/getting-started.html#global)

## Запуск шелла с другой версией рантайма

Для смены версии рантайма удобно (и быстро) использовать подкоманду `shell`:

```shell
asdf shell golang 1.16
```

Мой `powerlevel10k` даже подсвечивает текущую версию в строке приглашения:

```
... ✔ │ 1.16 Go │ at 17:18:05
```

## Пути для настройки рантайма в IDE

Среда разработки, к сожалению, не смотрит на `~/.toolversions` и не знает
про `asdf-vm`. Поэтому потребуется немного ручной настройки.

В общем виде:

```
~/.asdf/installs/<plugin>/<version>/go
```

На примере go 1.15.6:

```
~/.asdf/installs/golang/1.15.6/go
```
