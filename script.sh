sudo docker build -t asciiartweb .
sudo docker run --rm -p 8080:8080 --name server asciiartweb
sudo docker stop server
sudo docker rmi -f asciiartweb