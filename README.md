# devback

Бекэнд написанный на фреймворке gin

## endpoints

- `/` GET
- `/messages` GET
- `/message` POST
- `/user/login` POST
- `/user/register` POST

## Текущая схема бд

```mermaid
---
title: "devpage database scheme"
---

erDiagram
    users {
        varchar username PK
        varchar password
        timestamp created_at
    }

    user_messages {
        bigint userId PK, FK
        bigint messageId PK, FK
    }

    messages {
        bigint id PK
        varchar(2000) content
        bigint authorId
        bigint reply FK
        timestamp modified_at
        timestamp created_at
    }

    users only one to zero or many user_messages : ""
    messages only one to zero or many user_messages : ""
```

## Запланированная схема бд

```mermaid
---
title: "devpage database scheme"
---

erDiagram
    users {
        varchar username PK
        varchar password
        timestamp created_at
    }

    user_messages {
        bigint userId PK, FK
        bigint messageId PK, FK
    }

    messages {
        bigint id PK
        varchar(2000) content
        bigint authorId
        bigint reply FK
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
    users only one to zero or many user_messages : ""
    messages only one to zero or many user_messages : ""
    messages only one to zero or many attachments : ""
    links only one to zero or many attachments : ""
    links only one to zero or many projects : ""
    links only one to zero or many contacts : ""

```
