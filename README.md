# Возможные улучшения
> ***Самое очевидное улучшения для дальнейшего расширения - разбиение на микросервисы при увеличении количества сущностей***
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

#Установка Docker
$ sudo apt-get update
$ sudo apt-get install docker-ce docker-ce-cli containerd.io

#Установка docker-compose
$ sudo curl -L "https://github.com/docker/compose/releases/download/1.29.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
$ sudo chmod +x /usr/local/bin/docker-compose

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

#В случае, если порт для бд занят, необходимо убить процесс, его занимающий

#Для нахождения необходимого pid
$ sudo ss -lptn 'sport = :80' | grep pid

#Убить процесс по pid
$ kill {pid}

# Принцип работы
$ http://localhost:5000/11/11/2021  (любая дата в формате DD/MM/YYYY. Рассчеты ведутся за 90 дней от введенной даты)
```