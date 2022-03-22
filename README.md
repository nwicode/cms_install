# NWICODE v2 installation steps

Step-by-step instructions for installing Nwicode v2 CMS on hosting running Ubuntu 18 and higher.

### 1. Install "git"
    sudo apt-get update
	sudo apt-get -y install git

### 2. Clone installer from our repository
    git clone https://github.com/nwicode/cms_install.git

### 3. Install docker and  docker-compose
    cd cms_install
    chmod +x preinstall.sh
