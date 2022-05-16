if ! command -v migrate &>/dev/null; then
    go get -u -d github.com/golang-migrate/migrate/cmd/migrate
    cd $GOPATH/src/github.com/golang-migrate/migrate/cmd/migrate
    git checkout v4.15.2
    go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.15.2
fi
