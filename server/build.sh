TZ=Europe/Bratislava
BUILDTIME=$(date +"%d.%m.%Y_%H:%M")
FLAGS="-X main.BuildTime=${BUILDTIME}"
NAME=${1:-server}

go build -ldflags="$FLAGS" -o $NAME .
