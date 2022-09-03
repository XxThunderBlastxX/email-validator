# Email-Validator

> Made using Go and Fiber

### Prerequisite before running the application

> Create .env in root dir of the project and copy all the variables from .env.example and past in .env and assign the values of all the variables.

### How to run the application

- To run the application in development mode:

```shell
go run .
```

- To make a build of the application:

```shell
go build -o email-validator
```

- To make a build of the application for windows:

```shell
go build -o email-validator.exe
```

### How to send email

- You will have to ***POST*** request to `/valid` route with body as -
```json
{
    "email": "Email"
}
```

- Success Response we receive - 
```json
{
    "status": true,
    "error": "",
    "contains_mx": true,
    "mx_record": [
        "Array of MX Records"
    ],
    "contains_spf": true,
    "spf": "SPF Record",
    "contains_dmarc": true,
    "dmarc": "DMARC Record"
}
```

- Error Response - 
```json
{
    "status": false,
    "error": "Error Message"
}
```

