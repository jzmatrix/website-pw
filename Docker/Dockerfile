FROM debian:10
################################################################################
RUN apt update && \
    apt upgrade && \
    apt -y install openssh-server passwd postfix lsb-release apt-transport-https ca-certificates wget curl cron && \
    wget -O /etc/apt/trusted.gpg.d/php.gpg https://packages.sury.org/php/apt.gpg && \
    echo "deb https://packages.sury.org/php/ $(lsb_release -sc) main" | tee /etc/apt/sources.list.d/php.list && \
    apt update && \
    apt -y install php7.4 apache2

RUN apt-get -y install php7.4-bcmath php7.4-bz2 php7.4-cgi php7.4-common php7.4-curl php7.4-gd php7.4-geoip php7.4-gmp php7.4-imagick php7.4-intl php7.4-json php7.4-mbstring php7.4-mcrypt php7.4-memcache php7.4-memcached php7.4-mongodb php7.4-mysql php7.4-opcache php7.4-pspell php7.4-readline php7.4-snmp php7.4-tidy php7.4-xmlrpc php7.4-xml php7.4-xsl php7.4-zip
################################################################################
RUN /bin/rm -f /etc/localtime
RUN /bin/cp /usr/share/zoneinfo/America/New_York /etc/localtime
################################################################################
RUN ssh-keygen -t rsa -f /etc/ssh/ssh_host_rsa_key -N '' -y
################################################################################
ADD config/postfix/main.cf /etc/postfix/main.cf
RUN chmod 0666 /etc/postfix/main.cf
################################################################################
ADD config/crontab /var/spool/cron/root
RUN chmod 0600 /var/spool/cron/root
################################################################################
ADD config/php7/php.ini /etc/php/7.4/apache2/php.ini
RUN chmod 644 /etc/php/7.4/apache2/php.ini
################################################################################
ADD config/apache2/apache2.conf /etc/apache2/apache2.conf
RUN chmod 644 /etc/apache2/apache2.conf
################################################################################
ADD config/apache2/000-default.conf /etc/apache2/sites-available/000-default.conf
RUN chmod 644 /etc/apache2/sites-available/000-default.conf
################################################################################
ADD config/apache2/info.conf /etc/apache2/mods-available/info.conf
RUN chmod 644 /etc/apache2/mods-available/info.conf
################################################################################
ADD config/apache2/status.conf /etc/apache2/mods-available/status.conf
RUN chmod 644 /etc/apache2/mods-available/status.conf
################################################################################
RUN rm -rf /var/www/html/*
ADD website /var/www/html
################################################################################
RUN ln -f -s /etc/apache2/conf-available/php7.4-cgi.conf /etc/apache2/conf-enabled/ ; \
    ln -f -s /etc/apache2/mods-available/info.conf /etc/apache2/mods-enabled/ ; \
    ln -f -s /etc/apache2/mods-available/info.load /etc/apache2/mods-enabled/ ; \
    ln -f -s /etc/apache2/mods-available/remoteip.load /etc/apache2/mods-enabled/ ; \
    ln -f -s /etc/apache2/mods-available/rewrite.load /etc/apache2/mods-enabled/  ; \
    ln -f -s /etc/apache2/mods-available/status.conf /etc/apache2/mods-enabled/ ; \
    ln -f -s /etc/apache2/mods-available/status.load /etc/apache2/mods-enabled/ 
################################################################################
# ADD start.sh /start.sh
# RUN mkdir /var/run/sshd
# RUN chmod 755 /start.sh
# RUN ./start.sh
################################################################################
# Simple startup script to avoid some issues observed with container restart
ADD run-httpd.sh /run-httpd.sh
RUN chmod -v +x /run-httpd.sh
################################################################################
ADD startServices.sh /opt/startServices.sh
RUN chmod 755 /opt/startServices.sh
################################################################################
CMD [ "/opt/startServices.sh" ]
