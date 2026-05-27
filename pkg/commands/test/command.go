package test

import (
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
)

type options struct {
	config                      string
	testFile                    string
	applyTimeout                metav1.Duration
	assertTimeout               metav1.Duration
	errorTimeout                metav1.Duration
	deleteTimeout               metav1.Duration
	cleanupTimeout              metav1.Duration
	execTimeout                 metav1.Duration
	testDirs                    []string
	skipDelete                  bool
	template                    bool
	defaultCompiler             string
	failFast                    bool
	parallel                    int
	repeatCount                 int
	reportFormat                string
	reportPath                  string
	reportName                  string
	namespace                   string
	fastNamespaceDeletion       bool
	deletionPropagationPolicy   string
	fullName                    bool
	excludeTestRegex            string
	includeTestRegex            string
	noColor                     bool
	kubeConfigOverrides         clientcmd.ConfigOverrides
	forceTerminationGracePeriod metav1.Duration
	delayBeforeCleanup          metav1.Duration
	selector                    []string
	noCluster                   bool
	pauseOnFailure              bool
	values                      []string
	set                         []string
	setString                   []string
	clusters                    []string
	remarshal                   bool
	shardIndex                  int
	shardCount                  int
	quiet                       bool
}

func Command() *cobra.Command { _ = "STUB: not implemented"; return nil }

// helper function for conditional logging based on quiet flag

// if no config file was provided, give a chance to the default config name

// try to load configuration file

// flags take precedence over configuration file

// if pause on failure is set, force non concurrency

// load tests

// TODO: we may want to find a sort key here ?

// load values

// merge --set into values

// run tests

// setup test context

// setup testing flags

// setup runner and execute

//nolint:errcheck

// process report

// config

// timeouts options

// discovery options

// execution options

// namespace options

// templating options

// cleanup options

// deletion options

// error options
// reporting options

// multi-cluster options

// pause options

// no cluster options

// label selectors

// external values

// sharding

// others
