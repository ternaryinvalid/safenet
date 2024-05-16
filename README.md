# Welcome to SAFE.NET

## 1. About project:
* **SAFE.NET** is a project that is designed to provide a secure and reliable way to store and transfer data.
* The project is based on GoLang and uses the Diffie-Hellman key exchange algorithm to encrypt and decrypt data.
* The project is designed to be used in conjunction with the SAFE.NET client, which is a simple and easy-to-use application that allows users to store and transfer data securely.
* Users know only their own private key and public key of the recipient. The server does not store any data, only the public keys of the users.

## 2. Getting started:

To get started with the project, you will need to install the SAFE.NET client and server on your computer.
```sh
git clone https://github.com/ternaryinvalid/safenet.git
```

After that you need to install docker on your computer and run the following command:
```sh
docker-compose up -d
```

After that you can enjoy the project :)

## 3. API Documentation:

### 3.1 Methods

**3.1.1 Send Message**

**URL:**
    `http://localhost:8000/api/v1/message/send`

Request body:
```json5
{
  "message_to":"04a3c57ef403154129402c18ed243ab0d9ff79cbe229da39cbe229da308e079d965361ce2f3a771276f8e04692710dd393ef4b659784093ef4b6597840be2e2ac2a1", // публичный ключ получателя
  "message_from":"04a3ca44f40315412689f0825a59402c18ed243ab0d9ff79cbe229da308e079d965361ce2f3a771276f8e04692710dd393ef4b6597840be2e28781e18ae5f1aca5", // публичный ключ отправителя
  "message_data":"Hello, Bob!"
}
```

Response Body
```json5
{
  "message_id": 1 // уведомление об успешной записи
}
```

**3.1.2 Create Account**

**URL:**
`http://localhost:8000/api/v1/account/create`

**METHOD:** POST

Request body:
```json
{
  "name": "Alice",
  "private_key": "ff791d94aef0efb5f0466b21e8391f68157f798c4555282a996eef957482965a" // опциональное поле для генерации собственных ключей
}
```

**3.1.3 Get Messages**

**URL:**
    `http://localhost:8000/api/v1/message/get`

**METHOD:** GET

Response body:
```json
{
    "messages": [
      {
        "message_from": "0sad1ffe6c2bf123ew", // public key отправителя
        "message_data": "Hello, Bob!",
        "dt": "2024-05-10" // дата отправки сообщения
      },
      {
        "message_from": "aff123eee3981", // public key отправителя
        "message_data": "Hello, Alice!",
        "dt": "2024-05-12" // дата отправки сообщения
      }
    ]
}
```

## Creators:

* Daniil Khiznyakov - [Project Owner] - GO Developer
* Maxim Korobov – [Contributor] - GO Developer | System Architect

## Contributing

1. Fork it (<https://github.com/ternaryivalid/safenet/fork>)
2. Create your feature branch (`git checkout -b feature/fooBar`)
3. Commit your changes (`git commit -am 'Add some fooBar'`)
4. Push to the branch (`git push origin feature/fooBar`)
5. Create a new Pull Request