package par2

import(
	"os"
	//"fmt"
	//"strings"
	"bytes"
	//"strconv"
	"encoding/hex"
	//"crypto/md5"
	"encoding/binary"
)
const (
	_magic_sequence = "PAR2\000PKT"
	_typestring_main = "PAR 2.0\x00Main\x00\x00\x00\x00"
	_headersize = 64
)	
type Par2File struct{
	filename string
	fhandle * os.File
  Mainpacket * MainPacket
  Fdpackets []* FDPacket
}

type PacketHeader struct{
	Length uint8		// total length of the packet in bytes
	P_hash string		// md5 hash of the packet
	Rs_id string	  // md5 hash of the _body_ of the main packet
	P_type string	  // 16-byte string
}
type MainPacket struct{
	Header * PacketHeader
	Slice_size uint8
	Num_files uint8
	R_file_ids []string
	Nr_file_ids []string
}
type FDPacket struct{
	header PacketHeader
	file_id string
	file_hash string
	file_md516k string
	file_length int
	file_name string
}
type ISCSPacket struct{
	header PacketHeader
	file_id string
	md5crc_pairs [][]string
}
type Packet interface{
	GetHeader() (header PacketHeader)
}
func Open(filename string) (par2file * Par2File , err os.Error){	
	fhandle, err := os.Open(filename)
	if err != nil{
		return nil, err
	}
  par2file = &Par2File{filename:filename, fhandle:fhandle}
	info, err := par2file.fhandle.Stat()
	if err != nil{
		return nil, err
	}
	fsize := info.Size
	filedata := make([]byte, fsize)
	numbytes, _ := par2file.fhandle.Read(filedata)
	if int64(numbytes) != fsize{
		return nil, err
	}
  par2file.Mainpacket = extractMainPacket(filedata)//PopulateMainPacket()
	return par2file, nil
}
/*func ( par2file * Par2File) PopulateMainPacket(){
	info, err := par2file.fhandle.Stat()
	if err != nil{
		return //nil
	}
	fsize := info.Size
	filedata := make([]byte, fsize)
	numbytes, _ := par2file.fhandle.Read(filedata)
	if int64(numbytes) != fsize{
		return //nil
	}
	type_start := bytes.Index(filedata, []byte(_typestring_main))

	if type_start < 0{
		return //nil
	}
	mainpacket := new(MainPacket)
	//the start of the header is 48 bytes behind the type string
	header_start := type_start - 48
	mainpacket.Header = parseHeader(filedata[header_start:header_start + _headersize])
	mainpacket.parse(filedata[header_start + _headersize:header_start + _headersize + int(mainpacket.Header.Length)])
  par2file.Mainpacket = mainpacket
	//return mainpacket;
}
  */
func (par2file * Par2File) String() string{
	return par2file.filename
}
func extractMainPacket(filedata []byte) (mainpacket * MainPacket){
  //Find the start of main packet type string
	type_start := bytes.Index(filedata, []byte(_typestring_main))

	if type_start < 0{
		return nil
	}
	mainpacket = new(MainPacket)
	//the start of the header is 48 bytes in front of the type string
	header_start := type_start - 48
	mainpacket.Header = parseHeader(filedata[header_start:header_start + _headersize])
	//mainpacket.parse(filedata[header_start + _headersize:header_start + _headersize + int(mainpacket.Header.Length)])
  packet_body := filedata[header_start + _headersize:header_start + _headersize + int(mainpacket.Header.Length)]
	//read slice size (8 bytes)
	binary.Read(bytes.NewBuffer(packet_body[0:8]), binary.LittleEndian, &mainpacket.Slice_size)
	//read number of files in set (4 bytes)
	binary.Read(bytes.NewBuffer(packet_body[8:12]), binary.LittleEndian, &mainpacket.Num_files)
  mainpacket.R_file_ids = make([]string, mainpacket.Num_files)
  for i := 0; i < int(mainpacket.Num_files); i++{
    start := 12 + (i * 16)
    end := start + 16
	  mainpacket.R_file_ids[i] = hex.EncodeToString(packet_body[start:end])
  }
  return mainpacket

}
func parseHeader(header_data []byte) (header * PacketHeader){
	//var length uint8
	header = new(PacketHeader)
	//grab the size of the packet in bytes
	binary.Read(bytes.NewBuffer(header_data[8:16]), binary.LittleEndian, &header.Length)
	//grab the hash of the packet
	header.P_hash = hex.EncodeToString(header_data[16:32])
	//grab the recovery set id
	header.Rs_id = hex.EncodeToString(header_data[32:48])
	return header
}
/*func (mainpacket * MainPacket) parse(packet_body []byte){
	//read slice size (8 bytes)
	binary.Read(bytes.NewBuffer(packet_body[0:8]), binary.LittleEndian, &mainpacket.Slice_size)
	//read number of files in set (4 bytes)
	binary.Read(bytes.NewBuffer(packet_body[8:12]), binary.LittleEndian, &mainpacket.Num_files)
  mainpacket.R_file_ids = make([]string, mainpacket.Num_files)
  for i := 0; i < int(mainpacket.Num_files); i++{
    start := 12 + (i * 16)
    end := start + 16
	  mainpacket.R_file_ids[i] = hex.EncodeToString(packet_body[start:end])
  }
}
*/












