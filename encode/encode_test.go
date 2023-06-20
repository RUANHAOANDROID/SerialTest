package encode

import (
	"fmt"
	"testing"
)

func TestGenerateFrame(t *testing.T) {

	//01030000000AC5CD
	//03 05 00 01 00 00 9d e8 -> [3 5 0 1 0 0 157 232]
	frame01 := GenerateFrame(03, 05, 00, 01, 00, 00)
	fmt.Println(frame01)
	//01 03 00 00 00 03 05 CB -> [1 3 0 0 0 3 5 203]
	frame02 := GenerateFrame(01, 03, 00, 00, 00, 03)
	fmt.Println(frame02)
	//01 05 00 01 ff 00 dd fa -> [1 5 0 1 255 0 221 250]
	frame03 := GenerateFrame(01, 05, 00, 01, 0xff, 00)
	fmt.Println(frame03)
	//03 05 00 01 ff 00 dc 18 -> [3 5 0 1 255 0 220 24]
	frame04 := GenerateFrame(03, 05, 00, 01, 0xff, 00)
	fmt.Println(frame04)

}
