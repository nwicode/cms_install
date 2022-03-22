# NWICODE v2 installation steps

Step-by-step instructions for installing Nwicode v2 CMS on hosting running Ubuntu 18 and higher.

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

