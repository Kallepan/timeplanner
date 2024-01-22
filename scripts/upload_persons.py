"""
Scripts to upload persons into the database by using the backend API
"""

import os
import csv
import requests
import logging

URL = "http://localhost:8080/api/v1/planner/person"
FILE_PATH = os.path.join(os.path.dirname(__file__), "../tests/persons.csv")

logging.basicConfig(
    level=logging.INFO,
    format="%(asctime)s - %(levelname)s - %(message)s",
    handlers=[logging.StreamHandler()],
)


def main() -> None:
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

            response = requests.post(URL, json=data)
            if response.status_code != 201:
                logging.error(f"Error uploading person {id}: {response.text}")
                return

            # upload weekdays
            weekdays = person["present_weekdays"]
            if weekdays == "":
                continue

            weekday_request_url = f"{URL}/{id.lower()}/weekday"
            for weekday_id in weekdays.split(","):
                weekday_data = {"weekday_id": weekday_id}
                weekday_response = requests.post(weekday_request_url, json=weekday_data)
                if weekday_response.status_code != 201:
                    logging.info(
                        f"Error uploading weekday {weekday_id} for person {id}: {weekday_response.text}. Status: {weekday_response.status_code}"
                    )
                    continue

    logging.info("Done")


if __name__ == "__main__":
    main()
