
### Esse é um projeto com o intuito de entender mais sobre a linguagem Go utilizando 

> Docker
> Kafka
> Mysql
> Hexagonal Architecture

```bash
docker compose exec mysql bash
```

```bash
mysql -uroot -proot products
```

```mysql
create table products (id varchar(255), name varchar(255), price float);
```

### Criar os topicos do Kafka

```bash
docker compose exec kafka bash
```

```bash
kafka-topics --bootstrap-server=localhost:9092 --topic=products --create
```

### Subir a aplicação

```bash
docker compose exec goapp bash
```

```bash
go run cmd/app/main.go
```

### Realizar chamadas HTTP

> test.http

or

### Produzir mensagem via Kafka


```bash
docker compose exec kafka bash
```

```bash
kafka-console-producer --bootstrap-server=localhost:9092 --topic=products
```

```bash
{"name":"Product2","price":100}
```