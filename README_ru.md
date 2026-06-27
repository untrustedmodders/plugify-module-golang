[![English](https://img.shields.io/badge/English-%F0%9F%87%AC%F0%9F%87%A7-blue?style=for-the-badge)](README.md)

# Модуль языка Go для Plugify

Модуль языка Go для Plugify позволяет разработчикам писать плагины на Go для фреймворка Plugify. Этот модуль обеспечивает бесшовную интеграцию Go-плагинов, позволяя ядру Plugify динамически их загружать и управлять ими.

## Возможности

- **Поддержка плагинов на Go**: Пишите плагины на Go и легко интегрируйте их с фреймворком Plugify.
- **Автоматическая экспортируемость**: Легко экспортируйте и импортируйте методы между плагинами и языковым модулем.
- **Инициализация и завершение**: Обрабатывайте запуск, инициализацию и завершение плагина с помощью событий модуля.
- **Взаимодействие между языками**: Общение с плагинами на других языках через автоматически сгенерированные интерфейсы.
- **Режим единого рантайма**: Опционально запускайте все Go-плагины внутри одного общего Go-рантайма для снижения потребления памяти и прямых вызовов между плагинами (см. [managed/README_ru.md](managed/README_ru.md)).

## Начало работы

### Требования

- Компилятор Go  
- Установленный фреймворк Plugify

### Установка

#### Вариант 1: Установка через менеджер плагинов Plugify

Вы можете установить модуль языка Go с помощью менеджера пакетов Mamba, выполнив следующую команду:

```bash
mamba install -n your_env_name -c https://untrustedmodders.github.io/plugify-module-golang/ plugify-module-golang
```

#### Вариант 2: Ручная установка

1. Установите зависимости:

   a. Windows  
   > Настройка [CMake-инструментов через Visual Studio Installer](https://learn.microsoft.com/en-us/cpp/build/cmake-projects-in-visual-studio#installation)

   b. Linux:  
   ```sh
   sudo apt-get install -y build-essential cmake ninja-build
   ```

   c. Mac:  
   ```sh
   brew install cmake ninja
   ```

2. Клонируйте репозиторий:

   ```bash
   git clone https://github.com/untrustedmodders/plugify-module-golang.git --recursive
   cd plugify-module-golang
   ```

3. Соберите модуль языка Go:

   ```bash
   mkdir build && cd build
   cmake ..
   cmake --build .
   ```

### Использование

1. **Интеграция с Plugify**

   Убедитесь, что модуль языка Go находится в той же директории, что и ваша установка Plugify.

2. **Создание плагинов на Go**

   Разрабатывайте свои плагины на Go с использованием API Plugify. Подробности смотрите в [руководстве по созданию Go-плагинов](https://untrustedmodders.github.io/languages/golang/first-plugin).

3. **Сборка и установка плагинов**

   Скомпилируйте ваши Go-плагины и разместите скомпилированные библиотеки в директории, доступной для ядра Plugify.

4. **Запуск Plugify**

   Запустите фреймворк Plugify — он автоматически загрузит ваши Go-плагины.

## Режимы работы плагинов

Модуль языка Go поддерживает два режима запуска плагинов.

### Стандартный режим — C Shared Library

Каждый плагин компилируется как отдельная C-библиотека. Это стандартный режим, который работает на всех поддерживаемых платформах.

```sh
go build -buildmode=c-shared -ldflags="-X main.PluginName=example_plugin" -o my_plugin.so ./
```

### Режим единого рантайма — Go Native Plugin

Все плагины работают внутри единого общего Go-рантайма, размещённого в `libplugify-host-golang.so`. Это снижает потребление памяти и позволяет Go-плагинам вызывать друг друга напрямую, минуя слой маршалинга Plugify. Однако этот режим предъявляет жёсткие требования к совместимости версий тулчейна и зависимостей.

> **Рантайм-хост необходимо собрать самостоятельно**, чтобы он точно соответствовал вашим плагинам. Подробности сборки, развёртывания и описание компромиссов смотрите в [managed/README_ru.md](managed/README_ru.md).

Плагины для этого режима собираются с помощью:

```sh
go build -buildmode=plugin -tags=plugin -o my_plugin.so ./
```

Бинарный файл хоста (`libplugify-host-golang.so`) должен быть помещён в папку `bin` рядом с `libplugify-module-golang.so` и **предварительно загружен через `LD_PRELOAD`** перед запуском Plugify:

```sh
LD_PRELOAD=/path/to/bin/libplugify-host-golang.so ./plugify
```

## Пример

### Инициализация модуля

```sh
go mod init example.com/my-go-plugin
```

### Установка модуля go-plugify

Обратите внимание, что версия должна начинаться с `v`.

```sh
go get github.com/untrustedmodders/go-plugify
```

```go
// go build -buildmode="c-shared" -ldflags="-X main.PluginName=example_plugin" ./
package main

import (
	"fmt"
	"github.com/untrustedmodders/go-plugify"
)

var plugin plugify.Plugin
var PluginName string // должен соответсвовать имени в манифесте

func onPluginStart() error {
	fmt.Println("Go: OnPluginStart")
	return nil
}

func onPluginUpdate(dt float32) error {
	fmt.Println("Go: OnPluginUpdate")
	return nil
}

func onPluginEnd() error {
	fmt.Println("Go: OnPluginEnd")
	return nil
}

func init() {
	plugin = plugify.NewPlugin(PluginName, onPluginStart, onPluginUpdate, onPluginEnd)
}

func main() {}
```

## Документация

Полную документацию по созданию плагинов на Go для Plugify можно найти в [официальной документации Plugify](https://untrustedmodders.github.io).

## Участие

Вы можете внести вклад, открыв issue или отправив pull request. Мы будем рады вашим идеям и предложениям!

## Лицензия

Этот модуль языка Go для Plugify распространяется по лицензии [MIT](LICENSE).
