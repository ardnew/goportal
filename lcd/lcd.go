package lcd

import (
	"image/color"

	"github.com/ardnew/goportal/lcd/font"
	"github.com/ardnew/goportal/lcd/style"
	"tinygo.org/x/drivers/ili9341"
)

// Orientation defines screen rotation with respect to the USB port.
type Orientation ili9341.Rotation

// Constants defining all possible screen rotations.
const (
	OrientTop    = Orientation(ili9341.Rotation0)
	OrientLeft   = Orientation(ili9341.Rotation90)
	OrientBottom = Orientation(ili9341.Rotation180)
	OrientRight  = Orientation(ili9341.Rotation270)
)

// DrawChar writes the given ASCII char in given font at given coordinates on
// the receiver lcd.
func (lcd *Lcd) DrawChar(f *font.Font, x, y int16, char rune) {
	if bmap := f.Bitmap(char); nil != bmap {
		lcd.DrawRGBBitmap(x, y, bmap, f.Width(), f.Height())
	}
}

// DrawString writes the given ASCII string in given font at given coordinates
// on the receiver lcd.
func (lcd *Lcd) DrawString(f *font.Font, x, y int16, str string) {
	for _, c := range str {
		lcd.DrawChar(f, x, y, c)
		x += f.Width()
	}
}

// Display enables or disables the backlight in the receiver lcd.
func (lcd *Lcd) Display(enable bool) {

	lcd.backlight.Set(enable)
}

// Init initializes the receiver lcd.
func (lcd *Lcd) Init(background color.RGBA) error {

	lcd.Display(true)
	lcd.Device.FillScreen(background)

	lcd.Device.FillRectangle(100, 100, 10, 10, style.TextColorForeDefault)
	lcd.Device.FillRectangle(100, 115, 10, 10, style.TextColorForeDefault)
	lcd.Device.FillRectangle(100, 130, 10, 10, style.TextColorForeDefault)
	lcd.Device.FillRectangle(100, 145, 10, 10, style.Red)

	return nil
}
