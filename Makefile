.PHONY: dist

APP_NAME=home-bot
APP_SRC=github.com/szeist/home-bot
DIST_DIR=dist
RPI_BUILD_CONTAINER=home-bot-build

start:
	go run cmd/home-bot/main.go

build-rpi:
	docker run \
		--name $(RPI_BUILD_CONTAINER) \
		-v $(PWD):/go/src/$(APP_SRC) \
		-w /go/src/$(APP_SRC) \
		go-build-rpi:rpi2 \
		$(APP_SRC)/cmd/home-bot /dist/$(APP_NAME)
	docker cp $(RPI_BUILD_CONTAINER):dist/$(APP_NAME) $(DIST_DIR)/$(APP_NAME)
	docker rm $(RPI_BUILD_CONTAINER)

build-keybase-rpi:
	docker run \
		--name rpi2-keybase-build \
		go-build-rpi:rpi2 \
		github.com/keybase/client/go/keybase /dist/keybase
	docker cp rpi2-keybase-build:/dist/keybase $(DIST_DIR)/keybase
	docker rm rpi2-keybase-build

clean:
	rm -rf $(DIST_DIR)
	mkdir -p $(DIST_DIR)
	docker rm -f $(RPI_BUILD_CONTAINER) || true 2> /dev/null