# Возможные улучшения
> ***Самое очевидное улучшения для дальнейшего расширение - разбиение на микросервисы при увеличении количества сущностей***
> ***в проекте.***  

> ***Добавление обработки большего количества ошибок, вызов логгера на всех этапах проверок, а не только на уровне***
> ***entrypoint, что позволит грамотнее отслеживать места их возникновения***

> ***Добавить различные утилитарные функции в utils для дальнейшего реиспользования***

> ***Еще много разных других вещей...***


### Как запустить проект
Шаги для запуска приложения через `docker-compose`

```bash
#Переключиться на директорию
$ cd workspace

# Скопировать проект в свой $GOPATH/src
$ git clone https://github.com/Constantilation/CentralBankTask.git

#Перейти в папку проекта
$ cd CentralBankTask

# Создать образ докера
$ sudo make docker

# Запустить приложение
$ sudo make run

# проверить, запущено ли наше приложение
$ docker ps

# Остановить приложение
$ sudo make stop
```