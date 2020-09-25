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
