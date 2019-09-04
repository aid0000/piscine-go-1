
package main

import "fmt"

var sudoku = [9][9]int{
    {9, 0, 0, 1, 0, 0, 0, 0, 5},
    {0, 0, 5, 0, 9, 0, 2, 0, 1},
    {8, 0, 0, 0, 4, 0, 0, 0, 0},
    {0, 0, 0, 0, 8, 0, 0, 0, 0},
    {0, 0, 0, 7, 0, 0, 0, 0, 0},
    {0, 0, 0, 0, 2, 6, 0, 0, 9},
    {2, 0, 0, 3, 0, 0, 0, 0, 6},
    {0, 0, 0, 2, 0, 0, 9, 0, 0},
    {0, 0, 1, 9, 0, 4, 5, 7, 0},
  }

func main(){
  IsValid(&sudoku, 0)
  Display(sudoku)
}

func Display(sudoku[9][9] int){
  for row := 0; row < 9; row++ {
	
		for col := 0; col < 9; col++ {
			if col == 3 || col == 6 {		
			}
			fmt.Printf("%d ", sudoku[row][col])
			if col == 8 {		
			}
		}
		if row == 2 || row == 5 || row == 8 {
			fmt.Print("\n")
		} else {
			fmt.Println()
		}
	}

}

func AbsentOnLine(k int, sudoku [9][9]int, x int) bool {
  var y int
    for y=0; y < 9; y++ {
        if (sudoku[x][y] == k){
            return false
          }
        }
    return true
}

func AbsentOnRow(k int, sudoku [9][9]int, y int) bool {
  var x int
  for x=0; x < 9; x++{
       if (sudoku[x][y] == k){
           return false;
         }
       }
   return true;
}

func AbsentOnBloc(k int, sudoku [9][9]int, x int, y int) bool {
  var firstX, firstY int;
  firstX =  x-(x%3)
  firstY =  y-(y%3)
  for x = firstX; x < firstX+3; x++ {
        for y = firstY; y < firstY+3; y++ {
            if (sudoku[x][y] == k){
                return false;
              }
        }
      }
    return true;

}

func IsValid(sudoku *[9][9]int, position int) bool {

  if (position == 9*9){
        return true;
      }

      var x, y, k int

    x = position/9
    y = position%9

    if (sudoku[x][y] != 0){
        return IsValid(sudoku, position+1);
      }

    for k=1; k <= 9; k++ {
        if (AbsentOnLine(k,*sudoku,x) && AbsentOnRow(k,*sudoku,y) && AbsentOnBloc(k,*sudoku,x,y)){
            sudoku[x][y] = k;

            if (IsValid(sudoku, position+1)){
                return true;
              }
        }
    }
    sudoku[x][y] = 0;
    return false;

}






/*package main

import (
	"os"
	
  "fmt"

  "strings"
  piscine ".."
)

type CellID struct{ row, col int }

type Cell struct {
  ID           CellID
  Solution     int
  solved       bool
  alternatives []int
}
 var list []string
func main() {
  var square [9][9]Cell;

  
 
  
  
  
  
  

  
  
 

    square = parser(argumento(string(len(os.Args))))

 
  //printAlternatives(square);

  solution, _ := SolveSquare(square);
  
  PrintSquare(solution)
}

func CreateSquare() [9][9]Cell {
  var square [9][9]Cell;
  for i := 0; i < 9; i++ {
    for j := 0; j < 9; j++ {
      square[i][j] = Cell{
        ID: CellID{i, j},
        Solution:-1,
        solved:false,
        alternatives:[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
      };
    }
  }
  return square;
}
func argumento(a []string){
  for p,v:=range a {
    if p!=0{
      list=append(list, v)
    }
  }
}
func PrintSquare(square [9][9]Cell) {
  for i, row := range square {
    for j, cell := range row {
      if (cell.solved) {
        fmt.Printf(" %d ", cell.Solution)
      } 
      if j == 2 || j == 5 {
     
      }
    }
    fmt.Print("\n");
    if i == 2 || i == 5 {
      
    }
  }
}

func SolveSquare(square [9][9]Cell) ([9][9]Cell, bool) {
  for true {
    cell, min, found := lowerCellAlternatives(square)

    if !found {
      return square, true;
    }

    if min >= 2 {
      for _, alternative := range cell.alternatives {
        squareTemp, solved := SolveSquare(solveCell(square, cell.ID.row, cell.ID.col, alternative))
        if solved {
          return squareTemp, true
        }
      }
      return square, false
    }

    if ( min == 1) {
      square = solveCell(square, cell.ID.row, cell.ID.col, cell.alternatives[0]);
    }

    if (min == 0) {
      return square, false;
    }
  }

  return square, false;
}

func printAlternatives(square [9][9]Cell) {
  for _, row := range square {
    for _, cell := range row {
      fmt.Print(cell.alternatives)
      for i := len(cell.alternatives); i < 9; i++ {
        fmt.Print("  ")
      }
      if len(cell.alternatives) != 0 {
        fmt.Print(" ")
      }
    }
    fmt.Print("\n");
  }
}

func solveCell(square [9][9]Cell, row int, col int, solution int) [9][9]Cell {
  square[row][col].Solution = solution;
  square[row][col].solved = true;
  square[row][col].alternatives = nil;

  for j := 0; j < 9; j++ {
    square[row][j] = removeAlternative(square[row][j], solution)
  }

  for i := 0; i < 9; i++ {
    square[i][col] = removeAlternative(square[i][col], solution)
  }

  for _, cellID := range subSquare(row, col) {
    square[cellID.row][cellID.col] = removeAlternative(square[cellID.row][cellID.col], solution)
  }

  return square;
}

func removeAlternative(cell Cell, alternative int) Cell {
  var alternatives []int;
  for _, n := range cell.alternatives {
    if n != alternative {
      alternatives = append(alternatives, n);
    }
  }
  cell.alternatives = alternatives;
  return cell;
}

func lowerCellAlternatives(square [9][9]Cell) (Cell, int, bool) {
  min := 999;
  var lowerCell Cell;
  found := false;

  for _, row := range square {
    for _, cell := range row {
      if (!cell.solved && len(cell.alternatives) < min) {
        min = len(cell.alternatives);
        lowerCell = cell;
        found = true;
      }
    }
  }
  return lowerCell, min, found;
}

func between(min int, number int, max int) bool {
  return min <= number && number <= max;
}

func subSquare(row int, col int) [9]CellID {
  switch {
  //first col
  case between(0, row, 2) && between(0, col, 2):
    return [9]CellID{
      CellID{0, 0}, CellID{0, 1}, CellID{0, 2},
      CellID{1, 0}, CellID{1, 1}, CellID{1, 2},
      CellID{2, 0}, CellID{2, 1}, CellID{2, 2},
    }
  case between(3, row, 5) && between(0, col, 2):
    return [9]CellID{
      CellID{3, 0}, CellID{3, 1}, CellID{3, 2},
      CellID{4, 0}, CellID{4, 1}, CellID{4, 2},
      CellID{5, 0}, CellID{5, 1}, CellID{5, 2},
    }
  case between(6, row, 8) && between(0, col, 2):
    return [9]CellID{
      CellID{6, 0}, CellID{6, 1}, CellID{6, 2},
      CellID{7, 0}, CellID{7, 1}, CellID{7, 2},
      CellID{8, 0}, CellID{8, 1}, CellID{8, 2},
    }

  //center col
  case between(0, row, 2) && between(3, col, 5):
    return [9]CellID{
      CellID{0, 3}, CellID{0, 4}, CellID{0, 5},
      CellID{1, 3}, CellID{1, 4}, CellID{1, 5},
      CellID{2, 3}, CellID{2, 4}, CellID{2, 5},
    }
  case between(3, row, 5) && between(3, col, 5):
    return [9]CellID{
      CellID{3, 3}, CellID{3, 4}, CellID{3, 5},
      CellID{4, 3}, CellID{4, 4}, CellID{4, 5},
      CellID{5, 3}, CellID{5, 4}, CellID{5, 5},
    }
  case between(6, row, 8) && between(3, col, 5):
    return [9]CellID{
      CellID{6, 3}, CellID{6, 4}, CellID{6, 5},
      CellID{7, 3}, CellID{7, 4}, CellID{7, 5},
      CellID{8, 3}, CellID{8, 4}, CellID{8, 5},
    }

  //right col
  case between(0, row, 2) && between(6, col, 8):
    return [9]CellID{
      CellID{0, 6}, CellID{0, 7}, CellID{0, 8},
      CellID{1, 6}, CellID{1, 7}, CellID{1, 8},
      CellID{2, 6}, CellID{2, 7}, CellID{2, 8},
    }
  case between(3, row, 5) && between(6, col, 8):
    return [9]CellID{
      CellID{3, 6}, CellID{3, 7}, CellID{3, 8},
      CellID{4, 6}, CellID{4, 7}, CellID{4, 8},
      CellID{5, 6}, CellID{5, 7}, CellID{5, 8},
    }
  case between(6, row, 8) && between(6, col, 8):
    return [9]CellID{
      CellID{6, 6}, CellID{6, 7}, CellID{6, 8},
      CellID{7, 6}, CellID{7, 7}, CellID{7, 8},
      CellID{8, 6}, CellID{8, 7}, CellID{8, 8},
    }

  //default
  default:
    return [9]CellID{};
  }
}

func parser(text string) [9][9]Cell {
  square := CreateSquare();
  text = strings.TrimSpace(text);


  row, col := 0, 0;

  for _, char := range text {
    switch (char){
    case '1', '2', '3', '4', '5','6', '7', '8', '9':
      sol:= piscine.Atoi(string(char))
      square = solveCell(square, row, col, sol);
      col++

    case '\n':
      row++
      col = 0

    case '_':
      col++
    }
  }

  return square
}















































/*package main

int main(int argc, char* argv[]) {
    int grid[9][9];
    int input_error = 0;
    int i, j;
    if (argc != 1 + 9) { /* check number of rows in input 
        input_error = 1;
    } else {
        for (i = 0; i < 9; i++) { /* read each rows 
            if (strlen(argv[i + 1]) != 9) { /* check number of cols in input 
                input_error = 1;
                break;
            }
            for (j = 0; j < 9; j++) { /* read each cols 
                if (isdigit(argv[i + 1][j]) && argv[i + 1][j] != '0') {
                    /* digits except for 0 
                    grid[i][j] = argv[i + 1][j] - '0'; /* convert digit to integer 
                } else if (argv[i + 1][j] == '.') {
                    /* dot 
                    grid[i][j] = 0;
                } else {
                    /* invalid character 
                    input_error = 1;
                    break;
                }
            }
        }
    }
    /* check if some errors are detected 
    if (input_error) {
        fputs("invalid usage\n", stderr);
        return 1;
    }
    /* print what is read for testing 
    for (i = 0; i < 9; i++) {
        for(j = 0; j < 9; j++) {
            printf(" %d", grid[i][j]);
        }
        putchar('\n');
    }
    return 0;
}

/*import (
	"os"
	"log"
	"fmt"
	"strconv"
)

type SudokuBoard struct {
	squares [81]int
}

func main() {
	//Load the board
	var baseBoard SudokuBoard
	//baseBoard = LoadBoard()

	
		if len(os.Args) < 3 {
			fmt.Println("Invalid usage parameter:")
			fmt.Println("\tCorrect usage is: sudoku {size} {string}")
			fmt.Println("\t{size} denotes the board size, e.g. 9")
			fmt.Println("\t{string} denotes the initial solution and must be of length (size X size)")
			return
		}
		s, err :=strconv.Atoi(os.Args[1])
		if err != nil || (s != 2 && s != 4 && s != 9 && s != 16 && s != 25) {
			fmt.Printf("Invalid usage parameter: %s\n", os.Args[1])
			fmt.Println("\tCorrect usage is: sudoku {size} {string}")
			fmt.Println("\t{size} must be a valid integer of value {2,4,9,16,25}")
			return
		}
		init := os.Args[2]
		if len(init) != s*s {
			fmt.Println("Invalid usage parameter: {string}")
			fmt.Println("\tCorrect usage is: sudoku {size} {string}")
			fmt.Println("\t{string} must be of length (size X size)")
			return
		}
	
	
	
	//Check that it is valid
	if !ValidBoard(baseBoard) {
		log.Fatal("The board provided was invalid.")
	}

	//Create list of permutations, and one for solutions
	var permutationBoards []SudokuBoard //All boards must be valid, but not necessarily complete
	var solutionBoards []SudokuBoard

	permutationBoards = append(permutationBoards, baseBoard)

	for {
		if len(permutationBoards) == 0 {
			break
		}

		//Grab last board and remove from ones to work through
		currentBoard := permutationBoards[len(permutationBoards)-1]
		permutationBoards = permutationBoards[:len(permutationBoards)-1]

		filled := true
		//Find Empty square in this board
		for i:=0; i<81; i++ {
			if currentBoard.squares[i] == 0 {
				//fmt.Println("Ahh")
				filled = false

				//Fill empty square with all possible numbers, and check if its valid
				//if so, add it to the permutations
				for j:=1; j < 10; j++{
					newBoard := currentBoard
					newBoard.squares[i] = j

					//Check if we have broken any rules if so, skip this board
					if !ValidBoard(newBoard) {
						continue
					}

					//Board is valid so add it to the next ones we will look at
					permutationBoards = append( permutationBoards, newBoard )
					//fmt.Printf("%d boards to attempt!\n", len(permutationBoards))
				}
				break
			}
		}

		if filled {
			solutionBoards = append(solutionBoards, currentBoard)
		}
	}

	fmt.Printf("%d solutions found!\n", len(solutionBoards))
	for i := 0; i < len(solutionBoards); i++ {
		fmt.Printf("\nSolution: %d\n", i+1)
		for j := 0; j < 9; j++ {
			for k := 0; k < 9; k++ {
				fmt.Printf("%d ", solutionBoards[i].squares[9*j + k])
			}
			fmt.Println("")
		}
	}
}

func ValidBoard(board SudokuBoard) bool {
	//Check rows
	for i:=0; i<9; i++ {
		var numbersUsed [9]int
		for j:=0; j<9; j++ {
			if board.squares[9*i + j] != 0 {
				//fmt.Printf( "%d,%d value:%d\n", j, i, board.squares[9*i + j] )
				if numbersUsed[board.squares[9*i + j]-1] != 0 {
					//fmt.Printf( "Invalid at %d,%d, value:%d\n", j, i, board.squares[9*i + j] )
					return false
				}
				numbersUsed[board.squares[9*i + j]-1]++
			}
		}
	}

	//Check columns
	for i:=0; i<9; i++ {
		var numbersUsed [9]int
		for j:=0; j<9; j++ {
			if board.squares[9*j + i] != 0 {
				if numbersUsed[board.squares[9*j + i]-1] != 0 {
					//fmt.Printf( "Invalid at %d,%d, value:%d\n", j, i, board.squares[9*i + j] )
					return false
				}
				numbersUsed[board.squares[9*j + i]-1]++
			}
		}
	} 


	//Check squares
	for i:=0; i<9; i++ {
		var numbersUsed [9]int
		for j:=0; j<9; j++ {
			x:= 3*(i%3)+ j%3
			y:= 3*(i/3)+ j/3

			if board.squares[9*y + x] != 0 {
				if numbersUsed[board.squares[9*y + x]-1] != 0 {
					//fmt.Printf( "Invalid at %d,%d, value:%d\n", j, i, board.squares[9*i + j] )
					return false
				}
				numbersUsed[board.squares[9*y + x]-1]++
			}
		}
	}

	return true
}

/*func LoadBoard() SudokuBoard {
	var board SudokuBoard

	//Row 1
	board.squares[1] = 9
	board.squares[2] = 6
	board.squares[4] = 4
	board.squares[8] = 1
	//Row 2
	board.squares[9] = 1
	board.squares[13] = 6
	board.squares[15] = 1
	board.squares[17] = 4
	//Row 3
	board.squares[18] = 5
	board.squares[20] = 4
	board.squares[21] = 8
	board.squares[22] = 1
	board.squares[24] = 3
	board.squares[25] = 9
	//Row 4
	board.squares[29] = 7
	board.squares[30] = 9
	board.squares[31] = 5
	board.squares[34] = 4
	board.squares[35] = 3
	//Row 5
	board.squares[37] = 3
	board.squares[40] = 8
	//Row 6
	
	board.squares[45] = 4
	board.squares[47] = 7
	board.squares[49] = 6
	board.squares[51] = 6
	//Row 7
	board.squares[53] = 9
	board.squares[55] = 4
	board.squares[58] = 6
	//Row 8
	board.squares[60] = 1
	board.squares[62] = 4
	board.squares[64] = 6

	//Row 9
	board.squares[68] = 1
	board.squares[69] = 4
	board.squares[70] = 6
	board.squares[74] = 7


	return board
}*/