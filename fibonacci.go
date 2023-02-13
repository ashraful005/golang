package main

import "fmt"

func main(){

var i,j,temp,sum,n int64

i=1
j=2
sum=0
for n=1; j<=4000000; n++{
fmt.Println(i)
temp=i+j
if j%2==0{
sum=sum+j
}
i=j
j=temp
}
fmt.Println(sum)


}