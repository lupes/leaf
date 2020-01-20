all:

up:
	server="leaf"

    if [[ ! -f "./leaf/${server}}" ]]; then
      # shellcheck disable=SC2164
      cd leaf

      export GOROOT=/Users/klook/work/sdk/go1.13.4
      export GO111MODULE=on

      GOOS=linux /Users/klook/work/sdk/go1.13.4/bin/go build -v -o leaf
    fi

    docker-compose --compatibility up

stop: