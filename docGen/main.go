package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/khulnasoft-lab/cvedb-generator/menu"
)

var (
	Years []string

	misConfigurationMenu = menu.New("misconfig", "content/misconfig")
	complianceMenu       = menu.New("compliance", "content/compliance")
	runTimeSecurityMenu  = menu.New("runsec", "content/tracker")
)

type Clock interface {
	Now(format ...string) string
}

type realClock struct{}

func (realClock) Now(format ...string) string {
	formatString := time.RFC3339
	if len(format) > 0 {
		formatString = format[0]
	}

	return time.Now().Format(formatString)
}

func main() {

	firstYear := 1999

	for y := firstYear; y <= time.Now().Year(); y++ {
		Years = append(Years, strconv.Itoa(y))
	}

	generateChainBenchPages("../cvedb-repo/chain-bench-repo/internal/checks", "../cvedb-repo/content/compliance")
	generateKubeBenchPages("../cvedb-repo/kube-bench-repo/cfg", "../cvedb-repo/content/compliance")
	generateDefsecComplianceSpecPages("../cvedb-repo/defsec-repo/rules/specs/compliance", "../cvedb-repo/content/compliance")
	generateKubeHunterPages("../cvedb-repo/kube-hunter-repo/docs/_kb", "../cvedb-repo/content/misconfig/kubernetes")
	generateCloudSploitPages("../cvedb-repo/cloudsploit-repo/plugins", "../cvedb-repo/content/misconfig", "../cvedb-repo/remediations-repo/en")
	generateTrackerPages("../cvedb-repo/tracker-repo/signatures", "../cvedb-repo/content/tracker", realClock{})
	generateDefsecPages("../cvedb-repo/defsec-repo/cvedb_docs", "../cvedb-repo/content/misconfig")

	generateVulnPages()

	for _, year := range Years {
		generateReservedPages(year, realClock{}, "vuln-list", "content/nvd")
	}

	createTopLevelMenus()
}

func createTopLevelMenus() {
	if err := menu.NewTopLevelMenu("Misconfiguration", "toplevel_page", "content/misconfig/_index.md").
		WithHeading("Misconfiguration Categories").
		WithIcon("khulnasoft").
		WithCategory("misconfig").Generate(); err != nil {
		fail(err)
	}
	if err := menu.NewTopLevelMenu("Compliance", "toplevel_page", "content/compliance/_index.md").
		WithHeading("Compliance").
		WithIcon("khulnasoft").
		WithCategory("compliance").Generate(); err != nil {
		fail(err)
	}
	if err := menu.NewTopLevelMenu("Tracker", "toplevel_page", "content/tracker/_index.md").
		WithHeading("Runtime Security").
		WithIcon("tracker").
		WithCategory("runsec").
		Generate(); err != nil {
		fail(err)
	}

	if err := misConfigurationMenu.Generate(); err != nil {
		fail(err)
	}
	if err := runTimeSecurityMenu.Generate(); err != nil {
		fail(err)
	}
	if err := complianceMenu.Generate(); err != nil {
		fail(err)
	}
}

func fail(err error) {
	fmt.Println(err)
	os.Exit(1)
}
