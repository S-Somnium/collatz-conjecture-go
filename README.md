# collatz-conjecture-go
 A collactz conjecture service running in kubernets

# Instructions to setup
Clone the project
 ```sh
git clone https://github.com/S-Somnium/collatz-conjecture-go
```
Now go inside the project folter and run create docker image
 ```sh
docker build -t collatz_service .
```
Create minikube pod
 ```sh
kubectl create -f ./collatz_service.yaml
```
Attach port 5000 to 8081 of container
 ```sh
kubectl port-forward collatz_service 5000:8081
```
Done !

# How to test
Inside the project folder run this command
 ```sh
go run ./script
```
Now follow the instructions.
You can also just communicate via telnet on port 5000.