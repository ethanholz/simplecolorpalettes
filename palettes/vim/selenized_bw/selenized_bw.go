package selenized_bw

import (
	"sort"
	"sync"

	"github.com/taigrr/go-colorpalettes/simplecolor"
)

var (
	once   sync.Once
	colors = simplecolor.SimplePalette{
		0x0054cf,
		0x0064e4,
		0x008400,
		0x009a8a,
		0x00ad9c,
		0x181818,
		0x1d9700,
		0x252525,
		0x282828,
		0x368aeb,
		0x3b3b3b,
		0x3fc5b7,
		0x474747,
		0x4f9cfe,
		0x56d8c9,
		0x6b40c3,
		0x70b433,
		0x777777,
		0x7f51d6,
		0x83c746,
		0x878787,
		0xa580e2,
		0xaf8500,
		0xb891f5,
		0xb9b9b9,
		0xba3700,
		0xbf0000,
		0xc49700,
		0xc7008b,
		0xcdcdcd,
		0xd04a00,
		0xd6000c,
		0xdbb32d,
		0xdd0f9d,
		0xdedede,
		0xe67f43,
		0xeb6eb7,
		0xebebeb,
		0xed4a46,
		0xefc541,
		0xfa9153,
		0xff5e56,
		0xff81ca,
		0xffffff,
	}
)

func GetPalette() (sp simplecolor.SimplePalette) {
	once.Do(func() {
		sort.Sort(colors)
	})
	return colors
}
