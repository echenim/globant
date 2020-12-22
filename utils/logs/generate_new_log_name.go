package logs

import (
	
	"strconv"
	"time"
)


func generateName(kindoflog, logname string) string{	
	t:=time.Now().Local()
	return kindoflog+"_"+logname+"_"+t.Month().String()+"-"+strconv.Itoa(t.Day())+".log"
}