# /internal

Private application and library code. This is the code you don't want others importing in ther applications or libraries. Note that this layout pattern is enforced by the Go compiler itself. Note that you are not limited to the top level `internal` directory. You can have more than one `internal` directory at any level of your project tree.

You can aptionally add a bit of extra extructure to your internal packages to separate your shared and non-shared internal code. It's not required (specially for smaller projects), but it's nice to have visual clues showing the intended package use. Your actual application code can go in the `internal/app` directory (e.g. `/internal/app/myapp`) and the code shared by those apps in the `/internal/pkg` directory (e.g. `/internal/pkg/myprivlib`).

Examples:

* https://github.com/hashicorp/terraform/tree/main/internal
* https://github.com/influxdata/influxdb/tree/master/internal
* https://github.com/perkeep/perkeep/tree/master/internal
* https://github.com/jaegertracing/jaeger/tree/main/internal
* https://github.com/moby/moby/tree/master/internal
* https://github.com/satellity/satellity/tree/main/internal

## `/internal/pkg`

Examples:

* https://github.com/hashicorp/waypoint/tree/main/internal/pkg
