package main

// operadores de atribuicao e declaracao de variaveis
var var1 int = 1 // operador de atribuicao simples
var2 := 2 // operador de atribuicao por inferencia de tipo
var3, var4 := 3, 4 // operador de atribuicao multipla
var5, var6 int // declaracao de variaveis
var7, var8 int = 7, 8 // declaracao e atribuicao de variaveis
var9, var10 := 9, 10 // declaracao e atribuicao por inferencia de tipo

// operadores aritmeticos
var11 := var1 + var2 // soma
var12 := var1 - var2 // subtracao
var13 := var1 * var2 // multiplicacao
var14 := var1 / var2 // divisao
var15 := var1 % var2 // resto da divisao

// operadores relacionais
var16 := var1 == var2 // igualdade
var17 := var1 != var2 // diferenca
var18 := var1 > var2 // maior que
var19 := var1 < var2 // menor que
var20 := var1 >= var2 // maior ou igual que
var21 := var1 <= var2 // menor ou igual que

// operadores logicos
var22 := var1 && var2 // and
var23 := var1 || var2 // or
var24 := !var1 // not

// operadores de incremento e decremento
var25 := var1++ // incremento
var26 := var1-- // decremento

// operadores de atribuicao
var27 += var1 // atribuicao de soma
var28 -= var1 // atribuicao de subtracao
var29 *= var1 // atribuicao de multiplicacao
var30 /= var1 // atribuicao de divisao
var31 %= var1 // atribuicao de resto da divisao
var32 <<= var1 // atribuicao de deslocamento a esquerda
var33 >>= var1 // atribuicao de deslocamento a direita
var34 &= var1 // atribuicao de and
var35 |= var1 // atribuicao de or
var36 ^= var1 // atribuicao de xor

// operadores de ponteiros
var37 := &var1 // endereco de memoria
var38 := *var37 // valor do endereco de memoria

// operadores de deslocamento
var39 := var1 << var2 // deslocamento a esquerda
var40 := var1 >> var2 // deslocamento a direita

// operadores de tipo
var41 := var1.(int) // conversao de tipo

// operadores de canal
var42 := <-var1 // leitura de canal
var43 <- var1 // escrita de canal

// operadores de indexacao
var44 := var1[var2] // indexacao de array
var45 := var1[var2:var3] // indexacao de slice
var46 := var1[var2:var3:var4] // indexacao de slice com capacidade

// operadores de selecao
var47 := var1.var2 // selecao de campo
var48 := var1.var2() // chamada de metodo
var49 := var1[0] // indexacao de array
var50 := var1[0:1] // indexacao de slice
var51 := var1[0:1:2] // indexacao de slice com capacidade

