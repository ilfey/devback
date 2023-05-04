# devback

Бекэнд написанный на фреймворке gin

## endpoints

- `/` GET
- `/api/v1/messages` GET
- `/api/v1/user/login` POST
- `/api/v1/user/register` POST
- `/api/v1/user/delete` POST
- `/api/v1/user/message` POST
- `/api/v1/user/message/:id` PATCH, DELETE
- `/api/v1/<admin path>/user/message/:id` PATCH, DELETE
- `/api/v1/<admin path>/user/message/:id/permanently` DELETE
- `/api/v1/<admin path>/user/message/:id/restore` POST
- `/api/v1/<admin path>/user/:username` DELETE
- `/api/v1/<admin path>/user/:username/permanently` DELETE
- `/api/v1/<admin path>/user/:username/restore` POST

## Текущая схема бд

```mermaid
---
title: "devpage database scheme"
---

erDiagram
    users {
        varchar(16) username PK
        varchar(64) password
        boolean is_deleted
        timestamp created_at
    }

    messages {
        bigint id PK
        varchar(2000) content
        varchar(16) userId FK
        bigint reply FK
        boolean is_deleted
        timestamp modified_at
        timestamp created_at
    }

    users only one to zero or many messages : ""
```

## Запланированная схема бд

```mermaid
---
title: "devpage database scheme"
---

erDiagram
    users {
        varchar(16) username PK
        varchar(64) password
        bigint profile FK
        boolean is_deleted
        timestamp created_at
    }

    messages {
        bigint id PK
        varchar(2000) content
        varchar(16) userId FK
        bigint reply FK
        boolean is_deleted
        timestamp modified_at
        timestamp created_at
    }

    attachments {
        bigint messageId PK, FK
        bigint url PK, FK
    }

    projects {
        bigint id PK
        varchar title
        text description
        bigint source FK
        bigint url FK
        varchar language
        timestamp updatedAt
        timestamp createdAt
    }

    contacts {
        bigint id PK
        varchar title
        varchar logo
        bigint linkUrl FK
        bool isPrimary
    }

    links {
        bigint id PK
        varchar description
        varchar url
    }

    users only one to only one links : ""
    users only one to zero or many messages : ""
    messages only one to zero or many attachments : ""
    links only one to zero or many attachments : ""
    links only one to zero or many projects : ""
    links only one to zero or many contacts : ""

```
