package main
import(
  "fmt"
)

type st struct{
  x int
  y int
}

func main(){
  v := st{55, 18}
  fmt.Println(v.x)
}
