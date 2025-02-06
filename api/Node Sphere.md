---
title: Node Sphere
language_tabs:
  - shell: Shell
  - http: HTTP
  - javascript: JavaScript
  - ruby: Ruby
  - python: Python
  - php: PHP
  - java: Java
  - go: Go
toc_footers: []
includes: []
search: true
code_clipboard: true
highlight_theme: darkula
headingLevel: 2
generator: "@tarslib/widdershins v4.0.23"

---

# Node Sphere

Base URLs:

* <a href="http://localhost:8888">dms-be: http://localhost:8888</a>

# Authentication

# dms-be/users

## POST Create New User

POST /users

> Body Parameters

```json
{
  "Names": "{{$person.fullName}}",
  "Email": "{{$internet.email}}",
  "Password": "{{$internet.password}}"
}
```

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|body|body|object| no |none|

> Response Examples

> 201 Response

```json
{
  "data": {
    "User_Id": "string",
    "Name": "string",
    "Email": "string",
    "Created_At": "string"
  }
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|201|[Created](https://tools.ietf.org/html/rfc7231#section-6.3.2)|none|Inline|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|none|Inline|
|409|[Conflict](https://tools.ietf.org/html/rfc7231#section-6.5.8)|none|Inline|

### Responses Data Schema

HTTP Status Code **201**

|Name|Type|Required|Restrictions|Title|description|
|---|---|---|---|---|---|
|» data|object|true|none||none|
|»» User_Id|string|true|none||none|
|»» Name|string|true|none||none|
|»» Email|string|true|none||none|
|»» Created_At|string|true|none||none|

## GET Get New User

GET /users

> Response Examples

> 200 Response

```json
{
  "data": {
    "User_Id": "string",
    "Name": "string",
    "Email": "string",
    "Created_At": "string"
  }
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|
|404|[Not Found](https://tools.ietf.org/html/rfc7231#section-6.5.4)|none|Inline|

### Responses Data Schema

HTTP Status Code **200**

|Name|Type|Required|Restrictions|Title|description|
|---|---|---|---|---|---|
|» data|object|true|none||none|
|»» User_Id|string|true|none||none|
|»» Name|string|true|none||none|
|»» Email|string|true|none||none|
|»» Created_At|string|true|none||none|

## DELETE Delete User

DELETE /users

> Response Examples

> 204 Response

> 404 Response

```json
{}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|204|[No Content](https://tools.ietf.org/html/rfc7231#section-6.3.5)|none|Inline|
|404|[Not Found](https://tools.ietf.org/html/rfc7231#section-6.5.4)|none|Inline|

### Responses Data Schema

# Data Schema

<h2 id="tocS_Errors">Errors</h2>

<a id="schemaerrors"></a>
<a id="schema_Errors"></a>
<a id="tocSerrors"></a>
<a id="tocserrors"></a>

```json
{
  "errors": {
    "timestamp": "string",
    "status": 0,
    "error": "string",
    "message": "string",
    "path": "string"
  }
}

```

### Attribute

|Name|Type|Required|Restrictions|Title|Description|
|---|---|---|---|---|---|
|errors|object|true|none||none|
|» timestamp|string|true|none||none|
|» status|integer|true|none||none|
|» error|string|true|none||none|
|» message|string|true|none||none|
|» path|string|true|none||none|

