import logging

###
# @brief:   Create new logging object with the identity of `name`
#
# @arg:     name - The name to give the identity of the log
#
# @return:  Logger object
###
def Applogger(name):
    logging.basicConfig(level=logging.INFO,
                        filename='/tmp/cat-feeder.log',
                        format='%(asctime)s %(name)s - %(message)s',
                        datefmt='%Y/%m/%d %H:%M:%S'
                       )

    logger = logging.getLogger(name)

    return logger

###
# @brief:   Turn the feeding time input to the time format
#
# @arg:     feeding_time - Array holding the hour and minute
#
# @return:  String formated time
###
def PrettyTime(feeding_times):
    hour = feeding_times[0]
    minute = feeding_times[1]

    return str(hour) + ":" + (str(minute) if int(minute) > 9 else "0" + str(minute))

###
# @brief:   Parse the OS error to return the error message
#
# @arg:     err - OSError from try-except
#
# @return:  String formatted error message
###
def StrOSError(err):
    return str(err.args[0].reason)[str(err.args[0].reason).find(":") + 2:]
