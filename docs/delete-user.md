# Delete User

Used to delete user account.

**URL**: `/api/v1/settings/delete-user`

**Method**: `DELETE`

**Auth required**: Send User `JWT token` in `X-Token` header

### Request

```json
{
    "password": "password"
}
```

### Success Response

**CODE**: `200`

```json
{
    "message": "success"
}
```

### Invalid Input Response

**CODE**: `200`

```json
{
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
