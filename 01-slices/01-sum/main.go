package main
import "fmt"

func main(){
	a := []int{1,2,3,4,5}
	fmt.Println("Слайс:", a)
	fmt.Println("Сумма элементов:", sumSlice(a))

}
func sumSlice(s []int) int{
	total := 0
	for _, v := range s{
		total += v
	}
	return total
}