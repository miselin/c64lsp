/** Package reference includes data from the C64 Programmers Guide for the LSP to present. */
package reference

type BasicFunction struct {
	Name            string
	Type            string
	Format          string
	Match           string
	Action          string
	ExampleMarkdown string
}

var BasicFunctions = []BasicFunction{
	{
		Name:   "abs",
		Type:   "Function-Numeric",
		Format: "ABS(<expression>)",
		Match:  "abs(",
		Action: "Returns the absolute value of the number, which is its value without any signs. The absolute value of a negative number is that number multiplied by -1.",
		ExampleMarkdown: "```" + `
10 X = ABS ( Y )
10 PRINT ABS ( X * J )
10 IF X = ABS (X) THEN PRINT "POSITIVE"
` + "```",
	},
	{
		Name:            "rnd",
		Type:            "Floating-Point Function",
		Format:          "RND(<numeric>)",
		Match:           "rnd(",
		Action:          "RND creates a floating-point random from 0.0 to 1.0. The computer generates a sequence of random numbers byperforming calculations on a starting number, which in computer jargon is called a seed. TheRND function isseeded onsystem power-up. The <numeric> argument is a dummy, except for its sign (positive, zero, or negative).\nIf the <numeric> argument is positive, the same \"pseudorandom\" sequence of numbers is returned, starting froma given seed value. Different number sequences will result from different seeds, but any sequence is repeatable by starting from the same seed number. Having a known sequence of \"random\" numbers is useful in testing programs.\nIf you choose a <numeric> argument of zero, then RND generates a number directly from a free-running hardware clock (the system \"jiffy clock\"). Negative arguments cause the RNDfunction to be re-seeded with each function call.",
		ExampleMarkdown: "todo",
	},
}
