package logs

import (
	"io/ioutil"
	"os"
	"time"
)

//Log function
func (lb LogBook) Log(nameforthelog string){
	filename:= "tmp/"+generateName(lb.Kind,nameforthelog)
	data:="---------START---------\nSTATUS: "+lb.Kind+"\nMESSAGE: "+lb.Message+"\nDATE: "+lb.DateTime+"\n---------END---------\n\n\n"	
	if doesFileExist(filename){
		f,e:= os.OpenFile(filename,os.O_APPEND|os.O_WRONLY,0644)
		if e!=nil{
		}
	   defer f.Close()
	   _,er:=f.WriteString(data)
	   if er != nil{
	   }
    }else{
	   e:=ioutil.WriteFile(filename,[]byte(data),0644)
	   if e!=nil{
	   }
   }

}

//ComposeLogBody function to compose the log information
func (lb LogBook) ComposeLogBody(kind, message string) *LogBook{
  l:= LogBook{}
  l.Kind = kind
  l.Message = message
  l.DateTime = time.Now().String()
  return &l
}