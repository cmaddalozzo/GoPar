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
	//parfile.PopulateMainPacket()
  mainpacket := parfile.Mainpacket
	fmt.Printf("packet size: %d\n", mainpacket.Header.Length)
	fmt.Printf("packet hash: %s\n", mainpacket.Header.P_hash)
	fmt.Printf("slice size: %d\n", mainpacket.Slice_size)
	fmt.Printf("Num files: %d\n", mainpacket.Num_files)	
	fmt.Printf("File ids:\n")	
  for i := 0; i < len(mainpacket.R_file_ids); i++{
    fmt.Printf("\t%s\n", mainpacket.R_file_ids[i])
  }
}
