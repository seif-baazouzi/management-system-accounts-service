# Update Username

Used to update user username.

**URL**: `/api/v1/settings/update/username`

**Method**: `PUT`

**Auth required**: Send User `JWT token` in `X-Token` header

### Request

```json
{
    "newUsername": "username",
    "password": "password"
}
```

### Success Response

**CODE**: `200`

```json
{
    "message": "success",
}
```

### Invalid Input Response

**CODE**: `200`

```json
{
    "newUsername": "username error message",
    "password": "password error message"
}
```

### Server Error Response

**CODE**: `500`

```json
{
    "message": "server-error",
}
```
