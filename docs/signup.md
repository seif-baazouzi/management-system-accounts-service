# Signup

Used to create user and his token.

**URL**: `/api/v1/signup`

**Method**: `POST`

**Auth required**: no

### Request

```json
{
    "username": "username",
    "password": "password"
}
```

### Success Response

**CODE**: `200`

```json
{
    "message": "success",
    "token": "User JWT Token"
}
```

### Invalid Input Response

**CODE**: `200`

```json
{
    "username": "username error message",
    "password": "password error message"
}
```

### Server Error Response

**CODE**: `500`

```json
{
    "message": "server-error"
}
```
