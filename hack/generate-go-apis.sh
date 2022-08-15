#!/bin/bash

set -e

GO111MODULE=off
echo "Generate deepcopy..."
deepcopy-gen --input-dirs ../argoproj/argo-cd/v2/pkg/apis/... --go-header-file hack/boilerplate.go.txt-O zz_generated.deepcopy

echo "Done generating"