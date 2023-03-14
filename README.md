
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

Here is a list of features available in this project :
* Customer Account Sign-up
* Customer Account Sign-in
* Email Verification for Customer Accounts
* Customer Account Profile Editing

## Links

- Project Author: https://github.com/MirzaHilmi/
- Repository: https://github.com/MirzaHilmi/LABIREEN
- Issue tracker: https://github.com/MirzaHilmi/LABIREEN/issues
  - In case of sensitive bugs like security vulnerabilities, please contact
   exquisitemirza@gmail.com directly instead of using issue tracker. We value your effort
    to improve the security and privacy of this project!
- Related projects :
  - My other project : 
    - https://github.com/MirzaHilmi/LABIREEN-Customer
    - https://github.com/MirzaHilmi/LABIREEN-Merchant


## Licensing

This project is licensed under the terms of the MIT License.

**What does this mean?**

The MIT License is a permissive open source license that allows you to use, copy, modify, merge, publish, distribute, and/or sell copies of the software, as long as you include the original copyright and license notice in any copies or substantial portions of the software.

**Why did we choose the MIT License?**

We chose the MIT License because it is a widely recognized and well-understood open source license that provides a good balance between permissiveness and protection. We want to encourage the use, modification, and distribution of our software, while also protecting our contributors and users.

**Can I contribute?**

Yes! We welcome contributions from anyone, under the same terms as the license. Please read our Contributing Guidelines for more information.

[Git Website]: https://git-scm.com/
[Go Website]: https://go.dev/
[Postman Website]: https://www.postman.com/
