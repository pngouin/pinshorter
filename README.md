![Logo](./www/dist/assets/images/pinshorter.png "Logo")

# Pinshorter

A URL shortener in Go and Angular. [https://pinshorter.herokuapp.com/JW3b7](https://pinshorter.herokuapp.com/JW3b7)

![Login](https://i.imgur.com/Js0583F.png "Login")
![Site](https://i.imgur.com/Of5ndEE.png "Site")

## Getting started

Pinshorter only needs a Postgresql database to work.

### Environment variable

| Name | Description |
| :---: | :---: |
| DATABASE_URL | The connection querystring to authenticate to the postgresql database |
| PINSHORTER_SECRET | The secret to signing JWT tokens |
| PORT | Must be a number. The port where the apllication is listening |

### CLI

The CLI has only one argument

| Name | Description |
| :---: | :---: |
| dev | Start the server in dev mode, CORS allow all origin |

