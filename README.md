# Backend for IOICamp 2020 Website

## API

All **POST** parameters should be in JSON format.

Response will always be in JSON format.

All response contains `status` field. The value is either *success* or *failed*. When the value is *failed* there's also a field `error` in string format representing the error.

### Email verification request

> POST /api/get-verification-token

Request for a new email verification mail. Will be sent directly to the specified email address. The default rate limit is one request per email address per minute.

#### Parameters

| Name | Description |
|---|---|
| `email` | email address to be verified |

#### Response

N/A

### Register request

> POST /api/register

Register a new account.

#### Parameters

| Name | Description |
|---|---|
| `email` | email address |
|  `password` | password |
| `token` | email verification token |

#### Response

N/A

### Login

> POST /api/login

Login request.

#### Parameters

| Name | Description |
|---|---|
| `email` | email address |
|  `password` | password |

#### Response

| Name | Description |
|---|---|
| `token` | JWT Token |

### Apply Form

#### Form Fields

| Name |
|---|
| `name` |
| `gender` |
| `school` |
| `grade` |
| `code-time` |
| `cp-time` |
| `prize` |
| `oj` |
| `motivation` |
| `comment` |

#### Authorization

Put JWT Token in `Authorization` header (Bearer Authentication).

> GET /api/users/apply-form

No parameter needed. Response:

| Name | Descriptioin |
| --- | --- |
| `applyForm` | A JSON of apply form data |
| `email` | The user's email address |

> PUT /api/users/apply-form

Update form, no additional response.

### Change Password

> POST /api/users/change-password

#### Authorization

Put JWT Token in `Authorization` header (Bearer Authentication).

#### Parameters

| Name | Description |
|---|---|
| `old-password` | Original password |
| `new-password` | New password |

#### Response

N/A

### Request Password Reset Token

> POST /api/get-password-reset-token

#### Parameters

| Name | Description |
|---|---|
| `email` | Email address |

#### Response

N/A

### Password Reset

> POST /api/password-reset

#### Parameters

| Name | Description |
|---|---|
| `token` | Password reset token |
| `new-password` | New password |

#### Response

N/A
