package kit

import (
	"math/rand"
	"strconv"
	"strings"
)

var KeyUpMap = make(map[string]bool)

// For the param 'key string',You can refer to:
// https://github.com/go-vgo/robotgo/blob/master/docs/keys.md

// KeyPress is press key func
func KeyPress(key string) {
	KeyDown(key)
	Sleep(25 + rand.Intn(10))
	KeyUp(key)
}

// KeyDown is press down a key
func KeyDown(key string) {
	keyToggle(key, "down")
	KeyUpMap[key] = true
}

// KeyUp is press up a key
func KeyUp(key string) {
	keyToggle(key, "up")
	KeyUpMap[key] = false
}

func KeyDefer() {
	for key, isUp := range KeyUpMap {
		if isUp == true {
			KeyUp(key)
		}
	}
}

// KeyOutput output text
func KeyOutput(text string) {

}

func keyToggle(key, upOrDown string) {
	keyMap := make(map[string]string) // 例：a：AA60011c AA6002F01c
	keyMap["0"] = "AA600145 AA6002F045"
	keyMap["1"] = "AA600116 AA6002F016"
	keyMap["2"] = "AA60011e AA6002F01e"
	keyMap["3"] = "AA600126 AA6002F026"
	keyMap["4"] = "AA600125 AA6002F025"
	keyMap["5"] = "AA60012e AA6002F02e"
	keyMap["6"] = "AA600136 AA6002F036"
	keyMap["7"] = "AA60013d AA6002F03d"
	keyMap["8"] = "AA60013e AA6002F03e"
	keyMap["9"] = "AA600146 AA6002F046"
	keyMap["a"] = "AA60011c AA6002F01c"
	keyMap["b"] = "AA600132 AA6002F032"
	keyMap["c"] = "AA600121 AA6002F021"
	keyMap["d"] = "AA600123 AA6002F023"
	keyMap["e"] = "AA600124 AA6002F024"
	keyMap["f"] = "AA60012b AA6002F02b"
	keyMap["g"] = "AA600134 AA6002F034"
	keyMap["h"] = "AA600133 AA6002F033"
	keyMap["i"] = "AA600143 AA6002F043"
	keyMap["j"] = "AA60013b AA6002F03b"
	keyMap["k"] = "AA600142 AA6002F042"
	keyMap["l"] = "AA60014b AA6002F04b"
	keyMap["m"] = "AA60013a AA6002F03a"
	keyMap["n"] = "AA600131 AA6002F031"
	keyMap["o"] = "AA600144 AA6002F044"
	keyMap["p"] = "AA60014d AA6002F04d"
	keyMap["q"] = "AA600115 AA6002F015"
	keyMap["r"] = "AA60012d AA6002F02d"
	keyMap["s"] = "AA60011b AA6002F01b"
	keyMap["t"] = "AA60012c AA6002F02c"
	keyMap["u"] = "AA60013c AA6002F03c"
	keyMap["v"] = "AA60012a AA6002F02a"
	keyMap["w"] = "AA60011d AA6002F01d"
	keyMap["x"] = "AA600122 AA6002F022"
	keyMap["y"] = "AA600135 AA6002F035"
	keyMap["z"] = "AA60011a AA6002F01a"
	keyMap["`"] = "AA60010e AA6002F00e"
	keyMap["-"] = "AA60014e AA6002F04e"
	keyMap["="] = "AA600155 AA6002F055"
	keyMap["\\"] = "AA60015d AA6002F05d"
	keyMap["backspace"] = "AA600166 AA6002F066"
	keyMap["space"] = "AA600129 AA6002F029"
	keyMap["tab"] = "AA60010d AA6002F00d"
	keyMap["caps"] = "AA600158 AA6002F058"
	keyMap["l shift"] = "AA600112 AA6002F012"
	keyMap["l ctrl"] = "AA600114 AA6002F014"
	keyMap["l gui"] = "AA6002E01f AA6003E0F01f"
	keyMap["l  alt"] = "AA600111 AA6002F011"
	keyMap["r shift"] = "AA600159 AA6002F059"
	keyMap["r ctrl"] = "AA6002E014 AA6003E0F014"
	keyMap["r gui"] = "AA6002E027 AA6003E0F027"
	keyMap["r alt"] = "AA6002E011 AA6003E0F011"
	keyMap["apps"] = "AA6002E02f AA6003E0F02f"
	keyMap["enter"] = "AA60015a AA6002F05a"
	keyMap["esc"] = "AA600176 AA6002F076"
	keyMap["f1"] = "AA600105 AA6002F005"
	keyMap["f2"] = "AA600106 AA6002F006"
	keyMap["f3"] = "AA600104 AA6002F004"
	keyMap["f4"] = "AA60010c AA6002F00c"
	keyMap["f5"] = "AA600103 AA6002F003"
	keyMap["f6"] = "AA60010b AA6002F00b"
	keyMap["f7"] = "AA600183 AA6002F083"
	keyMap["f8"] = "AA60010a AA6002F00a"
	keyMap["f9"] = "AA600101 AA6002F001"
	keyMap["f10"] = "AA600109 AA6002F009"
	keyMap["f11"] = "AA600178 AA6002F078"
	keyMap["f12"] = "AA600107 AA6002F007"
	keyMap["prnt scrn"] = "AA6004e012e07c AA6006e0f012e0f07c"
	keyMap["["] = "AA600154 AA6002F054"
	keyMap["insert"] = "AA6002E070 AA6003E0F070"
	keyMap["home"] = "AA6002E06c AA6003E0F06c"
	keyMap["pg up"] = "AA6002E07d AA6003E0F07d"
	keyMap["delete"] = "AA6002E071 AA6003E0F071"
	keyMap["end"] = "AA6002E069 AA6003E0F069"
	keyMap["pg dn"] = "AA6002E07a AA6003E0F07a"
	keyMap["u arrow"] = "AA6002E075 AA6003E0F075"
	keyMap["l arrow"] = "AA6002E06b AA6003E0F06b"
	keyMap["d arrow"] = "AA6002E072 AA6003E0F072"
	keyMap["r arrow"] = "AA6002E074 AA6003E0F074"
	keyMap["num"] = "AA600177 AA6002F077"
	keyMap["kp /"] = "AA6002E04a AA6003E0F04a"
	keyMap["kp *"] = "AA60017c AA6002F07c"
	keyMap["kp -"] = "AA60017b AA6002F07b"
	keyMap["kp +"] = "AA600179 AA6002F079"
	keyMap["kp en"] = "AA6002E05a AA6003E0F05a"
	keyMap["kp ."] = "AA600171 AA6002F071"
	keyMap["kp 0"] = "AA600170 AA6002F070"
	keyMap["kp 1"] = "AA600169 AA6002F069"
	keyMap["kp 2"] = "AA600172 AA6002F072"
	keyMap["kp 3"] = "AA60017a AA6002F07a"
	keyMap["kp 4"] = "AA60016b AA6002F06b"
	keyMap["kp 5"] = "AA600173 AA6002F073"
	keyMap["kp 6"] = "AA600174 AA6002F074"
	keyMap["kp 7"] = "AA60016c AA6002F06c"
	keyMap["kp 8"] = "AA600175 AA6002F075"
	keyMap["kp 9"] = "AA60017d AA6002F07d"
	keyMap["]"] = "AA60015b AA6002F05b"
	keyMap[";"] = "AA60014c AA6002F04c"
	keyMap["‘"] = "AA600152 AA6002F052"
	keyMap[","] = "AA600141 AA6002F041"
	keyMap["."] = "AA600149 AA6002F049"
	keyMap["scroll"] = "AA60017e AA6002F07e"
	keyMap["pause"] = "AA6003E11477 AA6005e1f014f077"

	if str, ok := keyMap[key]; ok {
		keyStrByte := strings.Fields(str)
		var byteStr string
		if upOrDown == "down" {
			byteStr = keyStrByte[0]
		} else {
			byteStr = keyStrByte[1]
		}
		bytes := []byte{}
		for i := 0; i < len(byteStr); {
			codeStr := string(byteStr[i]) + string(byteStr[i+1])
			codeInt, _ := strconv.ParseInt(codeStr, 16, 32)
			bytes = append(bytes, byte(codeInt))
			i += 2
		}
		SerialWrite(bytes)
	}

}
