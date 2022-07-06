set -e

wget -O /tmp/migrate.linux-amd64.tar.gz https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz &&
tar -xvf /tmp/migrate.linux-amd64.tar.gz -C /tmp

mkdir -p $HOME/bin/
cp /tmp/migrate.linux-amd64 $HOME/bin/migrate
