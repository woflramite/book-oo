# Book-oo API
Oleh: Balya Ibnu Sulistiyono (18215005)

Sistem Book-oo API ini adalah sistem web service yang membantu user untuk mendata buku-buku yang member miliki. Sistem ini juga memungkinkan user baru untuk mendaftarkan akun baru. Selain itu, sistem ini juga membantu user untuk menemukan pemilik suatu buku yang sedang member cari.

Kebutuhan fungsional sistem ini adalah sebagai berikut:
* Memberikan respons data seluruh member
* Memberikan respons data seluruh buku
* Memberikan respons data suatu member spesifik
* Memberikan respons data suatu buku spesifik
* Memberikan respons data rak buku suatu member spesifik
* Memberikan respons data pemilik suatu buku spesifik
* Menerima request untuk memasukkan data member baru
* Menerima request untuk memasukkan data buku baru
* Menerima request untuk memasukkan data kepemilikan suatu buku

> Spesifikasi kebutuhan perangkat lunak Book-oo API bisa dilihat pada dokumen "SKPL Progif.pdf"

## Petunjuk Pengujian Book-oo API
Modul pengujian terbagi menjadi dua macam, yaitu modul mendapatkan data (GET) dan modul mengirimkan data (POST). Pada bagian ini, akan dibahas metode untuk menguji modul-modul tersebut.
GET Method.

> Gambar hasil pengujian bisa dilihat pada dokumen "Test Progif.pdf"

### Mendapatkan Data Seluruh Member
Cara untuk mendapatkan data seluruh member adalah dengan mengisi berikut di address bar.
```sh
http://167.205.67.244:9801/member
```
### Mendapatkan Data Seluruh Buku
Cara untuk mendapatkan data seluruh buku adalah dengan mengisi berikut di address bar.
```sh
http://167.205.67.244:9801/book
```
### Mendapatkan Data Member Ber-ID Tertentu
Cara untuk mendapatkan data member dengan ID tertentu adalah dengan mengisi berikut di address bar.
```sh
http://167.205.67.244:9801/member/{nomor id}
```
> Catatan: Nomor ID HARUS berupa integer. Input berupa string belum di-handle dan menyebabkan program tertutup.
### Mendapatkan Data Buku Ber-ISBN Tertentu
Cara untuk mendapatkan data member dengan ID tertentu adalah dengan mengisi berikut di address bar.
```sh
http://167.205.67.244:9801/book/{nomor isbn}
```
### Mendapatkan Data Buku-buku yang Dimiliki Member Ber-ID Tertentu
Cara untuk mendapatkan data member dengan ID tertentu adalah dengan mengisi berikut di address bar.
```sh
http://167.205.67.244:9801/member/{nomor id}/shelf
```
>Catatan: Nomor ID HARUS berupa integer. Input berupa string belum di-handle dan menyebabkan program tertutup.
### Mendapatkan Data Member-member yang Memiliki Buku Ber-ISBN Tertentu
Cara untuk mendapatkan data member dengan ID tertentu adalah dengan mengisi berikut di address bar.
```sh
http://167.205.67.244:9801/book/{nomor isbn}/owned-by
```
## POST Method
### Mengirimkan Data Member Baru
Cara untuk mengirimkan data JSON untuk menambahkan member baru adalah sebagai berikut.
```sh
curl -H "Content-Type: application/json" -X POST -d '{"students_id": "{id_input}", "name":"{nama_input}", "social_media_contact":"{kontak_input}"}' http://167.205.67.244:9801/member/add
```
>Catatan: Nomor ID HARUS berupa integer. Input berupa string belum di-handle dan menyebabkan program tertutup.
### Mengirimkan Data Buku Baru
Cara untuk mengirimkan data JSON untuk menambahkan member baru adalah sebagai berikut.
```sh
curl -H "Content-Type: application/json" -X POST -d '{"isbn": "{isbn_input}", "book_title":"{judul_input}"}' http://167.205.67.244:9801/book/add
```
>Catatan: Nomor ID HARUS berupa integer. Input berupa string belum di-handle dan menyebabkan program tertutup.
### Mengirimkan Data Kepemilikan Buku Baru
Cara untuk mengirimkan data JSON untuk menambahkan koleksi buku baru adalah sebagai berikut.
```sh
curl -H "Content-Type: application/json" -X POST -d '{"isbn": "{isbn_input}", "students_id":"{id_input}", "number_of_copy":"{jumlah_salinan}"}' http://167.205.67.244:9801/member/{id_input}/shelf/add
```
> Catatan: Nomor ID HARUS berupa integer. Input berupa string belum di-handle dan menyebabkan program tertutup.
