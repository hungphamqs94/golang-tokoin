package main

import (
	"fmt"
	"math"
)

func chuoidoixung(s string) bool {
	for i:=0; i<len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false	
		} 
	}
	return true
} 

func tinhtoan(a [][]int) int {
	sum := 0;
	for i:=0;i<len(a);i++ {
		if i==0 || i == len(a)-1 {
			for j:=0;j<len(a[i]);j++ {
				sum+=a[i][j]	
			}
		}else{
			sum += a[i][0]+a[i][len(a[i])-1]
		}
	}
	return sum;
}
func daochuoi(s string) string {
	mangdao := []byte(s)
	
	for i:=len(s)-1; i>=0; i=i-2 {
		mangdao[len(s)-1-i] = byte(s[i])
	}
	return string(mangdao)
}

func tinhSn(n int) int{
	sum := 0;
	for i:=1;i<=n; i++{
		sum += i;
	}
	return sum;
}

func tinhS2n(n int) int {
	sum:= 0;
	for i:=1;i<=n;i++{
		sum+=i*i;
	}
	return sum;
}

func tinh1n(n int) float64 {
	var sum float64
	sum=0
	for i:=1;i<=n;i++ {
		sum+=1/float64(i);
	}
	return sum;
}

func hammu(x float64, n float64) float64{
	return math.Pow(x,n);
}

func hamnhanhay(n int) int {
	sum:=1;
	
	for i:=1;i<=n;i++ {
		if i<n || i==1 {
			sum=sum*i
		} else {
			sum=sum*i+hamnhanhay(n-1)
		}
	}
	return sum;
}

func hamchiahay(n int) float64 {
	var sum float64
	sum=1
	for i:=1;i<=n;i++ {
		sum=1/float64(hamnhanhay(n))
	}
	return sum;
}




func main(){
	fmt.Println("check chuoi doi xung cua:", chuoidoixung("aaabbaaa"))
	fmt.Println("check chuoi doi xung cua:", chuoidoixung("adddbddda"))
	fmt.Println("check chuoi doi xung cua:", chuoidoixung("ajdjdjdjs"))
	
	a:= [][]int{{1,3,4,5},{1,3,4,5},{1,3,4,5},{1,3,4,5}}
	fmt.Println("gia tri tong tinh duoc la:", tinhtoan(a))
	fmt.Println("ta thu dao chuoi xem nao:", daochuoi("hungpham"))
	fmt.Println("gia tri cua Sn:",tinhSn(100))
	fmt.Println("the con gia tri nao cua s2n:", tinhS2n(50))
	fmt.Println("the thi tinh gia tri:", tinh1n(100))
	fmt.Println("ham so mu thi lam the nao:", hammu(6,3))
	fmt.Println("gia tri cua phep nhan lien tuc:", hamnhanhay(8))
	fmt.Println("roi thi se thap sang:", hamchiahay(8))
}