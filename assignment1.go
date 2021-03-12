package main

import "fmt"

type User struct {
	mobileno string
	address  string
}

type Product struct {
	brand string
	price int
	qty   int
}

//arrayUsers:=make([]User)
var arrayUsers = make([]User, 0)
var m = make(map[string][2]Product, 0)

//var qty  []int = make([]int,0)

var mapUsers = make(map[User][]Product)

func AddUser(user User) bool {
	if user.mobileno > "9999999999" || user.mobileno < "5999999999" {
		fmt.Println("Invalid mobile no")
	}

	for _, val := range arrayUsers {
		if val.mobileno == user.mobileno {
			fmt.Println("Mobile no already exists")
			return false
		}
	}

	return true
}

func GetProducts(prod string) {

	for _, p := range m[prod] {
		if p.brand != "nil" {
			fmt.Printf("Brand:%s\nPrice:%d\n", p.brand, p.price)
		}
	}

}

func changeOrder(user User, oldbrand string, newBrand string, qty int, prod string) {

	for i := 0; i < len(mapUsers[user]); i++ {
		if mapUsers[user][i].brand == oldbrand {
			mapUsers[user][i].brand = newBrand
			mapUsers[user][i].qty = qty
			for _, val := range m[prod] {
				if val.brand == newBrand {
					mapUsers[user][i].price = val.price
				}
			}
			break
		}
	}

}

func cancelOrder(user User, brand string, prod string) {

	for i := 0; i < len(mapUsers[user]); i++ {
		if mapUsers[user][i].brand == brand {
			mapUsers[user][i].brand = "nil"
			break
		}

	}

}

func GetOrders(user User) {
	fmt.Printf("Your Cart\n")
	fmt.Printf("Order          Quantity           Price\n")
	for _, order := range mapUsers[user] {
		if order.brand != "nil" {
			fmt.Printf("%s     %d     %d\n", order.brand, order.qty, order.price*order.qty)
		}
	}
}

func main() {

	m["Smart Phone"] = [2]Product{Product{"Redmi", 15000, 0}, Product{"Samsung", 20000, 0}}
	m["Laptop"] = [2]Product{Product{"Dell", 50000, 0}, Product{"HP", 80000, 0}}
	m["TV"] = [2]Product{Product{"Sony", 150000, 0}, Product{"LG", 200000, 0}}

	flag := true
	valid := true
	var ch, choice, quantity int
	var user User
	//var choice int
	var prod, brand string

	var products []Product

	for flag {

		fmt.Printf("1.Add User\n2.Get Products\n3.Place order\n4.Change order\n5.Cancel order\n6.Get Orders\n7.Exit\nEnter your choice:")
		fmt.Scanln(&ch)

		switch ch {
		case 1:
			fmt.Println("Enter your mobile number:")
			fmt.Scanln(&user.mobileno)
			//fmt.Scanf("")
			fmt.Println("Enter your address:")
			fmt.Scanln(&user.address)
			valid = AddUser(user)
			products = make([]Product, 0)
			if valid == true {
				arrayUsers = append(arrayUsers, user)
			}

		case 2:
			fmt.Println("PRODUCTS:\n1.Smart Phone\n2.TV\n3.Laptop\nEnter your choice:")
			fmt.Scanln(&choice)

			switch choice {
			case 1:
				prod = "Smart Phone"
			case 2:
				prod = "TV"
			case 3:
				prod = "Laptop"
			default:
				fmt.Println("Invalid Choice!")
			}

			GetProducts(prod)
		case 3:
			fmt.Printf("PRODUCTS:\n1.Smart Phone\n2.TV\n3.Laptop\nEnter your choice:")
			fmt.Scanln(&choice)

			switch choice {
			case 1:
				prod = "Smart Phone"
			case 2:
				prod = "TV"
			case 3:
				prod = "Laptop"
			default:
				fmt.Printf("Invalid Choice!\n")
			}

			GetProducts(prod)

			fmt.Printf("Enter the brand name:")
			fmt.Scanln(&brand)
			fmt.Printf("Enter the quantity:")
			fmt.Scanln(&quantity)
			var product Product
			for _, b := range m[prod] {
				if b.brand == brand {
					product.brand = b.brand
					product.price = b.price
					product.qty = quantity
					//fmt.Printf("%T", mapUsers[user])
					products = append(products, product)
					mapUsers[user] = products
				}
			}
			//qty = append(qty,quantity)

			//mapUsers[user] = products
		case 4:
			GetOrders(user)
			fmt.Printf("Enter the order to be changed:")
			fmt.Scanln(&brand)
			var newBrand string
			fmt.Printf("Enter the new brand:")
			fmt.Scanln(&newBrand)
			fmt.Printf("Enter the quantity:")
			fmt.Scanln(&quantity)

			changeOrder(user, brand, newBrand, quantity, prod)

		case 5:
			GetOrders(user)
			fmt.Printf("Enter the order you want to cancel:")
			fmt.Scanln(&brand)

			cancelOrder(user, brand, prod)

		case 6:
			GetOrders(user)

		case 7:
			flag = false

		}

	}
	//fmt.Println(arrayUsers)
	//fmt.Println(mapUsers)
}
