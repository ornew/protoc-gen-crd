```
out=./k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1/generated.proto
mkdir -p $(dirname $out)
curl -fsSLo $out https://raw.githubusercontent.com/kubernetes/apiextensions-apiserver/v0.27.3/pkg/apis/apiextensions/v1/generated.proto

out=./k8s.io/apimachinery/pkg/apis/meta/v1/generated.proto
mkdir -p $(dirname $out)
curl -fsSLo $out https://raw.githubusercontent.com/kubernetes/apimachinery/v0.27.3/pkg/apis/meta/v1/generated.proto

out=./k8s.io/apimachinery/pkg/runtime/generated.proto
mkdir -p $(dirname $out)
curl -fsSLo $out https://raw.githubusercontent.com/kubernetes/apimachinery/v0.27.3/pkg/runtime/generated.proto

out=./k8s.io/apimachinery/pkg/runtime/schema/generated.proto
mkdir -p $(dirname $out)
curl -fsSLo $out https://raw.githubusercontent.com/kubernetes/apimachinery/v0.27.3/pkg/runtime/schema/generated.proto
```
