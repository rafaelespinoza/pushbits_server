[![Build status](https://img.shields.io/travis/eikendev/pushbits/master)](https://travis-ci.com/github/eikendev/pushbits/builds/)
[![Docker Pulls](https://img.shields.io/docker/pulls/eikendev/pushbits)](https://hub.docker.com/r/eikendev/pushbits)
![License](https://img.shields.io/github/license/eikendev/pushbits)

## About

PushBits is a relay server for push notifications.
It enables your services to send notifications via a simple web API, and delivers them to you through [Matrix](https://matrix.org/).
This is similar to what [PushBullet](https://www.pushbullet.com/), [Pushover](https://pushover.net/), and [Gotify](https://gotify.net/) offer, but a lot less complex.

The vision is to have compatibility with Gotify on the sending side, while on the receiving side an established service is used.
This has the advantages that
- sending plugins written for Gotify (like those for [Watchtower](https://containrrr.dev/watchtower/) and [Jellyfin](https://jellyfin.org/)) as well as
- receiving clients written for the messaging service can be reused.

### Why Matrix instead of X?

For now, only [Matrix](https://matrix.org/) is supported, but support for different services like [Telegram](https://telegram.org/) could be added in the future.
[WhatsApp](https://www.whatsapp.com/) and [Signal](https://signal.org/) unfortunately do not have an API through which PushBits can interact.

I am myself experimenting with Matrix currently because I like the idea of a federated, synchronized but still end-to-end encrypted protocol.
If you haven't tried it yet, I suggest you to check it out.

## Usage

PushBits is meant to be self-hosted.
You are advised to install PushBits behind a reverse proxy and enable TLS.

At the moment, there is no front-end implemented.
New users and applications need to be created via the API.
Details will be made available once the interface is more stable.

To get started, here is a Docker Compose file you can use.
```yaml
version: '2'

services:
    server:
        image: eikendev/pushbits:latest
        ports:
            - 8080:8080
        environment:
            PUSHBITS_DATABASE_DIALECT: 'sqlite3' # Can use either 'sqlite3' or 'mysql'.
            PUSHBITS_ADMIN_MATRIXID: '@your/matrix/username:matrix.org' # The matrix account on which the admin will receive their notifications.
            PUSHBITS_ADMIN_PASSWORD: 'your/matrix/password' # The login password of the admin for PushBits. Default username is 'admin'.
            PUSHBITS_MATRIX_USERNAME: 'your/pushbits/username' # The matrix account from which PushBits notifications are sent to users.
            PUSHBITS_MATRIX_PASSWORD: 'your/pushbits/password' # The password of the above account.
        volumes:
            - /etc/localtime:/etc/localtime:ro
            - /etc/timezone:/etc/timezone:ro
            - ./mount/data:/data
```

## Development

PushBits is currently in alpha stage.
The API is neither stable, nor is provided functionality guaranteed to work.
Stay tuned! 😉

## Acknowledgments

The idea for this software and most parts of the initial source are heavily inspired by [Gotify](https://gotify.net/).
Many thanks to [jmattheis](https://jmattheis.de/) for his well-structured code.
