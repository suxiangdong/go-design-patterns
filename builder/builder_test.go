package builder

import "testing"

func TestComputerA(t *testing.T) {
	computerA := &ComputerBuilderA{}
	director := NewDirector(computerA)
	if r := director.CreateComputer(); r != "ram_AhardDisk_Amonitor_A" {
		t.Fatal("ComputerA test fail")
	}
}

func TestComputerB(t *testing.T) {
	computerB := &ComputerBuilderB{}
	director := NewDirector(computerB)
	if r := director.CreateComputer(); r != "ram_BhardDisk_Bmonitor_B" {
		t.Fatal("ComputerB test fail")
	}
}
