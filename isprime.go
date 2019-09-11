
package piscine


func IsPrime(value int) bool {
    if value<=1{
        return false
    }
      for i:=2; i<value; i++{
         if value % i == 0{
          return false
        }
    }
    return true
}


 