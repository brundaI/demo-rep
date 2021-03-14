package shopping

import "fmt"


type User struct{
    Mobileno int
    Place string 
}

type Product struct{
    Brand string
    Price int
    Qty   int
}
var Users_list = make([]User,0)
var Products_list = make(map[string][2]Product)
var Users_map map[User][]Product

func init(){
Products_list["Smartphone"] = [2]Product{Product{"Redmi",15000,0},Product{"Samsung",20000,0}}
Products_list["Laptop"] = [2]Product{Product{"Dell",50000,0},Product{"HP",80000,0}}
Products_list["TV"]=[2]Product{Product{"Sony",150000,0},Product{"LG",200000,0}}
}


func AddUser(mobileno int,place string) bool{
	if mobileno > 9999999999 || mobileno < 5999999999{
	fmt.Println("Invalid mobile no")
        return false
    }
    if len(Users_list)==0{
		return true
	}
    for _,val:= range Users_list{
        if val.Mobileno == mobileno{
	    fmt.Println("mobile no already exists")
            return false
        }
    }
	return true
}

func GetProducts(){
    
    for k,_:=range Products_list{
            fmt.Printf("%s\n",k)
        }
  }

func GetBrands(prod string){
    
    for _,p:=range Products_list[prod]{
            fmt.Printf("Brand:%s Price:%d\n",p.Brand,p.Price)
        }
    }
    


func PlaceOrder(user User,prod Product){
	if Users_map==nil{
		Users_map = make(map[User][]Product)
	}

	_,present := Users_map[user] 
	if present{
		Users_map[user] = append(Users_map[user],prod)
	}else{
		Users_map[user] = make([]Product,0)
		Users_map[user] = append(Users_map[user],prod)
	}
}

func ChangeOrder(user User,oldbrand string,newBrand string,qty int,prod string){

    for i:=0;i<len(Users_map[user]);i++{
        if Users_map[user][i].Brand == oldbrand{
             Users_map[user][i].Brand = newBrand
             Users_map[user][i].Qty = qty
             for j:=0;j<len(Products_list[prod]);j++{
                if Products_list[prod][j].Brand == newBrand{
                    Users_map[user][i].Price = Products_list[prod][j].Price
		    break
                }
            }
          break  
        }
    }
    
}

func CancelOrder(user User,brand string,prod string){

    for i:=0;i<len(Users_map[user]);i++{
        if Users_map[user][i].Brand == brand{
             Users_map[user][i].Brand = "nil"
             break   
            }
            
        }
    
}


func GetOrders(user User) {
	fmt.Printf("Your Cart\n")
	//fmt.Printf("Order          Quantity           Price\n")
	for _,order := range Users_map[user]{
	    if order.Brand != "nil"{
		    fmt.Printf("order:%s     Quantity:%d     Price:%d\n",order.Brand,order.Qty,order.Price*order.Qty)
	    }
    }
}

/*func main(){
	var user = User{9113259038,"Bangalore"}
	flag := AddUser(user.mobileno,user.place)
	if flag == false{
		fmt.Println("Invalid mobile no/mobile no already exists")
		
	}else{
		fmt.Println("User added")
		GetProducts("TV")
		prod:=Product{"LG",200000,1}
		placeOrder(user,prod)
		GetOrders(user)
		changeOrder(user,"LG","Sony",1,"TV")
		GetOrders(user)
		cancelOrder(user,"Sony","TV")
		GetOrders(user)
	}



}

*/