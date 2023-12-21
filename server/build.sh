BUILDTIME=$(date +"%d.%m.%Y")
FLAGS="-X main.BuildTime=$BUILDTIME"
NAME=${1:-server}

go build -ldflags="$FLAGS" -o $NAME .
