# Legends

-  ❓: To be discussed

# Global Standard and Terminologies

- Required Key: Required key should be exist. The value should not be `undefined` or `null`. However, empty string is allowed.
- Optional key: Optional key can be ommited. The value might be `undefined` or `null`.

    Frontend need to check the value before parsing the value:

    ```javascript
    if (key in obj && obj[key] !== undefined && obj[key] !== null) {
        console.log('key exists');
    }
    ```

    Backend however is enforced to not give null object:

    ```go
    // status is required, ValidationError is optional
    type ResponseStruct struct {
        Status             string            `json:"status"`
        ValidationError    *ValidationError  `json:"validation_error,omitempty"`
    }    
    ```

- Access Token and Refresh Token
    - Frontend uses Access Token to authenticate a user.
    - Access Token is used because we don't want to send username and password back and forth everytime.
    - Access Token has expiration, either embeded in the token (use non temperable mechanism) or mapped by backend.
    - Refresh Token has longer expiration than Access Token
    - When Frontend get 401 unauthorized from the backend, it will send refresh-token request, and Backend will send the new tokens (refresh and access)

    

# Common Response

Common response contains of the following fields:

- `status` (required, string): Error code, or `SUCCESS`.
- `errorMessage` (optional, string): User readable error message.
- `data` (required, any):
    - For `getting list of entities`, this should return a list.
    - For `getting a single entity`, `update`, or `delete`, this should return an object.
    - Otherwise, we will define this per template.
- `paging` (optional, required only when `getting list of entities`)
    - `page` (required, int): Page index
    - `total` (required, int): Total rows/data.

# Common CRUD

## GET `/api/v1/<entity>`

Getting list of entities.

Header

```
Authorization: Bearer <token>
```

Query Param

- `limit` (optional, int).
- `offset` (optional, int).
- ❓ `filters[<field>][<operator>]` (optional, any): Field filter (see: [strapi](https://docs.strapi.io/dev-docs/api/rest/filters-locale-publication)).
    - Example: `?filters[username][$eq]=ari&filters[age][$gt]=20`
    - List of operators:
        - `$eq`
        - `$ne`
        - `$gt`
        - `$gte`
        - `$lt`
        - `$lte`
        - `$contains`
        - `$notContains`
- ❓ `sortBy` (optional, string): Field name
- ❓ `sortOrder` (optional, string, default: `asc`): Either `asc` or `desc`.


Response

```json
{
    "status": "INVALID_SOMETHING",
    "errorMessage": "",
    "data": [
        {} // object
    ] 
}
```

## GET `/api/v1/<entity>/<id>`

Getting a single entity.

Header

```
Authorization: Bearer <token>
```

Query Param

- `<field>`: `str`


Response

```json
{
    "status": "INVALID_SOMETHING",
    "errorMessage": "",
    "data": {} // object 
}
```

## POST `/api/v1/<entity>`

Inserting a new entity.

Header

```
Authorization: Bearer <token>
```

Request

```json
{
    "<field>": "value",
}
```

Response

```json
{
    "status": "INVALID_SOMETHING",
    "errorMessage": "",
    "data": {
        "id": "<new-id>",
        "<field>": "value"
    }
}
```

## PUT `/api/v1/<entity>/<id>`

Updating an existing entity, must provide all fields.

Header

```
Authorization: Bearer <token>
```

Request

```json
{
    "<field>": "value",
}
```

Response

```json
{
    "status": "INVALID_SOMETHING",
    "errorMessage": "",
    "data": {
        "id": "<existing-id>",
        "<field>": "value"
    }
}
```

## PATCH `/api/v1/<entity>/<id>`

Updating an existing entity, doesn't have to provide all fields.

> Is this needed?

Header

```
Authorization: Bearer <token>
```

Request

```json
{
    "<field>": "value",
}
```

Response

```json
{
    "status": "INVALID_SOMETHING",
    "errorMessage": "",
    "data": {
        "id": "<existing-id>",
        "<field>": "value"
    },
}
```

## DELETE `/api/v1/<entity>/<id>`

Deleting an existing entity.

> Is this needed?

Header

```
Authorization: Bearer <token>
```

Response

```json
{
    "status": "INVALID_SOMETHING",
    "errorMessage": "",
    "data": {
        "id": "<existing-id>",
        "<field>": "value"
    },
}
```

# Login

## POST `/user/login`

Request

```json
{
    "user_id": "",
    "password": "", // RFC: hash??
}
```

Response

```json
{
    "status": "INVALID_SOMETHING",
    "errorMessage": "",
    "data": {
        "token" : "",
        "refreshToken": "",
    },
}
```

## POST `/user/refresh-token`

Header

```
Authorization: Bearer <token>
```

Request

```json
{
    "refreshToken": "",
}
```

Response

```json
{
    "status": "INVALID_SOMETHING",
    "errorMessage": "",
    "data": {
        "token" : "",
        "refreshToken": "",
    }
}
```

# User Matrix

## POST `/user/matrix`

> This should be `GET`, but currently we use `POST` since query parameter has size limitation.

Header

```
Authorization: Bearer <token>
```

Request

```json
{
    "refreshToken": "",
    "permissions": ["permission-1", "permission-2"],
}
```

Response

```json
{
    "status": "INVALID_SOMETHING",
    "errorMessage": "",
    "data": {
        "<permission-1>": true,
        "<permission-2>": false,
    }
}
```

# User, Role, and Permission

Definition:

- User representing a single user. 
- Role representing a user role (e.g., `administrator`, `tester`, etc.).
- Permission representing a granular authorization (e.g., `read_payment`, `create_payment`, etc.).

Ordinality:

- A single user might has multiple roles.
- A single role might be assigned to multiple users.
- A single role might contains multipe permissions.
- A single permission might be assigned to multiple roles.

## User

URL: 

- GET, POST, PUT, DELETE, PATCH: `/api/v1/users`: User CRUD.

Data:

```json
{
    "id": "",
    "username": "",
}
```

> To get user roles/permissions, use role/permission API

## Role

URL:

- GET, POST, PUT, DELETE, PATCH: `/api/v1/roles`: Role CRUD.

Data:

```json
{
    "id": "",
    "role": "",
}
```

## Permission

URL:

- GET, POST, PUT, DELETE, PATCH: `/api/v1/permissions`: Permission CRUD.

Data:

```json
{
    "id": "",
    "permission": "",
}
```
