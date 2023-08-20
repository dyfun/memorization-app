Developing a simple api with go and mysql.

## Installation for development
First, you need to install the dependencies
```bash
go mod download
```
Then, you need to create a .env file
```bash
cp .env.example .env
``` 
Finally, you can run the project
```bash
go run main.go start
```
note: You can use the "seed" and "migrate" arguments.
## Hot reload
If you want to use air for hot reload, you can run
```bash
air init
```
After that, you can add "start" argument to the .air.conf file
``` bash
[build]
    args_bin = ["start"]
    ...
```
And run
```bash
air
```
## Swagger
If you want to reload the swagger documentation, you can run following command

```bash
swag init
```

## Docker
```bash
docker-compose up --build
```