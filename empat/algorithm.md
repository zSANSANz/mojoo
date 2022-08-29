#algorithm
 untuk mengurutkan angka dari nilai terkecil sampai tertinggi ataupun juga sebaliknya, kita memerlukan sebuah variable temperatur yang berfungsi sebagai wadah sementara jika itu ditukar dengan isi index yang lebih besar dari index dengan nilai terkecil itu, selanjutnya index terkecil kedua akan disimpan pada variable tmp (sementara) lalu ditukar dengan index kedua dan berulang terus sampai semua selesai terurut begitupun juga sebalikanya, kalau yang diinginkan mengurutkan dengan ascending maupun descending

    baris 1 : mendeklarasikan file dalam package main
    baris 3 : mengimport library fmt untuk formatting
    baris 5 : membuat fungsi selectionSort yang menerima sebuah variable array
    baris 7 : melakukan perulangan i sebanyak isi element arrray variable arr – 1.
    baris 8 : mendeklarasikan variable minIndex dengan diisikan variable i.
    baris 9 : melakukan perulangan j sebanyak isi element arrray variable arr.
    baris 10 : melakukan pengecekan jika isi element array ke minIndex lebih besar dari isi element array ke j.
    baris 11 : maka minIndex diisikan dengan variable j.
    baris 12 : simbol tutup kurung kurawal dari pengecekan.
    baris 13 : simbol tutup kurung kurawal dari perulangan j.
    baris 14 : mendeklarasikan variable tmp diisikan dengan isi element array ke i.
    baris 15 : isi element array dari variable arr ke i diisikan dengan isi element array dari variable arr ke minIndex.
    baris 16 : dan isi element array dari variable arr index ke minIndex diisikan dengan isi variable tmp.
    baris 17 : simbol tutup kurung kurawal dari perulangan i
    baris 18 : simbol tutup kurung kurawal dari fungsi selectionSort
    baris 20 : membuat fungsi main
    baris 21 : mendeklarasikan variable arr berupa tipe data array float yang memiliki element array [4, -7, -5, 3, 3.3, 9, 0, 10, 0.2].
    baris 22 : memanggil fungsi selectionSort dengan menginputkan variable arr.
    baris 23 : melakukan perulangan sebanyak jumlah element array dari variable arr.
    baris 24 : akan menampilkan output dari isi element array dari variable arr.
    baris 25 : simbol tutup kurung kurawal dari perulangan
    baris 26 : menampilkan output baris baru.
    baris 27 : simbol tutup kurung kurawal dari fungsi main.