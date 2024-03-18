

rm -f go.sum
go get -u "github.com/Patrick-ring-motive/utils"
go build -ldflags "-g" -gcflags="-B -v -std"  -o String String.go