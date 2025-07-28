# stress_test

Ferramenta CLI em Go para realizar testes de carga em serviços web.
Permite especificar o número total de requisições, a quantidade de chamadas simultâneas e gera um relatório detalhado ao final da execução.

## 🧩 Funcionalidades
- Realiza testes de carga em uma URL definida.
- Suporte a execução com concorrência (goroutines).
- Relatório com:

  -Tempo total de execução.
  
  -Quantidade de requisições com sucesso (HTTP 200).

  -Distribuição de outros códigos de status (404, 500, etc.).

  -Tempo médio, mínimo e máximo de resposta.

  -Percentual de sucesso.

  -Listagem de erros (ex.: timeouts, DNS falhas).


### Requisitos
Go 1.21+
Docker (opcional, se quiser executar via container)

### instalação
Clone o repositório e compile:
```
git clone https://github.com/henriquedessen/stress_test.git
cd stress_test
go mod tidy
go build -o stress_test main.go
```

## 🚀 Como Executar
```
./stress_test --url=http://google.com --requests=1000 --concurrency=10
```

### Parâmetros

--url: URL do serviço a ser testado (obrigatório).

--requests: Número total de requisições (default: 1).

--concurrency: Número de chamadas simultâneas (defaul: 1).

### Saída de exemplo:

```
--- Relatório ---
Total de requests:     1000
Sucesso (200):         997
Percentual de sucesso: 99.70%
Status HTTP:
  200: 997
  0: 3

Tempo total:           12m15.598162974s
Tempo médio:           0.735s
Tempo mínimo:          0.134s
Tempo máximo:          3.521s
Desvio padrão:         0.218s

⚠️  Erros de request (status 0):
  [1] Get "http://example.com": context deadline exceeded
  [2] Get "http://example.com": dial tcp: lookup example.com: no such host
```

### Execução via Docker
Build da imagem
```
docker build -t stress_test .
```
### Execução do container
```
docker run --rm stress_test --url=http://google.com --requests=1000 --concurrency=10
```

### 📦 Estrutura do Projeto
```
stress_test/
├── cmd/
│   └── root.go          # CLI principal (flags e execução)
├── internal/
│   ├── runner/          # Coordena goroutines e requests
│   │   └── runner.go
│   ├── reporter/        # Geração de relatório
│   │   └── reporter.go
│   └── client/          # Lógica de requisições HTTP
│       └── client.go
├── Dockerfile
├── go.mod
└── main.go

```
