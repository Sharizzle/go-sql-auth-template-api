# GoLang Gorm, SQL, Gorm, Mux, JWT Auth REST API Boilerplate

Easy to follow boilerplate for a **Golang** webserver with authentication.

## Initialization

Run the `initialize_users.sql` SQL Command to initialize the database table.

## Packages Used

- Mux (Routing)
- Crypto (Hashig passwords)
- Gorm (Go ORM)
- JWT GO (Authentication)
- SQL Lite Drivers

## Routes

Basic CRUD Routes for User Data

- Show Users `GET /users`
- Create User `POST /users`
- Show User `GET /users/{userId}`
- Delete User `DELETE /users/{userId}`
- Update User `PUT /users/{userId}`
- User Login `POST /users/login`

## Configuration

Create a `.env` file with the following parameters. For SQL Lite only port and JWT secret are required.

PORT = ...  
ENVIRONMENT = ... \
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
