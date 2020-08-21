"""
An app to feed the cats
"""
import toga
import requests
import json
from toga.style import Pack
from toga.style.pack import COLUMN, ROW

BASE_URL = "http://localhost:6969"
CREATE_TIME = BASE_URL + "/feedingTime"
RECIEVE_A_TIME = BASE_URL + "/feedingTime"
RECIEVE_ALL_TIMES = BASE_URL + "/feedingTimes"

feeding_times = []

class CatFeeder(toga.App):
    # TODO: Come up with better variable names
    def startup(self):
        main_box = toga.Box()

        time_label = toga.Label('Enter feeding time:', style=Pack(padding=10, padding_top=20))

        # TODO: Check if input is valid, on_change=self.validate_input??
        self.hour_input = toga.TextInput(placeholder='Hour', style=Pack(padding=10))
        self.min_input = toga.TextInput(placeholder='Minute', style=Pack(padding=10))

        add_to_table_butt = toga.Button('Add time', on_press=self.addTime)
        send_butt = toga.Button('Send time', on_press=self.sendFeedingTime)
        self.time_table = toga.Table(
            headings=['Hour', 'Minute'],
            multiple_select=False
        )

        main_box.add(time_label)
        main_box.add(self.hour_input)
        main_box.add(self.min_input)
        main_box.add(add_to_table_butt)
        main_box.add(send_butt)
        main_box.add(self.time_table)

        self.main_window = toga.MainWindow(title=self.formal_name)
        self.main_window.content = main_box
        self.main_window.show()

    ###
    # @brief:   Add a new time to the table
    ###
    def addTime(self, widget):
        if self.hour_input.value == "":
            print("Hour time is missing")
            return
        if self.min_input.value == "":
            print("Minute time is missing")
            return

        feeding_times.append((self.hour_input.value,self.min_input.value))
        self.time_table.data.insert(0, *feeding_times[len(feeding_times) - 1])
        print("New Time:", feeding_times[len(feeding_times) - 1])

    ###
    # @brief:   Arrange the feeding time payload
    ###
    def sendFeedingTime(self, widget):
        if len(feeding_times) == 0:
            print("No times given")
            return

        send_feeding_times = []
        for ft in range(len(feeding_times)):
            data = {
                    "id": str(ft),
                    "hour": int(feeding_times[ft][0]),
                    "minute": int(feeding_times[ft][1])
                   }

            send_feeding_times.append(data)

        # TODO: Check for errors in the POST method
        req = requests.post(
            url = CREATE_TIME,
            data = str(json.dumps(send_feeding_times))
        )
        if req.status_code != 200:
            print("Uh oh! Looks like you are in trouble...", req.status_code)
            return

        resp = req.text
        print("Sent feeding time:", resp)

def main():
    return CatFeeder()
