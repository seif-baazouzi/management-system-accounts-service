# Update Password

Used to update user password.

**URL**: `/api/v1/settings/update/password`

**Method**: `PUT`

**Auth required**: Send User `JWT token` in `X-Token` header

### Request

```json
{
    "oldPassword": "old password",
    "newPassword": "new password"
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
    "oldPassword": "old password error message",
    "newPassword": "new password error message"
}
```

### Server Error Response

**CODE**: `500`

```json
{
    "message": "server-error"
}
```
