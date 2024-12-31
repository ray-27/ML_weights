#!bin/bash

#this script should only be used when the instance is started,

#updating the packeages
sudo apt upgrade -y
sudo apt-get update

#installing aws codedeploy
sudo apt install ruby-full wget -y
cd /home/ubuntu  # Replace '/home/ubuntu' with your user directory
wget https://aws-codedeploy-ap-south-1.s3.ap-south-1.amazonaws.com/latest/install
sudo chmod +x ./install
sudo ./install auto

# install golang
wget https://go.dev/dl/go1.21.1.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.21.1.linux-amd64.tar.gz
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.profile
source ~/.profile
