package assert

const ansiRed = "\033[31m"
const ansiYellow = "\033[33m"
const ansiCyan = "\033[36m"
const ansiGray = "\033[38;5;251m"
const ansiBold = "\033[1m"
const ansiUnderline = "\033[4m"
const ansiReset = "\033[0m"

var SectionTitleColor = ansiUnderline + ansiBold
var ContextColor = ansiGray
var LineColor = ansiRed
var ArgumentColor = ansiCyan
var StackColor = ansiGray

func applyAnsi(code string, text string) string {
	return code + text + ansiReset
}
