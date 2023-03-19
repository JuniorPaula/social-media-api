# DevGo social media API

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)
![MySQL](https://img.shields.io/badge/mysql-%2300f.svg?style=for-the-badge&logo=mysql&logoColor=white)
![Git](https://img.shields.io/badge/git-%23F05033.svg?style=for-the-badge&logo=git&logoColor=white)

A aplicação consiste em uma  **API** de interação baseada em uma rede social, desenvolvida em **Golang** com linguagem principal, **Docker** para a persistência de no banco de dados **Mysql**, além de **GIT** para versionamento do código. 

Requerimentos

- Golang
- Docker
- Git

### Como rodar o projeto
`git clone`

#### Subir o server
~~~go
go run main.go
~~~

#### Gerar o binário

~~~go
go build main.go
~~~

Configure as variáveis de ambientes seguindo: `env.example`


## Principais funcionalidades
- realizar login
- -------
- criar um usuário
- atualizar um usuário
- deletar um usuário
- seguir um usuário
- deixar de seguir um usuário
- atualizar senha do usuário
- listar todos os usuário
- listar um usuário por ID
- listar todos os seguidores de um usuário
- listar todos que o usuário segue
- --------------
- criar uma publicação
- atualizar uma publicação
- deletar uma publicação
- curtir uma publicação
- descurtir uma publicação
- listar todas as publicações
- listar uma publicação por ID
- listar todos as publicações de uma usuário específico 

## Endpoints da aplicação

### Rotas de usuários
#### Create
~~~go
[POST] /users
~~~

#### **Request body**
~~~json
{
	"name": string,
	"nickname": string,
	"email": string,
	"password": string
}
~~~

#### **Responses**
![Generic badge](https://img.shields.io/badge/OK-200-<COLOR>.svg)

~~~json
{
	"id": number,
	"name": string,
	"nickname": string,
	"email": string,
	"password": string,
	"createdAt": string
}
~~~
![Generic badge](https://img.shields.io/badge/bad%20request-400-red)

~~~json
{
    "error": "string"
}
~~~

![Generic badge](https://img.shields.io/badge/Internal%20Server%20Error-500-red)

~~~json
{
    "error": "string"
}
~~~

### Update
~~~go
[PUT] /users/{{userID}}
~~~

#### **Request body**
~~~json
{
	"name": string,
	"nickname": string,
	"email": string
}
~~~
#### **Request headers**
~~~json
{
	"headers": {
		"Authorization": "Bearer token"
	}
}
~~~

#### **Responses**
![Generic badge](https://img.shields.io/badge/NoContent-204-<COLOR>.svg)


![Generic badge](https://img.shields.io/badge/bad%20request-400-red)

~~~json
{
    "error": "string"
}
~~~

![Generic badge](https://img.shields.io/badge/anauthorized-401-orange)
~~~json
{
    "error": "string"
}
~~~

![Generic badge](https://img.shields.io/badge/Internal%20Server%20Error-500-red)

~~~json
{
    "error": "string"
}
~~~

### Delete
~~~go
[DELETE] /users/{{userID}}
~~~
#### **Request headers**
~~~json
{
	"headers": {
		"Authorization": "Bearer token"
	}
}
~~~

#### **Responses**
![Generic badge](https://img.shields.io/badge/NoContent-204-<COLOR>.svg)

![Generic badge](https://img.shields.io/badge/bad%20request-400-red)

~~~json
{
    "error": "string"
}
~~~

![Generic badge](https://img.shields.io/badge/anauthorized-401-orange)
~~~json
{
    "error": "string"
}
~~~

![Generic badge](https://img.shields.io/badge/Internal%20Server%20Error-500-red)

~~~json
{
    "error": "string"
}
~~~

### Update password
~~~go
[POST] /users/{{userID}}/update-password
~~~

#### **Request body**
~~~json
{
	"new_password": "string",
	"old_password": "string"
}
~~~
#### **Request headers**
~~~json
{
	"headers": {
		"Authorization": "Bearer token"
	}
}
~~~

#### **Responses**
![Generic badge](https://img.shields.io/badge/NoContent-204-<COLOR>.svg)

![Generic badge](https://img.shields.io/badge/bad%20request-400-red)

~~~json
{
    "error": "string"
}
~~~

![Generic badge](https://img.shields.io/badge/anauthorized-401-orange)
~~~json
{
    "error": "string"
}
~~~

![Generic badge](https://img.shields.io/badge/Internal%20Server%20Error-500-red)

~~~json
{
    "error": "string"
}
~~~

### Follower
~~~go
[POST] /users/{{userID}}/follower
~~~

#### **Request headers**
~~~json
{
	"headers": {
		"Authorization": "Bearer token"
	}
}
~~~

#### **Responses**
![Generic badge](https://img.shields.io/badge/NoContent-204-<COLOR>.svg)

![Generic badge](https://img.shields.io/badge/bad%20request-400-red)

~~~json
{
    "error": "string"
}
~~~

![Generic badge](https://img.shields.io/badge/anauthorized-401-orange)
~~~json
{
    "error": "string"
}
~~~

![Generic badge](https://img.shields.io/badge/Internal%20Server%20Error-500-red)

~~~json
{
    "error": "string"
}
~~~

### Unfollower
~~~go
[POST] /users/{{userID}}/unfollower
~~~

#### **Request headers**
~~~json
{
	"headers": {
		"Authorization": "Bearer token"
	}
}
~~~

#### **Responses**
![Generic badge](https://img.shields.io/badge/NoContent-204-<COLOR>.svg)

![Generic badge](https://img.shields.io/badge/bad%20request-400-red)

~~~json
{
    "error": "string"
}
~~~

![Generic badge](https://img.shields.io/badge/anauthorized-401-orange)
~~~json
{
    "error": "string"
}
~~~

![Generic badge](https://img.shields.io/badge/Internal%20Server%20Error-500-red)

~~~json
{
    "error": "string"
}
~~~

### Find users
~~~go
[GET] /users
~~~
#### **Request headers**
~~~json
{
	"headers": {
		"Authorization": "Bearer token"
	}
}
~~~

#### **Responses**
![Generic badge](https://img.shields.io/badge/OK-200-<COLOR>.svg)

~~~json
[
	{
		"id": number,
		"name": "string",
		"nickname": "string",
		"email": "string",
		"createdAt": "string"
	}
]
~~~
![Generic badge](https://img.shields.io/badge/bad%20request-400-red)

~~~javascript
{
    "error": "string"
}
~~~

![Generic badge](https://img.shields.io/badge/anauthorized-401-orange)
~~~json
{
    "error": "string"
}
~~~

![Generic badge](https://img.shields.io/badge/Internal%20Server%20Error-500-red)

~~~json
{
    "error": "string"
}
~~~

### Find an user by ID
~~~go
[GET] /users/{{userID}}
~~~
#### **Request headers**
~~~json
{
	"headers": {
		"Authorization": "Bearer token"
	}
}
~~~

#### **Responses**
![Generic badge](https://img.shields.io/badge/OK-200-<COLOR>.svg)

~~~json
{
	"id": number,
	"name": "string",
	"nickname": "string",
	"email": "string",
	"createdAt": "string"
}
~~~
![Generic badge](https://img.shields.io/badge/bad%20request-400-red)

~~~javascript
{
    "error": "string"
}
~~~

![Generic badge](https://img.shields.io/badge/anauthorized-401-orange)
~~~json
{
    "error": "string"
}
~~~

![Generic badge](https://img.shields.io/badge/Internal%20Server%20Error-500-red)

~~~json
{
    "error": "string"
}
~~~

### Find user followers
~~~go
[GET] /users/{{userID}}/followers
~~~
#### **Request headers**
~~~json
{
	"headers": {
		"Authorization": "Bearer token"
	}
}
~~~

#### **Responses**
![Generic badge](https://img.shields.io/badge/OK-200-<COLOR>.svg)

~~~json
[
	{
		"id": number,
		"name": "string",
		"nickname": "string",
		"email": "string",
		"createdAt": "string"
	}
]
~~~
![Generic badge](https://img.shields.io/badge/bad%20request-400-red)

~~~javascript
{
    "error": "string"
}
~~~

![Generic badge](https://img.shields.io/badge/anauthorized-401-orange)
~~~json
{
    "error": "string"
}
~~~

![Generic badge](https://img.shields.io/badge/Internal%20Server%20Error-500-red)

~~~json
{
    "error": "string"
}
~~~

### Find user following
~~~go
[GET] /users/{{userID}}/following
~~~
#### **Request headers**
~~~json
{
	"headers": {
		"Authorization": "Bearer token"
	}
}
~~~

#### **Responses**
![Generic badge](https://img.shields.io/badge/OK-200-<COLOR>.svg)

~~~json
[
	{
		"id": number,
		"name": "string",
		"nickname": "string",
		"email": "string",
		"createdAt": "string"
	}
]
~~~
![Generic badge](https://img.shields.io/badge/bad%20request-400-red)

~~~javascript
{
    "error": "string"
}
~~~

![Generic badge](https://img.shields.io/badge/anauthorized-401-orange)
~~~json
{
    "error": "string"
}
~~~

![Generic badge](https://img.shields.io/badge/Internal%20Server%20Error-500-red)

~~~json
{
    "error": "string"
}
~~~

## Rota de Login
### Login
~~~go
[POST] /login
~~~
#### **Request body**
~~~json
{
	"email": "string",
	"password": "string"
}
~~~

#### **Responses**
![Generic badge](https://img.shields.io/badge/OK-200-<COLOR>.svg)

~~~json
{
	"userId": "string",
	"token": "string"
}
~~~
![Generic badge](https://img.shields.io/badge/bad%20request-400-red)

~~~javascript
{
    "error": "string"
}
~~~

![Generic badge](https://img.shields.io/badge/Internal%20Server%20Error-500-red)

~~~json
{
    "error": "string"
}
~~~

## Rota de Publicações
### create
~~~go
[POST] /posts
~~~
#### **Request body**
~~~json
{
	"title": "string",
	"content": "string"
}
~~~
#### **Request headers**
~~~json
{
	"headers": {
		"Authorization": "Bearer token"
	}
}
~~~

#### **Responses**
![Generic badge](https://img.shields.io/badge/created-201-<COLOR>.svg)

~~~json
{
	"id": number,
	"title": "string",
	"content": "string",
	"authorId": number,
	"likes": number,
	"createdAt": "string"
}
~~~
![Generic badge](https://img.shields.io/badge/bad%20request-400-red)

~~~javascript
{
    "error": "string"
}
~~~

![Generic badge](https://img.shields.io/badge/anauthorized-401-orange)
~~~json
{
    "error": "string"
}
~~~

![Generic badge](https://img.shields.io/badge/Internal%20Server%20Error-500-red)

~~~json
{
    "error": "string"
}
~~~

### update
~~~go
[PUT] /posts/{{postID}}
~~~
#### **Request body**
~~~json
{
	"title": "string",
	"content": "string"
}
~~~
#### **Request headers**
~~~json
{
	"headers": {
		"Authorization": "Bearer token"
	}
}
~~~

#### **Responses**
![Generic badge](https://img.shields.io/badge/noContent-204-<COLOR>.svg)

![Generic badge](https://img.shields.io/badge/bad%20request-400-red)

~~~javascript
{
    "error": "string"
}
~~~

![Generic badge](https://img.shields.io/badge/anauthorized-401-orange)
~~~json
{
    "error": "string"
}
~~~

![Generic badge](https://img.shields.io/badge/Internal%20Server%20Error-500-red)

~~~json
{
    "error": "string"
}
~~~

### delete
~~~go
[DELETE] /posts/{{postID}}
~~~
#### **Request headers**
~~~json
{
	"headers": {
		"Authorization": "Bearer token"
	}
}
~~~

#### **Responses**
![Generic badge](https://img.shields.io/badge/noContent-204-<COLOR>.svg)

![Generic badge](https://img.shields.io/badge/bad%20request-400-red)

~~~javascript
{
    "error": "string"
}
~~~

![Generic badge](https://img.shields.io/badge/anauthorized-401-orange)
~~~json
{
    "error": "string"
}
~~~

![Generic badge](https://img.shields.io/badge/Internal%20Server%20Error-500-red)

~~~json
{
    "error": "string"
}
~~~

### like a post
~~~go
[POST] /posts/{{postID}}/like
~~~
#### **Request headers**
~~~json
{
	"headers": {
		"Authorization": "Bearer token"
	}
}
~~~

#### **Responses**
![Generic badge](https://img.shields.io/badge/noContent-204-<COLOR>.svg)

![Generic badge](https://img.shields.io/badge/bad%20request-400-red)

~~~javascript
{
    "error": "string"
}
~~~

![Generic badge](https://img.shields.io/badge/anauthorized-401-orange)
~~~json
{
    "error": "string"
}
~~~

![Generic badge](https://img.shields.io/badge/Internal%20Server%20Error-500-red)

~~~json
{
    "error": "string"
}
~~~

### unlike a post
~~~go
[POST] /posts/{{postID}}/unlike
~~~
#### **Request headers**
~~~json
{
	"headers": {
		"Authorization": "Bearer token"
	}
}
~~~

#### **Responses**
![Generic badge](https://img.shields.io/badge/noContent-204-<COLOR>.svg)

![Generic badge](https://img.shields.io/badge/bad%20request-400-red)

~~~javascript
{
    "error": "string"
}
~~~

![Generic badge](https://img.shields.io/badge/anauthorized-401-orange)
~~~json
{
    "error": "string"
}
~~~

![Generic badge](https://img.shields.io/badge/Internal%20Server%20Error-500-red)

~~~json
{
    "error": "string"
}
~~~

### Find posts
~~~go
[GET] /posts
~~~
#### **Request headers**
~~~json
{
	"headers": {
		"Authorization": "Bearer token"
	}
}
~~~

#### **Responses**
![Generic badge](https://img.shields.io/badge/OK-200-<COLOR>.svg)

~~~json
[
	{
		"id": number,
		"title": "string",
		"content": "string",
		"authorId": number,
		"likes": number,
		"createdAt": "string"
	}
]
~~~
![Generic badge](https://img.shields.io/badge/bad%20request-400-red)

~~~javascript
{
    "error": "string"
}
~~~

![Generic badge](https://img.shields.io/badge/anauthorized-401-orange)
~~~json
{
    "error": "string"
}
~~~

![Generic badge](https://img.shields.io/badge/Internal%20Server%20Error-500-red)

~~~json
{
    "error": "string"
}
~~~

### Find posts by ID
~~~go
[GET] /posts/{{postID}}
~~~
#### **Request headers**
~~~json
{
	"headers": {
		"Authorization": "Bearer token"
	}
}
~~~

#### **Responses**
![Generic badge](https://img.shields.io/badge/OK-200-<COLOR>.svg)

~~~json
	{
		"id": number,
		"title": "string",
		"content": "string",
		"authorId": number,
		"likes": number,
		"createdAt": "string"
	}
~~~
![Generic badge](https://img.shields.io/badge/bad%20request-400-red)

~~~javascript
{
    "error": "string"
}
~~~

![Generic badge](https://img.shields.io/badge/anauthorized-401-orange)
~~~json
{
    "error": "string"
}
~~~

![Generic badge](https://img.shields.io/badge/Internal%20Server%20Error-500-red)

~~~json
{
    "error": "string"
}
~~~

### Find all post from specific user
~~~go
[GET] /users/{{userID}}/posts
~~~
#### **Request headers**
~~~json
{
	"headers": {
		"Authorization": "Bearer token"
	}
}
~~~

#### **Responses**
![Generic badge](https://img.shields.io/badge/OK-200-<COLOR>.svg)

~~~json
[
	{
		"id": number,
		"title": "string",
		"content": "string",
		"authorId": number,
		"likes": number,
		"createdAt": "string"
	}
]
~~~
![Generic badge](https://img.shields.io/badge/bad%20request-400-red)

~~~javascript
{
    "error": "string"
}
~~~

![Generic badge](https://img.shields.io/badge/anauthorized-401-orange)
~~~json
{
    "error": "string"
}
~~~

![Generic badge](https://img.shields.io/badge/Internal%20Server%20Error-500-red)

~~~json
{
    "error": "string"
}
~~~