package maps

import "fmt"

var x =make(map[string]int)

func MapInit(){
	x["age"]=10
	x["salary"]=10000
	x["Cities"]=107
	fmt.Println(x)
}

//loop through map

func GetKeysValues(newmap map[string]int){
	for key,value:=range newmap{
		fmt.Printf("key:%s value:%d\n",key,value)
	}

}

func ReturnAMap()map[string]int{
	return map[string]int{
		"age":10,
		"salary":100000000,
		"cars":30}

	
}

func DeleteMapAtAnIndex(newmap map[string]int, key string){

	if _,ok := newmap[key];ok{
		delete(newmap,key)
	fmt.Println(newmap)


	}else{
		fmt.Println("Key does not exists")
	}
	
	
}