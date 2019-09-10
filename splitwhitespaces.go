 package piscine

func SplitWhiteSpaces(str string) []string {
	
	TextToString := ""
	t := []string{}
	for i, v := range str {
		if i == lent2(str)-1 && string(v) != " " && string(v) != "\t" && string(v) != "\n" {
			TextToString += string(v)
			t = append(t, TextToString)
		}else if string(v) != " " && string(v) != "\t" && string(v) != "\n" {
				TextToString += string(v)
		}else{
			if len(TextToString)>=1{
				t = append(t, TextToString)
			}
			TextToString = ""
		}
	}
	return t
}