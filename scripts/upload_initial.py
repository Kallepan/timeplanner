import csv
import logging
import json
import os
import requests

logging.basicConfig(
    level=logging.INFO,
    format="%(asctime)s - %(levelname)s - %(message)s",
    handlers=[logging.StreamHandler()],
)

# This is the URL of the API
URL = "http://localhost:8080/api/v1/planner"

# These are the paths to the data files
DEPARTMENTS = os.path.join(os.path.dirname(__file__), "../data/departments.csv")
WORKPLACES = os.path.join(os.path.dirname(__file__), "../data/workplaces.csv")
TIMESLOTS = os.path.join(os.path.dirname(__file__), "../data/timeslots.csv")

# This is due to the complexity of the data structure
OFFERED_ON = os.path.join(os.path.dirname(__file__), "../data/offered_on.json")


def create_departments() -> None:
    with open(DEPARTMENTS, "r", encoding="utf-8-sig") as f:
        reader = csv.DictReader(f)
        departments = list(reader)

        for i, department in enumerate(departments):
            logging.info(
                f"Uploading department {i+1}/{len(departments)}: {department['id']}"
            )

            department_id = department["id"]
            department_name = department["name"]

            data = {"id": department_id, "name": department_name}

            req = requests.post(
                URL + "/department", json={"id": department_id, "name": department_name}
            )

            if req.status_code != 201:
                logging.info(
                    f"Error uploading department {department_id}: {req.text}. Status: {req.status_code}"
                )
                continue


def create_workplaces() -> None:
    with open(WORKPLACES, "r", encoding="utf-8-sig") as f:
        reader = csv.DictReader(f)
        workplaces = list(reader)

        for i, workplace in enumerate(workplaces):
            logging.info(
                f"Uploading workplace {i+1}/{len(workplaces)}: {workplace['id']}"
            )

            workplace_id = workplace["id"]
            workplace_name = workplace["name"]
            department_id = workplace["department_id"]

            data = {
                "id": workplace_id,
                "name": workplace_name,
            }

            req = requests.post(
                f"{URL}/department/{department_id}/workplace",
                json={"id": workplace_id, "name": workplace_name},
            )

            if req.status_code != 201:
                logging.info(
                    f"Error uploading workplace {workplace_id}: {req.text}. Status: {req.status_code}"
                )
                continue


def create_timeslots() -> None:
    with open(TIMESLOTS, "r", encoding="utf-8-sig") as f:
        reader = csv.DictReader(f)
        timeslots = list(reader)

        for i, timeslot in enumerate(timeslots):
            logging.info(
                f"Uploading timeslot {i+1}/{len(timeslots)}: {timeslot['name']}"
            )

            timeslot_name = timeslot["name"]
            department_id = timeslot["department_id"]
            workplace_id = timeslot["workplace_id"]

            data = {
                "name": timeslot_name,
                "active": True,
            }

            req = requests.post(
                f"{URL}/department/{department_id}/workplace/{workplace_id}/timeslot",
                json=data,
            )

            if req.status_code != 201:
                logging.info(
                    f"Error uploading timeslot {timeslot_name}: {req.text}. Status: {req.status_code}"
                )
                continue


def create_timeslot_offered_on() -> None:
    # load json file
    with open(OFFERED_ON, "r") as f:
        offered_on = json.load(f)

        for i, timeslot in enumerate(offered_on):
            logging.info(
                f"Uploading timeslot offered on {i+1}/{len(offered_on)}: {timeslot['timeslot_name']}"
            )

            timeslot_name = timeslot["timeslot_name"]
            department_id = timeslot["department_id"]
            workplace_id = timeslot["workplace_id"]
            data = timeslot["data"]

            req = requests.post(
                f"{URL}/department/{department_id}/workplace/{workplace_id}/timeslot/{timeslot_name}/weekday/bulk",
                json=data,
            )

            if req.status_code != 201:
                logging.info(
                    f"Error uploading timeslot offered on {timeslot_name}: {req.text}. Status: {req.status_code}"
                )
                continue


def main():
    create_departments()
    create_workplaces()
    create_timeslots()
    create_timeslot_offered_on()


if __name__ == "__main__":
    main()
