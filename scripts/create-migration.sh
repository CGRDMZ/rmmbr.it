set -e

read -p "enter the name of the migration..." migname

if [[ "$migname" == "" ]]; then
    echo "what is wrong with you?"
fi



migrate create -ext sql -dir db/migrations -seq $migname
