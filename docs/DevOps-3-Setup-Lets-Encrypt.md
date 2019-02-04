# HOWTO: Setup Lets Encrypt for ComicsCantia Golang Web-App on DigitalOcean using CentOS 7 OS
## Description
This article assumes you've completed the first two articles with setting up Comics Cantina.

## Instruction
### Nginx
The following instructions are used to manually setup `letsencrypt` and automatically integrate with `nginx`.

1. Install our **Lets Encrypt** client.

  ```
  $ sudo yum install -y certbot-nginx
  ```


2. Generate our certificate.

  ```
  $ sudo certbot --nginx --agree-tos --server https://acme-v02.api.letsencrypt.org/directory --preferred-challenges dns -d "*.comicscantina.xyz" -d comicscantina.xyz
  ```


3. Follow the instructions and choose the most appropriate options.


4. **(Optional)** Please make a copy of the ``/etc/letsencrypt`` file. Instructions are below.


5. Restart ``nginx``.

  ```
  $ sudo systemctl restart nginx
  ```


6. Using your favourite browser, load up (https://comicscantina.xyz)[https://comicscantina.xyz] and it should work. If it does then congradulations!


### DigitalOcean
The problem with the above instructions is that **you are responsible for manually renewing within 90 days**. This manual renewing is tedious, can we automate? Turns out we can. The above instructions setup ``letsencrypt`` with our ``nginx`` so we have it working. In this section we will integrate our code with ``DigitalOcean`` and make auto-renewing taken care of by a script.

1. Install our ``DigitalOcean`` plugin.

  ```
  $ sudo yum install -y certbot-dns-digitalocean
  ```


2. Log into (DigitalOcean)[https://cloud.digitalocean.com/settings/api/tokens] and create an ``API key``.
https://certbot-dns-digitalocean.readthedocs.io/en/latest/#


3. Create our credentials file:

  ```
  $ sudo mkdir /etc/letsencrypt/digitalocean
  $ sudo cat > /etc/letsencrypt/digitalocean/credentials.ini
  ```


4. Populate the file with your key. Here is example:

  ```
  # DigitalOcean API credentials used by Certbot
  dns_digitalocean_token = 0000111122223333444455556666777788889999aaaabbbbccccddddeeeeffff
  ```


5. Finally run the code which will automatically generate our certificate.

  ```
  $ sudo certbot certonly --dns-digitalocean --dns-digitalocean-credentials /etc/letsencrypt/digitalocean/credentials.ini --dns-digitalocean-propagation-seconds 60 -d "*.comicscantina.xyz" -d comicscantina.xyz
  ```


6. (Optional) https://www.digitalocean.com/community/tutorials/how-to-secure-nginx-with-let-s-encrypt-on-centos-7

  ```bash
  sudo openssl dhparam -out /etc/ssl/certs/dhparam.pem 2048
  ```


7. Restart the server.

    ```
    sudo systemctl restart nginx
    ```


8. Would you like to know more?

* https://certbot.eff.org/lets-encrypt/centosrhel7-nginx
* https://certbot-dns-digitalocean.readthedocs.io/en/latest/


### HOW DO WE AUTO RENEW?
https://certbot.eff.org/lets-encrypt/centosrhel7-nginx.html

sudo crontab -e

# Add this to the crontab and save it:
0 0,12 * * * python -c 'import random; import time; time.sleep(random.random() * 3600)' && /usr/bin/certbot renew && systemctl restart nginx


### Nginx + SSL
If your SSL is not being populated at your address then follow these.




Congratulations! Your certificate and chain have been saved at:
  /etc/letsencrypt/live/comicscantina.xyz/fullchain.pem
  Your key file has been saved at:
  /etc/letsencrypt/live/comicscantina.xyz/privkey.pem
  Your cert will expire on 2019-05-05. To obtain a new or tweaked
  version of this certificate in the future, simply run certbot
  again. To non-interactively renew *all* of your certificates, run
  "certbot renew"
- If you like Certbot, please consider supporting our work by:

  Donating to ISRG / Let's Encrypt:   https://letsencrypt.org/donate
  Donating to EFF:                    https://eff.org/donate-le
