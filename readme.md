# Jokenpo API

Este projeto é uma API simples implementada em Go que simula o jogo clássico "Jokenpo" (Pedra, Papel e Tesoura). A API aceita entradas do usuário, escolhe aleatoriamente uma opção para a máquina, e retorna o resultado do jogo.

---

## 🛠️ Tecnologias Utilizadas

- [Go](https://go.dev/) - Linguagem de programação
- [Echo](https://echo.labstack.com/) - Framework web para Go

---

## 🚀 Configuração do Ambiente

### Pré-requisitos

- [Go](https://go.dev/dl/) instalado na versão 1.22.6 ou superior

### Instalação

1. Clone este repositório:
   ```bash
   git clone https://github.com/PatrickChagasTavares/jokenpo-api.git
   cd jokenpo-api
   ```

2. Instale as dependências:
   ```bash
   go mod tidy
   ```

---

## 🔧 Uso

### Iniciando o servidor

1. Execute o comando:
   ```bash
   go run main.go
   ```

2. O servidor estará disponível na porta `3002`. Certifique-se de que a porta esteja livre no seu ambiente.

---

### Rotas Disponíveis

#### POST `/`

Simula uma rodada de "Jokenpo" contra a máquina.

##### Corpo da Requisição (JSON)

```json
{
  "input": "pedra"
}
```

- O campo `input` deve conter uma das seguintes opções (insensível a maiúsculas/minúsculas):
  - `pedra`
  - `papel`
  - `tesoura`

##### Exemplo de Resposta (JSON)

```json
{
  "resultado": "Você venceu",
  "opção da maquina": "tesoura"
}
```

##### Códigos de Resposta

- `200 OK`: Resultado do jogo retornado com sucesso.
- `400 Bad Request`: O input é inválido (não corresponde a "pedra", "papel" ou "tesoura").
- `500 Internal Server Error`: Erro ao processar o resultado.

---

## 📝 Notas Técnicas

- O projeto utiliza o pacote `math/rand` para selecionar aleatoriamente a opção da máquina.
- A lógica do jogo é implementada por meio de um mapa `resultJokenpo`, onde as combinações possíveis de jogadas determinam os resultados.

---

## 🧪 Testando a API

Para testar a API, você pode usar ferramentas como [Postman](https://www.postman.com/) ou `curl`.

#### Exemplo com `curl`:

```bash
curl -X POST http://localhost:3002/ -H "Content-Type: application/json" -d '{"input": "pedra"}'
```
