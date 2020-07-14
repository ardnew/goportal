package main

import (
	"github.com/ardnew/goportal/lcd"
	"github.com/ardnew/goportal/lcd/font"
	"github.com/ardnew/goportal/lcd/style"
)

func main() {

	lcd := lcd.New(lcd.OrientLeft)

	lcd.Init(style.BackgroundDefault)

	mesg := []rune("Hello, from tinygo!  -  ")

	if text, _ := font.New(font.Medium, style.TextColorForeDefault, style.TextColorBackDefault); nil != text {
		lcd.DrawString(text, 10, 10, string(mesg))
		for {
			//			lcd.DrawString(text, 10, 10, string(mesg))
			//			c := mesg[0]
			//			mesg = mesg[1:]
			//			mesg = append(mesg, c)
			//			time.Sleep(50 * time.Millisecond)
		}
	}
}
