"""
An app to feed the cats
"""
import json
import requests
import toga
import utils
import TSLAdapter
from toga.style import Pack
from toga.style.pack import COLUMN, ROW

BASE_URL = "https://canescent-saola-6329.dataplicity.io"
CREATE_TIME = BASE_URL + "/feedingTime"
RECIEVE_A_TIME = BASE_URL + "/feedingTime"
RECIEVE_ALL_TIMES = BASE_URL + "/feedingTimes"

feeding_times = []
logger = utils.Applogger('cat-app')
cat_session = requests.Session()
cat_session.mount(BASE_URL, TLSAdapter.Tls12HttpAdapter())

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
            self.error_label.text = "Missing time for hour"
            return
        if self.min_input.value == "":
            self.error_label.text = "Missing time for minute"
            return

        self.send_butt.enabled = True
        feeding_times.append((self.hour_input.value,self.min_input.value))
        time = utils.PrettyTime(feeding_times[len(feeding_times) - 1])
        self.time_table.data.insert(0, time)
        logger.info("New Time: " + str(time))
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
            self.error_label.text = "No feeding times present"
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
            req = cat_session.post(
                url = CREATE_TIME,
                data = str(json.dumps(send_feeding_times))
            )
        except OSError as err:
            logger.error("Uh oh! Looks like trouble found you: error(" + utils.StrOSError(err) + ")")
            self.error_label.text = "Error on communicating with machine: error(" + utils.StrOSError(err) + ")"
            return

        resp = req.text
        logger.info("Sent feeding time:" + str(resp))
        self.error_label.text = "Feeding times sent"

    ###
    # @brief:   Get any existing feeding time payload
    ###
    def getFeedingTimes(self, widget):
        try:
            req = cat_session.get(url = RECIEVE_ALL_TIMES)
            if req.json() is None:
                logger.error("Uh oh! Received invalid feeding times: Feeding-Time(" + str(req) + ")")
                self.error_label.text = "Received invalid feeding times:", str(req)
                return
            if req.status_code != 200:
                logger.error("Uh oh! Request unsuccessful: HTTP(" + str(req.status_code) + ")")
                self.error_label.text = "Failed to receive request: HTTP(" + str(req.status_code) + ")"
                return

        except OSError as err:
            logger.error("Uh oh! Looks like trouble found you: error(" + utils.StrOSError(err) + ")")
            self.error_label.text = "Error on communicating with machine: error(" + utils.StrOSError(err) + ")"
            return

        feeding_times = []
        self.time_table.data.clear()
        resp = req.json()
        for i in range(len(resp)):
            feeding_times.append((resp[i]['hour'],resp[i]['minute']))
            time = utils.PrettyTime(feeding_times[len(feeding_times) - 1])
            self.time_table.data.insert(i, time)
        logger.info("Received feeding times:" + str(feeding_times))
        self.send_butt.enabled = True
        self.error_label.text = "Received feeding times"

def main():
    return CatFeeder()
