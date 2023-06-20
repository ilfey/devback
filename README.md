# devback

Бекэнд написанный на фреймворке gin

---

## endpoints

- `/` GET
- `api/v1/ping` GET
- `/api/v1/messages` GET
- `/api/v1/contacts` GET
- `/api/v1/contacts/:id` GET
- `/api/v1/users/login` POST
- `/api/v1/users/register` POST
- `/api/v1/users/delete` POST
- `/api/v1/users/message` POST
- `/api/v1/users/message/:id` PATCH, DELETE
- `/api/v1/<admin path>/users/messages/:id` PATCH, DELETE
- `/api/v1/<admin path>/users/messages/:id/permanently` DELETE
- `/api/v1/<admin path>/users/messages/:id/restore` POST
- `/api/v1/<admin path>/users/:username` DELETE
- `/api/v1/<admin path>/users/:username/permanently` DELETE
- `/api/v1/<admin path>/users/:username/restore` POST
- `/api/v1/<admin path>/links` GET, POST
- `/api/v1/<admin path>/links/:id` GET, DELETE
- `/api/v1/<admin path>/links/:id/permanently` DELETE
- `/api/v1/<admin path>/links/:id/restore` POST
- `/api/v1/<admin path>/contacts` POST

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

    links {
        bigint id PK
        varchar description
        varchar url
        boolean is_deleted
        timestamp modified_at
        timestamp created_at
    }

    contacts {
        bigint id PK
        varchar title
        bigint linkId FK
        boolean is_deleted
        timestamp modified_at
        timestamp created_at
    }

    
    links only one to zero or many contacts : ""
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

    links {
        bigint id PK
        varchar description
        varchar url
    }

    contacts {
        bigint id PK
        varchar title
        bigint linkId FK
        bool isPrimary
    }

    users only one to only one links : ""
    users only one to zero or many messages : ""
    messages only one to zero or many attachments : ""
    links only one to zero or many attachments : ""
    links only one to zero or many projects : ""
    links only one to zero or many contacts : ""

```


## Todo

- [ ] добавить PATCH для links
- [ ] добавить в find для юзера modified_at и мигрировать
- [ ] добавить остальные методы для contacts
- [ ] добавить уведомления по e-mail, хотя бы для админа
  - [ ] добавить возможность смены пароля пользователям
- [ ] добавить сущности user новый атрибут profile FK links(id)