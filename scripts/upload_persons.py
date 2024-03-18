"""
Scripts to upload persons into the database by using the backend API
"""

import os
import csv
import requests
import logging

BASE_URL = "http://localhost:8080"
URL = f"{BASE_URL}/api/v1/planner"
FILE_PATH = os.path.join(os.path.dirname(__file__), "../data/persons.csv")

logging.basicConfig(
    level=logging.INFO,
    format="%(asctime)s - %(levelname)s - %(message)s",
    handlers=[logging.StreamHandler()],
)


def create_person_relationships(
    department_id: str, workplace_ids: list[str], person_id: str
) -> None:
    department_request_url = f"{URL}/person/{person_id}/department"
    department_data = {"department_id": department_id}
    department_response = requests.post(department_request_url, json=department_data)
    if department_response.status_code != 201:
        logging.info(
            f"Error uploading department {department_id} for person {person_id}: {department_response.text}. Status: {department_response.status_code}"
        )
    # upload workplaces
    for workplace_id in workplace_ids:
        workplace_request_url = f"{URL}/person/{person_id}/workplace"
        request_data = {
            "department_id": department_id,
            "workplace_id": workplace_id,
        }

        workplace_response = requests.post(workplace_request_url, json=request_data)

        if workplace_response.status_code != 201:
            logging.info(
                f"Error uploading workplace {workplace_id} for person {person_id}: {workplace_response.text}. Status: {workplace_response.status_code}"
            )
            continue


def create_persons(department_id: str, workplace_ids: list[str]) -> None:
    # load file
    with open(FILE_PATH, "r", encoding="utf-8-sig") as f:
        reader = csv.DictReader(f)
        persons = list(reader)

        for i, person in enumerate(persons):
            logging.info(f"Uploading person {i+1}/{len(persons)}: {person['id']}")

            first_name = person["first_name"]
            last_name = person["last_name"]
            email = person["email"]
            working_hours = person["working_hours"]
            id = person["id"]

            data = {
                "first_name": first_name,
                "last_name": last_name,
                "email": email,
                "working_hours": float(working_hours),
                "id": id.lower(),
                "active": True,
            }

            response = requests.post(f"{URL}/person", json=data)
            if response.status_code != 201 and response.status_code != 409:
                logging.error(f"Error uploading person {id}: {response.text}")
                continue

            # upload department and workplaces
            create_person_relationships(department_id, workplace_ids, id.lower())

            # upload weekdays
            weekdays = person["present_weekdays"]
            if weekdays == "":
                continue

            weekday_request_url = f"{URL}/person/{id.lower()}/weekday"
            for weekday_id in weekdays.split(","):
                weekday_data = {"weekday_id": int(weekday_id)}
                weekday_response = requests.post(weekday_request_url, json=weekday_data)
                if (
                    weekday_response.status_code != 201
                    and weekday_response.status_code != 409
                ):
                    logging.info(
                        f"Error uploading weekday {weekday_id} for person {id}: {weekday_response.text}. Status: {weekday_response.status_code}"
                    )
                    continue

    logging.info("Done")


def main() -> None:
    department_id = "bak"
    workplaces = [
        "psl",
        "var",
        "anl",
        "arz",
        "stu",
        "ate",
        "mal",
        "vit",
        "tbl",
        "hyg",
        "spa",
        "cha",
        "jok",
        "azu",
        "not",
    ]
    create_persons(
        department_id=department_id,
        workplace_ids=workplaces,
    )


if __name__ == "__main__":
    main()
