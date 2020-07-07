package main

import (
	"github.com/filecoin-project/oni/lotus-soup/paych"
	"github.com/filecoin-project/oni/lotus-soup/testkit"

	"github.com/testground/sdk-go/run"
)

var cases = map[string]interface{}{
	"deals-e2e":         testkit.WrapTestEnvironment(dealsE2E),
	"deals-stress-test": testkit.WrapTestEnvironment(dealStressTest),
	"drand-halting":     testkit.WrapTestEnvironment(dealsE2E),
	"paych-stress":      testkit.WrapTestEnvironment(paych.Stress),
}

func main() {
	sanityCheck()

	run.InvokeMap(cases)
}

