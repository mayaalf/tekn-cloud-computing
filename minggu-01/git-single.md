# Installasi Git Windows

[ [<< Kembali](README.md) ]

1. Double click pada file download Git kemudian klik **Next** untuk melanjutkan installasi.

![01](images/1/1.png)

2. Setelah itu pilih lokasi instalasi. Secara default akan terisi _C:\Program Files\Git_. Gantilah lokasi jika menginginkannya.

![02](images/1/2.png)

3. Pilih komponen. Tidak perlu diubah-ubah, sesuaikan dengan default saja lalu klik **Next**.

![03](images/1/3.png)

4. Mengisi shortcut untuk menu start. Gunakan default **(Git)**, ganti jika ingin mengganti -misalnya Git VCS

![04](images/1/4.png)

5. Pilih edior yang akan digunakan bersama dengan Git. Pada pilihan ini, digunakan Visual Studio Code.

![05](images/1/5.png)

6. Pada saat installasi, Git menyediakan akses git melalui Bash maupun commad prompt. Pilih pilihan kedua supaya bisa menggunakan dua antarmuka.

![06](images/1/7.png)

7. Pilih OpenSSL untuk HTTPS. Git menggunakan https untuk akses ke repo GitHub atau repo-repo lain.

![07](images/1/9.png)

8.Pilih pilihan pertama untuk konversi akhir baris (CR-LF).

![08](images/1/10.png)

9. Pilih PuTTY untuk terminal yang digunakan untuk mengakses GitBash.

![09](images/1/11.png)

10. Opsi ekstra, pilih enable file system caching.

![10](images/1/14.png)

11. Setelah itu proses installasi akan dilakukan.

![11](images/1/15.png)

12. Jika sudah selesai akan muncul dialog pemberitahun lalu kilik **Finish**.

![12](images/1/16.png)

13. Menjalankan Git dari menu start ketikkan "Git", lalu akan muncul pilihan "Git Bash", "Git CMD" atau "Git GUI"

![13](images/1/17.png)

14. Tampilan jika menggunakan "Git Bash"

![14](images/1/18.png)

15. Tampilan jika menggunkan "Git GUI"

![15](images/1/19.png)

16. Mencoba dari command prompt lalu eksekusi "git --version" untuk melihat sudah berhasil terinstall atau belum. Jika sudah terinstall dengan benar maka akan muncul seperti ini :

![16](images/1/20.png)



# Konfigurasi Git

Konfigurasi menggunakan username dan email dengan perintah berikut :

```
$ git config --global user.name "Nama Anda di GutHub"
$ git config --global user.email email@domain.tld
```

![17](images/2/21.png)

Isian diatas harus sesuai dengan email yang digunakan untuk mendaftar GitHub. Untuk melihat konfigurasi yang sudah ada :

```
$ git config --list
user.name=mayaalf
user.email=maya.alif@students.utdi.ac.id
$
```

![18](images/2/22.png)

Langkah ini cukup dilakukan sekali saja, kecuali jika ingin melakukan perubahan nama dan email.


# Mengelola Repo Sendiri

## Mengelola Repo Sendiri di Account Sendiri
