package ct

import "fmt"

const (
	cE = "e"
	cW = "w"
	cS = "s"
	cI = "i"
)

const tmpl = " \033[38;2;%d;%d;%dm%v\033[0m"

var c = map[string][3]uint8{
	cE: {210, 82, 82},
	cW: {255, 198, 109},
	cS: {97, 150, 71},
	cI: {104, 151, 187},
}

func build(clr string, a ...interface{}) (res string) {
	for i := range a {
		res += fmt.Sprintf(tmpl, c[clr][0], c[clr][1], c[clr][2], a[i])
	}
	return res[1:]
}

func Error(a ...interface{}) {
	fmt.Print(build(cE, a...))
}

func Errorln(a ...interface{}) {
	Error(a...)
	println()
}

func Warning(a ...interface{}) {
	fmt.Print(build(cW, a...))
}

func Warningln(a ...interface{}) {
	Warning(a...)
	println()
}

func Success(a ...interface{}) {
	fmt.Print(build(cS, a...))
}

func Successln(a ...interface{}) {
	Success(a...)
	println()
}

func Info(a ...interface{}) {
	fmt.Print(build(cI, a...))
}

func Infoln(a ...interface{}) {
	Info(a...)
	println()
}