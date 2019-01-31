package FileHelper

import (
	"bytes"
	"compress/gzip"
	"io"
	"io/ioutil"
	"os"
)

func GzipEncode(in []byte) ([]byte, error) {
	var (
		buffer bytes.Buffer
		out    []byte
		err    error
	)
	writer := gzip.NewWriter(&buffer)
	_, err = writer.Write(in)
	if err != nil {
		writer.Close()
		return out, err
	}
	err = writer.Close()
	if err != nil {
		return out, err
	}

	return buffer.Bytes(), nil
}

func GzipDecode(in []byte) ([]byte, error) {
	reader, err := gzip.NewReader(bytes.NewReader(in))
	if err != nil {
		var out []byte
		return out, err
	}
	defer reader.Close()

	return ioutil.ReadAll(reader)
}

//压缩文件Src到Dst
func CompressFile(Dst string, Src string) error {
	newfile, err := os.Create(Dst)
	if err != nil {
		return err
	}
	defer newfile.Close()

	file, err := os.Open(Src)
	if err != nil {
		return err
	}

	zw := gzip.NewWriter(newfile)

	filestat, err := file.Stat()
	if err != nil {
		return nil
	}

	zw.Name = filestat.Name()
	zw.ModTime = filestat.ModTime()
	_, err = io.Copy(zw, file)
	if err != nil {
		return nil
	}

	zw.Flush()
	if err := zw.Close(); err != nil {
		return nil
	}
	return nil
}

//解压文件Src到Dst
/*
example:

	pwd, _ := os.Getwd()

	src := pwd + "/D175021_20180623_1442_bis.csv.gz"
	fmt.Println("src=", src)

	dst := pwd + "/D175021_20180623_1442_bis.csv"
	fmt.Println("dst=", dst)

	err := FileHelper.DeCompressFile(src, dst)
*/
func DeCompressFile(Src string, Dst string) error {
	file, err := os.Open(Src)
	if err != nil {
		return err
	}
	defer file.Close()

	newfile, err := os.Create(Dst)
	if err != nil {
		return err
	}
	defer newfile.Close()

	zr, err := gzip.NewReader(file)
	if err != nil {
		return err
	}

	filestat, err := file.Stat()
	if err != nil {
		return err
	}

	zr.Name = filestat.Name()
	zr.ModTime = filestat.ModTime()
	_, err = io.Copy(newfile, zr)
	if err != nil {
		return err
	}

	if err := zr.Close(); err != nil {
		return err
	}
	return nil
}
