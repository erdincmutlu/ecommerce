# Ecommerce
Ecommerce with Golang

```bash
# You can start the project with below commands
docker-compose up -d
go run main.go
```

- **SIGNUP**

POST http://localhost:8000/users/signup

```json
{
    "first_name": "Erdinc",
    "last_name": "Mutlu",
    "email": "erdinc@mutlu.com",
    "password": "erdincpassword",
    "phone": "+441234567890"
}
```

Response :"Successfully Signed Up!!"
