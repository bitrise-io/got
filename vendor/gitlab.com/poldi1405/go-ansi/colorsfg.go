package ansi

import "fmt"

// Black wraps the content in ANSI codes to make its foreground color black
func Black(content ...interface{}) string {
	return SetBlack() + fmt.Sprint(content...) + UnsetBlack()
}

// SetBlack sets the foreground color to black
func SetBlack() string {
	return escape + blackfg + set
}

// UnsetBlack resets the foreground color from black to default.
func UnsetBlack() string {
	return escape + resetfg + set
}

// Red wraps the content in ANSI codes to make its foreground color red
func Red(content ...interface{}) string {
	return SetRed() + fmt.Sprint(content...) + UnsetRed()
}

// SetRed sets the foreground color to red
func SetRed() string {
	return escape + redfg + set
}

// UnsetRed resets the foreground color from red to default.
func UnsetRed() string {
	return escape + resetfg + set
}

// Green wraps the content in ANSI codes to make its foreground color green
func Green(content ...interface{}) string {
	return SetGreen() + fmt.Sprint(content...) + UnsetGreen()
}

// SetGreen sets the foreground color to green
func SetGreen() string {
	return escape + greenfg + set
}

// UnsetGreen resets the foreground color from green to default.
func UnsetGreen() string {
	return escape + resetfg + set
}

// Yellow wraps the content in ANSI codes to make its foreground color yellow
func Yellow(content ...interface{}) string {
	return SetYellow() + fmt.Sprint(content...) + UnsetYellow()
}

// SetYellow sets the foreground color to yellow
func SetYellow() string {
	return escape + yellowfg + set
}

// UnsetYellow resets the foreground color from yellow to default.
func UnsetYellow() string {
	return escape + resetfg + set
}

// Blue wraps the content in ANSI codes to make its foreground color blue
func Blue(content ...interface{}) string {
	return SetBlue() + fmt.Sprint(content...) + UnsetBlue()
}

// SetBlue sets the foreground color to blue
func SetBlue() string {
	return escape + bluefg + set
}

// UnsetBlue resets the foreground color from blue to default.
func UnsetBlue() string {
	return escape + resetfg + set
}

// Magenta wraps the content in ANSI codes to make its foreground color magenta
func Magenta(content ...interface{}) string {
	return SetMagenta() + fmt.Sprint(content...) + UnsetMagenta()
}

// SetMagenta sets the foreground color to magenta
func SetMagenta() string {
	return escape + magentafg + set
}

// UnsetMagenta resets the foreground color from magenta to default.
func UnsetMagenta() string {
	return escape + resetfg + set
}

// Cyan wraps the content in ANSI codes to make its foreground color cyan
func Cyan(content ...interface{}) string {
	return SetCyan() + fmt.Sprint(content...) + UnsetCyan()
}

// SetCyan sets the foreground color to cyan
func SetCyan() string {
	return escape + cyanfg + set
}

// UnsetCyan resets the foreground color from cyan to default.
func UnsetCyan() string {
	return escape + resetfg + set
}

// White wraps the content in ANSI codes to make its foreground color white
func White(content ...interface{}) string {
	return SetWhite() + fmt.Sprint(content...) + UnsetWhite()
}

// SetWhite sets the foreground color to white
func SetWhite() string {
	return escape + whitefg + set
}

// UnsetWhite resets the foreground color from white to default.
func UnsetWhite() string {
	return escape + resetfg + set
}

// Color256 sets a Term256 color that is applied to the provided text
func Color256(color int, content ...interface{}) string {
	return SetColor256(color) + fmt.Sprint(content...) + UnsetColor256()
}

// SetColor256 writes the following text on the specified term256 color
func SetColor256(color int) string {
	return escape + fmt.Sprintf(fg256, color) + set
}

// UnsetColor256 resets the color
func UnsetColor256() string {
	return escape + resetfg + set
}

// ColorTrue sets a RGB-Color that is set as the background color, writes the
func ColorTrue(r, g, b int, content ...interface{}) string {
	return SetColorTrue(r, g, b) + fmt.Sprint(content...) + UnsetColorTrue()
}

// SetColorTrue sets a RGB-Color for the font
func SetColorTrue(r, g, b int) string {
	return escape + fmt.Sprintf(fgtrue, r, g, b) + set
}

// UnsetColorTrue removes the RGB-color
func UnsetColorTrue() string {
	return escape + resetfg + set
}
