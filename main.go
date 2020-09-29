package main

import "fmt"

func main(){
	var mes int
	var dia int
	fmt.Scan(&dia)
	fmt.Scan(&mes)

	switch{
		case mes == 1:
			if dia < 21{
				fmt.Println("capricornio")
			}else{
				fmt.Println("acuario")
			}
		case mes == 2:
			if dia < 21{
				fmt.Println("acuario")
			}else{
				fmt.Println("piscis")
			}
		case mes == 3:
			if dia < 21{
				fmt.Println("piscis")
			}else{
				fmt.Println("aries")
			}
		case mes == 4:
			if dia < 21{
				fmt.Println("aries")
			}else{
				fmt.Println("tauro")
			}
		case mes == 5:
			if dia < 21{
				fmt.Println("tauro")
			}else{
				fmt.Println("geminis")
			}
		case mes == 6:
			if dia < 21{
				fmt.Println("geminis")
			}else{
				fmt.Println("cancer")
			}
		case mes == 7:
			if dia < 21{
				fmt.Println("cancer")
			}else{
				fmt.Println("leo")
			}
		case mes == 8:
			if dia < 21{
				fmt.Println("leo")
			}else{
				fmt.Println("virgo")
			}
		case mes == 9:
			if dia < 21{
				fmt.Println("virgo")
			}else{
				fmt.Println("libra")
			}
		case mes == 10:
			if dia < 21{
				fmt.Println("libra")
			}else{
				fmt.Println("escorpio")
			}
		case mes == 11:
			if dia < 21{
				fmt.Println("escorpio")
			}else{
				fmt.Println("sagitario")
			}
		case mes == 12:
			if dia < 21{
				fmt.Println("sagitario")
			}else{
				fmt.Println("capricornio")
			}
	}

}