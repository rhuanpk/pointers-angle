# Pointers Angle API

Essa _API_ calcula o menor ângulo arredondado entre os ponteiros da hora e minuto de um relógio.

Ela recebe os valores via _param paths_ (parâmetros de caminho), ou seja, os valores são informados como se fossem parte do _endpoint_. Seu retorno é um objeto _json_ com um atributo nomeado de `angle` e seu conteúdo é o resultado do cálculo.

## Utilização

A requisição pode ser feita por qualquer cliente que consiga abrir uma conexão _HTTP_:

- Browser's: Qualquer navegador moderno deverá ser capaz de requisitar a rota, basta digitar a _URL_ na barra de endereços e buscar.

- [Swagger](https://swagger.io/): Esta _API_ conta com a implementação do _swagger_, na qual, por dentro da própria interface dele é possível realizar as requisições (depois de executar a _API_, o _swagger_ se encontrará disponível na _URL_: <http://localhost:8080/swagger/index.html>).

- [cURL](https://curl.se/): Com o comando `curl <url>` também é possível fazer a requisição via terminal ou em _scripts_. Ou com qualquer outra ferramenta _CLI_ que possa ser um cliente _HTTP_.

### URL

A _URL_ para realizar as requisições via _browser_ ou _CLI_:

```
[http://]localhost:8080/v1/rest/clock/<hour>[/<minute>]
```

### Retorno

O modelo de _json_ a ser retornado:

```json
{
	"angle": <value>
}
```

## Execução

A infraestrura da _API_ foi projetada para rodar agnosticamente a qualquer "_software_ externo" ou sistema operacional, sendo necessário somente a presença do [**Docker**](https://www.docker.com/) e [**Docker Compose**](https://docs.docker.com/compose/). Sendo assim, basta clonar o [repositório do projeto](https://github.com/rhuanpk/pointers-angle) e estando na raiz, executar:

```sh
docker-compose up -d
```

Além de colocar a _API_ de pé, este _docker compose file_ criará a instância de um banco de dados [**PostgreSQL**](https://www.postgresql.org/) e também de um leve gerenciador de _DBMS_ ([**Adminer**](https://www.adminer.org/)).

Se desejar subir por conta própria o seu banco de dados e gerênciador gráfico, bastará trocar os valores do arquivo de configuração (`config.yml`) localizado na raiz do projeto e na hora de executar o `compose`, chamar somente a _API_:

```sh
docker-compose up -d api
```

### DBMS

Para entrar no gerênciador gráfico, acesse:

```
[http://]localhost:8888
```

E no formulário de login defina os valores:

- System: `PostgreSQL`;
- Server: `database`;
- Username: `root`;
- Password: `root`.

---

_OBSERVAÇÕES_:

- O que está envolvido em colchetes (`[]`) é opcional.
