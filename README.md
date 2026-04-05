# Documentação Técnica: DSQL Motor Lite

O **DSQL Motor Lite** é uma implementação em linha de comando (CLI) do motor de banco de dados DSQL. Ele permite a manipulação de dados estruturados utilizando a API do Discord como camada de persistência, mapeando conceitos de bancos de dados relacionais para elementos de interface do Discord.

## 1. Arquitetura de Mapeamento

O motor opera convertendo comandos SQL customizados em ações dentro de uma guilda (servidor) do Discord.



| Conceito Relacional | Implementação Discord |
| :--- | :--- |
| **Tabela** | Canal de Texto (Text Channel) |
| **Registro (Row)** | Mensagem (Message) |
| **Chave Primária (ID)** | Snowflake ID da Mensagem |
| **Esquema (Schema)** | Tabela ASCII dentro de blocos de código `text` |

---

## 2. Operações Suportadas (Versão Atual)

A versão Lite foca na estabilidade das operações fundamentais de leitura e escrita.

### 2.1. DSELECT
Realiza a leitura e o parsing de dados de um canal.
* **Comando:** `DSELECT <nome_da_tabela>`
* **Processamento:** O motor varre o histórico de mensagens do canal especificado, extrai o conteúdo de blocos de código formatados e reconstrói os pares de chave-valor no terminal.

### 2.2. DINSERT
Cria um novo registro persistente.
* **Comando:** `DINSERT <tabela> COLUMNS(col1, col2) VALUES('val1', 'val2')`
* **Processamento:** O motor valida os argumentos via Expressões Regulares (Regex), gera uma representação visual em ASCII e envia para o canal correspondente. O ID único da mensagem é retornado como confirmação de sucesso.

---

## 3. Configuração de Ambiente

O motor utiliza variáveis de ambiente para autenticação, garantindo que credenciais sensíveis não sejam expostas no código-fonte.

### Variáveis Necessárias (PowerShell):
```powershell
$env:DSQL_TOKEN = "TOKEN_DO_BOT"
$env:DSQL_GUILD = "ID_DO_SERVIDOR"
```

### Inicialização:
```bash
go run lite.go
```

---

## 4. Especificações do Parser

* **Motor de Regex:** Utiliza a biblioteca padrão `regexp` do Go para isolar comandos e argumentos.
* **Tratamento de Strings:** Suporta valores envolvidos em aspas simples ou duplas.
* **Limites:** Sujeito aos limites de taxa (rate limiting) da API do Discord e ao máximo de 2000 caracteres por registro.

---

