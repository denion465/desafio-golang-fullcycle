# Desafio Golang Imersão Fullcycle 9 - Codebank
<a href="https://fullcycle.com.br/">
  <img src="https://events-fullcycle.s3.amazonaws.com/events-fullcycle/static/site/img/grupo_4417.png" />
</a>

## Instruções
#### Para levantar a aplicação use:
```shell
  docker-compose up
````
> API rodando na porta 8080, certifique-se que está porta não esteja sendo usada por outra aplicação.

## End-points
### Criar uma conta:
```
  POST /bank-accounts
```
### Payload:
```json
  {
    "account_number": "1111-11"
  }
```
> Conta ja inicia com 1.000 na conta, para melhorar a experiência dos testes.
#
### Fazer uma transferência:
```
  POST /bank-accounts/transfer
```
### Payload:
```json
  {
    "from": "1111-11",
    "to": "1111-12",
    "amount": 100
  }	
```
> Para uma melhor experiência, crie duas contas, e faça a transferência entre elas.

## 🛠 Tecnologias usadas:
<div align="start">
  <a href="https://go.dev" target="_blank" style="text-decoration: none;">
    <img
      src="https://img.shields.io/static/v1?label=|&message=Golang&color=informational&style=plastic&logo=GO"
    />
  </a>
  <a href="https://www.sqlite.org" target="_blank" style="text-decoration: none;">
    <img
      src="https://img.shields.io/static/v1?label=|&message=SQLite&color=informational&style=plastic&logo=sqlite"
    />
  </a>
</div>

## Author
### Daniel Vidal

* GitHub: https://github.com/denion465
* Linkedin: https://www.linkedin.com/in/daniel-vidal465
