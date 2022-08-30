Test Case : Backend
1. Terdapat ERD seperti berikut:
Dalam team ini anda ditugaskan membuat API fungsi-fungsi berikut:
    A. Membuat fungsi login (5 point)
    B. Untuk authorization pada point A gunakan JWT (6 point)
    C. Laporan nama merchant, omzet per hari dalam pada bulan november mulai tanggal 1
    sampai dengan tanggal 30 dengan pagination. Apabila tidak ada transaksi pada tanggal itu
    omzet akan bernilai 0 (6 point)
    D. API untuk menampilkan laporan nama merchant, nama outlet, omzet per hari pada bulan
    november mulai tanggal 1 sampai dengan tanggal 30 dengan pagination. Apabila tidak ada
    transaksi pada tanggal itu omzet akan bernilai 0 (6 point)
    E. Pada poin C pastikan user tidak bisa melakukan akses pada merchant_id yang bukan
    miliknya (10 point)
    F. Pada poin D pastikan user tidak bisa melakukan akses laporan pada outlet_id yang bukan
    miliknya (5 point)
    G. Dari test case pada point C dan point D, apakah struktur ERD yang dibentuk sudah optimal
    ? berikan penjelasannya (9 point)
    H. Dokumen teknis Data Manipulation Language (DML) (3 point)

Dari test case pada point C dan point D, apakah struktur ERD yang dibentuk sudah optimal
    - belum optimal, karena transaction mempunyai relationship dengan merhcant dan outlet sekaligus, padahal sebenarnya outlet sudah mempunyai relationship dengan merchant, jadi sebenarnya transaction cukup memanggil outlet, maka merchant akan terpanggil secara otomatis  