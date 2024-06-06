package invitation_controller

import (
	"net/http"

	"github.com/fikrirnurhidayat/dhasar"
	invitation_service "github.com/fikrirnurhidayat/ulemulem/internal/domain/invitation/service"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type InvitationWebController interface {
	Register(*echo.Echo)
	Show(echo.Context) error
}

type InvitationWebControllerImpl struct {
	GetInvitationService invitation_service.GetInvitationService
}

// Register implements InvitationWebController.
func (i *InvitationWebControllerImpl) Register(e *echo.Echo) {
	renderer := &InvitationTemplate{}
	e.Renderer = renderer.Init()
	e.GET("/", i.Show)
}

func (i *InvitationWebControllerImpl) WithGetInvitationService(svc invitation_service.GetInvitationService) *InvitationWebControllerImpl {
	i.GetInvitationService = svc
	return i
}

type InvitationWebControllerBuilder interface {
	WithGetInvitationService(invitation_service.GetInvitationService) *InvitationWebControllerImpl
}

func (i *InvitationWebControllerImpl) Show(c echo.Context) error {
	id, err := uuid.Parse(c.QueryParam("id"))
	if err != nil {
		return dhasar.ErrInvalidUUID
	}

	ctx := c.Request().Context()

	result, err := i.GetInvitationService.Call(ctx, &invitation_service.GetInvitationParams{
		ID: id,
	})

	if err != nil {
		return err
	}

	return c.Render(http.StatusOK, "invitation.show.html", map[string]any{
		"Groom":           "Fikri Rahmat Nurhidayat",
		"Bride":           "Dhea Arintiara",
		"Date":            "Sabtu, 24 Agustus",
		"Year":            "2024",
		"Time":            "Bada Isya",
		"LocationName":    "Ndalem Danarhadi",
		"LocationAddress": "Jl. Bhayangkara No.55, Panularan, Kec. Laweyan, Kota Surakarta, Jawa Tengah",
		"RecipientName":   result.Invitation.RecipientName,
	})
}

func NewInvitationWebController() *InvitationWebControllerImpl {
	return &InvitationWebControllerImpl{
		GetInvitationService: nil,
	}
}

var _ InvitationWebController = (*InvitationWebControllerImpl)(nil)
var _ InvitationWebControllerBuilder = (*InvitationWebControllerImpl)(nil)
