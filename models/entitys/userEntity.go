package entitys

type UserEntity struct {
	IdUser int64
	WechatName string
	WechatAvatar string
	Openid string
	CreatedAt int64
	LoginAt int64
	Sex int
}

func (this *UserEntity) GetIdUser() int64 {
	if this.IdUser != 0 {
		return this.IdUser
	}
	return 0
}

func (this *UserEntity) GetWechatName() string {
	if this.WechatName != "" {
		return this.WechatName
	}
	return ""
}

func (this *UserEntity) GetWechatAvatar() string {
	if this.WechatAvatar != "" {
		return this.WechatAvatar
	}
	return ""
}


func (this *UserEntity) GetOpenid() string {
	if this.Openid != "" {
		return this.Openid
	}
	return ""
}

func (this *UserEntity) GetCreatedAt() int64 {
	if this.CreatedAt != 0 {
		return this.CreatedAt
	}
	return 0
}

func (this *UserEntity) GetLoginAt() int64 {
	if this.LoginAt != 0 {
		return this.LoginAt
	}
	return 0
}

func (this *UserEntity) GetSex() int {
	if this.Sex != 0 {
		return this.Sex
	}
	return 0
}

