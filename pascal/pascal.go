package main

import("fmt")

func main(){
	var rows int
	fmt.Print("Enter the number of rows: ")
	fmt.Scan(&rows)
	arr:=make([][]int,rows)
	for i:=range arr{
		arr[i]=make([]int,i+1)
	}
	for i:=0;i<rows;i++{
		for j:=0;j<i+1;j++{
			if i==0{
				arr[i][j]=1
				continue;
			}
			if i ==1{
				arr[i][j]=1
				continue;
			}
			if i>1{
				if j==0||j==i{
					arr[i][j]=1
				}else{
					arr[i][j]=arr[i-1][j-1]+arr[i-1][j]
				}
			}
		}
		
	
	}
	for i:=0;i<rows;i++{
		for k:=1;k<=(rows-(i+1));k++{
			fmt.Print(" ")
		}
		for j:=0;j<i+1;j++{
             fmt.Print(arr[i][j]," ")
		}
		fmt.Println()
	}
                                                                                                                                                                                                                                                                                                                                                                                                      


}