package option

var Opt Options
const STDIN = "stdin"
const STDOUT = "stdout"

func init() {
	Opt = NewOption()
}

func NewOption() Options {
	return Options{
		InputFile: STDIN,
		OutputFile: STDOUT,
	}
}

type Options struct {
	InputFile	string
	OutputFile string
}
