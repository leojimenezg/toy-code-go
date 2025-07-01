package slices

// URL: https://go.dev/tour/moretypes/18
import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	outter_slice := make([][]uint8, dy)
	for i := range outter_slice {
		inner_slice := make([]uint8, dx)
		for j := range inner_slice {
			inner_slice[j] = uint8(((dx - i) + (dx + j)) / 2)
		}
		outter_slice[i] = inner_slice
	}
	// return [dy * [dx * uint8]]
	return outter_slice
}

func main() {
	pic.Show(Pic)
}
