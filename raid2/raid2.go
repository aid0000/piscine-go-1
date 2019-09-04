package main

import (
	"fmt"
	"os"
	piscine".."
 
)
var list [][]string
func main() {

argumento(os.Args)
IsValid(list,0)
Display(list)
}

func argumento(a []string){
  for p,v:=range a {
     if p!=0 {
    var ar[] string
    
      for o:=0; o<len(v) ; o++{
        
        ar=append(ar,string(v[o]))   
      }
      list=append(list,ar)
    
    }
  }
}

// display  of the board 
func Display(list [][]string){
  for row := 0; row < 9; row++ {

		for col := 0; col < 9; col++ {
			fmt.Printf("%s ",list[row][col])
		}
			fmt.Print("\n")	
	}
}

//converts the number that  string to an int
//then verifies the line 
//convert list to int 
func AbsentOnLine(k int, list [][]string, x int) bool {
  var y int
    for y=0; y < 9; y++ {
        if piscine.Atoi(list [x][y]) == k{
            return false
          }
        }
    return true
  }

  func AbsentOnRow(k int, list [][]string, y int) bool {
    var x int
    for x=0; x < 9; x++{
         if piscine.Atoi(list[x][y]) == k{
             return false;
           }
         }
     return true;
  }
//search if a number is on the blok by /3 and giving the rest
  func AbsentOnBloc(k int, list [][]string, x int, y int) bool {
    var firstX, firstY int;
    firstX =  x-(x%3)
    firstY =  y-(y%3)
    for x = firstX; x < firstX+3; x++ {
          for y = firstY; y < firstY+3; y++ {
              if piscine.Atoi(list[x][y]) == k{
                  return false;
                }
          }
        }
      return true;
  
  } 
  //if the number is equal to last then return true
  //else go check if is not a dot 
  //then check line row and blok 

  func IsValid(list [][]string, position int) bool {

    if (position == 9*9){
          return true;
        }
  
        var x, y, k int
  
      x = position/9
      y = position%9
  
      if list[x][y] != "."{
          return IsValid(list, position+1);
        }
  
      for k=1; k <= 9; k++ {
          if AbsentOnLine(k,list,x) && AbsentOnRow(k,list,y) && AbsentOnBloc(k,list,x,y){
              list[x][y] = convertInttostring(k)
  
              if IsValid(list, position+1){
                  return true;
                }
          }
      }
      list[x][y] = ".";
      return false;
  
  }
  


func convertInttostring(n int) string{
	t:=1
		if n!=0{
			f:=(n/10)*t
			if f!=0{
				convertInttostring(f)
			}
		  	k:=((n%10*t))+'0'
	     	return (string(k))
		 }else{
		    return "0"
		}
}
