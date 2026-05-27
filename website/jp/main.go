package main

import (
	"embed"
	"fmt"

	jpfunctions "github.com/jmespath-community/go-jmespath/pkg/functions"
	chainsawfunctions "github.com/kyverno/chainsaw/pkg/engine/functions"
	"github.com/kyverno/kyverno-json/pkg/jp/functions"
	kyvernofunctions "github.com/kyverno/kyverno-json/pkg/jp/kyverno"
)

//go:embed examples
var examples embed.FS

func main() {
	fmt.Println("# Functions")
	fmt.Println()
	fmt.Println(`!!! warning "Experimental functions"`)
	fmt.Println()
	fmt.Println("    Experimental functions are denoted by the `x_` prefix.")
	fmt.Println()
	fmt.Println("    These are functions that are subject to signature change in a future version.")
	fmt.Println()
	fmt.Println("## built-in functions")
	fmt.Println()
	printFunctions(jpfunctions.GetDefaultFunctions()...)
	fmt.Println()
	fmt.Println("## kyverno-json functions")
	fmt.Println()
	printFunctions(functions.GetFunctions()...)
	fmt.Println()
	fmt.Println("## kyverno functions")
	fmt.Println()
	printFunctions(kyvernofunctions.GetBareFunctions()...)
	fmt.Println()
	fmt.Println("## chainsaw functions")
	fmt.Println()
	printFunctions(chainsawfunctions.GetFunctions()...)
	fmt.Println()
}

func printFunctions(funcs ...jpfunctions.FunctionEntry) { _ = "STUB: not implemented"; return }

func functionString(f jpfunctions.FunctionEntry) string { _ = "STUB: not implemented"; return "" }
