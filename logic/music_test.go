package logic

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestSearchByKeywordmain(t *testing.T) {
	//mysql.SetEngine()
	res := GetTopByKeyword("止战之殇")
	t.Log(res)
}
func TestUnM(t *testing.T) {
	strjson := "{\n  \"code\": 200,\n  \"cover\": \"https://y.qq.com/music/photo_new/T002R300x300M000003DFRzD192KKD_1.jpg\",\n  \"name\": \"止战之殇\",\n  \"singer\": \"周杰伦\",\n  \"quality\": \"SQ无损\",\n  \"url\": \"http://aqqmusic.tc.qq.com/amobile.music.tc.qq.com/F000000FNKvf4dGDjm.flac?guid=B69D8BC956E47C2B65440380380B7E9A&vkey=5D587A9DC7A03C732F97A31FB90BADF19E8916797B7E241D14742CA3DF6F27827810B69476833CBF1F728A0E1538D0463C6DED100D076DAC&uin=1828222534&fromtag=119117\",\n  \"tips\": \"慕名API：http://xiaoapi.cn\"\n}"
	var m Music
	err := json.Unmarshal([]byte(strjson), &m)
	if err != nil {
		fmt.Println("error")
	}
	t.Logf("%+v\n", m)
}
