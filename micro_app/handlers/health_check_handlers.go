package handlers
import (
	"net/http"

	"github.com/labstack/echo"

	"micro_apps/micro_app/models"
)

func HealthCheck(c echo.Context) error {

	resp := models.HealthCheckResponse{
		Message: "Everything is good!",
	}
	return c.JSON(http.StatusOK, resp)
}
