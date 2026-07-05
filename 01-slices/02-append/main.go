package main
import "fmt"
func main(){
	a := []int{1,2,3} // cap 3 len 3
	fmt.Printf("%v len: %d cap: %d\n\n", a, len(a), cap(a))
	a = append(a, 4,5,6,7) // cap вырос с 3 до 8 — Go автоматически выделил новый массив
	// с запасом когда  не хватило места. Алгоритм роста не всегда
	// ровно x2 (зависит от текущего размера и версии Go)
	fmt.Print(a, " len: ", len(a), " cap: ", cap(a), "\n")
}