import logging

def Applogger(name):
    logging.basicConfig(level=logging.WARNING,
                        filename='/tmp/cat-feeder.log',
                        format='%(asctime)s %(name)s - %(message)s',
                        datefmt='%Y/%m/%d %H:%M:%S'
                       )

    logger = logging.getLogger(name)

    return logger

def PrettyTime(feeding_times):
    hour = feeding_times[0]
    minute = feeding_times[1]

    return str(hour) + ":" + (str(minute) if int(minute) > 9 else "0" + str(minute))

def StrError(err):
    return str(err.args[0].reason)[str(err.args[0].reason).find(":") + 2:]
