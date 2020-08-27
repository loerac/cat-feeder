import json

class CommonLibs():
    ###
    # @brief:   Parse the common-names.json file
    ###
    def __init__(self):
        self.common_names = {}

        with open('/usr/local/share/.cat-feeder/common-names.json') as f:
            self.common_names = json.load(f)

    ###
    # @brief:   Return the path for the logfile
    ###
    def GetLogFilepath(self):
        return self.common_names['log_filepath']
