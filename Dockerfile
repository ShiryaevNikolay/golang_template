# использует линукс alpine
FROM golang:1.17.4-alpine

# добавляем рабочую директорию
WORKDIR /opt/code/
# копируем директорию, из которой запускается докер
ADD ./ /opt/code/

# устанавливаем git, т.к. в alpine нет его по умолчанию
RUN apk update && apk upgrade && \
    apk add --no-cache git

# скачает зависимости из go mod файла
RUN go mod download

# собираем бинарник в папку bin/example (путь контенера) 
# cmd/ (путь, к тому файлу, который будем компилировать)
RUN go build -o bin/exmaple cmd/exmaple/main.go

# инструкция к тому, где докер контейнер 
# к какому месту он должен обратиться, когда он запускается
ENTRYPOINT [ "bin/exmaple" ]