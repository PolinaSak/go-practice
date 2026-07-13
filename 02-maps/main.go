package main
import(
	"fmt"
	"errors"
)

func main(){
	res, err := sum(4,-5)
	if err != nil{
		fmt.Print("oshibka", err)
	} else{
		fmt.Print(res)
	}
}
func sum(a, b int)(int, error){
	if a>0 && b>0{
		return a+b, nil
	} else{
		return 0, errors.New("ыыыы ошиибка")
	}
	
}
