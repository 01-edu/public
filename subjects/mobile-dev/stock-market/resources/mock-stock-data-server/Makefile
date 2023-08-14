IMAGE_NAME="stock-server"
SERVER_PORT=5001
PUBLIC_PORT=5001

sample-stocks:
	tar xf sample-stocks.zip 

build-image:
	docker rmi ${IMAGE_NAME} || echo "Clean"
	docker build -t ${IMAGE_NAME} .

run: sample-stocks build-image 
	docker run -d -p  ${PUBLIC_PORT}:${SERVER_PORT} --name ${IMAGE_NAME} --rm ${IMAGE_NAME} 

stop:
	docker stop ${IMAGE_NAME}

.PHONY: run stop unzip-sample-stock build-image
