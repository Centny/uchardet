uchardet binding by go
======

### Install

<br/>
`linux/unix/window`

* install `uchardet` by `yum` or compile source

* install gosigar by command

```
export CGO_CPPFLAGS="-I/usr/local/include"
export CGO_LDFLAGS="-L/usr/local/lib -luchardet"
go get github.com/Centny/uchardet
```

* test uchardet

```
go test -v github.com/Centny/uchardet
```

note: adding libuchardet install path to LD_LIBRARY_PATH




### Example

```
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
```