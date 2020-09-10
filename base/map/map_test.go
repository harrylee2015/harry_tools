package _map

import "testing"

func TestInitMap(t *testing.T){
	bigMap := make(map[int]int,10)
	for i:=0;i<20;i++{
		bigMap[i]=i
	}
	for k,v :=range bigMap{
		t.Log(k,v)
	}
}

func BenchmarkMap(b *testing.B){
	bigMap := make(map[int]int)
	b.ResetTimer()
	for i:=0;i<b.N;i++{
		bigMap[i]=i
	}
}
func BenchmarkInitMap(b *testing.B){
	size :=b.N
	bigMap := make(map[int]int,size)
	b.ResetTimer()
	for i:=0;i<size;i++{
		bigMap[i]=i
	}
}