# oidc-client

Repositório para POC de uso de OpenID Connect

Este projeto tem como objetivo demonstrar na prática o uso do protocolo OpenID Connect (OIDC) em aplicações escritas em Go. Trata-se de uma Prova de Conceito (POC) simples e direta, servindo como base para quem deseja implementar autenticação moderna baseada em OIDC.

## ✨ Funcionalidades

- Integração com provedores OIDC para autenticação de usuários.
- Fluxo de autenticação seguro utilizando tokens.
- Estrutura simples, ideal para aprendizado e customização.

## 🚀 Instalação

Certifique-se de ter o [Go](https://golang.org/dl/) instalado em sua máquina.

Clone o repositório:

```bash
git clone https://github.com/italo13d/oidc-client.git
cd oidc-client
```

Instale as dependências (se houver):

```bash
go mod tidy
```

## ⚡ Uso

Compile e execute o projeto:

```bash
go run main.go
```

> **Nota:** Adapte o comando conforme o nome do arquivo principal do projeto.

## 🛠️ Como funciona?

Este projeto faz a autenticação do usuário utilizando um provedor OIDC. Os fluxos principais envolvem:

1. Redirecionamento do usuário para o provedor de autenticação.
2. Retorno do provedor para sua aplicação com o token de autenticação.
3. Validação e uso do token para autenticar o usuário na aplicação.

## 📚 Saiba mais

- [Documentação oficial do OpenID Connect](https://openid.net/connect/)
- [Documentação do Go](https://golang.org/doc/)

## 🤝 Contribuição

Sinta-se à vontade para abrir issues ou enviar pull requests!

---

Feito com 💙 por [italo13d](https://github.com/italo13d)
