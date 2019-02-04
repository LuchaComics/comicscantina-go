# HOWTO: Setup ComicsCantian Data on DigitalOcean using CentOS 7 OS

## Instructions

### Setup Web-App Database

1. While being logged in as ``techops`` run the following:

    ```
    $ sudo -i -u postgres
    $ psql
    ```

2. Then run the following.

    ```sql
    drop database comicscantina_db;
    create database comicscantina_db;
    \c comicscantina_db;
    CREATE USER golang WITH PASSWORD 'YOUR_PASSWORD';
    GRANT ALL PRIVILEGES ON DATABASE comicscantina_db to golang;
    ALTER USER golang CREATEDB;
    ALTER ROLE golang SUPERUSER;
    CREATE EXTENSION postgis;
    ```


### Setup Web-App from GitHub
Please run the following commands as the ``lucha`` user account.

1. Get the project.

    ```
    $ go get github.com/luchacomics/comicscantina-go
    ```

2. Install the dependencies.

    ```
    $ cd /opt/lucha/go/src/github.com/luchacomics/comicscantina-go/;
    $ ./requirements.sh;
    ```

3. Run the the following environment variables. **Please change the variables to meet your own.**

    ```
    export COMICSCANTINA_GORM_CONFIG="host=localhost port=5432 user=golang dbname=comicscantina_db password=YOUR_PASSWORD sslmode=disable"
    export COMICSCANTINA_SECRET="YOUR_SECRET_RANDOM_STRING"
    export COMICSCANTINA_ADDRESS="127.0.0.1:8080"  # Do not change!!!
    export COMICSCANTINA_UNIT_TEST_GORM_CONFIG="host=localhost port=5432 user=golang dbname=comicscantina_test_db password=YOUR_PASSWORD sslmode=disable"
    ```

5. Build our project.

   ```
   $ go install github.com/luchacomics/comicscantina-go
   ```

6. Enable permission and security while you are a ``techops`` user.

    ```
    $ sudo setcap 'cap_net_bind_service=+ep' /opt/lucha/go/bin/comicscantina-go
    $ sudo setsebool -P httpd_can_network_connect 1
    $ sudo semanage permissive -a httpd_t
    $ sudo chcon -Rt httpd_sys_content_t /opt/django/workery-django/workery/static
    ```

### Integrate Nginx with Golang
Please run the following commands as the ``techops`` user account.

1. Load up ``Nginx``.

   ```
   sudo vi /etc/nginx/nginx.conf
   ```

2. Replace with the following code.

    ```
    server {
        listen       80;
        server_name  SERVER_DOMAIN_NAME_OR_IP;

        charset utf-8;

        location / {
            proxy_set_header X-Forwarded-For $remote_addr;
            proxy_set_header Host            $http_host;

            proxy_pass http://127.0.0.1:8080;
        }
    }
    ```

3. Restart ``Nginx``.

    ```
    sudo systemctl restart nginx
    ```

4. Run your go app manually.

    ```
    # go run github.com/luchacomics/comicscantina-go
    ```

5. Now in your browser go to ``http://SERVER_DOMAIN_NAME_OR_IP`` and you should see the app!

6. Special thanks to:

* https://beego.me/docs/deploy/nginx.md

### Integrate Systemd with Golang

This section explains how to integrate our project with ``systemd`` so our operating system will handle stopping, restarting or starting.

1. (OPTIONAL) If you cannot access the server, please stop and review the steps above. If everything is working proceed forward.

2. While you are logged in as a ``techops`` user, please write the following into the console.

    ```
    $ sudo vi /etc/systemd/system/comicscantina-go.service
    ```

3. Implement

    ```
    [Unit]
    Description=ComicsCantina Data Microservice
    Wants=network.target
    After=network.target

    [Service]
    Type=simple
    DynamicUser=yes
    WorkingDirectory=/opt/lucha/go/bin
    ExecStart=/opt/lucha/go/bin/comicscantina-go
    Restart=always
    RestartSec=3
    SyslogIdentifier=comicscantina_data

    [Install]
    WantedBy=multi-user.target
    ```

4. Grant access.

   ```
   $ sudo chmod 755 /etc/systemd/system/comicscantina-go.service
   ```

5. (Optional) If you've updated the above, you will need to run the following before proceeding.

    ```
    $ systemctl daemon-reload
    ```

6. We can now start the Gunicorn service we created and enable it so that it starts at boot:

    ```
    $ sudo systemctl start comicscantina-go
    $ sudo systemctl enable comicscantina-go
    ```

7. Confirm our service is running.

    ```
    $ systemctl status comicscantina-go.service
    $ journalctl -f -u comicscantina-go.service
    ```

8. And verify the URL works in the browser.

    ```text
    http://SERVER_DOMAIN_NAME_OR_IP/en/
    ```
