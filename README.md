# go-goodgame

Первым делом после импортирование пакета Вам нужно инициализировать новый клиент для работы с сервисом.
```go
gg := goodgame.NewClient()
```

## Примеры
Список параметров для фильтрации трансляций
```go
options := &goodgame.StreamOptions{
    IDs: "27585,1141",
    Page: 1,
    OnlyGG: true,
    Adult: true,
    Hidden: true,
}
```

Получение списка трансляций
```go
streams, err := gg.Stream.All(options)
```

Получение конкретной трансляции
```go
stream, err := gg.Stream.Get("Yudzy")
```