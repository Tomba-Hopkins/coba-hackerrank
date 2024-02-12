package algorithm

import "fmt"

// Question :
// Sample Input
// 1 2 3 4 5

// Sample Output
// 10 14
// Answer :

func miniMaxSum(arr []int32) {
    small := arr[0]
    big := arr[0]

    someRest := []int32{}
    var valRest int32

    for i := 1; i < len(arr); i++{
        if arr[i] < small {
            small = arr[i]
        }
        if arr[i] > big {
            big = arr[i]
        }
        if arr[i] == big {
            someRest = append(someRest, arr[i]) // nambahnya len(arr) - 1, soalnya mulainya dari index 1
        }
    }

    for _, sr := range someRest{
        valRest += sr
    }

    var min, max int64
	kotakTanpaTerkecil := []int64{}
	kotakTanpaTerbesar := []int64{}

    for _, ar := range arr{
        if ar != small {
            kotakTanpaTerkecil = append(kotakTanpaTerkecil, int64(ar))
			min += int64(ar)
        }
        if ar != big {
            kotakTanpaTerbesar = append(kotakTanpaTerbesar, int64(ar))
			max += int64(ar)
        } else if ar == small {
            kotakTanpaTerkecil = append(kotakTanpaTerkecil, int64(ar))
            kotakTanpaTerbesar = append(kotakTanpaTerbesar, int64(ar))
            min = int64(valRest)
            max = int64(valRest)
            
        }
    }

    fmt.Println("Kotak tanpa angka terkecil : ",kotakTanpaTerkecil,"\nKotak tanpa angka terbesar : ",kotakTanpaTerbesar)

	fmt.Println("Sum tanpa angka terkecil: ", min,"\nSUm tanpa angka terbesar : ", max)

    //hasil jangan int32, biar lolos penjumlahan yang digit nya banyak, pake int 64 aja

}


// atau :

// func miniMaxSum(arr []int32) {
//     small := arr[0]
//     big := arr[0]
    
//     someRest := []int32{}
       
//        for i := 1; i < len(arr); i++{
//            if arr[i] < small {
//                small = arr[i] //1
//            }
//            if arr[i] > big {
//                big = arr[i] //9
//            }
//            if arr[i] == big {
//                someRest = append(someRest, arr[i])
//            }
//        }
//        var valRest int32
//        for _, sr := range someRest {
//            valRest += sr
//        }
       
//        var min, max int64
       
//        for _, ar := range arr{
//            if ar != small {
//                min += int64(ar)
//            }
//            if ar != big {
//                max += int64(ar)
//             } else if ar == small {
//                 min = int64(valRest)
//                 max = int64(valRest)
//             }
//        }
       
//       fmt.Println(max, min)

// }