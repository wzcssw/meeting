package models

// CcUser CcUser
type CcUser struct {
	Model
	ConsultationCenterID uint   `json:"consultation_center_id"`
	Username             string `json:"username"`
	PasswordDigest       string `json:"password_digest"`
	Realname             string `json:"realname"`
	Phone                string `json:"phone"`
	Email                string `json:"email"`
	IsUse                bool   `json:"is_use"`
	CcRoleID             uint   `json:"cc_role_id"`
	Rank                 uint   `json:"rank"`
}

// GetAllUser 查询所有用户
func GetAllUser(page int, size int) Result {
	result := NewResult(page, size)
	users := make([]CcUser, result.PageSize)
	Orm.Offset(result.OffSet()).Limit(result.PageSize).Find(&users).Count(&result.AllCount)
	result.Data = users
	return result
}
