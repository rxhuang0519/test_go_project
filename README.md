# test_go_project

## Description

homework
## Required

- `ngrok`
- `docker`

## Setup
#### ENV
- `configs/config.env`
    fill in the following environments.
    - `LINE_CHANNEL_SECRET`
    - `LINE_CHANNEL_TOKEN`

```
ENV="local"
DB_USER="admin"
DB_PASSWORD="root"
DB_HOST="localhost"
DB_PORT="27017"
LINE_CHANNEL_SECRET=""
LINE_CHANNEL_TOKEN=""
LINE_USER_ID="
```
#### Webhook

1. setting ngrok
    start forwarding to `http://localhost:8080`.
    execute the following command:

    ```
    ngrok http 8080
    ```
2.  update the webhook url in line developer console.

## Start
#### Loacl

```
make run
```
#### Query Messages
##### Get all messages

- GET http://localhost:8080/messages
##### Get messages by userId

- GET http://localhost:8080/messages/usr/{userId}
- GET http://localhost:8080/messages?usr={userId}
##### Get messages by messageId

- GET http://localhost:8080/messages/{messageId}
- GET http://localhost:8080/messages?id={messageId}

# Demo Video

- https://youtu.be/ft4sU-A5pyQ