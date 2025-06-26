package main

import (
	"fmt"

	v3loader "helm.sh/helm/v3/pkg/chart/loader"
)

func main() {
	chart, _ := v3loader.Load("testdata/chart-benchmark")

	fmt.Println(chart)
}

func inALoop(times int) {
	for i := 0; i < times; i++ {
		chart, _ := v3loader.Load("testdata/chart-benchmark")
		chart.Metadata.AppVersion = "1.0.0"
	}
}
