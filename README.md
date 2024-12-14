## Golang CRUD Sem ORM

Projeto criado para praticar Golang.

### Como executar

1 - Rodar o comando: 

`docker compose up`

2 - Conectar no banco e criar a tabela com o seguinte SQL

```
CREATE TABLE public.produtos (
	id varchar NOT NULL,
	"name" varchar NOT NULL,
	amount int NOT NULL,
	CONSTRAINT produtos_pk PRIMARY KEY (id)
);
```

3 - Executar o arquivo main.go com o seguinte comando

`go run main.go`

### Melhorias para aplicar

- Colocar a aplicação go para rodar no docker
- Mudar logger para que ele não precise ser injetado em todas as funções
- Entender como passar somente um campo da struct como parametro no SelectOneProduct repository