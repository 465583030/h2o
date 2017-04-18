package web

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/kapmahc/h2o/web/i18n"
	"github.com/unrolled/render"
)

// Wrap wrap
type Wrap struct {
	Render *render.Render `inject:""`
}

// FORM wrap form handler
func (p *Wrap) FORM(fm interface{}, fn func(*gin.Context, interface{}) (interface{}, error)) gin.HandlerFunc {
	return p.JSON(func(c *gin.Context) (interface{}, error) {
		if err := c.Bind(fm); err != nil {
			return nil, err
		}
		return fn(c, fm)
	})
}

// HTML wrap html render
func (p *Wrap) HTML(t string, f func(*gin.Context, string) (gin.H, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		lang := c.MustGet(i18n.LOCALE).(string)
		if v, e := f(c, lang); e == nil {
			v["lang"] = lang
			p.Render.HTML(c.Writer, http.StatusOK, t, v)
		} else {
			log.Error(e)
			c.String(http.StatusInternalServerError, e.Error())
		}
	}
}

// XML wrap xml render
func (p *Wrap) XML(f func(*gin.Context) (interface{}, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		if v, e := f(c); e == nil {
			c.XML(http.StatusOK, v)
		} else {
			log.Error(e)
			c.String(http.StatusInternalServerError, e.Error())
		}
	}
}

// JSON wrap json render
func (p *Wrap) JSON(f func(*gin.Context) (interface{}, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		if v, e := f(c); e == nil {
			c.JSON(http.StatusOK, v)
		} else {
			log.Error(e)
			c.String(http.StatusInternalServerError, e.Error())
		}
	}
}
