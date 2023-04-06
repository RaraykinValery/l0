# L0

Запустить сервис:\
```docker-compose up -d```

Скомпилировать publisher:\
```go build ./cmd/publisher/```

Запустить publisher:\
```./publisher```

Скрипт publisher раз в 10 секунд публикует order из model.json с уникальным uuid в канал nats-streaming, к которому подключен сервис.

Проверить работу можно перейдя по ссылке http://localhost:8080
