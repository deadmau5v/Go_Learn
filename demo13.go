package main

func main13() {
	var map1 map[int]string = map[int]string{
		1: "1",
		2: "hello world",
	}
	println(map1)
	if map1 == nil {
		println("map为空")
	}
	for i := 0; i < 1000; i++ {
		map1[i] = "hello world"
	}

	var map2 map[string]string = map[string]string{
		"首都":  "北京",
		"一线":  "上海",
		"乡里别": "长沙",
	}

	for k, v := range map2 {
		println(k, " : ", v)
		delete(map2, k)
	}
	updateMap(map2)
	println(len(map2))

}

func updateMap(map1 map[string]string) map[string]string {
	map1["已修改"] = "true"
	return map1
}
