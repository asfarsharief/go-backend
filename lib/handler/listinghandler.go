package handler

import (
	"net/http"

	"github.com/Karkhana-Studio/Backend-server/common/logingservice"
	"github.com/Karkhana-Studio/Backend-server/lib/config"
	"github.com/Karkhana-Studio/Backend-server/lib/payload"
	"github.com/Karkhana-Studio/Backend-server/lib/service"
	"github.com/labstack/echo/v4"
)

type ListingHandler struct {
	*config.ServerConfiguration
	ListingService service.ListingServiceInterface
}

func NewListingHandler() {
}

func (lh *ListingHandler) ListWomenPageDetails(c echo.Context) error {
	logingservice.Info("Women Listing Request Reached")

	productMinis := lh.ListingService.GetDataForWomenListingPage()
	apiResponse := payload.WomenListingResponsePayload{
		APIResponse:  "success",
		APIMsgCode:   200,
		ProductMinis: productMinis,
	}
	return c.JSON(http.StatusOK, apiResponse)
}
