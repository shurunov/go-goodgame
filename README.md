# go-goodgame

## Подключение
```go
import "github.com/shurunov/go-goodgame"
```

После импортирования пакета требуется инициализировать новый клиент для работы с сервисом:
```go
gg := goodgame.NewClient(nil)
```

Некоторые запросы требуют Вашего **Access Token**'а. В этом случае инициализировать новый клиент нужно с передачей первого параметра в ввиде ссылки на структуру **Settings**, содержащую Ваш **Access Token**:
```go
settings := &goodgame.Settings{"Ваш Access Token"}
gg := goodgame.NewClient(settings)
```

## Примеры
Список параметров для фильтрации трансляций:
```go
options := &goodgame.StreamOptions{
    IDs: "27585,1141",
    Page: 1,
    OnlyGG: true,
    Adult: true,
    Hidden: true,
}
```

Получение списка трансляций:
```go
streams, err := gg.Stream.All(options)
```

Получение конкретной трансляции:
```go
stream, err := gg.Stream.Get("Yudzy")
```

## TODO
- [X] Стримы
- [X] Игры
- [X] Смайлы
- [X] Плеер
- [X] Донаты
- [ ] Премиум подписчики
- [ ] Информация по AccessToken
- [ ] Токен чата