FROM php:7.2.3-fpm-stretch

# LABEL version=1.0

WORKDIR /var/www/html

ENV server_env=product

ADD docker-file/miniprogram-php/composer/composer /usr/local/bin/ 

RUN echo 'deb http://mirrors.163.com/debian/ jessie main non-free contrib \
    deb http://mirrors.163.com/debian/ jessie-updates main non-free contrib \
    deb http://mirrors.163.com/debian-security/ jessie/updates main non-free contrib' > /etc/apt/sources.list

RUN echo "nameserver 223.5.5.5" >> /etc/resolv.conf \
    && echo "nameserver 223.6.6.6" >> /etc/resolve.conf \
    && apt-get update \
    && apt-get install -y \
    git \
    vim \
    wget \
    # && docker-php-ext-configure gd --with-freetype-dir=/usr/include/ --with-jpeg-dir=/usr/include/ \
    # && docker-php-ext-install -j$(nproc) gd \
    && docker-php-ext-install mysqli pdo_mysql \
    # && pecl install swoole \
    && pecl install redis \
    # && docker-php-ext-enable swoole redis
    && docker-php-ext-enable redis
RUN composer config -g repo.packagist composer https://packagist.phpcomposer.com
# RUN git clone 