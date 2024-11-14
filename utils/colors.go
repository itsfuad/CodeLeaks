package utils

import "fmt"

type COLOR string

const (
	RED    COLOR = "\033[31m"
	GREEN  COLOR = "\033[32m"
	YELLOW COLOR = "\033[33m"
	BLUE   COLOR = "\033[34m"
	PURPLE COLOR = "\033[35m"
	GREY   COLOR = "\033[37m"
	CYAN   COLOR = "\033[36m"
	WHITE  COLOR = "\033[97m"
	RESET  COLOR = "\033[0m"
)

func (c COLOR) String() string {
	return string(c)
}

//Example: fmt.Println(RED.Sprintf("Hello %s", "World"))
func (c COLOR) Sprintf(format string, a ...interface{}) string {
	return c.String() + fmt.Sprintf(format, a...) + RESET.String()
}

func (c COLOR) Printf(format string, a ...interface{}) {
	fmt.Printf(c.String()+format+RESET.String(), a...)
}

func (c COLOR) Println(a ...interface{}) {
	fmt.Print(c.String())
	fmt.Println(a...)
	fmt.Print(RESET.String())
}

func (c COLOR) Print(a ...interface{}) {
	fmt.Print(c.String())
	fmt.Print(a...)
	fmt.Print(RESET.String())
}

func ColorPrint(color COLOR, a ...interface{}) {
	fmt.Print(color.String())
	fmt.Print(a...)
	fmt.Print(RESET.String())
}
