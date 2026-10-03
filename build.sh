   #!/usr/bin/env bash
   set -xe
   export GOCACHE=/tmp/gocache
   export HOME=${HOME:-/root}
   go build -o bin/application main.go
