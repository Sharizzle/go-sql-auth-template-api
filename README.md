# GoLang Gorm, SQL, Gorm, Mux, JWT Auth REST API Boilerplate

Easy to follow boilerplate for a **Golang** webserver with authentication.

## Packages Used

- Mux (Routing)
- Crypto (Hashig passwords)
- Gorm (Go ORM)
- JWT GO (Authentication)
- SQL Lite Drivers

For a postgres version, see the [Postgres]() Branch

## What's included

Basic CRUD Routes for User Data

- Show Users `GET /users`
- Show User `GET /users/{userId}`
- Create User `POST /users`
- User Login `POST /users/login`
- Delete User `DELETE /users/{userId}`
- Update User `PUT /users/{userId}`

## Configuration

Create a `.env` file with the following parameters

ENVIRONMENT = ... \
PORT = ... \
DB_HOST = ...\
DB_NAME = ...\
DB_USERNAME = ...\
DB_PASSWORD = ...\
JWT_SECRET = ...

## Installation

Run the command to install all dependancies.

```go
go mod download
```

## Authors

- **Sharif Kanaan** - [GitHub](https://github.com/Sharizzle) | | [Website](https://sharif.thekanaan.com/) | | [Email](mailto:sharif@thekanaan.com) || [Linkedin](https://www.linkedin.com/in/SharifKanaan/)

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
