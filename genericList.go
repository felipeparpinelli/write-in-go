package main
import "fmt"

// interface: allows any type 
type GenericList [] interface {}

func (list *GenericList) RemoveIndex(index int) interface{} {
    l := *list
    removed := l[index]
    *list = append(l[0:index], l[index+1:]...)
    return removed
}

func (list *GenericList) RemoveFirstIndex() interface{} {
    return list.RemoveIndex(0)
}

func (list *GenericList) RemoveLastIndex() interface{} {
    return list.RemoveIndex(len(*list)-1)
}

func main() {
    list := GenericList {
        1, "Coffee", 42, true, 23, "Ball", 3.14, false,
    }
    fmt.Printf("Original List:\n%v\n\n", list)
    fmt.Printf(
        "Removing from the beginning: %v, after remove:\n%v\n",
        list.RemoveFirstIndex(), list)
    fmt.Printf(
        "Removing from the end: %v, afer remove:\n%v\n",
        list.RemoveLastIndex(), list)
    fmt.Printf(
        "Removing from index 3: %v, after remove:\n%v\n",
        list.RemoveIndex(3), list)
    fmt.Printf(
        "removing from index 0: %v, after remove:\n%v\n",
        list.RemoveIndex(0), list)
    fmt.Printf(
        "Removing from the last index: %v, after remove:\n%v\n",
        list.RemoveIndex(len(list)-1), list)
}