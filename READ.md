Criar tabela

CREATE TABLE public.produtos (
	id varchar NOT NULL,
	"name" varchar NOT NULL,
	amount int NOT NULL,
	CONSTRAINT produtos_pk PRIMARY KEY (id)
);

Para refatorar e aprender

1 - Mudar logger para que ele não precise ser injetado em todas as funções
2 - Entender como passar somente um campo da struct como parametro no SelectOneProduct repository