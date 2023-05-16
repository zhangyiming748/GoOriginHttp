package model

import (
	"GoOriginHttp/mysql"
	"golang.org/x/exp/slog"
	"time"
)

type Music struct {
	Id         int       `xorm:"not null pk autoincr comment('主键id') INT(11)" json:"id"`
	Code       string    `xorm:"comment('状态码') VARCHAR(255)" json:"code"`
	Cover      string    `xorm:"comment('专辑封面') VARCHAR(255)" json:"cover"`
	Name       string    `xorm:"comment('歌曲名') VARCHAR(255)" json:"name"`
	Singer     string    `xorm:"comment('歌手') VARCHAR(255)" json:"singer"`
	Quality    string    `xorm:"comment('音质') VARCHAR(255)" json:"quality"`
	Url        string    `xorm:"comment('播放链接') TEXT" json:"url"`
	Tips       string    `xorm:"comment('赞助商') VARCHAR(255)" json:"tips"`
	UpdateTime time.Time `xorm:"updated comment('更新时间') DateTime" json:"update_time"`
	CreateTime time.Time `xorm:"created comment('创建时间') DateTime" json:"create_time"`
	DeleteTime time.Time `xorm:"deleted comment('删除时间') DateTime" json:"delete_time"`
}

func SyncMusic() {
	err := mysql.GetSession().Sync2(new(Music))
	if err != nil {
		slog.Error("同步数据表出错", slog.Any("错误原文", err))
		return
	}
}
func (m Music) InsertOne() (int64, error) {
	return mysql.GetSession().Insert(m)
}
