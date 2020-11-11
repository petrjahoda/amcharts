package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"time"
)

type DowntimeUsersInput struct {
	Start        int64
	End          int64
	DowntimeName string
}

type DowntimeUsersOutput struct {
	DowntimeUsers []DowntimeUser
	Result        string
}

type DowntimeUser struct {
	Name  string
	Value time.Duration
}

type DowntimesInput struct {
	Start int64
	End   int64
	Data  string
}

type DowntimesOutput struct {
	Downtimes []Downtime
	Result    string
}

type Downtime struct {
	Name  string
	Value time.Duration
}

func getDowntimes(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var data DowntimesInput
	err := json.NewDecoder(request.Body).Decode(&data)
	if err != nil {
		logError("Bar", "Error parsing data from page: "+err.Error())
		var responseData DowntimesOutput
		responseData.Result = "nok"
		writer.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(writer).Encode(responseData)
		return
	}

	startTime := time.Unix(data.Start/1000, 0)
	endTime := time.Unix(data.End/1000, 0)
	logInfo("Bar", "User asked data from "+startTime.String()+" to "+endTime.String()+" for "+data.Data)
	var downtimes []Downtime
	downtimes = append(downtimes, Downtime{"Smoking", (1*time.Hour + 27*time.Minute) / 1000000000})
	downtimes = append(downtimes, Downtime{"Resting", (27 * time.Minute) / 1000000000})
	downtimes = append(downtimes, Downtime{"Sleeping", (57 * time.Minute) / 1000000000})
	downtimes = append(downtimes, Downtime{" ", (0 * time.Minute) / 1000000000})
	downtimes = append(downtimes, Downtime{"  ", (0 * time.Minute) / 1000000000})
	downtimes = append(downtimes, Downtime{"   ", (0 * time.Minute) / 1000000000})
	downtimes = append(downtimes, Downtime{"    ", (0 * time.Minute) / 1000000000})
	downtimes = append(downtimes, Downtime{"     ", (0 * time.Minute) / 1000000000})
	downtimes = append(downtimes, Downtime{"      ", (0 * time.Minute) / 1000000000})
	downtimes = append(downtimes, Downtime{"       ", (0 * time.Minute) / 1000000000})
	logInfo("Bar", "Sending data")
	var responseData DowntimesOutput
	responseData.Result = "ok"
	responseData.Downtimes = downtimes
	writer.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(writer).Encode(responseData)
	return
}

func getDowntimeUsers(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var data DowntimeUsersInput
	err := json.NewDecoder(request.Body).Decode(&data)
	if err != nil {
		logError("Bar", "Error parsing data from page: "+err.Error())
		var responseData DowntimeUsersOutput
		responseData.Result = "nok"
		writer.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(writer).Encode(responseData)
		return
	}
	logInfo("Bar", "User asked for "+data.DowntimeName)
	var downtimeUsers []DowntimeUser
	downtimeUsers = append(downtimeUsers, DowntimeUser{"Petr Jahoda", 27 * time.Minute / 1000000000})
	downtimeUsers = append(downtimeUsers, DowntimeUser{"Karel Gott", 17 * time.Minute / 1000000000})
	downtimeUsers = append(downtimeUsers, DowntimeUser{"James Bond", 45 * time.Minute / 1000000000})
	downtimeUsers = append(downtimeUsers, DowntimeUser{"Josef Nedopil", 12 * time.Minute / 1000000000})
	downtimeUsers = append(downtimeUsers, DowntimeUser{" ", (0 * time.Minute) / 1000000000})
	downtimeUsers = append(downtimeUsers, DowntimeUser{"  ", (0 * time.Minute) / 1000000000})
	downtimeUsers = append(downtimeUsers, DowntimeUser{"   ", (0 * time.Minute) / 1000000000})
	downtimeUsers = append(downtimeUsers, DowntimeUser{"    ", (0 * time.Minute) / 1000000000})
	downtimeUsers = append(downtimeUsers, DowntimeUser{"     ", (0 * time.Minute) / 1000000000})
	downtimeUsers = append(downtimeUsers, DowntimeUser{"      ", (0 * time.Minute) / 1000000000})
	logInfo("Bar", "Sending data")
	var responseData DowntimeUsersOutput
	responseData.Result = "ok"
	responseData.DowntimeUsers = downtimeUsers
	writer.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(writer).Encode(responseData)
	return
}

func getUsersDowntimes(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var data DowntimesInput
	err := json.NewDecoder(request.Body).Decode(&data)
	if err != nil {
		logError("Bar", "Error parsing data from page: "+err.Error())
		var responseData DowntimesOutput
		responseData.Result = "nok"
		writer.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(writer).Encode(responseData)
		return
	}

	startTime := time.Unix(data.Start/1000, 0)
	endTime := time.Unix(data.End/1000, 0)
	logInfo("Bar", "User asked data from "+startTime.String()+" to "+endTime.String()+" for "+data.Data)
	var downtimes []Downtime
	downtimes = append(downtimes, Downtime{"Hiking", (3*time.Hour + 27*time.Minute) / 1000000000})
	downtimes = append(downtimes, Downtime{"Hoping", (2*time.Hour + 27*time.Minute) / 1000000000})
	downtimes = append(downtimes, Downtime{"  ", (0 * time.Minute) / 1000000000})
	downtimes = append(downtimes, Downtime{"   ", (0 * time.Minute) / 1000000000})
	downtimes = append(downtimes, Downtime{"    ", (0 * time.Minute) / 1000000000})
	downtimes = append(downtimes, Downtime{"     ", (0 * time.Minute) / 1000000000})
	downtimes = append(downtimes, Downtime{"      ", (0 * time.Minute) / 1000000000})
	downtimes = append(downtimes, Downtime{"       ", (0 * time.Minute) / 1000000000})
	downtimes = append(downtimes, Downtime{"        ", (0 * time.Minute) / 1000000000})
	downtimes = append(downtimes, Downtime{"         ", (0 * time.Minute) / 1000000000})
	logInfo("Bar", "Sending data")
	var responseData DowntimesOutput
	responseData.Result = "ok"
	responseData.Downtimes = downtimes
	writer.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(writer).Encode(responseData)
	return
}
