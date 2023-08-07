package clock

import (
	"net/http"

	"github.com/rhuanpk/pointers-angle/internal/infra/db"
	"github.com/rhuanpk/pointers-angle/internal/models"
	"github.com/rhuanpk/pointers-angle/pkg/logger"
	"github.com/rhuanpk/pointers-angle/pkg/utils/api"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func angle(ctx *gin.Context) {
	var (
		clockGet *models.Clock
		result   int16
	)

	clock := models.NewClock()

	if err := ctx.ShouldBindUri(&clock); err != nil {
		api.Return(
			ctx,
			http.StatusBadRequest,
			"wrong RANGE or TYPE for hour or minute values",
			err.Error(),
		)
		return
	}

	data := db.Tx.Where("hour", *clock.Hour).Where("minute", *clock.Minute).First(&clockGet)
	// The second validation is necessary compare string with string because the "gorm.ErrRecordNotFound" not work with "errors.Is()".
	if data.Error != nil && data.Error.Error() != gorm.ErrRecordNotFound.Error() {
		api.Return(
			ctx,
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
			data.Error.Error(),
		)
		return
	}

	if data.RowsAffected > 0 {
		result = *clockGet.Angle
	} else {
		*clock.Angle = int16(((11 * int(*clock.Minute)) - (60 * int(*clock.Hour))) / 2)
		if *clock.Angle > 180 {
			*clock.Angle -= 360
		} else if *clock.Angle < -180 {
			*clock.Angle -= -360
		}
		if *clock.Angle < 0 {
			*clock.Angle *= -1
		}
		if err := db.Tx.Create(&clock).Error; err != nil {
			logger.Error.Println(err.Error())
		}
		result = *clock.Angle
	}

	ctx.JSON(http.StatusOK, gin.H{"angle": result})
}
