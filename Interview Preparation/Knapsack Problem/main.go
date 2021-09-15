package main

import (
	"fmt"
)

//type serverUsage defined with important project parameters - Server capacity,Server Users,Users Priority,Users RAM
type serverUsage struct {
	serverCapacity int
	serverUsers    int
	usersPriority  []int
	usersRAM       []int
}

//main - Driver Function
func main() {

	var servers int
	totalOptimalValue:= 0

	fmt.Println("Enter the no of Servers")
	fmt.Scanln(&servers)
    //Loop to find the optimal & total value for each project
	for i:=0;i<servers;i++{
		var project serverUsage
		(&project).createProject()
		optimalValue:=(&project).optimalValue()
		fmt.Println("Optimal Value can be added to the Server ",i," is ",optimalValue)
		fmt.Println("total Value of the each projet ",i," is ",(&project).totalValue())
		totalOptimalValue=totalOptimalValue+optimalValue

		}
        //Print the sum of optimal values of each project
		fmt.Println("total optimal Value of all projects for the week is ",totalOptimalValue)

}

//createProject - Create projects with user Input
func (s *serverUsage)createProject(){

	fmt.Println("Inputs")
	fmt.Println("1.Server Capacity")
	fmt.Println("2.No of users")

    // to get Server Capacity & Server Users
	fmt.Scanln(&(s.serverCapacity))
	fmt.Scanln(&(s.serverUsers))
    
	//to get priority of each user
	fmt.Println("Enter the users Priority")
	s.usersPriority=make([]int,s.serverUsers+1)
	for i:=1;i<=s.serverUsers;i++{
		fmt.Scanln(&(s.usersPriority[i]))	
	}

	//to get RAM of each user
	fmt.Println("Enter the users RAM")
	s.usersRAM=make([]int,s.serverUsers+1)
	for j:=1;j<s.serverUsers+1;j++{
		fmt.Scanln(&(s.usersRAM[j]))			
	}
	

}

//optimalValue - find the optimal value for the project.
//Optimal Value = Max value can be added to the project without exceeding server capacity
//Solved using dynamic programming 
func(s *serverUsage)optimalValue()int{

	var i,j,k int	

	//Allocate memory for 2D slice of size 
	// Row = serverusers+1 (0 to server users)
	// Column = Servercapacity+1 (0 to server capacity)
	DPmatrix:= make([][]int,s.serverUsers+1)
	for i=0;i<=s.serverUsers;i++{
		DPmatrix[i]= make([]int,s.serverCapacity+1)
	}

	//Build Dynamic Programming matrix,by calculating optimal value for the given user - under the given capacity based on Priority
	for j=0;j<s.serverUsers+1;j++{
		for k=0;k<=s.serverCapacity;k++ {
		
			if (j==0 || k == 0){
				DPmatrix[j][k]=0
				
			}else if (s.usersRAM[j]<= s.serverCapacity && (k-s.usersRAM[j])>=0){
				userPicked:=s.usersPriority[j]+DPmatrix[j-1][k-s.usersRAM[j]]
				userNotPicked:=DPmatrix[j-1][k]
				DPmatrix[j][k]= Max(userPicked,userNotPicked)
				
			}else {			
				DPmatrix[j][k]=DPmatrix[j-1][k] 
				
			}

		
		}
	}

	//Print the derived Dynamic Programming matrix
	for l:=0;l<=s.serverUsers;l++{
	fmt.Println(DPmatrix[l])
	}

	//Return the Last row,Last Column value - Which should be the optimal vlaue
	return DPmatrix[j-1][k-1]
}

//totalValue - total value of the project
func (s *serverUsage)totalValue()int{
	totalSum:= 0
    //Loop to Calculate the total priority of the project
	for _,value:= range s.usersPriority{
		totalSum=totalSum+ value
	}

	return totalSum

}

//Max - to find the maximum of two integer value
func Max(x, y int) int {
    if x < y {
        return y
    }
    return x
}

 