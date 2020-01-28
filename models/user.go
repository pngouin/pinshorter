package models

type UserConn struct {
	Name     string `json:"name"`
	Password string `json:"password"`

	Shared
}

type UserInfo struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

func (uc UserConn) ToUserInfo() UserInfo {
	return UserInfo{
		Id:   uc.Id,
		Name: uc.Name,
	}
}
