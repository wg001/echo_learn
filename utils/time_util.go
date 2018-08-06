package utils

import "time"

const TIME_FAMMAT = "2006-01-02 15:04:05"

func FormatTime(time time.Time) string {
	if time.IsZero(){
		return ""
	}

	return time.Format(TIME_FAMMAT)
}

func FormatTimeToLocal(timeval time.Time) time.Time{
	tempTime,parseErr:=time.ParseInLocation(TIME_FAMMAT, timeval.Format(TIME_FAMMAT),time.Local)
	if parseErr!=nil{
		panic(parseErr.Error())
	}
	return tempTime
}
func FormatStrTimeToLocal(timeval string) time.Time{
	tempTime,parseErr:=time.ParseInLocation(TIME_FAMMAT, timeval,time.Local)
	if parseErr!=nil{
		panic(parseErr.Error())
	}
	return tempTime
}