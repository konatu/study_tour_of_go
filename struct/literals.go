package main

import(
"fmt"
)

type st struct{
  X int
  Y int
}

var (
  red = st{15,25}
  bull= st{X: 1}
  monster= st{}
  repo = &st{77,88}
)

func main(){
    fmt.Println(red, bull, monster, repo)
}
