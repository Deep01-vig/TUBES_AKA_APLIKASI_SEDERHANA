1. Deskripsi Proyek
Proyek ini merupakan implementasi aplikasi desktop sederhana berbasis Command Line Interface (CLI) yang bertujuan untuk menganalisis efisiensi algoritma Binary Search menggunakan dua pendekatan, yaitu iteratif dan rekursif.

Analisis dilakukan dengan membandingkan total waktu eksekusi dari kedua algoritma pada berbagai ukuran data, sehingga dapat diamati perbedaan performa serta kesesuaiannya dengan teori kompleksitas algoritma.

2. Tujuan
Tujuan dari proyek ini adalah:

  - Mengimplementasikan algoritma Binary Search secara iteratif dan rekursif
  - Membandingkan efisiensi waktu eksekusi kedua pendekatan
  - Mengamati pengaruh ukuran data terhadap performa algoritma
  - Memverifikasi kesesuaian hasil pengujian dengan kompleksitas waktu O(log n)

3. Spesifikasi Aplikasi

  - Jenis Aplikasi: Desktop sederhana (CLI / CMD)
  - Bahasa Pemrograman: Go (Golang)

Input:

- Jumlah data

- Data NIM terurut

- Nilai target yang dicari

Output:

- Indeks data yang ditemukan

- Total waktu eksekusi algoritma iteratif

- Total waktu eksekusi algoritma rekursif

Cara Kerja Aplikasi:

1. Pengguna memasukkan jumlah data dan data numerik yang sudah terurut
2. Pengguna memasukkan nilai target yang ingin dicari
3. Program menjalankan:
  - Binary Search Iteratif
  - Binary Search Rekursif

5. Setiap algoritma diuji sebanyak 100.000 iterasi

6. Program menampilkan hasil pencarian dan total waktu eksekusi dalam satuan mikrodetik (µs)
