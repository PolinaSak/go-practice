//Задача: Создай слайс a := []int{1, 2, 2, 3, 4, 4, 4, 5} (отсортированный, с повторами подряд). Напиши функцию, которая удаляет соседние дубликаты in-place (без создания нового слайса через make), и возвращает укороченный слайс. Результат должен быть [1 2 3 4 5].
package main
import(
	"fmt"
)

func main(){
	a := []int{1, 2, 2, 3, 4, 4, 4, 5}
	fmt.Println(dedup(a))

}

func dedup(a []int)[]int{
	slow :=0
	for fast:=1; fast< len(a); fast++{
		if a[fast] != a[slow]{
			slow++
			a[slow]= a[fast]
		}
	}
	return a[:slow+1]
}