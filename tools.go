//go:build tools
// +build tools

package tools

import (
	_ "github.com/prometheus/alertmanager/cmd/amtool"
	_ "github.com/prometheus/prometheus/cmd/promtool"
)
