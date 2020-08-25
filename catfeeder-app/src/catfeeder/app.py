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
        time_box = toga.Box()
        butt_box = toga.Box()
        table_box = toga.Box()
        info_box = toga.Box()

        time_label = toga.Label('Enter feeding time:', style=Pack(padding=10, padding_top=20))

        # TODO: Check if input is valid, on_change=self.validate_input??
        self.hour_input = toga.TextInput(placeholder='Hour', style=Pack(padding=10))
        self.min_input = toga.TextInput(placeholder='Minute', style=Pack(padding=10))
        time_box.add(time_label)
        time_box.add(self.hour_input)
        time_box.add(self.min_input)

        add_to_table_butt = toga.Button('Add time', on_press=self.addTime, style=Pack(padding=10))
        clear_table_butt = toga.Button('Clear times', on_press=self.clearTable, style=Pack(padding=10))
        self.send_butt = toga.Button('Send time', on_press=self.sendFeedingTime, style=Pack(padding=10))
        get_butt = toga.Button('Get times', on_press=self.getFeedingTimes, style=Pack(padding=10))
        self.send_butt.enabled = False
        butt_box.add(add_to_table_butt)
        butt_box.add(clear_table_butt)
        butt_box.add(self.send_butt)
        butt_box.add(get_butt)

        self.time_table = toga.Table(
            headings=['Feeding Times'],
            multiple_select=False
        )
        table_box.add(self.time_table)

        self.error_label = toga.Label('', style=Pack(padding=10, padding_top=20))
        info_box.add(self.error_label)

        main_box.add(time_box)
        main_box.add(butt_box)
        main_box.add(table_box)
        main_box.add(info_box)

        main_box.style.update(direction=COLUMN, padding_top=10)
        time_box.style.update(direction=ROW, padding=5)
        butt_box.style.update(direction=ROW, padding=5)
        table_box.style.update(direction=ROW, padding=5)
        info_box.style.update(direction=ROW, padding=5)

        self.main_window = toga.MainWindow(title=self.formal_name)
        self.main_window.content = main_box
        self.main_window.show()

    ###
    # @brief:   Add a new feeding time to the table
    ###
    def addTime(self, widget):
        if self.hour_input.value == "":
            print("Hour time is missing")
            return
        if self.min_input.value == "":
            print("Minute time is missing")
            return

        self.send_butt.enabled = True
        feeding_times.append((self.hour_input.value,self.min_input.value))
        time = self.prettyTime(feeding_times[len(feeding_times) - 1][0], feeding_times[len(feeding_times) - 1][1])
        self.time_table.data.insert(0, time)
        print("New Time:", feeding_times[len(feeding_times) - 1])
        self.error_label.text = ""

    ###
    # @brief:   Clear the table of the existing feeding time(s)
    ###
    def clearTable(self, widget):
        feeding_times = []
        self.time_table.data.clear()
        self.send_butt.enabled = False

    ###
    # @brief:   Send the feeding time payload
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

        try:
            req = requests.post(
                url = CREATE_TIME,
                data = str(json.dumps(send_feeding_times))
            )
        except OSError as err:
            print("OSError: Uh oh! Looks like you are in trouble...", err)
            self.error_label.text = "Error on communicating with machine"
            return

        resp = req.text
        print("Sent feeding time:", resp)
        self.error_label.text = ""

    ###
    # @brief:   Get any existing feeding time payload
    ###
    def getFeedingTimes(self, widget):
        feeding_times = []
        self.time_table.data.clear()

        try:
            req = requests.get(url = RECIEVE_ALL_TIMES)
            if req == None:
                print("Uh oh! Things not looking good")
                self.error_label.text = "Received invalid feeding times:", req
                return
            if req.status_code != 200:
                print("Uh oh! Not success:", req.status_code)
                self.error_label.text = "Failed to recieve request:" + str(req.status_code)
                return

        except OSError as err:
            print("OSError: Uh oh! Looks like you are in trouble...", err)
            self.error_label.text = "Error on communicating with machine"
            return

        resp = req.json()
        for i in range(len(resp)):
            feeding_times.append((resp[i]['hour'],resp[i]['minute']))
            time = self.prettyTime(feeding_times[len(feeding_times) - 1][0], feeding_times[len(feeding_times) - 1][1])
            self.time_table.data.insert(i, time)
        print("Recieved feeding times:", feeding_times)
        self.send_butt.enabled = True
        self.error_label.text = ""

    def prettyTime(self, hour, minute):
        return str(hour) + ":" + (str(minute) if int(minute) > 9 else "0" + str(minute))

def main():
    return CatFeeder()
