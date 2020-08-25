# All
- [ ] loera - 007: Create a logger for both the machine and app
  * [ ] It should rollover any existing logs and then create a new log
    > Currently have the machine logs completed, still working on the app portion
- [ ] loera - 008: Update the configuration to install what is needed and where executables/logs should be stored in

# Machine
- [x] loera - 001: Monitor the time 
  * [x] Only feed cat specified time, then keep feeder closed for the rest of the time
  * [x] Make sure to lock/unlock `[]feeder_times` since it's used in both monitoring and REST API thread
- [ ] loera - 002: Create REST API
  * [x] POST new feeding times
  * [x] GET all/single feeding times
  * [ ] POST to feed cat
- [ ] loera - 003: Get them servo motors working
  * [ ] Get the motor to open and close
  * [ ] Add it to the monitoring
- [ ] loera - 005: Communication from machine and app
  * Over WiFi?
  * Bluetooth?
  * Host app on a webserver and let the user navigate to webpage?
- [ ] 009: Update `/feedNow` REST API to use the servo motors

# App
- [ ] 004: Continue development: 
  * [ ] Delete time in table
  * [ ] Edit time in table
  * [ ] Check if time is valid
  * [ ] Allow 24 hour and 12 hour (AM/PM)
    > Currently, only has 24 hour
- [ ] 006: Add button to feed
  * Send command to feed cat
  * Tell Christian to finish up the REST API to start this
- [ ] 010: Display last time fed on a label
