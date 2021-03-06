package vpn

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kapmahc/h2o/web"
)

func (p *Plugin) indexUsers(c *gin.Context, _ string) (interface{}, error) {
	var items []User
	if err := p.Db.Order("updated_at DESC").Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

type fmUserNew struct {
	FullName             string `form:"fullName" binding:"required,max=255"`
	Email                string `form:"email" binding:"required,email"`
	Password             string `form:"password" binding:"min=6,max=32"`
	PasswordConfirmation string `form:"passwordConfirmation" binding:"eqfield=Password"`
	Details              string `form:"details"`
	Enable               bool   `form:"enable"`
	StartUp              string `form:"startUp"`
	ShutDown             string `form:"shutDown"`
}

func (p *Plugin) createUser(c *gin.Context, _ string, o interface{}) (interface{}, error) {
	fm := o.(*fmUserNew)
	startUp, err := time.Parse(web.FormatDateInput, fm.StartUp)
	if err != nil {
		return nil, err
	}
	shutDown, err := time.Parse(web.FormatDateInput, fm.ShutDown)
	if err != nil {
		return nil, err
	}
	user := User{
		FullName: fm.FullName,
		Email:    fm.Email,
		Details:  fm.Details,
		Enable:   fm.Enable,
		StartUp:  startUp,
		ShutDown: shutDown,
	}
	if err := user.SetPassword(fm.Password); err != nil {
		return nil, err
	}
	if err := p.Db.Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

type fmUserEdit struct {
	FullName string `form:"fullName" binding:"required,max=255"`
	Details  string `form:"details"`
	Enable   bool   `form:"enable"`
	StartUp  string `form:"startUp"`
	ShutDown string `form:"shutDown"`
}

func (p *Plugin) updateUser(c *gin.Context, _ string, o interface{}) (interface{}, error) {
	fm := o.(*fmUserEdit)
	startUp, err := time.Parse(web.FormatDateInput, fm.StartUp)
	if err != nil {
		return nil, err
	}
	shutDown, err := time.Parse(web.FormatDateInput, fm.ShutDown)
	if err != nil {
		return nil, err
	}
	if err := p.Db.Model(&User{}).
		Where("id = ?", c.Param("id")).
		Updates(map[string]interface{}{
			"full_name": fm.FullName,
			"enable":    fm.Enable,
			"start_up":  startUp,
			"shut_down": shutDown,
			"details":   fm.Details,
		}).Error; err != nil {
		return nil, err
	}

	return gin.H{}, nil
}

type fmUserResetPassword struct {
	Password             string `form:"password" binding:"min=6,max=32"`
	PasswordConfirmation string `form:"passwordConfirmation" binding:"eqfield=Password"`
}

func (p *Plugin) resetUserPassword(c *gin.Context, _ string, o interface{}) (interface{}, error) {
	var item User
	if err := p.Db.Where("id = ?", c.Param("id")).First(&item).Error; err != nil {
		return nil, err
	}
	fm := o.(*fmUserResetPassword)

	if err := item.SetPassword(fm.Password); err != nil {
		return nil, err
	}
	if err := p.Db.Model(&item).
		Updates(map[string]interface{}{
			"password": item.Password,
		}).Error; err != nil {
		return nil, err
	}

	return gin.H{}, nil
}

func (p *Plugin) destroyUser(c *gin.Context, _ string) (interface{}, error) {
	err := p.Db.
		Where("id = ?", c.Param("id")).
		Delete(User{}).Error
	return gin.H{}, err
}

type fmUserChangePassword struct {
	Email                string `form:"email" binding:"required,email"`
	CurrentPassword      string `form:"currentPassword" binding:"required"`
	NewPassword          string `form:"newPassword" binding:"min=6,max=32"`
	PasswordConfirmation string `form:"passwordConfirmation" binding:"eqfield=NewPassword"`
}

func (p *Plugin) changeUserPassword(c *gin.Context, l string, o interface{}) (interface{}, error) {
	fm := o.(*fmUserChangePassword)
	var user User
	if err := p.Db.Where("email = ?", fm.Email).First(&user).Error; err != nil {
		return nil, err
	}
	if !user.ChkPassword(fm.CurrentPassword) {
		return nil, p.I18n.E(l, "ops.vpn.users.email-password-not-match")
	}
	if err := user.SetPassword(fm.NewPassword); err != nil {
		return nil, err
	}

	if err := p.Db.Model(user).
		Updates(map[string]interface{}{
			"password": user.Password,
		}).Error; err != nil {
		return nil, err
	}

	return gin.H{}, nil
}
