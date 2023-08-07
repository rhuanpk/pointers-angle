# Pointers Angle API

Essa API calcula o menor ângulo arredondado entre os ponteiros da hora e minuto de um relógio.

Ela recebe os valores via _param paths_ (parâmetros de caminho), ou seja, os valores são informados como se fossem parte do _endpoint_ e o seu retorno é um objeto _json_ com um atributo chamado `angle` e que seu conteúdo é o resultado do cálculo.

## Utilização

As requisiçẽos podem ser feitas por qualquer cliente que consiga se conectar via HTTP:

- Browser's: Qualquer navegador moderno deverá ser capaz de requisitar a rota, basta digitar a _URL_ na barra de endereços e buscar.

- Swagger: Depois de executando, esta API conta com uma documentação em _swagger_ na qual, por dentro da própria interface dele é possível realizar requisições (o _swagger_ se encontra no link: <http://localhost:8080/swagger/index.html>).

- cURL: Com o comando `curl <url>` também é possível fazer a requisição via terminal ou em _scripts_. Qualquer outra ferramenta _CLI_ que possa ser um cliente HTTP, também é possível fazer a requisição.

### URL

A _URL_ para realizar as requisições via _browser_ ou _CLI_:

```
[http://]localhost:8080/v0/rest/clock/<hour>[/<minute>]
```

OBS: o _scheme_ da da _URL_ está envolvido em colchetes (`[]`) pois é opcional.

### Retorno

O modelo de _json_ a ser retornado:

```json
{
	"angle": \<value\>
}
```

## Execução

A infraestrura da API foi projetada para rodar de forma independente via _docker compose_, sendo assim, estando na raiz do projeto, basta executar:

```bash
docker-compose up -d
```

Logo tudo estará de pé. Além de colocar a api de pé, ele cria a instância de um banco de dados (**postgres**) e também de um leve gerenciador de _DBMS_ (**adminer**).

Se desejar subir por conta própria o seu banco de dados e seu gerênciador gráfico. Bastará trocar os valores do arquivo de configuração (`config.yml`) localizado na raiz do projeto e na hora de executar e na hora de executar o _compose_ chamar somente a api:

```bash
docker-compose up -d api
```

### DBMS

Para entrar no gerênciador, acesse:

```
[http://]localhost:8888
```

E no formulário de login defina os valores:

- System: `PostgreSQL`;
- Server: `database`;
- Username: `root`;
- Password: `root`.
