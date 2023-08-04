# Scheduler - Backend
> Repository ini merupakan bagian front-end dari sebuah web-app repository untuk mencari kombinasi mata kuliah paling optimal.
> Kombinasi dari mata kuliah yang paling optimal adalah kombinasi yang menghasilkan IPK tertinggi dengan jumlah SKS tertinggi.
> Pemilihan mata kuliah dibatasi oleh jurusan, semester, sks minimal, dan sks maksimal.
> Repository back-end dapat dilihat di [sini](https://github.com/Marthenn/scheduler-frontend)

## Cara Menjalankan
1. Clone repository
2. Jalankan perintah `docker-compose up -d` pada terminal

## Framework
- pgx (PostgreSQL driver)
- gorilla/mux (HTTP router)
- net/http (HTTP server)

## Dynamic Programming
- Dynamic programming merupakan optimisasi dari rekursi
- Dynamic programming memanfaatkan memoization untuk mengurangi kompleksitas waktu dari rekursi
- Dynamic programming memanfaatkan tabel untuk menyimpan hasil dari rekursi
- Terdapat 2 jenis dynamic programming, yaitu top-down dan bottom-up
- Top-down merupakan dynamic programming yang memanfaatkan rekursi
- Bottom-up merupakan dynamic programming yang memanfaatkan iterasi

## Analisis Algoritma
- Algoritma yang digunakan adalah algoritma dynamic programming bottom-up
- Algoritma menggunakan tabel dua baris untuk menyimpan hasil dari iterasi
- Algoritma menggunakan pendekatan yes no untuk setiap mata kuliah
- Algoritma menggunakan bitmasking untuk menyimpan mata kuliah yang sudah diambil
- Indeks x pada setiap barisnya merepresentasikan hasil dari iterasi pada mata kuliah dengan bit tertentu
- Contoh: indeks 3 merepresentasikan hasil bila mata kuliah pertama dan kedua diambil (011)
- Algoritma mengembalikan hasil pengecekan pada baris terakhir
- Digunakan pendekatan ini karena dynamic programming 1/0 knapsack classic tidak dapat menangani kasus penambahan constraint minimum weight

## Referensi
- [Dynamic Programming](https://www.geeksforgeeks.org/dynamic-programming/)
- [1/0 Knapsack Classic](https://www.geeksforgeeks.org/0-1-knapsack-problem-dp-10/)
- [GoLang PostGreSQL](https://www.cockroachlabs.com/)