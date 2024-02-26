package repository

import (
	"context"
	"fmt"
	"planner-backend/app/domain/dao"
	"time"

	"github.com/docker/go-connections/nat"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func NewTestDBInstance(ctx context.Context) (*neo4j.DriverWithContext, error) {
	/**
	* Function to start a neo4j container for testing
	* This function creates an isolated neo4j database which can be accessed by the test cases
	* The function returns a neo4j driver which can be used to interact with the database
	**/
	req := testcontainers.ContainerRequest{
		Image: "neo4j:latest",
		Env: map[string]string{
			"NEO4J_AUTH": "neo4j/test",
			"NEO4J_dbms_security_auth__minimum__password__length": "1",
		},
		WaitingFor:   wait.ForLog("Started."),
		ExposedPorts: []string{"7687/tcp"},
	}

	neo4jContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, err
	}

	// Terminate the container when the context is done
	go func() {
		<-ctx.Done()

		// we need a new context here because the original one is already done
		if err := neo4jContainer.Terminate(context.Background()); err != nil {
			panic(err)
		}
	}()

	neo4jPort, err := neo4jContainer.MappedPort(ctx, nat.Port("7687"))
	if err != nil {
		return nil, err
	}

	dbSN := fmt.Sprintf("bolt://localhost:%s", neo4jPort.Port())

	db, err := neo4j.NewDriverWithContext(dbSN, neo4j.BasicAuth("neo4j", "test", ""))
	if err != nil {
		return nil, err
	}

	return &db, nil
}

type Creator interface {
	Create(db *neo4j.DriverWithContext, ctx context.Context) error
}

type DeparmentCreatorImpl struct {
	name string
	id   string
}

func (c *DeparmentCreatorImpl) Create(db *neo4j.DriverWithContext, ctx context.Context) error {
	d := DepartmentRepositoryImpl{
		db:  db,
		ctx: ctx,
	}

	if _, err := d.Save(&dao.Department{
		Name: c.name,
		ID:   c.id,
	}); err != nil {
		return err
	}

	return nil
}

type WorkplaceCreatorImpl struct {
	departmentID   string
	departmentName string
	name           string
	id             string
}

func (c *WorkplaceCreatorImpl) Create(db *neo4j.DriverWithContext, ctx context.Context) error {
	d := DepartmentRepositoryImpl{
		db:  db,
		ctx: ctx,
	}
	wp := WorkplaceRepositoryImpl{
		db:  db,
		ctx: ctx,
	}

	if _, err := d.Save(&dao.Department{
		Name: c.departmentName,
		ID:   c.departmentID,
	}); err != nil {
		return err
	}

	if _, err := wp.Save(
		c.departmentID,
		&dao.Workplace{
			Name: c.name,
			ID:   c.id,
		},
	); err != nil {
		return err
	}

	return nil
}

type TimeslotCreatorImpl struct {
	departmentID   string
	departmentName string
	workplaceID    string
	workplaceName  string
	id             string
	name           string

	weekdays []struct {
		id        string
		startTime time.Time
		endTime   time.Time
	}
}

func (c *TimeslotCreatorImpl) Create(db *neo4j.DriverWithContext, ctx context.Context) error {
	d := DepartmentRepositoryImpl{
		db:  db,
		ctx: ctx,
	}
	wp := WorkplaceRepositoryImpl{
		db:  db,
		ctx: ctx,
	}
	tr := TimeslotRepositoryImpl{
		db:  db,
		ctx: ctx,
	}
	wd := WeekdayRepositoryImpl{
		db:  db,
		ctx: ctx,
	}

	if _, err := d.Save(&dao.Department{
		ID:   c.departmentID,
		Name: c.departmentName,
	}); err != nil {
		return err
	}

	if _, err := wp.Save(
		c.departmentID,
		&dao.Workplace{
			ID:   c.workplaceID,
			Name: c.workplaceName,
		},
	); err != nil {
		return err
	}

	if _, err := tr.Save(c.departmentID, c.workplaceID, &dao.Timeslot{
		ID:   c.id,
		Name: c.name,
	}); err != nil {
		return err
	}

	for _, weekday := range c.weekdays {
		if _, err := wd.AddWeekdayToTimeslot(
			&dao.Timeslot{
				ID:           c.id,
				DepartmentID: c.departmentID,
				WorkplaceID:  c.workplaceID,
			}, &dao.OnWeekday{
				ID:        weekday.id,
				StartTime: weekday.startTime,
				EndTime:   weekday.endTime,
			}); err != nil {
			return err
		}
	}

	return nil
}

type PersonCreatorImpl struct {
	departments []struct {
		id   string
		name string
	}
	workplaces []struct {
		id           string
		name         string
		departmentID string
	}
	weekdayIDs []string
	person     struct {
		id           string
		email        string
		active       bool
		lastName     string
		firstName    string
		workingHours float64
	}
}

func (c *PersonCreatorImpl) Create(db *neo4j.DriverWithContext, ctx context.Context) error {
	d := DepartmentRepositoryImpl{
		db:  db,
		ctx: ctx,
	}
	wp := WorkplaceRepositoryImpl{
		db:  db,
		ctx: ctx,
	}
	pr := PersonRepositoryImpl{
		db:  db,
		ctx: ctx,
	}
	prl := PersonRelRepositoryImpl{
		db:  db,
		ctx: ctx,
	}

	for _, department := range c.departments {
		if _, err := d.Save(&dao.Department{
			ID:   department.id,
			Name: department.name,
		}); err != nil {
			return err
		}
	}

	for _, workplace := range c.workplaces {
		if _, err := wp.Save(
			workplace.departmentID,
			&dao.Workplace{
				ID:   workplace.id,
				Name: workplace.name,
			},
		); err != nil {
			return err
		}
	}

	person := &dao.Person{
		ID:           c.person.id,
		FirstName:    c.person.firstName,
		LastName:     c.person.lastName,
		Email:        c.person.email,
		Active:       c.person.active,
		WorkingHours: c.person.workingHours,
	}
	if _, err := pr.Save(person); err != nil {
		return err
	}

	for _, weekdayID := range c.weekdayIDs {
		if err := prl.AddWeekdayToPerson(*person, weekdayID); err != nil {
			return err
		}
	}

	for _, department := range c.departments {
		if err := prl.AddDepartmentToPerson(*person, department.id); err != nil {
			return err
		}
	}

	for _, workplace := range c.workplaces {
		if err := prl.AddWorkplaceToPerson(*person, workplace.departmentID, workplace.id); err != nil {
			return err
		}
	}

	return nil
}
