#!/bin/bash

set -e

echo "Generate deepcopy..."
go run vendor/k8s.io/code-generator/cmd/deepcopy-gen/main.go \
  --go-header-file ./hack/boilerplate.go.txt \
  --input-dirs ./loft-sh/admin-services/pkg/server,./argoproj/argo-cd/v2/pkg/apis,./argoproj/argo-cd/v2/pkg/apis/application,./argoproj/argo-cd/v2/pkg/apis/application/v1alpha1 \
  -O zz_generated.deepcopy \
  -o ./

echo "Done generating"