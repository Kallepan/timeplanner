package service

import (
	"log/slog"
	"net/http"
	"planner-backend/app/constant"
	"planner-backend/app/domain/dao"
	"planner-backend/app/domain/dco"
	"planner-backend/app/pkg"
	"planner-backend/app/repository"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

type PersonService interface {
	// Function Used by the controller
	GetAllPersons(c *gin.Context)
	GetPersonByID(c *gin.Context)
	AddPerson(c *gin.Context)
	UpdatePerson(c *gin.Context)
	DeletePerson(c *gin.Context)
}

type PersonServiceImpl struct {
	PersonRepository repository.PersonRepository
}

func (p PersonServiceImpl) GetAllPersons(c *gin.Context) {
	/* GetAllPersons is a function to get all persons
	 * @param c is gin context
	 * @return void
	 */

	defer pkg.PanicHandler(c)
	slog.Info("start to execute program get all persons")

	departmentID := c.Query("department")
	if departmentID == "" {
		pkg.PanicException(constant.InvalidRequest)
	}

	rawData, err := p.PersonRepository.FindAllPersons(departmentID)
	switch err {
	case nil:
		break
	case pkg.ErrNoRows:
		break
	default:
		slog.Error("Error when fetching data from database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	data := mapPersonListToPersonResponseList(rawData)

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (p PersonServiceImpl) GetPersonByID(c *gin.Context) {
	/* GetPersonByID is a function to get person by id
	 * @param c is gin context
	 * @return void
	 */

	defer pkg.PanicHandler(c)
	slog.Info("start to execute program get person by id")

	personID := c.Param("personID")
	if personID == "" {
		pkg.PanicException(constant.InvalidRequest)
	}

	rawData, err := p.PersonRepository.FindPersonByID(personID)
	switch err {
	case nil:
		break
	case pkg.ErrNoRows:
		pkg.PanicException(constant.DataNotFound)
	default:
		slog.Error("Error when fetching data from database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	data := mapPersonToPersonResponse(rawData)

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (p PersonServiceImpl) AddPerson(c *gin.Context) {
	/* AddPerson is a function to add person
	 * @param c is gin context
	 * @return void
	 */

	defer pkg.PanicHandler(c)
	slog.Info("start to execute program add person")

	var personRequest dco.PersonRequest
	if err := c.ShouldBindJSON(&personRequest); err != nil {
		slog.Error("Error when binding json", "error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	_, err := p.PersonRepository.FindPersonByID(personRequest.ID)
	switch err {
	case nil:
		pkg.PanicException(constant.Conflict)
	case pkg.ErrNoRows:
		break
	default:
		slog.Error("Error when fetching data from database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	person := mapPersonRequestToPerson(personRequest)
	rawData, err := p.PersonRepository.Save(&person)
	if err != nil {
		slog.Error("Error when saving data to database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	data := mapPersonToPersonResponse(rawData)

	c.JSON(http.StatusCreated, pkg.BuildResponse(constant.Success, data))
}

func (p PersonServiceImpl) UpdatePerson(c *gin.Context) {
	/* UpdatePerson is a function to update person
	 * @param c is gin context
	 * @return void
	 */

	defer pkg.PanicHandler(c)
	slog.Info("start to execute program update person")

	personID := c.Param("personID")
	if personID == "" {
		pkg.PanicException(constant.InvalidRequest)
	}

	person, err := p.PersonRepository.FindPersonByID(personID)
	switch err {
	case nil:
		break
	case pkg.ErrNoRows:
		pkg.PanicException(constant.DataNotFound)
	default:
		slog.Error("Error when fetching data from database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	var personRequest dco.PersonRequest
	if err := c.ShouldBindJSON(&personRequest); err != nil {
		pkg.PanicException(constant.InvalidRequest)
	}

	person.FirstName = personRequest.FirstName
	person.LastName = personRequest.LastName
	person.Email = personRequest.Email
	person.Active = *personRequest.Active
	person.WorkingHours = personRequest.WorkingHours

	rawData, err := p.PersonRepository.Save(&person)
	if err != nil {
		slog.Error("Error when saving data to database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	data := mapPersonToPersonResponse(rawData)

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (p PersonServiceImpl) DeletePerson(c *gin.Context) {
	/* DeletePerson is a function to delete person
	 * @param c is gin context
	 * @return void
	 */

	defer pkg.PanicHandler(c)
	slog.Info("start to execute program delete person")

	personID := c.Param("personID")
	if personID == "" {
		pkg.PanicException(constant.InvalidRequest)
	}

	person, err := p.PersonRepository.FindPersonByID(personID)
	switch err {
	case nil:
		break
	case pkg.ErrNoRows:
		pkg.PanicException(constant.DataNotFound)
	default:
		slog.Error("Error when fetching data from database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	err = p.PersonRepository.Delete(&person)
	if err != nil {
		slog.Error("Error when deleting data from database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null()))
}

func mapPersonToPersonResponse(person dao.Person) dco.PersonResponse {
	/* mapPersonToPersonResponse is a function to map person to person response
	 * @param person is dao.Person
	 * @return dco.PersonResponse
	 */

	return dco.PersonResponse{
		ID:           person.ID,
		FirstName:    person.FirstName,
		LastName:     person.LastName,
		Email:        person.Email,
		Active:       &person.Active,
		WorkingHours: person.WorkingHours,

		Base: dco.Base{
			CreatedAt: person.CreatedAt,
			UpdatedAt: person.UpdatedAt,
			DeletedAt: person.DeletedAt,
		},

		Workplaces:  mapWorkplaceInPersonToWorkplaceInPersonResponseList(person.Workplaces),
		Departments: mapDepartmentInPersonToDepartmentInPersonResponseList(person.Departments),
		Weekdays:    mapWeekdayListToWeekdayResponseList(person.Weekdays),
	}
}

func mapDepartmentInPersonToDepartmentInPersonResponseList(departments []dao.DepartmentInPerson) []dco.DepartmentInPersonResponse {
	/* mapDepartmentInPersonToDepartmentInPersonResponseList is a function to map department in person to department in person response list
	 * @param departments is []dao.DepartmentInPerson
	 * @return []dco.DepartmentInPersonResponse
	 */

	var departmentInPersonResponseList []dco.DepartmentInPersonResponse
	for _, department := range departments {
		departmentInPersonResponseList = append(departmentInPersonResponseList, dco.DepartmentInPersonResponse{
			ID:   department.ID,
			Name: department.Name,
		})
	}

	return departmentInPersonResponseList
}

func mapWorkplaceInPersonToWorkplaceInPersonResponseList(workplaces []dao.WorkplaceInPerson) []dco.WorkplaceInPersonResponse {
	/* mapWorkplaceInPersonToWorkplaceInPersonResponseList is a function to map workplace in person to workplace in person response list
	 * @param workplaces is []dao.WorkplaceInPerson
	 * @return []dco.WorkplaceInPersonResponse
	 */

	var workplaceInPersonResponseList []dco.WorkplaceInPersonResponse
	for _, workplace := range workplaces {
		workplaceInPersonResponseList = append(workplaceInPersonResponseList, dco.WorkplaceInPersonResponse{
			ID:           workplace.ID,
			Name:         workplace.Name,
			DepartmentID: workplace.DepartmentID,
		})
	}

	return workplaceInPersonResponseList
}

func mapWeekdayListToWeekdayResponseList(weekdays []dao.Weekday) []dco.WeekdayResponse {
	/* mapWeekdayListToWeekdayResponseList is a function to map weekday list to weekday response list
	 * @param weekdays is []dao.Weekday
	 * @return []dco.WeekdayResponse
	 */

	var weekdayResponseList []dco.WeekdayResponse
	for _, weekday := range weekdays {
		weekdayResponseList = append(weekdayResponseList, dco.WeekdayResponse{
			ID:   weekday.ID,
			Name: weekday.Name,
		})
	}

	return weekdayResponseList
}

func mapPersonListToPersonResponseList(persons []dao.Person) []dco.PersonResponse {
	/* mapPersonListToPersonResponseList is a function to map person list to person response list
	 * @param persons is []dao.Person
	 * @return []dco.PersonResponse
	 */

	var personResponseList []dco.PersonResponse
	for _, person := range persons {
		personResponseList = append(personResponseList, mapPersonToPersonResponse(person))
	}

	return personResponseList
}

func mapPersonRequestToPerson(personRequest dco.PersonRequest) dao.Person {
	/* mapPersonRequestToPerson is a function to map person request to person
	 * @param personRequest is dco.PersonRequest
	 * @return dao.Person
	 */

	var active bool
	if personRequest.Active != nil {
		active = *personRequest.Active
	}

	return dao.Person{
		ID:           personRequest.ID,
		FirstName:    personRequest.FirstName,
		LastName:     personRequest.LastName,
		Email:        personRequest.Email,
		Active:       active,
		WorkingHours: personRequest.WorkingHours,
	}
}

var personServiceSet = wire.NewSet(
	wire.Struct(new(PersonServiceImpl), "*"),
	wire.Bind(new(PersonService), new(*PersonServiceImpl)),
)
