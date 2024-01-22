package service

import (
	"log/slog"
	"net/http"
	"planner-backend/app/constant"
	"planner-backend/app/domain/dao"
	"planner-backend/app/domain/dco"
	"planner-backend/app/pkg"
	"planner-backend/app/repository"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

type PersonRelService interface {
	// Function Used by the controller
	AddAbsencyToPerson(c *gin.Context)
	RemoveAbsencyFromPerson(c *gin.Context)
	FindAbsencyForPerson(c *gin.Context)

	AddDepartmentToPerson(c *gin.Context)
	RemoveDepartmentFromPerson(c *gin.Context)

	AddWorkplaceToPerson(c *gin.Context)
	RemoveWorkplaceFromPerson(c *gin.Context)

	AddWeekdayToPerson(c *gin.Context)
	RemoveWeekdayFromPerson(c *gin.Context)
}

type PersonRelServiceImpl struct {
	PersonRelRepository  repository.PersonRelRepository
	PersonRepository     repository.PersonRepository
	DepartmentRepository repository.DepartmentRepository
	WorkplaceRepository  repository.WorkplaceRepository
}

/** Absency */
func (p PersonRelServiceImpl) AddAbsencyToPerson(c *gin.Context) {
	/* AddAbsencyToPerson is a function to add absency to a person
	 * @param c is gin context
	 * @return void
	 */
	defer pkg.PanicHandler(c)
	slog.Info("start to execute program add absency to person")

	personID := c.Param("personID")
	if personID == "" {
		pkg.PanicException(constant.InvalidRequest)
	}
	personID = strings.ToUpper(personID)

	var request dco.AbsenceRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		pkg.PanicException(constant.InvalidRequest)
	}

	person, err := p.PersonRepository.FindPersonByID(personID)
	switch err {
	case nil:
		break
	case pkg.ErrNoRows:
		pkg.PanicException(constant.InvalidRequest)
	default:
		slog.Error("Error when fetching data from database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	absence := mapAbsenceRequestToAbsence(request)

	if err := p.PersonRelRepository.AddAbsencyToPerson(person, absence); err != nil {
		slog.Error("Error when saving data to database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null()))
}

func (p PersonRelServiceImpl) RemoveAbsencyFromPerson(c *gin.Context) {
	/* RemoveAbsencyFromPerson is a function to remove absency from a person
	 * @param c is gin context
	 * @return void
	 */
	defer pkg.PanicHandler(c)
	slog.Info("start to execute program remove absency from person")

	personID := c.Param("personID")
	if personID == "" {
		pkg.PanicException(constant.InvalidRequest)
	}
	personID = strings.ToUpper(personID)

	date := c.Param("date")
	if date == "" {
		pkg.PanicException(constant.InvalidRequest)
	}

	person, err := p.PersonRepository.FindPersonByID(personID)
	switch err {
	case nil:
		break
	case pkg.ErrNoRows:
		pkg.PanicException(constant.InvalidRequest)
	default:
		slog.Error("Error when fetching data from database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	absence, err := p.PersonRelRepository.FindAbsencyForPerson(personID, date)
	switch err {
	case nil:
		break
	case pkg.ErrNoRows:
		pkg.PanicException(constant.InvalidRequest)
	default:
		slog.Error("Error when fetching data from database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	if err := p.PersonRelRepository.RemoveAbsencyFromPerson(person, absence); err != nil {
		slog.Error("Error when saving data to database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null()))
}

func (p PersonRelServiceImpl) FindAbsencyForPerson(c *gin.Context) {
	/* FindAbsencyForPerson is a function to find a given absencies by person
	 * @param c is gin context
	 * @return void
	 */
	defer pkg.PanicHandler(c)
	slog.Info("start to execute program find all absencies by person")

	personID := c.Param("personID")
	if personID == "" {
		pkg.PanicException(constant.InvalidRequest)
	}
	personID = strings.ToUpper(personID)

	date := c.Param("date")
	if date == "" {
		pkg.PanicException(constant.InvalidRequest)
	}

	rawData, err := p.PersonRelRepository.FindAbsencyForPerson(personID, date)
	switch err {
	case nil:
		break
	case pkg.ErrNoRows:
		pkg.PanicException(constant.DataNotFound)
	default:
		slog.Error("Error when fetching data from database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	data := mapAbsenceToAbsenceResponse(rawData)

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func mapAbsenceToAbsenceResponse(absence dao.Absence) dco.AbsenceResponse {
	/** Maps an absence to an absence response */

	return dco.AbsenceResponse{
		PersonID:  absence.PersonID,
		Date:      absence.Date,
		Reason:    absence.Reason,
		CreatedAt: absence.CreatedAt,
	}
}

func mapAbsenceRequestToAbsence(absenceRequest dco.AbsenceRequest) dao.Absence {
	/** Maps an absence request to an absence */

	// null check
	var reason string
	if absenceRequest.Reason == nil {
		reason = ""
	} else {
		reason = *absenceRequest.Reason
	}

	return dao.Absence{
		Date:   absenceRequest.Date,
		Reason: reason,
	}
}

/** Department */
func (p PersonRelServiceImpl) AddDepartmentToPerson(c *gin.Context) {
	/* AddDepartmentToPerson is a function to add department to a person
	 * @param c is gin context
	 * @return void
	 */
	defer pkg.PanicHandler(c)
	slog.Info("start to execute program add department to person")

	personID := c.Param("personID")
	if personID == "" {
		pkg.PanicException(constant.InvalidRequest)
	}
	personID = strings.ToUpper(personID)

	person, err := p.PersonRepository.FindPersonByID(personID)
	switch err {
	case nil:
		break
	case pkg.ErrNoRows:
		pkg.PanicException(constant.InvalidRequest)
	default:
		slog.Error("Error when fetching data from database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	var request dco.RelAddDepartmentRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		pkg.PanicException(constant.InvalidRequest)
	}
	_, err = p.DepartmentRepository.FindDepartmentByID(request.DepartmentID)
	switch err {
	case nil:
		break
	case pkg.ErrNoRows:
		pkg.PanicException(constant.InvalidRequest)
	default:
		slog.Error("Error when fetching data from database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	if err := p.PersonRelRepository.AddDepartmentToPerson(person, request.DepartmentID); err != nil {
		slog.Error("Error when saving data to database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null()))
}

func (p PersonRelServiceImpl) RemoveDepartmentFromPerson(c *gin.Context) {
	/* RemoveDepartmentFromPerson is a function to remove department from a person
	 * @param c is gin context
	 * @return void
	 */
	defer pkg.PanicHandler(c)
	slog.Info("start to execute program remove department from person")

	personID := c.Param("personID")
	if personID == "" {
		pkg.PanicException(constant.InvalidRequest)
	}
	personID = strings.ToUpper(personID)

	departmentID := c.Param("departmentID")
	if departmentID == "" {
		pkg.PanicException(constant.InvalidRequest)
	}

	person, err := p.PersonRepository.FindPersonByID(personID)
	switch err {
	case nil:
		break
	case pkg.ErrNoRows:
		pkg.PanicException(constant.InvalidRequest)
	default:
		slog.Error("Error when fetching data from database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	if err := p.PersonRelRepository.RemoveDepartmentFromPerson(person, departmentID); err != nil {
		slog.Error("Error when saving data to database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}
	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null()))
}

/** Workplace */
func (p PersonRelServiceImpl) AddWorkplaceToPerson(c *gin.Context) {
	/* AddWorkplaceToPerson is a function to add workplace to a person
	 * @param c is gin context
	 * @return void
	 */
	defer pkg.PanicHandler(c)
	slog.Info("start to execute program add workplace to person")

	personID := c.Param("personID")
	if personID == "" {
		pkg.PanicException(constant.InvalidRequest)
	}
	personID = strings.ToUpper(personID)

	person, err := p.PersonRepository.FindPersonByID(personID)
	switch err {
	case nil:
		break
	case pkg.ErrNoRows:
		pkg.PanicException(constant.InvalidRequest)
	default:
		slog.Error("Error when fetching data from database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	var request dco.RelAddWorkplaceRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		pkg.PanicException(constant.InvalidRequest)
	}

	_, err = p.WorkplaceRepository.FindWorkplaceByID(request.DepartmentID, request.WorkplaceID)
	switch err {
	case nil:
		break
	case pkg.ErrNoRows:
		pkg.PanicException(constant.InvalidRequest)
	default:
		slog.Error("Error when fetching data from database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	if err := p.PersonRelRepository.AddWorkplaceToPerson(person, request.DepartmentID, request.WorkplaceID); err != nil {
		slog.Error("Error when saving data to database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null()))
}

func (p PersonRelServiceImpl) RemoveWorkplaceFromPerson(c *gin.Context) {
	/* RemoveWorkplaceFromPerson is a function to remove workplace from a person
	 * @param c is gin context
	 * @return void
	 */
	defer pkg.PanicHandler(c)
	slog.Info("start to execute program remove workplace from person")

	personID := c.Param("personID")
	if personID == "" {
		pkg.PanicException(constant.InvalidRequest)
	}
	personID = strings.ToUpper(personID)

	request := dco.RelRemoveWorkplaceRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		pkg.PanicException(constant.InvalidRequest)
	}

	person, err := p.PersonRepository.FindPersonByID(personID)
	switch err {
	case nil:
		break
	case pkg.ErrNoRows:
		pkg.PanicException(constant.InvalidRequest)
	default:
		slog.Error("Error when fetching data from database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	if err := p.PersonRelRepository.RemoveWorkplaceFromPerson(person, request.DepartmentID, request.WorkplaceID); err != nil {
		slog.Error("Error when saving data to database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null()))
}

/** Weekday */
func (p PersonRelServiceImpl) AddWeekdayToPerson(c *gin.Context) {
	/* AddWeekdayToPerson is a function to add weekday to a person
	 * @param c is gin context
	 * @return void
	 */
	defer pkg.PanicHandler(c)
	slog.Info("start to execute program add weekday to person")

	personID := c.Param("personID")
	if personID == "" {
		pkg.PanicException(constant.InvalidRequest)
	}
	personID = strings.ToUpper(personID)

	person, err := p.PersonRepository.FindPersonByID(personID)
	switch err {
	case nil:
		break
	case pkg.ErrNoRows:
		pkg.PanicException(constant.InvalidRequest)
	default:
		slog.Error("Error when fetching data from database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	request := dco.RelAddWeekdayRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		pkg.PanicException(constant.InvalidRequest)
	}
	if err := request.Validate(); err != nil {
		pkg.PanicException(constant.InvalidRequest)
	}

	if err := p.PersonRelRepository.AddWeekdayToPerson(person, request.WeekdayID); err != nil {
		slog.Error("Error when saving data to database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null()))
}

func (p PersonRelServiceImpl) RemoveWeekdayFromPerson(c *gin.Context) {
	/* RemoveWeekdayFromPerson is a function to remove weekday from a person
	 * @param c is gin context
	 * @return void
	 */
	defer pkg.PanicHandler(c)
	slog.Info("start to execute program remove weekday from person")

	personID := c.Param("personID")
	if personID == "" {
		pkg.PanicException(constant.InvalidRequest)
	}
	personID = strings.ToUpper(personID)

	weekdayID := c.Param("weekdayID")
	if weekdayID == "" {
		pkg.PanicException(constant.InvalidRequest)
	}
	if weekdayID != "MON" && weekdayID != "TUE" && weekdayID != "WED" && weekdayID != "THU" && weekdayID != "FRI" && weekdayID != "SAT" && weekdayID != "SUN" {
		pkg.PanicException(constant.InvalidRequest)
	}

	person, err := p.PersonRepository.FindPersonByID(personID)
	switch err {
	case nil:
		break
	case pkg.ErrNoRows:
		pkg.PanicException(constant.InvalidRequest)
	default:
		slog.Error("Error when fetching data from database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	if err := p.PersonRelRepository.RemoveWeekdayFromPerson(person, weekdayID); err != nil {
		slog.Error("Error when saving data to database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null()))
}

var personRelServiceSet = wire.NewSet(
	wire.Struct(new(PersonRelServiceImpl), "*"),
	wire.Bind(new(PersonRelService), new(PersonRelServiceImpl)),
)
