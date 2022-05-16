FROM debian:10
################################################################################
RUN apt update && \
    apt upgrade && \
    apt -y install openssh-server passwd lsb-release apt-transport-https ca-certificates wget curl && \
    apt update && \
    apt -y install libcgi-pm-perl apache2 

################################################################################
ADD config/apache2/apache2.conf /etc/apache2/apache2.conf
ADD config/apache2/000-default.conf /etc/apache2/sites-available/000-default.conf
ADD config/apache2/info.conf /etc/apache2/mods-available/info.conf
ADD config/apache2/status.conf /etc/apache2/mods-available/status.conf
# Simple startup script to avoid some issues observed with container restart
ADD config/run-httpd.sh /run-httpd.sh  
ADD config/startServices.sh /opt/startServices.sh
################################################################################
RUN rm -rf /var/www/html/*
ADD www /var/www/html
################################################################################
RUN /bin/rm -f /etc/localtime && \
    /bin/cp /usr/share/zoneinfo/America/New_York /etc/localtime && \
    ssh-keygen -t rsa -f /etc/ssh/ssh_host_rsa_key -N '' -y && \
    chmod 644 /etc/apache2/apache2.conf && \
    chmod 644 /etc/apache2/sites-available/000-default.conf && \
    chmod 644 /etc/apache2/mods-available/info.conf && \
    chmod 644 /etc/apache2/mods-available/status.conf && \
    ln -f -s /etc/apache2/mods-available/cgi.load /etc/apache2/mods-enabled/ && \
    ln -f -s /etc/apache2/mods-available/info.conf /etc/apache2/mods-enabled/ && \
    ln -f -s /etc/apache2/mods-available/info.load /etc/apache2/mods-enabled/  && \
    ln -f -s /etc/apache2/mods-available/remoteip.load /etc/apache2/mods-enabled/ && \
    ln -f -s /etc/apache2/mods-available/rewrite.load /etc/apache2/mods-enabled/ && \
    ln -f -s /etc/apache2/mods-available/status.conf /etc/apache2/mods-enabled/ && \
    ln -f -s /etc/apache2/mods-available/status.load /etc/apache2/mods-enabled/  && \
    chmod -v +x /run-httpd.sh && \
    chmod 755 /opt/startServices.sh && \
    chmod -R 755 /var/www/html/*
################################################################################
CMD [ "/opt/startServices.sh" ]
