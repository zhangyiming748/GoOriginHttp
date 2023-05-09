package logic

import (
	mysql2 "GoOriginHttp/mysql"
	"os"
	"testing"
	"time"
)

func TestGetAmap(t *testing.T) {
	mysql2.SetEngine()
	for {
		var tam = &toAmapWeather{
			Key:        "080bdfdad2be2b55e883361c1de6fbf8",
			City:       "110107",
			Extensions: "all",
			Output:     "JSON",
		}
		b := getFromAmap(tam)
		file, err := os.OpenFile("all.json", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0777)
		if err != nil {
			return
		}
		defer file.Close()
		file.WriteString(string(b))
		time.Sleep(10 * time.Minute)
	}

}
