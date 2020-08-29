package main

import (
    "io/ioutil"
    "encoding/json"
    "os"
)

type CommonNames struct {
    Common_Names    string `json:"log_filepath"`
    FQDN            string `json:"FQDN"`
    PORT            string `json:"PORT"`
}

var common_names CommonNames

/***
 * @brief:  Read the common-names.json file. This holds values
 *          that both the app and machine use
 *
 * @return: nil on success, else error
 ***/
func InitCommon() error {
    common_names_os, err := os.Open("/usr/local/share/.cat-feeder/common-names.json")
    if err != nil {
        return err
    }
    defer common_names_os.Close()

    common_names_io, _ := ioutil.ReadAll(common_names_os)
    json.Unmarshal(common_names_io, &common_names)

    return err
}

/***
 * @brief:  Return the location of the log file
 *
 * @return: See @brief
 ***/
func GetLogFilepath() string {
    return common_names.Common_Names
}

/***
 * @brief:  Return the Fully Qualified Domain Name
 *
 * @return: See @brief
 ***/
func GetFQDN() string {
    return common_names.FQDN
}

/***
 * @brief:  Return the port that the server is listening on
 *
 * @return: See @brief
 ***/
func GetPORT() string {
    return common_names.PORT
}
