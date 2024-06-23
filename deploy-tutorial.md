# Deploy this Service

- install postgres in fedora ssh
    - user: postgres
    - password: postgres
    - sudo -u postgres psql
    - source: https://docs.fedoraproject.org/en-US/quick-docs/postgresql/


- install golang 
    - source: https://developer.fedoraproject.org/tech/languages/go/go-installation.html
    ------- doesnt work, because it have to go 17, and if download with dnf u got 16.5 -------

    - install wget to download go source from official website
    - download with wget `wget https://go.dev/dl/go1.22.2.src.tar.gz`
    - set environtment variable ('ve been noted in notion)

- setup environtment variable on project

- Postgres Problem
    - cannot connect, it should enable ssl/tls connection
    - Ident authentincation failed for user '...'
        - https://serverfault.com/questions/406606/postgres-error-message-fatal-ident-authentication-failed-for-user]

- running go app
