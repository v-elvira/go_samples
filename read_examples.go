// os.File.Read

file, err := os.Open("awesome.txt") // go is awesome
if err != nil {
	panic(err)
}
defer file.Close()

buf := make([]byte, 5)
for {
	n, err := file.Read(buf)
	fmt.Println(n, err)
	if err == io.EOF {
		break
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("read %d bytes: %q\n", n, buf[:n])
	// 5 <nil>
	// read 5 bytes: "go is"
	// 5 <nil>
	// read 5 bytes: " awes"
	// 3 <nil>
	// read 3 bytes: "ome"
	// 0 EOF
}

//-------------------------------------------------------------------
// bufio.Reader (disk -> inner buffer 4096 b ->)

file, err := os.Open("awesome.txt")
if err != nil {
	panic(err)
}
defer file.Close()

reader := bufio.NewReader(file) // (1)
buf := make([]byte, 5)
for {
	n, err := reader.Read(buf)
	if err == io.EOF {
		break
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("read %d bytes: %q\n", n, buf[:n])
}
// read 5 bytes: "go is"
// read 5 bytes: " awes"
// read 3 bytes: "ome"

//---------------------------------------------------
// bufio.Scanner for text

file, err := os.Open("people.txt")
if err != nil {
	panic(err)
}
defer file.Close()

scanner := bufio.NewScanner(file)
// scanner.Split(bufio.ScanWords) // default
// scanner.Split(bufio.ScanWords)
// scanner.Split(bufio.ScanRunes)
// scanner.Split(bufio.ScanBytes)
for scanner.Scan() {
	line := scanner.Text()
	fmt.Printf("%#v\n", line)
}

if err := scanner.Err(); err != nil {
	panic(err)
}
