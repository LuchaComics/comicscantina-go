# HOWTO: Setup ComicsCantian Data on DigitalOcean using CentOS 7 OS

## Instructions

### Setup from GitHub
Please run the following commands as the ``lucha`` user account.

1. Clone the project.

    ```
    $ mkdir /opt/lucha/go/src;
    $ cd /opt/lucha/go/src;
    $ git clone https://github.com/LuchaComics/comicscantina-data.git
    $ cd comicscantina-data
    ```

2. Install the dependencies.

    ```
    $ ./requirements.sh
    ```

3. Open the config file again.

    ```
    $ vi /opt/lucha/go/bin/.env
    ```

4. Append the following environment variables. **Please change the variables to meet your own.**

    ```
    COMICS_WS_GORM_CONFIG='host=localhost port=5432 user=lucha dbname=comics_db password=123password sslmode=disable'
    COMICS_WS_SECRET='YOUR_APPS_SECRET_KEY'
    COMICS_WS_ADDRESS='127.0.0.1:8080'  # Do not change!
    ```

5. Build our project.

   ```
   $ cd /opt/lucha/go/src;
   $ go install comicscantina-data;
   ```

6. Enable permission and security while you are a ``techops`` user.

    ```
    $ sudo setcap 'cap_net_bind_service=+ep' /opt/lucha/go/bin/comicscantina-data
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
    systemctl restart nginx
    ```

4. Run your go app manually.

    ```
    # go run comicscantina-data
    ```

5. Now in your browser go to ``http://SERVER_DOMAIN_NAME_OR_IP`` and you should see the app!

6. Special thanks to:

* https://beego.me/docs/deploy/nginx.md

### Integrate Systemd with Golang

This section explains how to integrate our project with ``systemd`` so our operating system will handle stopping, restarting or starting.

1. (OPTIONAL) If you cannot access the server, please stop and review the steps above. If everything is working proceed forward.

2. While you are logged in as a ``techops`` user, please write the following into the console.

    ```
    $ sudo vi /etc/systemd/system/comicscantina-data.service
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
    ExecStart=/opt/lucha/go/bin/comicscantina-data
    Restart=always
    RestartSec=3
    SyslogIdentifier=comicscantina_data

    [Install]
    WantedBy=multi-user.target
    ```

4. Grant access.

   ```
   $ sudo chmod 755 /etc/systemd/system/comicscantina-data.service
   ```

5. (Optional) If you've updated the above, you will need to run the following before proceeding.

    ```
    $ systemctl daemon-reload
    ```

6. We can now start the Gunicorn service we created and enable it so that it starts at boot:

    ```
    $ sudo systemctl start comicscantina-data
    $ sudo systemctl enable comicscantina-data
    ```

7. Confirm our service is running.

    ```
    $ systemctl status comicscantina-data.service
    $ journalctl -f -u comicscantina-data.service
    ```

8. And verify the URL works in the browser.

    ```text
    http://SERVER_DOMAIN_NAME_OR_IP/en/
    ```
