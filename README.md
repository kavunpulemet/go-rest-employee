## Запуск с использованием Docker

### Docker-Compose

`docker-compose up --build app`

### Миграции БД

`goose -dir ./schema postgres "postgres://postgres:qwerty@db:5432/postgres?sslmode=disable" up`

## Тесты мапперов

`go test ./pkg/service/mappers`

## API Endpoints

### Create Employee

![Create Employee](https://i.imgur.com/y9eO8N3.png)

![Create Employee BD](https://i.imgur.com/9yohwcj.png)

### Get employee by company id

![Get employee by company id](https://i.imgur.com/Ct82MY4.png)

### Get employee by department

![Get employee by department](https://i.imgur.com/eWX3ADq.png)

### Update employee

![Update employee](https://i.imgur.com/wdCVOpa.png)

![Update employee result](https://i.imgur.com/FrlpwMp.png)

### Delete employee

![Delete employee](https://i.imgur.com/cGB4t8S.png)

![Delete employee result](https://i.imgur.com/ElfHRBP.png)