package main

import (
	"encoding/json"
	"fmt"
	"os"

	"flag"

	"github.con/et-zone/ecod/template"
)

func main() {
	var Name string
	flag.StringVar(&Name, "n", "stu.json", "please input config file name ,like stu.json, cmd = ecod -n xx.json .")
	flag.Parse()
	d := template.GetFile(Name)
	dt := map[string]interface{}{}
	json.Unmarshal([]byte(d), &dt)
	// tmpData := template.GetFile("temp.tml")
	temp, err := template.GetTemplate(template.TMPDATA)
	if err != nil {
		fmt.Println("GetTemplate err = ", err.Error())
	}

	s, err := template.ExecTemplate(temp, dt)
	if err != nil {
		fmt.Println("ExecTemplate err = ", err.Error())
	}
	err = os.Chdir("./dao")
	if err != nil {
		os.Mkdir("dao", 0766)
	}
	os.Chdir("./dao")
	fileName := dt["sName"].(string)
	fileName = template.Jsondata(fileName)
	err = template.WriteFile(fileName+".go", s)
	if err != nil {
		fmt.Println("WriteFile err = ", err.Error())
	}

	// cmd := exec.Command("go", "fmt")
	// err = cmd.Run()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(string(opBytes))

}

/*json file

{
    "tname": "teach",
    "sName": "teacher",
    "data": [
        {
            "fname": "id",
            "ftype": "int64"
        },
        {
            "fname": "name",
            "ftype": "string"
        },
        {
            "fname": "v_fname",
            "ftype": "string"
        }
    ]
}
*/
