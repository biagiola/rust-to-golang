func countdown(seconds int) {
	if seconds == 0 {
		fmt.Println("Blastoff..!!")
	} else {
		fmt.Printf("%d seconds to blastoff..!!")
		countdown(seconds - 1)
	}
}
func main() {
	countdown(5)
}