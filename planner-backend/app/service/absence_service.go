package service

import (
	"database/sql"
	"log/slog"
	"net/http"
	"planner-backend/app/constant"
	"planner-backend/app/pkg"
	"planner-backend/app/repository"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

type AbsenceService interface {
	GetAllAbsencies(c *gin.Context)
}

type AbsenceServiceImpl struct {
	AbsenceRepository repository.AbsenceRepository
}

func (a AbsenceServiceImpl) GetAllAbsencies(c *gin.Context) {
	/* GetAllAbsencies is a function to get all absencies
	 * @param c is gin context
	 * @return void
	 */
	defer pkg.PanicHandler(c)
	slog.Info("start to execute program get all absencies")

	departmentID := c.Param("departmentID")
	if departmentID == "" {
		pkg.PanicException(constant.InvalidRequest)
	}

	date := c.Query("date")
	if date == "" {
		pkg.PanicException(constant.InvalidRequest)
	}
	if _, err := time.Parse("2006-01-02", date); err != nil {
		pkg.PanicException(constant.InvalidRequest)
	}

	rawData, err := a.AbsenceRepository.FindAllAbsencies(departmentID, date)
	switch err {
	case nil:
		break
	case sql.ErrNoRows:
		c.JSON(http.StatusOK, pkg.BuildResponse(constant.DataNotFound, pkg.Null()))
		return
	default:
		slog.Error("Error when fetching data from database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	data := mapAbsenciesToAbsenciesResponse(rawData)

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

var absencyServiceSet = wire.NewSet(
	wire.Struct(new(AbsenceServiceImpl), "*"),
	wire.Bind(new(AbsenceService), new(*AbsenceServiceImpl)),
)
