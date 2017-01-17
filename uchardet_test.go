package uchardet

import "testing"

import "io/ioutil"

func TestChardet(t *testing.T) {
	det := NewChardet()
	//
	bys, _ := ioutil.ReadFile("utf8.txt")
	println(string(bys))
	det.Handle(bys)
	enc := det.End()
	if enc != "UTF-8" {
		t.Error("error")
		println(enc)
		return
	}
	println(enc)
	//
	bys, _ = ioutil.ReadFile("gbk.txt")
	println(string(bys))
	det.Reset()
	det.Handle(bys)
	enc = det.End()
	if enc != "GB18030" {
		t.Error("error")
		println(enc)
		return
	}
	println(enc)
}
