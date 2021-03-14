package main

import (
	"fmt"
	 "urbanKart/shopping"
	 "urbanKart/admin"
	)

func main(){

var cont = true
var choice,mobile,quantity int
var p,product,brand,newbrand string
var user shopping.User
var prod shopping.Product

fmt.Println("Enter your mobileno")
fmt.Scanln(&mobile)
fmt.Println("Enter your place")
fmt.Scanln(&p)
shopping.AddUser(mobile,p)
user = shopping.User{mobile,p}
shopping.Users_list = append(shopping.Users_list,user)


for cont{
fmt.Println("1.Get Products\n2.Place Order\n3.Change order\n4.Cancel order\n5.Get Orders\n6.Add another account\n7.Admin\n8.Exit\nEnter your choice")
fmt.Scanln(&choice)
switch choice{
	case 1:shopping.GetProducts()
	       fmt.Println("Enter the product")
	       fmt.Scanln(&product)
	       shopping.GetBrands(product)

	case 2:fmt.Println("Enter the brand of "+product+" you want")
	       fmt.Scanln(&brand)
	       fmt.Println("Enter the quantity")
	       fmt.Scanln(&quantity)
	       for _,val:=range shopping.Products_list[product]{
			if val.Brand == brand{
				prod = shopping.Product{brand,val.Price,quantity}
				}
		}
	       shopping.PlaceOrder(user,prod)
	case 3:fmt.Println("Enter the product you want to change")
	       fmt.Scanln(&product)
	       fmt.Println("Enter the brand you want to change")
	       fmt.Scanln(&brand)
	       fmt.Println("Enter the new item that you want")
	       fmt.Scanln(&newbrand)
	       fmt.Println("Enter the quantity")
	       fmt.Scanln(&quantity)
	       shopping.ChangeOrder(user,brand,newbrand,quantity,product)

	case 4:fmt.Println("Enter the product you want to cancel")
	       fmt.Scanln(&product)
	       fmt.Println("Enter the brand you want to cancel")
	       fmt.Scanln(&brand)
	       shopping.CancelOrder(user,brand,product)
	case 5:shopping.GetOrders(user)
	case 6:fmt.Println("Enter your mobileno")
	       fmt.Scanln(&mobile)

	       fmt.Println("Enter your place")
	       fmt.Scanln(&p)

	       if shopping.AddUser(mobile,p){
			user = shopping.User{mobile,p}
			shopping.Users_list = append(shopping.Users_list,user)
			fmt.Println("User was added")
	       }else{
			fmt.Println("User was not added")
		}

	case 7: var ch int
		fmt.Println("1.Add representative\n2.Assign Representative to a customer\n3.View assigned Representatives\nEnter your choice")
		fmt.Scanln(&ch)
		if ch==1{
			admin.AddRep()
			}
		if ch==2{
			var mobile int
			var i int
			fmt.Println("Enter mobileno of the user to be assigned a rep")
			fmt.Scanln(&mobile)
			//fmt.Println("Users List")
			for i=0;i<len(shopping.Users_list);i++{
					if shopping.Users_list[i].Mobileno == mobile{
						if admin.AssignRep(shopping.Users_list[i]){
							fmt.Println("A rep has been assigned")
						}else{
							fmt.Println("No matching Rep found")
						 }
					break
					}
				}
			if i==len(shopping.Users_list){
					fmt.Println("User does not exist")
				}
			
			}
		if ch==3{
			 admin.ViewAssignments()
			}
	case 8:cont=false

	       
}

}

}