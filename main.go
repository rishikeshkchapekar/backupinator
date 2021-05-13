package main
 
import (
	"os/exec"
	"bytes"
	"fmt"
	"log"
	"strings"
	"io/ioutil"
	"os"
	"io"
	"time"

)

var user string
var backupDevice string
var toCopyDevice []string

var connectedDevices []string

func main(){
	copiedFlag := false

	for{
		uname := getUsername()
		media :=listMedia(uname)
		copiedFlag = checkForFlag(media)
		
		if copiedFlag==false{
			setBackupDevice(media)
			copyFiles()
			fmt.Println("files copied")
		}

		time.Sleep(5*time.Second)
	}
}

func checkForFlag(media []string) bool{
	copied := true
	if len(connectedDevices)==0 || connectedDevices==nil{
		connectedDevices=media
		copied = false
	}else{
		for i:=0;i<len(media);i++{
			exists := checkItemExists(media[i],connectedDevices)
			if exists==false{
				connectedDevices = append(connectedDevices,media[i])
				copied=false
				break
			}
		}
	}

	return copied
}

func copyFiles(){
	backupPath := "/media/"+user+"/"+backupDevice+"/"

	if backupDevice!=""{
		for i:=0;i<len(toCopyDevice);i++{
			devicePath := "/media/"+user+"/"+toCopyDevice[i]+"/"

			files,err := ioutil.ReadDir(devicePath)
			if err!=nil{
				log.Fatal(err)
			}

			for _,f := range files{
				fName := f.Name()
				fileExists := true
				// fName = strings.Replace(fName," ","\ ")
				from,err := os.Open(devicePath+fName)
				if err != nil {
					_,err := os.Create(backupPath+fName)
					if err!=nil{
						fmt.Println(err)
					}else{
						fileExists = false
					}
				}
				defer from.Close()

				to,err := os.OpenFile(backupPath+fName,os.O_RDWR|os.O_CREATE, 0666)
				if err != nil {
					log.Fatal(err)
				}
				defer to.Close()

				if fileExists==false{
					_,err = io.Copy(to,from)
					if err!=nil{
						log.Fatal(err)
					}
				}
			}	
		}

	}

}

func setBackupDevice(media []string){
	detected := false
	if len(media) < 2{
		fmt.Println("Not enough media devices")
		return
	}else{
		for i:=0;i<len(media);i++{
			backupDevBool:=detectBackup(media[i])
			if backupDevBool==true{
				backupDevice = media[i]
				detected=true
			}else{
				toCopyDevice = append(toCopyDevice,media[i])
			}
		}
	}

	if detected==false{
		backupDevice = ""
	}
}

func detectBackup(device string) bool{
	identifierFile := ".identifier"
	dir := "/media/"+user+"/"+device

	backupDev := false

	files,err := ioutil.ReadDir(dir)
	if err!=nil{
		log.Fatal(err)
	}

	for _,f := range files{
		fName := f.Name()
		fmt.Println(fName)
		if fName==identifierFile{
			backupDev=true
			break
		}
	}

	return backupDev
}

func listMedia(username string) []string{
	dir := "/media/"+username
	mediaNames := []string{}

	files,err := ioutil.ReadDir(dir)
	if err!=nil{
		log.Fatal(err)
	}

	for _,f := range files{
		mediaNames = append(mediaNames,f.Name())
	}

	return mediaNames
}

func getUsername() string{
	var username bytes.Buffer
	
	cmd := exec.Command("whoami")
	
	cmd.Stdout = &username
	
	err := cmd.Run()
	if err!=nil{
		log.Fatal(err)
	}
	uString := username.String()
	uString = strings.Trim(uString,"\n")
	user = uString
	return uString

}

func checkItemExists(item string,arr []string) bool{
	exists := false

	for i:=0;i<len(arr);i++{
		if item==arr[i]{
			exists=true
			break
		}
	}

	return exists
}
