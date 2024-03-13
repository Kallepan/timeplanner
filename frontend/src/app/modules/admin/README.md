# Admin Module

This module is responsible for managing the underlying dabatase and data of the application. It is only accessible to the admin user.

## Funtionalities

### Modify the schema

As the schema is used to generate the plan for a given number of days in advance, it is important to be able to modify it. The schema is composed of the following entities:

- Select a department, modify its associated workplace: e.g.: Add workplace, remove workplace, modify workplace
- Select a workplace and modify its associated timeslots: e.g.: Add timeslot, remove timeslot, modify timeslot
- Select a timeslot and modify the dates it is offered on.

These funtionalities are supposed to be offered one after another. E.g.: First select a department, then select a workplace, then select a timeslot and finally you can modify the dates it is offered on.

At last, a kind of confirmation is necessary as well as a date after which the changes take effect. The pre-generated schema after that date (including) are deleted in the database and the regular synchronization job is triggered (not sure about this maybe dont trigger it). Afterwards after the given date, only the new schema should be returned by workday queries. With confirmation field where the user has to type in the word "confirm" to confirm the changes.

### Modify the workdays

The workdays themselves should also be abled to be queried and turned on or off. Here we need to query all workdays on a given date and the user needs to be able to turn the ones he does not need to be off.

This should set the workdays to be inactive in the database (field active = false). Still during the query itself all workdays should be returned (active and inactive ones).

### Modify the persons of a Department

The admin should be able to add or remove persons from a department. Removal marks the person as inactive in the database (field deleted_at = current_timestamp). The person should not be able to login anymore and should not be able to be assigned to a workplace anymore.

Here we also need a confirmation field where the user has to type in the word "confirm" to confirm the changes. Furthermore persons which are inactive should be returned in this module as well. So that the admin can reactivate them if necessary.

Test Cases:

- The admin is able to modify the schema
- Test the schema modification process i.e. do dates after X get deleted and the regular synchronization job is triggered
- The admin is able to modify the workdays themselves
- The admin is able to modify the persons of a department
  - Test the reactivation of a person
  - Test the deactivation of a person
- Test the confirmation field
