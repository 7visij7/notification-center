package model

import (
	"os"
	"io"
	"net"
	"fmt"
	"errors"
	"io/ioutil"
	"github.com/hirochachacha/go-smb2"	
    "git.puls.ru/devops1/sre/notification-center/types"
)

func CheckDns(dnsrecord types.DnsRecord) (error) {
	// err := CheckIPAddress(dnsrecord.Ip)
	// if err != nil {
	//   fmt.Println(err)
	//   return err
	// }
	// fmt.Println(dnsrecord.Action)
	err := CheckAction(dnsrecord.Action)
	if err != nil {
	  fmt.Println(err)
	  return err
	}
	return nil
  }

// func FileExists(filename string) bool {
// 	_, err := os.Stat(filename)
// 	if os.IsNotExist(err) {
// 		return true
// 	}
// 	return false
//   }
  
func WriteCvs (data string) error{
	conn, err := net.Dial("tcp", types.SmbServer)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer conn.Close()
	d := &smb2.Dialer{
		Initiator: &smb2.NTLMInitiator{
			User:     types.SmbUser,
			Password: types.SmbPassword,
		},
	}
	s, err := d.Dial(conn)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer s.Logoff()
	fs, err := s.Mount(types.SmbCatalog)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer fs.Umount()
	f, err := fs.Open(types.Filecsv)
	if os.IsNotExist(err) {
	    err = fs.WriteFile(types.Filecsv, []byte(data), 0644)
		if err != nil {
			fmt.Println(err)
			return err
		}
		return nil
	}
	_, err = f.Seek(0, io.SeekStart)
	if err != nil {
		fmt.Println(err)
		return err
	}
	csv, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		return err
	}
	data += string(csv)
	err = fs.WriteFile(types.Filecsv, []byte(data), 0644)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil

	// if FileExists(types.Filecvs) {
	// 	file, err := os.Create(types.Filecvs)
	// 	if err != nil {
	// 		// fmt.Println(err)
	// 		return err
	// 	}
	// 	defer file.Close()    
	// }
	// file, err := os.OpenFile(types.Filecvs, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	// if err != nil {
	// 	// fmt.Println(err)
	// 	return err
	// }
	// defer file.Close()    
	// writer := csv.NewWriter(file)
	// defer writer.Flush()    
	// err = writer.Write(data)
	// 	if err != nil {
	// 	// fmt.Println(err)
	// 	return err
	// 	}
	// return nil
}
  

func DoDns(dnsrecords types.DnsRecord) (error){
	var data string
	data = fmt.Sprintf("%s,%s,%s\n",dnsrecords.Name, dnsrecords.Ip, dnsrecords.Action)
	// data = string{dnsrecords.Name, dnsrecords.Ip, dnsrecords.Action}
	err := WriteCvs(data)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
  }
  
func CheckIPAddress(ip string) error{
	if net.ParseIP(ip) == nil {
		fmt.Println("Invalid IP address")
		return errors.New("Invalid IP address")
	} else {
		return nil
	}
  }
  
func CheckAction(action string) error{
	if (action == "a" || action == "d")  {
		return nil
	} else {
		fmt.Println("Invalid actions")
		return errors.New("Invalid actions")
	}
  }