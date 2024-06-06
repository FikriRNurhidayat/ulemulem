package invitation_controller

import (
	"html/template"
	"io"

	echo "github.com/labstack/echo/v4"
)

type InvitationTemplate struct {
  Templates *template.Template
}

func (t *InvitationTemplate) Init() *InvitationTemplate {
  t.Templates = template.Must(template.ParseGlob("internal/domain/invitation/web/template/*.html"))

  return t
}

func (t *InvitationTemplate) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return t.Templates.ExecuteTemplate(w, name, data)
}
