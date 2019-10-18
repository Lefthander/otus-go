package main

import (
	"flag"
	"io"
	"log"
	"os"

	"github.com/cheggaaa/pb/v3"
)

var (
	source      string
	destination string
	offset      int64
	limit       int64
)

func init() {

	flag.StringVar(&source, "from", "", "Source file")
	flag.StringVar(&destination, "to", "", "Destination file")
	flag.Int64Var(&offset, "offset", 0, "Offset in the source file")
	flag.Int64Var(&limit, "limit", 0, "Limit of bytes to be copied")

}

func validateArgs() {
	flag.Parse()
	if source == "" && destination == "" && source == destination {
		log.Println("Wrong parameters...")
		flag.Usage()
		os.Exit(1)
	}

	if limit < 0 {
		log.Println("Error - negative limit does not accepted")
		flag.Usage()
		os.Exit(1)
	}
}

const (
	// DefaultBlockSize for dd tool is 512 but modern HDD/SDD use 4096
	DefaultBlockSize = 4096
)

func goCopy(source string, destination string, offset int64, limit int64) (int64, error) {
	// It is assumed that input parameters are validated outside the goCopy function
	// seek to the desired offset in the source file

	log.Println("Offset=", offset)
	log.Println("Limit=", limit)

	src, err := os.Open(source)

	if err != nil {
		return 0, err
	}
	defer src.Close()

	dst, err := os.Create(destination)

	if err != nil {

		return 0, err
	}

	defer dst.Close()

	if offset > 0 {
		if _, err := src.Seek(offset, 0); err != nil {
			return 0, err
		}
	} else if offset < 0 {
		if _, err := src.Seek(offset, 0); err != nil {
			return 0, err
		}
	}
	sourcefileinfo, err := src.Stat()
	var totalamountofbytes int64
	if sourcefileinfo.Size() < limit {
		totalamountofbytes = sourcefileinfo.Size() - offset
	} else if limit != 0 {
		totalamountofbytes = limit - offset
	} else {
		totalamountofbytes = sourcefileinfo.Size() - offset
	}

	log.Println("Souce file size", totalamountofbytes)

	var blocksize int64

	blocksize = DefaultBlockSize // Set the Default value 4096

	if limit-offset < blocksize {
		blocksize = limit - offset
	}

	log.Println("Block size", blocksize)

	buffer := make([]byte, blocksize)

	actuallimit := limit - offset

	if limit == 0 {
		actuallimit = totalamountofbytes
	}

	log.Println("Actual Limit", actuallimit)

	lreader := io.LimitReader(src, actuallimit)

	pbar := pb.Start64(totalamountofbytes)
	pbar.SetWidth(100)
	pbar.Set(pb.Bytes, true)

	for {

		bytesread, errRead := lreader.Read(buffer)

		//log.Println(bytesread)

		if bytesread > 0 {
			_, errWrite := dst.Write(buffer[:bytesread])
			if errWrite != nil {
				return 0, errWrite
			}
			pbar.Add(bytesread)

		}
		if errRead != nil {
			pbar.Finish()
			return 0, errRead
		}

	}

}

func main() {

	validateArgs()
	cb, err := goCopy(source, destination, offset, limit)

	if err != nil {
		if err == io.EOF {
			log.Println("Succesfully done...", cb)
			os.Exit(0)
		}
		log.Println("Error occured during copy...", err)
		os.Exit(1)
	}

}
