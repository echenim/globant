package logs

import "os"

func doesFileExist(name string) (bool){
  _,e:= os.Stat(name)
  if e!=nil{
	  return false
  }
  return true
}