# Use Nginx Reverse Proxy to serve Go Services
Golang served by Nginx reverse proxy.


#  <font color='red'>Installation</font>

* Install on of the latest stable version of [Docker](https://docs.docker.com/install/linux/docker-ce/ubuntu/#install-docker-ce-1), and install [Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)
* Install [Docker Compose](https://docs.docker.com/compose/install/#install-compose)
* Browse the repository's root
* Build the images 
    - `docker-compose build`
* Start containers 
    - `docker-compose up -d`
* The application will now be ready to use go ahead and hit http://localhost/api/user in your browser and you’ll presented with this lovely screen.
![](https://i.imgur.com/vODQaOl.png)