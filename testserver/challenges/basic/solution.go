package basic
import "fmt"

func ReverseString(str string) string {
	r := make([]byte, len(str))
	for i := 0; i < len(str); i++ {
		r[i] = str[len(str)-i-1]
	}
	fmt.Println("Gofcc")

	return ""
}