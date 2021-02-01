package database

import (
	sgJwt "salmon-go/jwt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

//TableModel 包含基础字段
type TableModel struct {
	ID         int64  `json:"id,string" db:"id"`
	CreatorID  int64  `json:"creator_id" db:"creator_id"`
	Creator    string `json:"creator" db:"creator"`
	CreatorIP  string `json:"creator_ip" db:"creator_ip"`
	CreateTime string `json:"create_time" db:"create_time"`
	ModifierID int64  `json:"modifier_id" db:"modifier_id"`
	Modifier   string `json:"modifier" db:"modifier"`
	ModifierIP string `json:"modifier_ip" db:"modifier_ip"`
	ModifyTime string `json:"modify_time" db:"modify_time" gorm:"default:NULL"`
	DeleteTag  int    `json:"delete_tag" db:"delete_tag"`
}

//Create ...
func (t *TableModel) Create(c *gin.Context) {

	t.ID = NewSnowflake(1).Generate()
	userID := c.Request.Header.Get(sgJwt.KeyTokenClaim)
	t.CreatorID, _ = strconv.ParseInt(userID, 10, 64)
	t.CreatorIP = c.ClientIP()
	t.CreateTime = time.Now().String()[0:19]
}

//Modify ...
func (t *TableModel) Modify(c *gin.Context) {
	userID := c.Request.Header.Get(sgJwt.KeyTokenClaim)
	t.ModifierID, _ = strconv.ParseInt(userID, 10, 64)
	t.ModifierIP = c.ClientIP()
	t.ModifyTime = time.Now().String()[0:19]
}
