package admin

import "fmt"
import  "urbanKart/shopping"

type Rep struct{
	 Name string
	Mobileno int 
	 Place string
}

var ArrayReps = make([]Rep,0)
var AssignedReps map[Rep]shopping.User

func AddRep() bool{
	var name string
	var mobile int
	var place string
	fmt.Println("Enter the name of the representative")
	fmt.Scanln(&name)
	fmt.Println("Enter the mobile no of the representative")
	fmt.Scanln(&mobile)
	fmt.Println("Enter the place of the representative")
	fmt.Scanln(&place)
	
	if mobile > 9999999999 || mobile < 5999999999{
			fmt.Println("Invalid mobile no!")
		}
	for _,val := range ArrayReps{
		if val.Mobileno==mobile{
			fmt.Println("Mobile no already exists!")
			return false
		}
	}

	rep:=Rep{name,mobile,place}
	ArrayReps = append(ArrayReps,rep)
	return true
}

func AssignRep(user shopping.User) bool{
	for _,v := range ArrayReps{
		if v.Place == user.Place{
			if AssignedReps == nil{
				AssignedReps = make(map[Rep]shopping.User)
				AssignedReps[v]=user
				return true
				}
			_,assigned :=  AssignedReps[v]
			if assigned{
					return false
			}

		        AssignedReps[v]=user
			return true
		 }
	}
	return false
}

func ViewAssignments(){
	for k,v := range AssignedReps{
		fmt.Println(k," has been assigned to the user ",v)
	}
}
