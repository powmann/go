package main

func main() {

	// exemplo de if
	if 1 == 1 {
		println("1 == 1")
	}

	// exemplo de if else
	if 1 == 2 {
		println("1 == 2")
	} else {
		println("1 != 2")
	}

	// exemplo de switch
	switch "go" {
	case "java":
		println("java")
	case "go":
		println("go")
	default:
		println("default")
	}

	// exemplo de switch sem condicao
	switch {
	case 1 == 1:
		println("1 == 1")
	case 1 == 2:
		println("1 == 2")
	default:
		println("default")
	}

	// exemplo de for
	for i := 0; i < 3; i++ {
		println(i)
	}

	// exemplo de while
	i = 0
	for i < 3 {
		println(i)
		i++
	}
	println()

	// exemplo de for sem condicao equivalente ao while true
	i := 0
	for {
		if i == 3 {
			break
		}
		println(i)
		i++
	}
	println()

	// exemplo de for equivalente ao until
	i := 0
	for i != 3 {
		println(i)
		i++
	}
	println()

	// exemplo de for range
	var1 := []string{"a", "b", "c"}
	for k, v := range var1 {
		println(k, v)
	}
	println()

	// exemplo de for range sem key
	var2 := []string{"a", "b", "c"}
	for _, v := range var2 {
		println(v)
	}
	println()

	// exemplo de for range sem value
	var3 := []string{"a", "b", "c"}
	for k, _ := range var3 {
		println(k)
	}
	println()

	// exemplo de for range sem key e value
	var4 := []string{"a", "b", "c"}
	for range var4 {
		println("oi")
	}
	println()

	// exemplo de for range em map
	var5 := map[string]string{"a": "alpha", "b": "beta"}
	for k, v := range var5 {
		println(k, v)
	}
	println()

	// exemplo de for range em map sem key
	var6 := map[string]string{"a": "alpha", "b": "beta"}
	for _, v := range var6 {
		println(v)
	}
	println()

	// exemplo de for range em map sem value
	var7 := map[string]string{"a": "alpha", "b": "beta"}
	for k, _ := range var7 {
		println(k)
	}
	println()

	// exemplo de for range em map sem key e value
	var8 := map[string]string{"a": "alpha", "b": "beta"}
	for range var8 {
		println("yohey")
	}
	println()

	// exemplo de for range em string
	var9 := "abc"
	for k, v := range var9 {
		println(k, string(v))
	}
	println()

	// exemplo de for range em string sem key
	var10 := "abc"
	for _, v := range var10 {
		println(string(v))
	}
	println()

	// exemplo de for range em string sem value
	var11 := "abc"
	for k, _ := range var11 {
		println(k)
	}
	println()

	// exemplo de for range em string sem key e value
	var12 := "abc"
	for range var12 {
		println("yey")
	}
	println()

	// exemplo de for range em array
	var13 := [3]string{"a", "b", "c"}
	for k, v := range var13 {
		println(k, v)
	}
	println()

	// exemplo de for range em array sem key
	var14 := [3]string{"a", "b", "c"}
	for _, v := range var14 {
		println(v)
	}
	println()

	// exemplo de for range em array sem value
	var15 := [3]string{"a", "b", "c"}
	for k, _ := range var15 {
		println(k)
	}
	println()

	// exemplo de for range em array sem key e value
	var16 := [3]string{"a", "b", "c"}
	for range var16 {
		println("yey")
	}
	println()

}
