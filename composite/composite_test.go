package composite

import "testing"

func TestComposite(t *testing.T) {
	root := &Directory{
		Name: "根目录",
	}

	n1 := &Directory{Name: "文件夹1"}
	n2 := &Directory{Name: "文件夹2"}
	f := &File{Name: "文件"}
	f1 := &File{Name: "文件1"}
	root.Add(n1)
	root.Add(f)
	root.Add(n2)
	n1.Add(f1)
	root.Display(0)
}
