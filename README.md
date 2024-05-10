# Welcome to SAFE.NET

## About project:
* **SAFE.NET** is a project that is designed to provide a secure and reliable way to store and transfer data.
* The project is based on GoLang and uses the Diffie-Hellman key exchange algorithm to encrypt and decrypt data.
* The project is designed to be used in conjunction with the SAFE.NET client, which is a simple and easy-to-use application that allows users to store and transfer data securely.
* Users know only their own private key and public key of the recipient. The server does not store any data, only the public keys of the users.

## Getting started:

To get started with the project, you will need to install the SAFE.NET client and server on your computer.
```sh
git clone https://github.com/ternaryinvalid/safenet.git
```

After that you need to install docker on your computer and run the following command:
```sh
docker-compose up -d
```

After that you can enjoy the project :)

## API Documentation:

### Methods

**Save Message**

**URL:**
    `http://localhost:8081/api/v1/send`

**METHOD:** POST

Request body:
```json
{
    "public_key_to": "0x4283104b22a688f347b946462cd62711ef68151deab79845f77fb365f15c0be4",
    "message": "Hello, World!"
}
```

**Get Messages**

**URL:**
    `http://localhost:8081/api/v1/messages`

**METHOD:** POST

Request body:
```json
{
    "public_key": "0x4283104b22a688f347b946462cd62711ef68151deab79845f77fb365f15c0be4",
    "limit": 10 // Optional, if you dont set limit, it will return all messages
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