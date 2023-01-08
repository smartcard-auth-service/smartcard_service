# smartcard_service

Данное приложение реализует логику сохранения/получения объекта двумя разными способами:
1) Использование rpc протокола
2) Использование простого tls сервера

В качестве хранилища используется mongodb
## Запуск:

```
docker-compose up -d
```

## Генерация объекта:

1) Для генерации нового объекта можно выполнить следующий get запрос:
```
curl -k --cert internal/tls/tls_client/client.crt --key internal/tls/tls_client/client.key -X GET https://localhost:1443/generate
```

2) Или подключиться к grpc клиенту с помощью https://github.com/ktr0731/evans:

```
evans --proto internal/grpc/pb/grpc_pb.proto -p 8082 repl
```

И вызвать метод ```RegisterCardData```. Параметром можно передать моковый объект:
```
    {"_id": "6339b520cbb4b87871767346","type_of_card": "credit","owner": "Dima Robinson","cvc": "991","number": "5026-5061-5734-4658","date": "2018-12-10T13:45:00.000Z"}
```
![image](https://user-images.githubusercontent.com/74458701/211200325-8ff1f093-f9ac-44d8-8645-aa0ac14c3168.png)


В обоих случая должен быть получен ответ с id сохранённого объекта

## Получение объекта

1) Для получения объекта можно выполнить следующий get запрос:
```
curl -k --cert internal/tls/tls_client/client.crt --key internal/tls/tls_client/client.key -X GET https://localhost:1443/get?id=<object_id>
```

2) Или подключиться к grpc клиенту с помощью https://github.com/ktr0731/evans:

```
evans --proto internal/grpc/pb/grpc_pb.proto -p 8082 repl
```

И вызвать метод ```GetCardData```. Параметром нужно передать id искомого объекта:
## Остановка:

```
docker-compose down
```
