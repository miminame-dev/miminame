package controller

import (
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/labstack/echo/v4"
)

type PostProcessVideoRequest struct {
	VideoID string `json:"videoID"`
}

func (r *PostProcessVideoRequest) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.VideoID, validation.Required),
	)
}

type PostProcessVideoResponse struct{}

func (ctrl *Controller) PostProcessVideo(c echo.Context) error {
	req := new(PostProcessVideoRequest)
	res := new(PostProcessVideoResponse)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, res)
	}

	if err := req.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, res)
	}

	if err := ctrl.props.ProcessVideoService.StartProcess(req.VideoID); err != nil {
		c.Logger().Errorf("failed to start process: %+v", err)
		return c.JSON(http.StatusInternalServerError, res)
	}

	return c.JSON(http.StatusOK, res)
}
