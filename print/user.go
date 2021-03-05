package print 

import (
	"fmt"
	"reflect"
	models "challenge/models"
)

func PrintlnUser(u models.User) {
	s := reflect.ValueOf(&u).Elem()
	typeOfT := s.Type()

	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		// if typeOfT.Field(i).Name == "Tags" {
		// 	 tag := fmt.Sprintf("%v", f.Interface())
		// 	 fmt.Printf("%30s", typeOfT.Field(i).Name);
		// 	 fmt.Printf("\t\t\t");
		// 	 fmt.Printf("%v", tag)
		// 	 fmt.Printf("\n");
		// }else {
			
			if typeOfT.Field(i).Name == "Tags" {
				var data []string = f.Interface().([]string)
				fmt.Printf("%30s", typeOfT.Field(i).Name);
				fmt.Printf("\t\t\t");
				fmt.Printf("[")
				for j:=0;j<len(data);j++ {
					fmt.Printf("\"")
					fmt.Printf("%v",data[j])
					fmt.Printf("\"")
					if j<len(data)-1{
						fmt.Printf(",")
					}
				}
				fmt.Printf("]")
				fmt.Printf("\n")
			}else{
				fmt.Printf("%30s", typeOfT.Field(i).Name);
				fmt.Printf("\t\t\t");
				fmt.Printf("%v", f.Interface())
				fmt.Printf("\n");
			}
		// }
	}
}