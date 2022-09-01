# Get User

Used to get user userID and username.

**URL**: `/api/v1/user`

**Method**: `GET`

**Auth required**: Send User `JWT token` in `X-Token` header

### Success Response

**CODE**: `200`

```json
{
    "userID": "user uuid",
    "username": "username"
}
```

### Server Error Response

**CODE**: `500`

```json
{
    "message": "server-error"
}
```
