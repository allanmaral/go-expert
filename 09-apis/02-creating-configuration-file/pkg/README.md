# /pkg

Library code that's ok to use by externa applications (e.g. `/pkg/mypubliclib`). Others projects will import these libraries expecting them to work, so think twice before you put something here. Note that the `internal` directory is a better way to ensure your private packages anot not importalble becouse it's enforces by Go.

Examples:

* https://github.com/containerd/containerd/tree/main/pkg
* https://github.com/slimtoolkit/slim/tree/master/pkg
* https://github.com/telepresenceio/telepresence/tree/release/v2/pkg
* https://github.com/jaegertracing/jaeger/tree/master/pkg
* https://github.com/istio/istio/tree/master/pkg
* https://github.com/GoogleContainerTools/kaniko/tree/master/pkg
* https://github.com/google/gvisor/tree/master/pkg
* https://github.com/google/syzkaller/tree/master/pkg
* https://github.com/perkeep/perkeep/tree/master/pkg
* https://github.com/minio/minio/tree/master/pkg
