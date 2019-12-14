# go-web-hello-world
I worked on this project at home. 
So, first of all, I need to download the Oracle VM VirtualBox as well as ubuntu-18.04.3-live-server-amd64.iso
After obtained a VM from VirtualBox, I could finish the following tasks:

Task 1. login and updates the system to the latest
    sudo apt-get update
Task 2. install gitlab-ce version in the host https://about.gitlab.com/install/#ubuntu?version=ce, by following the following instructions:
    sudo apt-get install -y curl openssh-server ca-certificates
    sudo apt-get install -y postfix
    curl -sS https://packages.gitlab.com/install/repositories/gitlab/gitlab-ce/script.deb.sh | sudo bash
    sudo EXTERNAL_URL="https://" apt-get install gitlab-ce
Task 3. create a demo group/project in gitlab, named demo/go-web-hello-world, use golang to build a helo world web app(listen to 8081 port) and check-in the code to mainline.
   step 1: install golang by 
       sudo snap install go
       sudo apt install golang-go
   step 2: create group/project and write a go code to build a hello world web app
       under my home directory, 
       mkdir demo; cd demo; mkdir go-web-hello-world; cd go-web-hello-world;
       write a go program called "hello-world.go" as follows:
       package main
       
       import ("fmt"
               "net/http")
       func hello_world( w http.ResponseWriter, r *http.Request) {
               fmt.Fprint(w, "Go Web Hello World!\n")
       }
       func main() {
               http.HandleFunc("/", hello_world)
               http.ListenAndServe(":8081", nil)
       }
   Step 3: test the code
       sudo go run hello-world.go 
       after input the password, the app is running, and from other terminal, run
       curl http://localhost:8081
       will get the following output:
       Go Web Hello World!
   Step 4: use git to create a new local repository under demo/go-web-hello-world
        git init
        And then 
        git add hello-world.go
        git commit -m"First Version of the hello-world"
 Task 4. build the app and expost the services to 8081 port, this was finished in Step 2 of the above task 3.
 
 Task 5. install docker by folowing the following steps:
      sudo apt-get remove docker docker-engine docker.io containerd runc
      sudo apt-get update
      sudo apt-get install \
                   apt-transport-https \
                   ca-certificates \
                   curl \
                   gnupg-agent \
                   software-properties-common
      curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
      sudo apt-key fingerprint 0EBFCD88
      sudo add-apt-repository \
           "deb [arch=amd64] https://download.docker.com/linux/ubuntu \
           $(lsb_release -cs) \
           stable"
      sudo apt-get update
      sudo apt-get install docker-ce docker-ce-cli containerd.io
      apt-cache madison docker-ce
      sudo apt-get install docker-ce=5:18.09.1~3-0~ubuntu-bionic
      
      Util now, docker is installed. Try to run an empty hello-world image, it works:
      sudo docker run hello-world
      
 Task 6. build a docker image ( docker build ) and run the web app in the container ( docker run), espose the service to 8082
       step 1. create Dockerfile as follows:
       FROM golang:lastest
       
       ADD . .
       
       RUN go build
       
       EXPOSE 8081
       
       CMD ["./go-web-hello-world"]
       
       Step 2. build the image:
       sudo docker build -t go-web-hello-world .
       
       Step 3. check the image:
       sudo docker imges
       I could find:
       REPOSITORY            TAG        IMAGE ID
       go-web-hello-world   latest      58df83d0ba3c  
       Step 4. run the app:
       
       sudo docker run go-web-hello-world:latest &
       curl http://localhost:8081
       output is:
       Go Web Hello World!
       
       Step 5 expose the service to 8082:
       sudo docker run -p 8082:8081 go-web-hello-world&
       curl http://localhost:8082
       output is:
       Go Web Hello World!
  Task 7. tag the docker image using gxdockerhub/go-web-hello-world:v0.1 and push it to docker hub
       step 1: create my docker called gxdockerhub
       Step 2: create the repository called go-web-hello-world and add the description
       Step 3: push the tag of the image into depository
          sudo docker push gxdockerhub/go-web-hello-world:v0.1
       Step 4: go to my docker hub, and find this tage
       
 Task 8.  create a README.md file and add the above technical procedures into this file.

