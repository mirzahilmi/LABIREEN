
![LABIREEN_README](https://user-images.githubusercontent.com/118604764/223939003-13621201-21db-402d-8eb5-775420ec699d.png)

# LABIREEN
> Food Ordering Management System Application

LABIREEN is an acronym of Colaboration and Integration of FILKOM Canteen

## Installing / Getting started

Before starting this project on your local environment, make sure you have [Git][Git Website] and [Go][Go Website] programming language installed.

```shell
git clone https://github.com/MirzaHilmi/LABIREEN.git
cd server/internal/auth
go run cmd/main.go
```

You can now send HTTP requests to your localhost using [Postman][Postman Website] or any other API platform.

## Endpoint Reference
The following endpoints are available through

`https://mirzahlm.aenzt.tech/`

### Authentication
**Description** : Customer account authentication that includes registration, login, and email verification

#### POST `{url}/auth/register`
**Parameters**
| Parameter        | Type   | Required | Description                                                     |
|------------------|--------|----------|-----------------------------------------------------------------|
| name             | string | YES      | Customer full name                                              |
| email            | string | YES      | Customer Customer email (should be a valid email)                          |
| password         | string | YES      | Customer Account Password (should be at least 8 character long) |
| password_confirm | string | YES      | The value should be the same as "password" field                |
| phone_number     | string | YES      | Customer Account phone number (max 15 character long)           |

**Response**
```json
{
    "status": "success",
    "code": 200,
    "message": "User successfuly created, please check your email for email verification",
    "data": {
        "name": " ",
        "email": " ",
        "password": " ",
        "password_confirm": " ",
        "phone_number": " ",
        "verification_code": " "
    }
}
```
```json
{
    "status": "error",
    "code": 500,
    "message": "Failed to register user",
    "data": "Error 1062 (23000): Duplicate entry '123456789' for key 'customers.phone_number'"
}
```

#### POST `{url}/auth/login`
**Parameters**
| Parameter | Type   | Required | Description                            |
|-----------|--------|----------|----------------------------------------|
| email     | string | YES      | Customer registered and verified email |
| password  | string | YES      | Customer registered Account password   |

**Response**
```json
{
    "status": "success",
    "code": 200,
    "message": "Login Successful",
    "data": " here is the jwt token "
}
```

```json
{
    "status": "error",
    "code": 401,
    "message": "Failed to logged in",
    "data": "user has not verified"
}
```
## Features

What's all the bells and whistles this project can perform?
* What's the main functionality
* You can also do another thing
* If you get really randy, you can even do this

## Configuration

Here you should write what are all of the configurations a user can enter when
using the project.

#### Argument 1
Type: `String`  
Default: `'default value'`

State what an argument does and how you can use it. If needed, you can provide
an example below.

Example:
```bash
awesome-project "Some other value"  # Prints "You're nailing this readme!"
```

#### Argument 2
Type: `Number|Boolean`  
Default: 100

Copy-paste as many of these as you need.

## Contributing

When you publish something open source, one of the greatest motivations is that
anyone can just jump in and start contributing to your project.

These paragraphs are meant to welcome those kind souls to feel that they are
needed. You should state something like:

"If you'd like to contribute, please fork the repository and use a feature
branch. Pull requests are warmly welcome."

If there's anything else the developer needs to know (e.g. the code style
guide), you should link it here. If there's a lot of things to take into
consideration, it is common to separate this section to its own file called
`CONTRIBUTING.md` (or similar). If so, you should say that it exists here.

## Links

Even though this information can be found inside the project on machine-readable
format like in a .json file, it's good to include a summary of most useful
links to humans using your project. You can include links like:

- Project homepage: https://your.github.com/awesome-project/
- Repository: https://github.com/your/awesome-project/
- Issue tracker: https://github.com/your/awesome-project/issues
  - In case of sensitive bugs like security vulnerabilities, please contact
    my@email.com directly instead of using issue tracker. We value your effort
    to improve the security and privacy of this project!
- Related projects:
  - Your other project: https://github.com/your/other-project/
  - Someone else's project: https://github.com/someones/awesome-project/


## Licensing

One really important part: Give your project a proper license. Here you should
state what the license is and how to find the text version of the license.
Something like:

"The code in this project is licensed under MIT license."

[Git Website]: https://git-scm.com/
[Go Website]: https://go.dev/
[Postman Website]: https://www.postman.com/
