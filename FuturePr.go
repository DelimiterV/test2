// FuturePr.go
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/charmap"
)

func FirstByteCompare(src []byte, mask []byte) bool {
	var rez bool = true
	if len(src) >= len(mask) {
		for i := 0; i < len(mask) && rez == true; i++ {
			if src[i] != mask[i] {
				rez = false
			}
		}
	} else {
		rez = false
	}
	return rez
}

func copyBuf(dst []byte, src []byte, offset int, len int) {
	for i := 0; i < len; i++ {
		dst[offset+i] = src[i]
	}
}

type ffl struct {
	filename string
}

var files []ffl

func DecodeWindows1251(ba []uint8) []uint8 {
	dec := charmap.Windows1251.NewDecoder()
	out, _ := dec.Bytes(ba)
	return out
}

func EncodeWindows1251(ba []uint8) []uint8 {
	enc := charmap.Windows1251.NewEncoder()
	out, _ := enc.String(string(ba))
	return []uint8(out)
}

const NPLACE = 8

type pos struct {
	kolvo int
}
type lenter struct {
	rstr    string
	nstr    int
	pg      int
	vid     string
	id      int
	artikul string
	brand   string
	model   string
	hit     string
	li      int
	data    [NPLACE]pos
}

var myenter []lenter
var kenter []lenter

func KIniEnter(pathToFile string) bool {
	var ex int
	var rez bool = true
	var un lenter
	_, err := os.Stat(pathToFile)
	if err == nil {
		f, err := os.Open(pathToFile)
		if err == nil {
			fi, _ := f.Stat()
			buf := make([]uint8, fi.Size()+100)

			nr, _ := f.Read(buf)
			k := 1
			l := 0

			for i := 0; i < nr && ex == 0; i++ {
				switch buf[i] {
				case 0x0d:
					if i > 0 {
						k = i
					}
				case 0x0a:
					if k > l {
						b := buf[l:k]
						bUTF := DecodeWindows1251(b)
						mstr := string(bUTF)
						// ----------- обработка ----------------
						rstr := strings.Split(mstr, "\t")
						if len(rstr) > 8 {
							un.rstr = mstr
							un.nstr, _ = strconv.Atoi(rstr[0])
							if rstr[1] == "PG" {
								un.pg = -1
							} else {
								un.pg, _ = strconv.Atoi(rstr[1])
							}
							un.nstr, _ = strconv.Atoi(rstr[0])
							un.artikul = rstr[2]
							un.brand = rstr[3]
							un.model = rstr[4]
							//fmt.Println(un.model)
							un.vid = rstr[5]
							un.hit = rstr[6]
							un.li, _ = strconv.Atoi(rstr[7])
							un.id, _ = strconv.Atoi(rstr[8])
							for v := 0; v < NPLACE; v++ {
								un.data[v].kolvo = 0
							}
							kenter = append(kenter, un)
							//fmt.Print(".")
						} else {

						}

					}
					l = i + 1
				default:
				}
			}
			f.Close()
		} else {
			panic(err)
			rez = false
		}
	} else {
		fmt.Println("Enter отсутствует или открыт другой программой")
		rez = false
	}
	return rez
}
func IniEnter(pathToFile string, enter *[]lenter) bool {
	var ex int
	var rez bool = true
	var un lenter
	_, err := os.Stat(pathToFile)
	if err == nil {
		f, err := os.Open(pathToFile)
		if err == nil {
			fi, _ := f.Stat()
			buf := make([]uint8, fi.Size()+100)

			nr, _ := f.Read(buf)
			k := 1
			l := 0

			for i := 0; i < nr && ex == 0; i++ {
				switch buf[i] {
				case 0x0d:
					if i > 0 {
						k = i
					}
				case 0x0a:
					if k > l {
						b := buf[l:k]
						bUTF := DecodeWindows1251(b)
						mstr := string(bUTF)
						// ----------- обработка ----------------
						rstr := strings.Split(mstr, "\t")
						if len(rstr) > 8 {
							un.rstr = mstr
							un.nstr, _ = strconv.Atoi(rstr[0])
							if rstr[1] == "PG" {
								un.pg = -1
							} else {
								un.pg, _ = strconv.Atoi(rstr[1])
							}
							un.nstr, _ = strconv.Atoi(rstr[0])
							un.artikul = rstr[2]
							un.brand = rstr[3]
							un.model = rstr[4]
							un.vid = rstr[5]
							un.hit = rstr[6]
							un.li, _ = strconv.Atoi(rstr[7])
							un.id, _ = strconv.Atoi(rstr[8])
							for v := 0; v < NPLACE; v++ {
								un.data[v].kolvo = 0
							}
							*enter = append(*enter, un)
							//fmt.Print(".")
						} else {

						}

					}
					l = i + 1
				default:
				}
			}
			f.Close()
			fmt.Println("Проверка дубликатов!")
			fmt.Println(">>>>>>>>>>>>>>>>>>>>>")
			for i := 0; i < len(*enter)-1; i++ {
				if (*enter)[i].id > 0 {
					for j := i + 1; j < len(*enter); j++ {
						if (*enter)[i].id == (*enter)[j].id {
							fmt.Println((*enter)[i].nstr, "-", (*enter)[j].nstr)
						}
					}
				}
			}
			fmt.Println("<<<<<<<<<<<<<<<<<<<<<")
		} else {
			panic(err)
			rez = false
		}
	} else {
		fmt.Println("Enter отсутствует или открыт другой программой")
		rez = false
	}
	return rez
}
func checkArtikul(st string) bool {
	var cnt int = 0
	for i := 0; i < len(st); i++ {
		if st[i] == '.' {
			cnt++
		}
	}
	if cnt == 2 {
		return true
	} else {
		return false
	}
}
func ClearSymb(src string) string {
	var dst string = ""
	for i := 0; i < len(src); i++ {
		switch src[i] {
		case '1':
			dst += string(src[i])
		case '2':
			dst += string(src[i])
		case '3':
			dst += string(src[i])
		case '4':
			dst += string(src[i])
		case '5':
			dst += string(src[i])
		case '6':
			dst += string(src[i])
		case '7':
			dst += string(src[i])
		case '8':
			dst += string(src[i])
		case '9':
			dst += string(src[i])
		case '0':
			dst += string(src[i])
		case '.':
			dst += string(src[i])
		case ',':
			dst += string('.')
		}
	}
	return dst
}

func changeDot(src string) string {
	var dst string = ""
	for i := 0; i < len(src); i++ {
		if src[i] == '.' {
			dst += ","
		} else {
			dst += string(src[i])
		}
	}
	return dst
}

func CheckComma(src string) (brand string, model string, artikul string, hit string, id int, r bool) {
	r = true
	fstr := strings.Split(src, ",")
	if len(fstr) > 4 {

		brand = strings.TrimSpace(fstr[1])
		model = strings.TrimSpace(fstr[0])
		artikul = strings.TrimSpace(fstr[2])
		hit = strings.TrimSpace(fstr[3])
		id, _ = strconv.Atoi(strings.TrimSpace(fstr[4]))

	} else {
		r = false
	}
	return
}
func fillEnterId(enter string, data string) {
	fmt.Println("Заполнение данных Enter.")

}

func CheckData(pathToFile string) (int, bool) {
	var rezult, ex2 int
	var r bool = true
	var flag int = 0
	var startPos int = 3760
	fmt.Println("Разбор данных.")

	_, err := os.Stat(pathToFile)
	if err == nil {
		f, err := os.Open(pathToFile)
		if err == nil {
			fi, _ := f.Stat()
			buf := make([]uint8, fi.Size()+100)

			nr, _ := f.Read(buf)
			k := 1
			l := 1
			ex := 0
			for i := 0; i < nr && ex == 0; i++ {
				switch buf[i] {
				case 0x0d:
					if i > 0 {
						k = i
					}
				case 0x0a:
					if k > l {
						b := buf[l:k]
						bUTF := DecodeWindows1251(b)
						mstr := string(bUTF)
						// ----------- обработка ----------------
						rstr := strings.Split(mstr, "\t")
						if flag == 1 {
							brand, model, artikul, hit, id, r := CheckComma(rstr[0])

							if r == true {

								//								ex3 = 0

								ex2 = 0
								for m := 0; m < len(myenter) && ex2 == 0; m++ {
									if id == myenter[m].id {
										myenter[m].pg = 1
										myenter[m].artikul = artikul
										myenter[m].brand = brand
										myenter[m].hit = hit
										myenter[m].model = model
										for k := 0; k < NPLACE; k++ {
											myenter[m].data[k].kolvo, _ = strconv.Atoi(rstr[mypos[k].position])
										}
										ex2 = 1
										fmt.Print(",")
									}
								}
								if ex2 == 0 {
									for m := 0; m < len(myenter) && ex2 == 0; m++ {
										if brand == myenter[m].brand && model == myenter[m].model {
											myenter[m].pg = 1
											myenter[m].id = id
											myenter[m].artikul = artikul
											myenter[m].brand = brand
											myenter[m].hit = hit
											myenter[m].model = model
											for k := 0; k < NPLACE; k++ {
												myenter[m].data[k].kolvo, _ = strconv.Atoi(rstr[mypos[k].position])
											}
											ex2 = 1
											fmt.Print("/")
										}
									}
								}

								if ex2 == 0 {
									for m := 0; m < len(myenter) && ex2 == 0; m++ {
										if artikul == myenter[m].artikul {
											myenter[m].pg = 1
											myenter[m].id = id
											myenter[m].artikul = artikul
											myenter[m].brand = brand
											myenter[m].hit = hit
											myenter[m].model = model
											for k := 0; k < NPLACE; k++ {
												myenter[m].data[k].kolvo, _ = strconv.Atoi(rstr[mypos[k].position])
											}
											ex2 = 1
											fmt.Print(";")
										}
									}
								}

								if ex2 == 0 {
									fmt.Println("")
									rezult++
									fmt.Println("Отсутствует позиция Check:", brand, model)
									myenter[startPos].model = model
									myenter[startPos].brand = brand
									myenter[startPos].artikul = artikul
									myenter[startPos].pg = 2
									myenter[startPos].hit = hit
									myenter[startPos].id = id
									for k := 0; k < NPLACE; k++ {
										myenter[startPos].data[k].kolvo, _ = strconv.Atoi(rstr[mypos[k].position])
									}
									startPos++
								}
							}
						} else {
							if rstr[0] == "Номенклатура, Производитель номенклатуры, Артикул , Хит продаж, Код" {
								flag = 1
							}
						}
					}
					l = i + 1
				default:
				}
			}
			fmt.Println("Ok")
			f.Close()
		} else {
			panic(err)
			fmt.Println(err)
			r = false
		}
	} else {
		fmt.Println("Файл data.txt отсутствует!")
		r = false
	}
	return rezult, r
}
func WriteRezult(FileRez string) {
	var rstr string
	f, err := os.Create(FileRez)
	defer f.Close()
	if err == nil {
		for i := 0; i < len(myenter); i++ {
			rstr = string(EncodeWindows1251([]uint8(myenter[i].rstr))) + "\t"
			switch myenter[i].pg {
			case 0:

				rstr += "\r\n"
			case -1:
				zstr := strings.Split(rstr, "\t")
				mus := 0
				for m := 0; m < len(zstr); m++ {
					if zstr[m] == "Id" {
						mus = m
					}
				}
				if mus > 0 {
					rstr = strings.Join(zstr[:mus+1], "\t")
					rstr += "\t"
				}
				for j := 0; j < len(mypos); j++ {
					d := EncodeWindows1251([]uint8(mypos[j].header))
					rstr += string(d)
					rstr += "\t"
				}
				rstr += "\r\n"
			case 1:
				a := strconv.Itoa(myenter[i].nstr)
				nstr := string(EncodeWindows1251([]uint8(a)))
				pg := "1"
				artikul := string(EncodeWindows1251([]uint8(myenter[i].artikul)))
				brand := string(EncodeWindows1251([]uint8(myenter[i].brand)))
				model := string(EncodeWindows1251([]uint8(myenter[i].model)))
				vid := string(EncodeWindows1251([]uint8(myenter[i].vid)))
				hit := string(EncodeWindows1251([]uint8(myenter[i].hit)))
				rstr = nstr + "\t" + pg + "\t" + artikul + "\t" + brand + "\t" + model + "\t" + vid + "\t" + hit + "\t" + strconv.Itoa(myenter[i].li) + "\t" + strconv.Itoa(myenter[i].id) + "\t"
				for j := 0; j < len(mypos); j++ {
					rstr += strconv.Itoa(myenter[i].data[j].kolvo)
					rstr += "\t"
				}
				rstr += "\r\n"
			case 2:
				a := strconv.Itoa(myenter[i].nstr)
				nstr := string(EncodeWindows1251([]uint8(a)))
				pg := "2"
				artikul := string(EncodeWindows1251([]uint8(myenter[i].artikul)))
				brand := string(EncodeWindows1251([]uint8(myenter[i].brand)))
				model := string(EncodeWindows1251([]uint8(myenter[i].model)))
				vid := string(EncodeWindows1251([]uint8(myenter[i].vid)))
				hit := string(EncodeWindows1251([]uint8(myenter[i].hit)))
				rstr = nstr + "\t" + pg + "\t" + artikul + "\t" + brand + "\t" + model + "\t" + vid + "\t" + hit + "\t" + strconv.Itoa(myenter[i].li) + "\t" + strconv.Itoa(myenter[i].id) + "\t"
				for j := 0; j < len(mypos); j++ {
					rstr += strconv.Itoa(myenter[i].data[j].kolvo)
					rstr += "\t"
				}
				rstr += "\r\n"

			}
			f.WriteString(rstr)
		}

	}
}

type mps struct {
	name     string
	header   string
	position int
}

var mypos [8]mps

func iniMyPos(pathToFile string) {
	mypos[0].name = "Км - Склад"
	mypos[0].header = "Косм"
	mypos[1].name = "Р - Склад"
	mypos[1].header = "Рост"
	mypos[2].name = "Магнит (Переславль) - Склад"
	mypos[2].header = "Магн(П)"
	mypos[3].name = "П - Склад"
	mypos[3].header = "Перес"
	mypos[4].name = "У - Склад"
	mypos[4].header = "Углич"
	mypos[5].name = "Б - Склад 1"
	mypos[5].header = "Б1"
	mypos[6].name = "Б - Склад 2"
	mypos[6].header = "Б2"

	mypos[7].name = "Доставка товаров"
	mypos[7].header = "Доставка"
	fmt.Println("Конфигурация данных.")
	_, err := os.Stat(pathToFile)
	if err == nil {
		f, err := os.Open(pathToFile)
		if err == nil {
			fi, _ := f.Stat()
			buf := make([]uint8, fi.Size()+100)

			nr, _ := f.Read(buf)
			k := 1
			l := 1
			ex := 0
			for i := 0; i < nr && ex == 0; i++ {
				switch buf[i] {
				case 0x0d:
					if i > 0 {
						k = i
					}
				case 0x0a:
					if k > l {
						b := buf[l:k]
						bUTF := DecodeWindows1251(b)
						mstr := string(bUTF)
						// ----------- обработка ----------------
						rstr := strings.Split(mstr, "\t")
						if strings.TrimSpace(rstr[0]) == "Номенклатура, Производитель номенклатуры, Артикул , Хит продаж, Код" {
							for j := 0; j < len(rstr); j++ {
								ex1 := 0
								for k := 0; k < NPLACE && ex1 == 0; k++ {
									if mypos[k].name == rstr[j] {
										mypos[k].position = j
										ex1 = 1
										fmt.Print(".")
									}
								}

							}
							ex = 1
							//fmt.Println("!Header!")
						}
					}
					l = i + 1
				default:
				}
			}
			f.Close()
		} else {
			panic(err)
			fmt.Println(err)
		}
	} else {
		fmt.Println("Файл data.txt отсутствует!")
	}
}
func fillKenter(DataFile string) bool {
	var r bool = true
	var flag int = 0
	var startPos int = 3760
	fmt.Println("Заполнение kenter.", DataFile)
	fmt.Println("Kenter длина :", len(kenter))
	_, err := os.Stat(DataFile)
	if err == nil {
		f, err := os.Open(DataFile)
		if err == nil {
			fi, _ := f.Stat()
			buf := make([]uint8, fi.Size()+100)

			nr, _ := f.Read(buf)
			k := 1
			l := 1
			ex := 0
			for i := 0; i < nr && ex == 0; i++ {
				switch buf[i] {
				case 0x0d:
					if i > 0 {
						k = i
					}
				case 0x0a:
					if k > l {
						b := buf[l:k]
						bUTF := DecodeWindows1251(b)
						mstr := string(bUTF)
						// ----------- обработка ----------------
						rstr := strings.Split(mstr, "\t")
						if flag == 1 {
							brand, model, artikul, hit, id, r := CheckComma(rstr[0])
							if r == true {

								//fmt.Println(brand, ",", model)
								ex2 := 0
								for m := 0; m < len(kenter) && ex2 == 0; m++ {
									if model == kenter[m].model && brand == kenter[m].brand {
										kenter[m].artikul = artikul
										//fmt.Println(id)
										kenter[m].hit = hit
										kenter[m].id = id
										ex2 = 1
									}
								}
								if ex2 == 0 {
									for m := 0; m < len(kenter) && ex2 == 0; m++ {
										if id == kenter[m].id {
											kenter[m].artikul = artikul
											kenter[m].model = model
											kenter[m].brand = brand
											kenter[m].hit = hit
											kenter[m].id = id
											ex2 = 1
										}
									}
								}
								if ex2 == 0 {
									for m := 0; m < len(kenter) && ex2 == 0; m++ {
										if artikul == kenter[m].artikul {
											kenter[m].artikul = artikul
											kenter[m].model = model
											kenter[m].brand = brand
											kenter[m].hit = hit
											kenter[m].id = id
											ex2 = 1
										}
									}
								}
								if ex2 == 0 {
									fmt.Println("Отсутствует позиция в энтер:", brand, model)
									//fmt.Println("StartPos:", startPos)
									kenter[startPos].model = model
									kenter[startPos].brand = brand
									kenter[startPos].artikul = artikul
									kenter[startPos].pg = 2
									kenter[startPos].hit = hit
									kenter[startPos].id = id

									startPos++
								}
							} else {

							}
						} else {
							if rstr[0] == "Номенклатура, Производитель номенклатуры, Артикул , Хит продаж, Код" {
								flag = 1
							}
						}
					}
					l = i + 1
				default:
				}
			}
			f.Close()
		} else {
			panic(err)
			fmt.Println(err)
			r = false
		}
	} else {
		fmt.Println("Файл data.txt отсутствует!")
		r = false
	}
	return r
}
func WriteKenter(newEnter string) {
	var rstr string
	f, err := os.Create(newEnter)
	defer f.Close()
	if err == nil {
		for i := 0; i < len(kenter); i++ {
			rstr = string(EncodeWindows1251([]uint8(kenter[i].rstr))) + "\t"
			switch kenter[i].pg {
			case 0:

				rstr += "\r\n"
			case -1:

				rstr += "\r\n"
			case 1:
				a := strconv.Itoa(kenter[i].nstr)
				nstr := string(EncodeWindows1251([]uint8(a)))
				pg := "1"
				artikul := string(EncodeWindows1251([]uint8(kenter[i].artikul)))
				brand := string(EncodeWindows1251([]uint8(kenter[i].brand)))
				model := string(EncodeWindows1251([]uint8(kenter[i].model)))
				vid := string(EncodeWindows1251([]uint8(kenter[i].vid)))
				hit := string(EncodeWindows1251([]uint8(kenter[i].hit)))
				rstr = nstr + "\t" + pg + "\t" + artikul + "\t" + brand + "\t" + model + "\t" + vid + "\t" + hit + "\t" + strconv.Itoa(kenter[i].li) + "\t" + strconv.Itoa(kenter[i].id)

				rstr += "\r\n"
			case 2:
				a := strconv.Itoa(kenter[i].nstr)
				nstr := string(EncodeWindows1251([]uint8(a)))
				pg := "2"
				artikul := string(EncodeWindows1251([]uint8(kenter[i].artikul)))
				brand := string(EncodeWindows1251([]uint8(kenter[i].brand)))
				model := string(EncodeWindows1251([]uint8(kenter[i].model)))
				vid := string(EncodeWindows1251([]uint8(kenter[i].vid)))
				hit := string(EncodeWindows1251([]uint8(kenter[i].hit)))
				rstr = nstr + "\t" + pg + "\t" + artikul + "\t" + brand + "\t" + model + "\t" + vid + "\t" + hit + "\t" + strconv.Itoa(kenter[i].li) + "\t" + strconv.Itoa(kenter[i].id)

				rstr += "\r\n"
			}
			f.WriteString(rstr)
		}

	}
}
func WriteTEnter(MyOldEnter string) {
	var rstr string
	f, err := os.Create(MyOldEnter)
	defer f.Close()
	if err == nil {
		for i := 0; i < len(myenter); i++ {
			rstr = string(EncodeWindows1251([]uint8(myenter[i].rstr))) + "\t"
			switch myenter[i].pg {
			case 0:
				astr := myenter[i].rstr
				arstr := strings.Split(astr, "\t")
				n := 0
				for i := 0; i < len(arstr) && n == 0; i++ {
					if arstr[i] == "-1" {
						n = i
					}
				}
				fstr := arstr[:]
				if n > 0 {
					fstr = append(arstr[:n-1], arstr[n:]...)
				}
				rstr = string(EncodeWindows1251([]uint8(strings.Join(fstr, "\t"))))
				rstr += "\r\n"
			case -1:
				astr := myenter[i].rstr
				arstr := strings.Split(astr, "\t")
				n := 0
				//fmt.Println(astr)
				for i := 0; i < len(arstr) && n == 0; i++ {
					if arstr[i] == "Бранд" {
						n = i
					}
				}
				fstr := append(arstr[:n], arstr[n+1:]...)
				rstr = string(EncodeWindows1251([]uint8(strings.Join(fstr, "\t"))))
				rstr += "\r\n"
			case 1:
				a := strconv.Itoa(myenter[i].nstr)
				nstr := string(EncodeWindows1251([]uint8(a)))
				pg := "1"
				artikul := string(EncodeWindows1251([]uint8(myenter[i].artikul)))
				brand := string(EncodeWindows1251([]uint8(myenter[i].brand)))
				model := string(EncodeWindows1251([]uint8(myenter[i].model)))
				Tmodel := "0"
				if brand != "0" && model != "0" {
					Tmodel = brand + " " + model
				}

				vid := string(EncodeWindows1251([]uint8(myenter[i].vid)))
				hit := string(EncodeWindows1251([]uint8(myenter[i].hit)))
				if hit == "" {
					hit = "P"
				}
				rstr = nstr + "\t" + pg + "\t" + artikul + "\t" + Tmodel + "\t" + vid + "\t" + hit + "\t" + strconv.Itoa(myenter[i].li) + "\t" + strconv.Itoa(myenter[i].id)

				rstr += "\r\n"
			case 2:
				a := strconv.Itoa(myenter[i].nstr)
				nstr := string(EncodeWindows1251([]uint8(a)))
				pg := "2"
				artikul := string(EncodeWindows1251([]uint8(myenter[i].artikul)))
				brand := string(EncodeWindows1251([]uint8(myenter[i].brand)))
				model := string(EncodeWindows1251([]uint8(myenter[i].model)))

				Tmodel := brand + " " + model

				vid := string(EncodeWindows1251([]uint8(myenter[i].vid)))
				hit := string(EncodeWindows1251([]uint8(myenter[i].hit)))
				if hit == "" {
					hit = "k"
				}
				rstr = nstr + "\t" + pg + "\t" + artikul + "\t" + Tmodel + "\t" + vid + "\t" + hit + "\t1\t" + strconv.Itoa(myenter[i].id)

				rstr += "\r\n"
			}
			f.WriteString(rstr)
		}

	}
}
func WriteSEnter(MyOldEnter string) {
	var rstr string
	f, err := os.Create(MyOldEnter)
	defer f.Close()
	if err == nil {
		for i := 0; i < len(myenter); i++ {
			rstr = string(EncodeWindows1251([]uint8(myenter[i].rstr))) + "\t"
			switch myenter[i].pg {
			case 0:

				rstr = string(EncodeWindows1251([]uint8(myenter[i].rstr)))
				rstr += "\r\n"
			case -1:

				rstr = string(EncodeWindows1251([]uint8(myenter[i].rstr)))
				rstr += "\r\n"
			case 1:
				a := strconv.Itoa(myenter[i].nstr)
				nstr := string(EncodeWindows1251([]uint8(a)))
				pg := "1"
				artikul := string(EncodeWindows1251([]uint8(myenter[i].artikul)))
				brand := string(EncodeWindows1251([]uint8(myenter[i].brand)))
				model := string(EncodeWindows1251([]uint8(myenter[i].model)))

				vid := string(EncodeWindows1251([]uint8(myenter[i].vid)))
				hit := string(EncodeWindows1251([]uint8(myenter[i].hit)))
				if hit == "" {
					hit = "k"
				}
				rstr = nstr + "\t" + pg + "\t" + artikul + "\t" + brand + "\t" + model + "\t" + vid + "\t" + hit + "\t" + strconv.Itoa(myenter[i].li) + "\t" + strconv.Itoa(myenter[i].id)

				rstr += "\r\n"
			case 2:
				a := strconv.Itoa(myenter[i].nstr)
				nstr := string(EncodeWindows1251([]uint8(a)))
				pg := "2"
				artikul := string(EncodeWindows1251([]uint8(myenter[i].artikul)))
				brand := string(EncodeWindows1251([]uint8(myenter[i].brand)))
				model := string(EncodeWindows1251([]uint8(myenter[i].model)))

				vid := string(EncodeWindows1251([]uint8(myenter[i].vid)))
				hit := string(EncodeWindows1251([]uint8(myenter[i].hit)))
				if hit == "" {
					hit = "k"
				}
				rstr = nstr + "\t" + pg + "\t" + artikul + "\t" + brand + "\t" + model + "\t" + vid + "\t" + hit + "\t" + strconv.Itoa(myenter[i].li) + "\t" + strconv.Itoa(myenter[i].id)

				rstr += "\r\n"
			}
			f.WriteString(rstr)
		}

	}
}
func WriteMyEnter(MyOldEnter string) {
	var rstr string
	f, err := os.Create(MyOldEnter)
	defer f.Close()
	if err == nil {
		for i := 0; i < len(kenter); i++ {
			rstr = string(EncodeWindows1251([]uint8(kenter[i].rstr))) + "\t"
			switch kenter[i].pg {
			case 0:
				astr := kenter[i].rstr
				arstr := strings.Split(astr, "\t")
				n := 0
				for i := 0; i < len(arstr) && n == 0; i++ {
					if arstr[i] == "-1" {
						n = i

					}
				}
				fstr := append(arstr[:n-1], arstr[n:]...)
				rstr = string(EncodeWindows1251([]uint8(strings.Join(fstr, "\t"))))
				rstr += "\r\n"
			case -1:
				astr := kenter[i].rstr
				arstr := strings.Split(astr, "\t")
				n := 0

				for i := 0; i < len(arstr) && n == 0; i++ {
					if arstr[i] == "Брэнд" {
						n = i

					}
				}
				fstr := append(arstr[:n], arstr[n+1:]...)
				rstr = string(EncodeWindows1251([]uint8(strings.Join(fstr, "\t"))))
				rstr += "\r\n"
			case 1:
				a := strconv.Itoa(kenter[i].nstr)
				nstr := string(EncodeWindows1251([]uint8(a)))
				pg := "1"
				artikul := string(EncodeWindows1251([]uint8(kenter[i].artikul)))
				brand := string(EncodeWindows1251([]uint8(kenter[i].brand)))
				model := string(EncodeWindows1251([]uint8(kenter[i].model)))

				Tmodel := brand + " " + model

				vid := string(EncodeWindows1251([]uint8(kenter[i].vid)))
				hit := string(EncodeWindows1251([]uint8(kenter[i].hit)))
				if hit == "" {
					hit = "k"
				}
				rstr = nstr + "\t" + pg + "\t" + artikul + "\t" + Tmodel + "\t" + vid + "\t" + hit + "\t" + strconv.Itoa(kenter[i].li) + "\t" + strconv.Itoa(kenter[i].id)

				rstr += "\r\n"
			case 2:
				a := strconv.Itoa(kenter[i].nstr)
				nstr := string(EncodeWindows1251([]uint8(a)))
				pg := "2"
				artikul := string(EncodeWindows1251([]uint8(kenter[i].artikul)))
				brand := string(EncodeWindows1251([]uint8(kenter[i].brand)))
				model := string(EncodeWindows1251([]uint8(kenter[i].model)))

				Tmodel := brand + " " + model

				vid := string(EncodeWindows1251([]uint8(kenter[i].vid)))
				hit := string(EncodeWindows1251([]uint8(kenter[i].hit)))
				if hit == "" {
					hit = "k"
				}
				rstr = nstr + "\t" + pg + "\t" + artikul + "\t" + Tmodel + "\t" + vid + "\t" + hit + "\t1\t" + strconv.Itoa(kenter[i].id)

				rstr += "\r\n"
			}
			f.WriteString(rstr)
		}

	}
}
func main() {
	var input string
	myenter = make([]lenter, 0)
	kenter = make([]lenter, 0)
	DataFile := "./data.txt"
	EnterPath := "E:\\Rostatki\\Enter\\UEnter.txt"
	MyRezFile := "./rez.txt"
	MyOldEnter := "E:\\Rostatki\\Enter\\Enter.txt"
	//------------------------------
	REnterPath := "E:\\Rostatki\\Enter\\EnterLab.txt"

	b := KIniEnter(REnterPath)

	if b == true {
		//fillKenter("C:\\Go\\my_files\\Roslan\\info.txt")
		fillKenter("./data.txt")
		WriteKenter(EnterPath)

	}

	//------------------------------

	iniMyPos(DataFile)
	b = IniEnter(EnterPath, &myenter)
	if b == true {
		n, _ := CheckData(DataFile)

		WriteRezult(MyRezFile)
		if n == 0 {
			WriteTEnter(MyOldEnter)
			WriteSEnter(REnterPath)
		}

		fmt.Println("Press any key and Enter!")
		fmt.Scanf("%s\r\n", &input)

	} else {
		fmt.Println("Press any key and Enter!")
		fmt.Scanf("%s\r\n", &input)
	}
}
