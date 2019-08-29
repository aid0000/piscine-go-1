package piscine
func AppendRange(min, max int) []int {

		var resultado []int

			for i:=min-1;i< max-1; i++{

				resultado=append(resultado, i+1)
			}
					return resultado
}
