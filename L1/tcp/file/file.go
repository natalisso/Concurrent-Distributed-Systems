package main
import "os";


func main() {
	// open output file
    fo, err := os.Create("input.txt")
    if err != nil {
        panic(err)
    }
    // close fo on exit and check for its returned error
    defer func() {
        if err := fo.Close(); err != nil {
            panic(err)
        }
	}()
	
	for i:= 0; i < 1E4; i++{
		// write a chunk
        if _, err := fo.Write([]byte("Hi\n")); err != nil {
            panic(err)
        }
	}
}