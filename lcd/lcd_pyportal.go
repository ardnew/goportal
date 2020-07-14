// +build pyportal

package lcd

import (
	"machine"

	"tinygo.org/x/drivers/ili9341"
)

// Lcd encapsulates the ILI9341 on an Adafruit PyPortal with methods for drawing
// a user interface.
type Lcd struct {
	*ili9341.Device
	backlight machine.Pin
}

// New creates a new Lcd using the default pins connecting the 8-bit parallel
// RGB interface of the PyPortal to the ILI9341 TFT LCD.
func New(orientation Orientation) *Lcd {

	lcd := ili9341.NewParallel(
		machine.LCD_DATA0,
		machine.TFT_WR,
		machine.TFT_DC,
		machine.TFT_CS,
		machine.TFT_RESET,
		machine.TFT_RD,
	)

	backlight := machine.TFT_BACKLIGHT

	backlight.Configure(machine.PinConfig{machine.PinOutput})

	lcd.Configure(ili9341.Config{Rotation: ili9341.Rotation(orientation)})

	return &Lcd{lcd, backlight}
}
