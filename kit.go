package kit

import (
	"crypto/md5"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"
	"time"

	log "github.com/go-ozzo/ozzo-log"
	"github.com/go-vgo/robotgo"
)

// Area is a screen area,
// X,Y is the start point
// W,H is the area's width and hight
type Area struct {
	X, Y, W, H int
}

var (
	Logger      log.Logger
	LogFileName = ""
	Screen      Area
)

func init() {
	w, h := robotgo.GetScreenSize()
	Screen = Area{0, 0, w, h}
}

// InitLogger the func Log() will initialize it.
func InitLogger() {
	if Logger.Category != "" {
		return
	}
	loggerTmp := log.NewLogger()
	t1 := log.NewConsoleTarget()
	t2 := log.NewFileTarget()
	if LogFileName == "" {
		t2.FileName = "app" + string(time.Now().Format(".2006-01-02")) + ".log"
	} else {
		t2.FileName = LogFileName
	}
	t2.MaxLevel = log.LevelNotice
	loggerTmp.Targets = append(loggerTmp.Targets, t1, t2)
	loggerTmp.Open()
	Logger = *loggerTmp
	Logger.Info("InitLogger")
}

// Sleep wait x millisecond
func Sleep(x int) {
	time.Sleep(time.Duration(x) * time.Millisecond)
}

// Fmt is print any variable
func Fmt(desc string, args ...interface{}) string {
	fmt.Println(Logger.Category, desc, args)
	return ""
}

// Log is write log and output any variable easily.
func Log(desc string, args ...interface{}) string {
	InitLogger()
	argsDesc := ""
	for _, val := range args {
		argsDesc += " " + strings.TrimSpace(fmt.Sprintln(val))
	}
	_, file, line, _ := runtime.Caller(1)
	fileName := path.Base(file)
	fileLine := fileName + " " + strconv.Itoa(line) + " "
	Logger.Notice(fileLine + "  " + desc + argsDesc)
	return desc + argsDesc
}

// Exit exit the program.
func Exit(desc string, args ...interface{}) {
	InitLogger()
	argsDesc := ""
	for _, val := range args {
		argsDesc += " " + strings.TrimSpace(fmt.Sprintln(val))
	}
	_, file, line, _ := runtime.Caller(1)
	fileName := path.Base(file)
	fileLine := fileName + " " + strconv.Itoa(line) + " "
	Logger.Error(fileLine + "  " + desc + argsDesc)
	Logger.Close()
	robotgo.ShowAlert("DNF GO Error!", desc)
	os.Exit(1)
}

// Mkdirs
func Mkdirs(imgpath string) error {
	_, err := os.Stat(imgpath)
	if err == nil {
		return err
	}
	errm := os.MkdirAll(imgpath, 0755)
	if errm == nil {
		return errm
	}
	return nil
}

// Authcode is a rc4 crypto func from Discuz!.
// Interoperability with Discuz! authcode().
func Authcode(str string, operation string, key string, expiry int) (string, error) {

	// keya, keyb, keyc
	ckeyLength := 4
	newKey := Md5(key)
	keya := Md5(string([]byte(newKey)[:16]))
	keyb := Md5(string([]byte(newKey)[16:32]))
	keyc := ""
	if operation == "DECODE" {
		keyc = string([]byte(str)[:ckeyLength])
	} else {
		microTime := strconv.FormatInt(time.Now().UnixNano(), 10)
		keyc = string([]byte(Md5(microTime))[2:6])
	}

	// crypt key
	cryptKey := keya + Md5(keya+keyc)
	keyLength := len(cryptKey)

	// string
	strNew := ""
	if operation == "DECODE" {
		if ckeyLength > len(str) {
			ckeyLength = 0
		}
		strTmp := string([]byte(str)[ckeyLength:len(str)])
		strNew = Base64Decode(strTmp)
	} else {
		var expiryStr int64
		if expiry > 0 {
			expiryStr = time.Now().Unix() + int64(expiry)
		}
		strPre := fmt.Sprintf("%010d", expiryStr)
		strNew = strPre + string([]byte(Md5(str + keyb))[:16]) + str
	}
	strNewLength := len(strNew)
	phpChr := [256]string{"AA==", "AQ==", "Ag==", "Aw==", "BA==", "BQ==", "Bg==", "Bw==", "CA==", "CQ==", "Cg==", "Cw==", "DA==", "DQ==", "Dg==", "Dw==", "EA==", "EQ==", "Eg==", "Ew==", "FA==", "FQ==", "Fg==", "Fw==", "GA==", "GQ==", "Gg==", "Gw==", "HA==", "HQ==", "Hg==", "Hw==", "IA==", "IQ==", "Ig==", "Iw==", "JA==", "JQ==", "Jg==", "Jw==", "KA==", "KQ==", "Kg==", "Kw==", "LA==", "LQ==", "Lg==", "Lw==", "MA==", "MQ==", "Mg==", "Mw==", "NA==", "NQ==", "Ng==", "Nw==", "OA==", "OQ==", "Og==", "Ow==", "PA==", "PQ==", "Pg==", "Pw==", "QA==", "QQ==", "Qg==", "Qw==", "RA==", "RQ==", "Rg==", "Rw==", "SA==", "SQ==", "Sg==", "Sw==", "TA==", "TQ==", "Tg==", "Tw==", "UA==", "UQ==", "Ug==", "Uw==", "VA==", "VQ==", "Vg==", "Vw==", "WA==", "WQ==", "Wg==", "Ww==", "XA==", "XQ==", "Xg==", "Xw==", "YA==", "YQ==", "Yg==", "Yw==", "ZA==", "ZQ==", "Zg==", "Zw==", "aA==", "aQ==", "ag==", "aw==", "bA==", "bQ==", "bg==", "bw==", "cA==", "cQ==", "cg==", "cw==", "dA==", "dQ==", "dg==", "dw==", "eA==", "eQ==", "eg==", "ew==", "fA==", "fQ==", "fg==", "fw==", "gA==", "gQ==", "gg==", "gw==", "hA==", "hQ==", "hg==", "hw==", "iA==", "iQ==", "ig==", "iw==", "jA==", "jQ==", "jg==", "jw==", "kA==", "kQ==", "kg==", "kw==", "lA==", "lQ==", "lg==", "lw==", "mA==", "mQ==", "mg==", "mw==", "nA==", "nQ==", "ng==", "nw==", "oA==", "oQ==", "og==", "ow==", "pA==", "pQ==", "pg==", "pw==", "qA==", "qQ==", "qg==", "qw==", "rA==", "rQ==", "rg==", "rw==", "sA==", "sQ==", "sg==", "sw==", "tA==", "tQ==", "tg==", "tw==", "uA==", "uQ==", "ug==", "uw==", "vA==", "vQ==", "vg==", "vw==", "wA==", "wQ==", "wg==", "ww==", "xA==", "xQ==", "xg==", "xw==", "yA==", "yQ==", "yg==", "yw==", "zA==", "zQ==", "zg==", "zw==", "0A==", "0Q==", "0g==", "0w==", "1A==", "1Q==", "1g==", "1w==", "2A==", "2Q==", "2g==", "2w==", "3A==", "3Q==", "3g==", "3w==", "4A==", "4Q==", "4g==", "4w==", "5A==", "5Q==", "5g==", "5w==", "6A==", "6Q==", "6g==", "6w==", "7A==", "7Q==", "7g==", "7w==", "8A==", "8Q==", "8g==", "8w==", "9A==", "9Q==", "9g==", "9w==", "+A==", "+Q==", "+g==", "+w==", "/A==", "/Q==", "/g==", "/w=="}

	box := [256]int{}
	for i := 0; i < 256; i++ {
		box[i] = i
	}
	rndkey := [256]int{}
	for i := 0; i < 256; i++ {
		rndkey[i] = Ord(string(cryptKey[i%keyLength]))
	}

	for j, i := 0, 0; i < 256; i++ {
		j = (j + box[i] + rndkey[i]) % 256
		tmp := box[i]
		box[i] = box[j]
		box[j] = tmp
	}
	result := ""
	for a, j, i := 0, 0, 0; i < strNewLength; i++ {
		a = (a + 1) % 256
		j = (j + box[a]) % 256
		tmp := box[a]
		box[a] = box[j]
		box[j] = tmp
		ordNum := Ord(string([]byte(strNew)[i:i+1])) ^ (box[(box[a]+box[j])%256])
		result += Base64Decode(phpChr[ordNum])
	}

	if operation == "DECODE" {
		timePref := string([]byte(result)[:10])
		verifyTime, _ := strconv.Atoi(timePref)
		verifyStr := string([]byte(result)[10:26])
		proofStr := string([]byte(Md5(string([]byte(result)[26:strNewLength]) + keyb))[:16])
		if (verifyTime == 0 || int64(verifyTime)-time.Now().Unix() > 0) && verifyStr == proofStr {
			return string([]byte(result)[26:strNewLength]), nil
		} else {
			return "", errors.New("DECODE Error")
		}
	} else {
		return keyc + strings.Replace(Base64Encode(result), "=", "", -1), nil
	}
}

// Md5 is just md5. ^ ^
func Md5(str string) string {
	h := md5.New()
	io.WriteString(h, str)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// Chr a chr func, 97 => a
func Chr(num int) string {
	return string(rune(num))
}

// Ord a ord func, a => 97.
func Ord(str string) int {
	return int(str[0])
}

// Base64Encode ^ ^, just base64 encode.
func Base64Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

// Base64Decode is just base64 decode.
func Base64Decode(str string) string {
	for i := 0; i < 3; i++ {
		decodeBytes, err := base64.StdEncoding.DecodeString(str)
		if err == nil {
			return string(decodeBytes)
		}
		str += "="
	}
	return ""
}
