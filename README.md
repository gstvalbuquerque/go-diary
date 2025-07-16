# Daily Diary

Um aplicativo de diário pessoal em linha de comando, escrito em Go, com autenticação de usuário e armazenamento local em JSON.

## Funcionalidades

- Cadastro e login de usuários com senha criptografada
- Adição de entradas diárias divididas em manhã, tarde e noite
- Visualização de entradas por data
- Listagem de entradas por período, últimos 7 dias ou todas as entradas
- Armazenamento seguro dos dados do diário por usuário

## Estrutura do Projeto

```
.
├── main.go
├── go.mod
├── go.sum
├── auth/
│   └── auth.go
├── diary/
│   └── diary.go
└── ui/
    └── ui.go
```

## Como usar

1. **Instale as dependências:**

   ```sh
   go mod tidy
   ```

2. **Compile o projeto:**

   ```sh
   go build -o diary
   ```

3. **Execute o aplicativo:**

   ```sh
   ./diary
   ```

4. **Siga o menu interativo para registrar-se, fazer login e gerenciar seu diário.**

## Armazenamento

- Os dados dos usuários são salvos em `~/.diary/users.json`
- As entradas do diário de cada usuário são salvas em `~/.diary/<usuario>-diary.json`

## Requisitos

- Go 1.18 ou superior

---

Desenvolvido por Gustavo Sales
