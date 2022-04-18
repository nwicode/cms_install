# NWICODE v2 installation steps ( temporarily unavailable )

Step-by-step instructions for installing Nwicode v2 CMS on hosting running Ubuntu 18 and higher.

### Requirements
-  The domain name to be used. The domain must be configured so that the domain is revoked from the server used to install
-  The operating system used for CMS - Ubuntu 20+.
-  Server requirements: 4 GB of RAM, at least 40 GB of hard disk space (preferably SSD), 2+ CORE.
-  Since additional packages will need to be installed on the server, you must have root access for these operations.


### 1. Install "git"
    sudo apt-get update
	sudo apt-get -y install git

### 2. Clone installer from our repository
    git clone https://github.com/nwicode/cms_install.git

### 3. Install docker and  docker-compose
If the system asks for confirmation when installing any package, confirm the choice by pressing "Y" and Enter

    cd cms_install
    chmod +x preinstall.sh
	./preinstall.sh

### 4. Generate SSL Certificate
For CMS to work, you need to install a SSL certificate. You can use the Letsencrypt service.  To install, use the code below, just replace {DOMAIN_NAME} with the real domain that you will use to work with the CMS.

    sudo apt install letsencrypt
    letsencrypt certonly -d {DOMAIN_NAME}
During installation, use confirmation with the "standalone" server.
**Be sure, before these you must bind the domain used to the server.**

### 5. Install CMS
    chmod +x installer
    ./installer
    docker-compose pull
    docker-compose up -d
    docker-compose exec app_creator_api php artisan migrate
    docker-compose exec app_creator_api php artisan db:seed

./installer prompts the user for data and stores it in ENV

Before installation, you must obtain a license key, it will need to be specified during installation, otherwise the installation will be interrupted.

### 6. After install
For first login use this credentials:

	admin@admin.com
    admin
    
Don't forget to setup Email SMTP credentials and  change default admin login and password.
