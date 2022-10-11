package material

import (
	"sort"

	"github.com/taigrr/simplecolorpalettes/simplecolor"
)

// The names and hex values for the material palettes are sourced from
// material.io: https://material.io/design/color/the-color-system.html#tools-for-picking-colors

var (
	red = simplecolor.NamedPalette{
		"50": 0xffebee, "100": 0xffcdd2, "200": 0xef9a9a, "300": 0xe57373,
		"400": 0xef5350, "500": 0xf44336, "600": 0xe53935, "700": 0xd32f2f,
		"800": 0xc62828, "900": 0xb71c1c, "A100": 0xff8a80, "A200": 0xff5252,
		"A400": 0xff1744, "A700": 0xd50000,
	}
	pink = simplecolor.NamedPalette{
		"50": 0xfce4ec, "100": 0xf8bbd0, "200": 0xf48fb1, "300": 0xf06292,
		"400": 0xec407a, "500": 0xe91e63, "600": 0xd81b60, "700": 0xc2185b,
		"800": 0xad1457, "900": 0x880e4f, "A100": 0xff80ab, "A200": 0xff4081,
		"A400": 0xf50057, "A700": 0xc51162,
	}
	purple = simplecolor.NamedPalette{
		"50": 0xf3e5f5, "100": 0xe1bee7, "200": 0xce93d8, "300": 0xba68c8,
		"400": 0xab47bc, "500": 0x9c27b0, "600": 0x8e24aa, "700": 0x7b1fa2,
		"800": 0x6a1b9a, "900": 0x4a148c, "A100": 0xea80fc, "A200": 0xe040fb,
		"A400": 0xd500f9, "A700": 0xaa00ff,
	}
	// Removed as not used, leaving in for future use if needed.
	// deepPurple = simplecolor.NamedPalette{
	// 	"50": 0xede7f6, "100": 0xd1c4e9, "200": 0xb39ddb, "300": 0x9575cd,
	// 	"400": 0x7e57c2, "500": 0x673ab7, "600": 0x5e35b1, "700": 0x512da8,
	// 	"800": 0x4527a0, "900": 0x311b92, "A100": 0xb388ff, "A200": 0x7c4dff,
	// 	"A400": 0x651fff, "A700": 0x6200ea,
	// }
	// indigo = simplecolor.NamedPalette{
	// 	"50": 0xe8eaf6, "100": 0xc5cae9, "200": 0x9fa8da, "300": 0x7986cb,
	// 	"400": 0x5c6bc0, "500": 0x3f51b5, "600": 0x3949ab, "700": 0x303f9f,
	// 	"800": 0x283593, "900": 0x1a237e, "A100": 0x8c9eff, "A200": 0x536dfe,
	// 	"A400": 0x3d5afe, "A700": 0x304ffe,
	// }
	blue = simplecolor.NamedPalette{
		"50": 0xe3f2fd, "100": 0xbbdefb, "200": 0x90caf9, "300": 0x64b5f6,
		"400": 0x42a5f5, "500": 0x2196f3, "600": 0x1e88e5, "700": 0x1976d2,
		"800": 0x1565c0, "900": 0x0d47a1, "A100": 0x82b1ff, "A200": 0x448aff,
		"A400": 0x2979ff, "A700": 0x2962ff,
	}
	lightBlue = simplecolor.NamedPalette{
		"50": 0xe1f5fe, "100": 0xb3e5fc, "200": 0x81d4fa, "300": 0x4fc3f7,
		"400": 0x29b6f6, "500": 0x03a9f4, "600": 0x039be5, "700": 0x0288d1,
		"800": 0x0277bd, "900": 0x01579b, "A100": 0x80d8ff, "A200": 0x40c4ff,
		"A400": 0x00b0ff, "A700": 0x0091ea,
	}
	cyan = simplecolor.NamedPalette{
		"50": 0xe0f7fa, "100": 0xb2ebf2, "200": 0x80deea, "300": 0x4dd0e1,
		"400": 0x26c6da, "500": 0x00bcd4, "600": 0x00acc1, "700": 0x0097a7,
		"800": 0x00838f, "900": 0x006064, "A100": 0x84ffff, "A200": 0x18ffff,
		"A400": 0x00e5ff, "A700": 0x00b8d4,
	}
	teal = simplecolor.NamedPalette{
		"50": 0xe0f2f1, "100": 0xb2dfdb, "200": 0x80cbc4, "300": 0x4db6ac,
		"400": 0x26a69a, "500": 0x009688, "600": 0x00897b, "700": 0x00796b,
		"800": 0x00695c, "900": 0x004d40, "A100": 0xa7ffeb, "A200": 0x64ffda,
		"A400": 0x1de9b6, "A700": 0x00bfa5,
	}
	green = simplecolor.NamedPalette{
		"50": 0xe8f5e9, "100": 0xc8e6c9, "200": 0xa5d6a7, "300": 0x81c784,
		"400": 0x66bb6a, "500": 0x4caf50, "600": 0x43a047, "700": 0x388e3c,
		"800": 0x2e7d32, "900": 0x1b5e20, "A100": 0xb9f6ca, "A200": 0x69f0ae,
		"A400": 0x00e676, "A700": 0x00c853,
	}
	lightGreen = simplecolor.NamedPalette{
		"50": 0xf1f8e9, "100": 0xdcedc8, "200": 0xc5e1a5, "300": 0xaed581,
		"400": 0x9ccc65, "500": 0x8bc34a, "600": 0x7cb342, "700": 0x689f38,
		"800": 0x558b2f, "900": 0x33691e, "A100": 0xccff90, "A200": 0xb2ff59,
		"A400": 0x76ff03, "A700": 0x64dd17,
	}
	lime = simplecolor.NamedPalette{
		"50": 0xf9fbe7, "100": 0xf0f4c3, "200": 0xe6ee9c, "300": 0xdce775,
		"400": 0xd4e157, "500": 0xcddc39, "600": 0xc0ca33, "700": 0xafb42b,
		"800": 0x9e9d24, "900": 0x827717, "A100": 0xf4ff81, "A200": 0xeeff41,
		"A400": 0xc6ff00, "A700": 0xaeea00,
	}
	yellow = simplecolor.NamedPalette{
		"50": 0xfffde7, "100": 0xfff9c4, "200": 0xfff59d, "300": 0xfff176,
		"400": 0xffee58, "500": 0xffeb3b, "600": 0xfdd835, "700": 0xfbc02d,
		"800": 0xf9a825, "900": 0xf57f17, "A100": 0xffff8d, "A200": 0xffff00,
		"A400": 0xffea00, "A700": 0xffd600,
	}
	amber = simplecolor.NamedPalette{
		"50": 0xfff8e1, "100": 0xffecb3, "200": 0xffe082, "300": 0xffd54f,
		"400": 0xffca28, "500": 0xffc107, "600": 0xffb300, "700": 0xffa000,
		"800": 0xff8f00, "900": 0xff6f00, "A100": 0xffe57f, "A200": 0xffd740,
		"A400": 0xffc400, "A700": 0xffab00,
	}
	orange = simplecolor.NamedPalette{
		"50": 0xfff3e0, "100": 0xffe0b2, "200": 0xffcc80, "300": 0xffb74d,
		"400": 0xffa726, "500": 0xff9800, "600": 0xfb8c00, "700": 0xf57c00,
		"800": 0xef6c00, "900": 0xe65100, "A100": 0xffd180, "A200": 0xffab40,
		"A400": 0xff9100, "A700": 0xff6d00,
	}
	deepOrange = simplecolor.NamedPalette{
		"50": 0xfbe9e7, "100": 0xffccbc, "200": 0xffab91, "300": 0xff8a65,
		"400": 0xff7043, "500": 0xff5722, "600": 0xf4511e, "700": 0xe64a19,
		"800": 0xd84315, "900": 0xbf360c, "A100": 0xff9e80, "A200": 0xff6e40,
		"A400": 0xff3d00, "A700": 0xdd2c00,
	}
	brown = simplecolor.NamedPalette{
		"50": 0xefebe9, "100": 0xd7ccc8, "200": 0xbcaaa4, "300": 0xa1887f,
		"400": 0x8d6e63, "500": 0x795548, "600": 0x6d4c41, "700": 0x5d4037,
		"800": 0x4e342e, "900": 0x3e2723,
	}
	grey = simplecolor.NamedPalette{
		"50": 0xfafafa, "100": 0xf5f5f5, "200": 0xeeeeee, "300": 0xe0e0e0,
		"400": 0xbdbdbd, "500": 0x9e9e9e, "600": 0x757575, "700": 0x616161,
		"800": 0x424242, "900": 0x212121,
	}
	blueGrey = simplecolor.NamedPalette{
		"50": 0xeceff1, "100": 0xcfd8dc, "200": 0xb0bec5, "300": 0x90a4ae,
		"400": 0x78909c, "500": 0x607d8b, "600": 0x546e7a, "700": 0x455a64,
		"800": 0x37474f, "900": 0x263238,
	}
)

type ColorName int

const (
	Red ColorName = iota
	Pink
	Purple
	DeepPurple
	Blue
	LightBlue
	Cyan
	Teal
	Green
	LightGreen
	Lime
	Yellow
	Amber
	Orange
	DeepOrange
	Brown
	Grey
	BlueGrey
	Gray     = Grey
	BlueGray = BlueGrey
)

var spectrum = map[ColorName]simplecolor.NamedPalette{
	Red:    red,
	Pink:   pink,
	Purple: purple,
	// These colors render too similarly to Purple for most use cases
	// "DeepPurple": DeepPurple,
	Blue:       blue,
	LightBlue:  lightBlue,
	Cyan:       cyan,
	Teal:       teal,
	Green:      green,
	LightGreen: lightGreen,
	Lime:       lime,
	Yellow:     yellow,
	Amber:      amber,
	Orange:     orange,
	DeepOrange: deepOrange,
	Brown:      brown,
	Grey:       grey,
	BlueGrey:   blueGrey,
}

func GetPalette() (colors simplecolor.SimplePalette) {
	for _, cp := range spectrum {
		for _, c := range cp {
			colors = append(colors, c)
		}
	}
	sort.Sort(colors)
	return
}

func GetColorsForShade(shade string) (colors simplecolor.SimplePalette) {
	for _, cp := range spectrum {
		if hue, ok := cp[shade]; ok {
			colors = append(colors, hue)
		}
	}
	sort.Sort(colors)
	return
}

// Parameter must be one of the following:
func GetShadesForColorName(color ColorName) (colors simplecolor.SimplePalette) {
	if c, ok := spectrum[color]; ok {
		for _, cp := range c {
			colors = append(colors, cp)
		}
	}
	sort.Sort(colors)
	return
}
