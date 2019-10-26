package main

import (
	"flag"
	"io"
	"log"
	"os"

	"github.com/cheggaaa/pb/v3"
)

// Create a random file with specific size - Powershell Windows
// $out = new-object byte[] 50G; (new-object Random).NextBytes($out); [IO.File]::WriteAllBytes('.\test.dat', $out)
// Create a random file with specific size - Linux / Mac OS
// head -c 8388608 </dev/urandom >myfile

var (
	source      string // sorce file
	destination string // destination file
	offset      int64  // offset in the source file
	limit       int64  // limit of bytes to copy
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

func copier(rs io.ReadSeeker, w io.Writer, srcSize int64, ofs int64, lm int64) error {

	if ofs > 0 {
		if _, err := rs.Seek(ofs, io.SeekStart); err != nil {
			log.Println("Seek Error", err)
			return err
		}
	} else if ofs < 0 {
		if _, err := rs.Seek(ofs, io.SeekEnd); err != nil {
			return err
		}
	}

	var totalamountofbytes int64
	// srcSize = 10, lm = 5 , ofs = 5
	switch {
	case ofs < 0 && lm == 0:
		totalamountofbytes = srcSize + ofs
	case ofs < 0 && lm > 0:
		totalamountofbytes = lm
	case srcSize < lm:
		totalamountofbytes = srcSize - offset
	case lm != 0 && ofs != 0 && lm == ofs:
		totalamountofbytes = lm
	case lm != 0 && lm > ofs:
		totalamountofbytes = lm - ofs
	default:
		totalamountofbytes = srcSize - offset
	}
	log.Println("Size of source", srcSize)
	log.Println("Offset", ofs)
	log.Println("Limit", lm)
	log.Println("totalamount of bytes", totalamountofbytes)
	buffer := make([]byte, DefaultBlockSize)

	lmr := io.LimitReader(rs, totalamountofbytes)

	pbar := pb.Start64(totalamountofbytes)
	pbar.SetWidth(100)
	pbar.Set(pb.Bytes, true)

	for {

		bytesread, errRead := lmr.Read(buffer)

		if bytesread > 0 {
			_, errWrite := w.Write(buffer[:bytesread])
			if errWrite != nil {
				return errWrite
			}
			pbar.Add(bytesread)

		}
		if errRead != nil {
			pbar.Finish()
			return errRead
		}

	}

}

func goCopy(source string, destination string, offset int64, limit int64) (int64, error) {
	// It is assumed that input parameters are validated outside the goCopy function
	// seek to the desired offset in the source file
	// 1. Open Source file & check err
	// 2. Open Destination file & check err
	// 3. Calculate the total amount of bytes to be copied
	// 4. Set offset in the source file in accordance with input data
	// 5. Initilize the progress bar
	// 6. Initilize the io.reader with specific limit in accordance with input parameters
	// 7. Read in loop the source file until EOF or error
	// 8. Finalize the progress bar

	// Debug printouts will be removed.
	log.Println("Offset=", offset)
	log.Println("Limit=", limit)

	// Open the source file
	src, err := os.Open(source)

	if err != nil {
		return 0, err
	}
	defer src.Close() // Don't forget to close it

	// Open the destination file
	dst, err := os.Create(destination)

	if err != nil {

		return 0, err
	}

	defer dst.Close() // Don't forget to close it
	// Get sourece file information struct in order to determine the size of file
	sourcefileinfo, err := src.Stat()

	sourcefilesize := sourcefileinfo.Size()

	// Run the copier within input data...
	err = copier(src, dst, sourcefilesize, offset, limit)

	if err != nil {
		return 0, err
	}

	return 0, nil
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
