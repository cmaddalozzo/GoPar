package main

import(
  "fmt"
  "./par2"
	"flag"
)

func main(){
	flag.Parse()
	fname := flag.Arg(0)
	parfile, err := par2.Open(fname)
	if err != nil{
		fmt.Printf("error: ", err.String())
		return
	}
	mainpacket := parfile.GetMainPacket()
	fmt.Printf("packet size: %d\n", mainpacket.Header.Length)
	fmt.Printf("packet hash: %d\n", mainpacket.Header.P_hash)
	fmt.Printf("slice size: %d\n", mainpacket.Slice_size, "\n")
	fmt.Printf("num files: %d\n", mainpacket.Num_files, "\n")	
}
