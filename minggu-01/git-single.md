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

### Langkah-langkah

1. Buat repo kosong di GitHub, bisa public maupun private
2. Clone repo kosong tersebut di komputer lokal
3. Perintah berikutnya terkait dengan perubahan repo serta sinkronisasi antara GitHub dengan lokal.

# Membuat Repo

1. Klik tanda + pada bagian atas setelah login, pilih **New repository**

![19](images/3/22-1.png)

2. Isikan nama, keterangan serta lisensi. Jika dikehendaki bisa membuat repo **Private**

![20](images/3/23.png)

3. Klik ```Create Repository```

Setelah langkah-langkah tersebut repo akan dibuat dan bisa diakses menggunakan pola ```https://github.com/username/reponame```. Pada repo tersebut, hanya akan muncul 1 file, yaitu LICENSE. Jika memilih membuat README pada saat langkah ke 2, juga akan muncul README.md. Ada atau tidak ada README.md tidak mempunyai efek apapun pada langkah ini.

![21](images/3/23-1.png)

## Clone Repo
Proses ```clone``` untuk duplikasi remote repo di GitHub ke komputer lokal, perintah yang digunakan :
```
git clone https://github.com/mayaalf/tekn-cloud-computing
Cloning into 'tekn-cloud-computing'...
remote: Enumerating objects: 3, done.
remote: Counting objects: 100% (3/3), done.
remote: Compressing objects: 100% (2/2), done.
remote: Total 3 (delta 0), reused 0 (delta 0), pack-reused 0
Unpacking objects: 100% (3/3), done.
```

![22](images/3/24.png)

Setelah perintah ini, di direktori ```tekn-cloud-computing``` akan disimpan isi repo yang sama dengan di GitHub. Perbedaannya, di komputer lokal terdapat direktori ```.git``` yang digunakan secara internal oleh Git.


## Mengelola Repo

### Mengubah Isi - Push Tanpa Branching dan Merging

Perubahan isi bisa terjadi karena satu atau kombinasi beberapa hal berikut:

1. File dihapus
2. File diedit
3. Membuat file / direktori baru
4. Menghapus direktori

Untuk kasus-kasus tersebut, lakukan perubahan di komputer lokal, setelah itu push ke repo.

![23](images/3/26.1.png)

Cara ini lebih mudah tetapi mempunyai resiko jika terjadi kesalahan dalam edit. Cara yang lebih aman tetapi memerlukan langkah yang lebih panjang adalah branching and merging

### Mengubah Isi dengan Branching and Merging
Dengan menggunakan cara ini, setiap kali akan melakukan perubaham, perubahan itu dilakukan di komputer lokal dengan membuat suatu cabang yang nantinya digunakan untuk menampung perubahan-perubahan tersebut. Setelah itu, cabang itu yang akan dikirim ke repo GitHub untuk dimintai review kemudian digabungkan (merge) ke master. Secara umum, repo yang dibuat biasanya sudah mempunyai satu branch yang disebut dengan master. Cara ini lebih aman, terstruktur, terkendali, dan mempunyai history yang lebih baik. Jika perubahan yang kita buat sudah terlalu kacau dan kita menyesal, maka ada cara untuk "membersihkan" repo lokal kita. Secara umum, langkahnya adalah sebagai berikut:

1. Buat branch untuk menampung perubahan-perubahan
2. Lakukan perubahan-perubahan
3. Add dan commit perubahan-perubahan tersebut ke branch
4. Kembali ke repo master
5. Buat pull request di GitHub
6. Merge pull request di GitHub
7. Merge branch untuk menampung perubahan-perubahan tersebut ke master.
8. Selesai.

![24](images/3/27.1.png)

Setelah itu, kirim pull request (PR):

![25](images/3/31.png)

Setelah membuat PR, PR bisa di-*merge*
![26](images/3/32.1.png)

Setelah itu, ```Confirm Merge```, branch yang kita kirimkan tadi sudah dimasukkan ke branch ```main```. Setelah itu, merge di komputer lokal:

![27](images/3/33.1.png)

### Sinkronisasi
Suatu saat, bisa saja terjadi kita menggunakan komputer lain dan mengedit repo melalui repo lokal di komputer lain, setelah itu pindah ke kamputer lain lagi. Saat itu, kita perlu melakukan sinkronisasi ke kemputer lokal. Perintah untuk sinkronisasi adalah:
```
$ git pull
```

### Membatalkan Perubahan

```
$ git checkout -b edit-readme-2
Switched to a new branch 'edit-readme-2'
$ vim README.md
$ cat README.md
# My Awesome Project

Ini isi proyek. Jadi agak kacau nih
$ git checkout master
M	README.md
Switched to branch 'master'
Your branch is up to date with 'origin/master'.
$ cat README.md
# My Awesome Project

Ini isi proyek. Jadi agak kacau nih
$ git branch -D edit-readme-2
Deleted branch edit-readme-2 (was 7e546b0).
$ git branch
* master
$ cat README.md
# My Awesome Project

Ini isi proyek. Jadi agak kacau nih
$ git reset --hard
HEAD is now at 7e546b0 Merge pull request #1 from oldstager/edit-readme-1
$ cat README.md
# My Awesome Project

Ini isi proyek
$
```

![01](images/3/34.2.png)

### Undo Commit Terakhir

Suatu saat, mungkin kita sudah terlanjur mem-push perubahan ke repo GitHub, setelah itu kita baru menyadari bahwa perubahan tersebut salah. Untuk itu, kita bisa melakukan ```git revert```.

```
$ cat README.md
# My Awesome Project

Ini isi proyek
$ git log --oneline
7e546b0 (HEAD -> master, origin/master, origin/HEAD) Merge pull request #1 from oldstager/edit-readme-1
032d079 (origin/edit-readme-1) Add: isi README.md
2ab2e28 Add: README.md
8dd68d4 Initial commit
$ vim README.md
$ git add -A
$ git commit -m "Add: contents"
[master c55fd06] Add: contents
 1 file changed, 4 insertions(+), 1 deletion(-)
$ git push origin master 
Username for 'https://github.com': oldstager
Password for 'https://oldstager@github.com': 
Enumerating objects: 5, done.
Counting objects: 100% (5/5), done.
Delta compression using up to 8 threads
Compressing objects: 100% (3/3), done.
Writing objects: 100% (3/3), 335 bytes | 335.00 KiB/s, done.
Total 3 (delta 0), reused 0 (delta 0)
To https://github.com/oldstager/awesome-project
   7e546b0..c55fd06  master -> master
$ vim README.md
$  git add -A
$  git commit -m "Add: contents - 2"
[master fed7e79] Add: contents - 2
 1 file changed, 1 insertion(+)
$ git push origin master
Username for 'https://github.com': oldstager
Password for 'https://oldstager@github.com': 
Enumerating objects: 5, done.
Counting objects: 100% (5/5), done.
Delta compression using up to 8 threads
Compressing objects: 100% (3/3), done.
Writing objects: 100% (3/3), 338 bytes | 338.00 KiB/s, done.
Total 3 (delta 0), reused 0 (delta 0)
To https://github.com/oldstager/awesome-project
   c55fd06..fed7e79  master -> master
$ cat README.md
# My Awesome Project

Ini isi proyeka

Ini isi 1

Ini isi 2
$ git status
On branch master
Your branch is up to date with 'origin/master'.

nothing to commit, working tree clean
$
```

![02](images/3/35.1.png)

Contoh di atas adalah contoh untuk mengubah README.md dengan beberapa commit. Setelh itu, kita akan mengembalikan ke posisi terakhir sebelum commit terakhir.

```$ git revert HEAD```

Perintah di atas akan membuka editor. Pada editor tersebut kita bisa mengetikkan pesan revert ( = pesan commit untuk pembatalan). Setelah selesai, simpan:

```
[master f800ced] Revert "Add: contents - 2"
 1 file changed, 1 deletion(-)
$
```
Selanjutnya, tinggal di-push ke repo GitHub.
```
$ git status
On branch master
Your branch is ahead of 'origin/master' by 1 commit.
  (use "git push" to publish your local commits)

nothing to commit, working tree clean
$ git push origin master
Username for 'https://github.com': oldstager
Password for 'https://oldstager@github.com': 
Enumerating objects: 5, done.
Counting objects: 100% (5/5), done.
Delta compression using up to 8 threads
Compressing objects: 100% (3/3), done.
Writing objects: 100% (3/3), 347 bytes | 347.00 KiB/s, done.
Total 3 (delta 1), reused 0 (delta 0)
remote: Resolving deltas: 100% (1/1), completed with 1 local object.
To https://github.com/oldstager/awesome-project
   fed7e79..f800ced  master -> master
$ cat README.md
# My Awesome Project

Ini isi proyeka

Ini isi 1

$
```
![24](images/3/36.1.png)

Jika commit sudah dilakukan, tetapi belum di-push ke repo GitHub (masih berada di lokal), cara membatalkannya:
```
$ vim README.md
$ cat README.md
# My Awesome Project

Ini isi proyeka

Ini isi 1

Ini isi tambahan 1

$ git add -A
$ git commit -m "Add: isi tambahan 1"
[master 42e6bf0] Add: isi tambahan 1
 1 file changed, 2 insertions(+)
$ git status
On branch master
Your branch is ahead of 'origin/master' by 1 commit.
  (use "git push" to publish your local commits)

nothing to commit, working tree clean
$ git log --oneline
42e6bf0 (HEAD -> master) Add: isi tambahan 1
f800ced (origin/master, origin/HEAD) Revert "Add: contents - 2"
fed7e79 Add: contents - 2
c55fd06 Add: contents
7e546b0 Merge pull request #1 from oldstager/edit-readme-1
032d079 (origin/edit-readme-1) Add: isi README.md
2ab2e28 Add: README.md
8dd68d4 Initial commit
$ git reset --hard HEAD^
HEAD is now at f800ced Revert "Add: contents - 2"
$ git status
On branch master
Your branch is up to date with 'origin/master'.

nothing to commit, working tree clean
$ cat README.md
# My Awesome Project

Ini isi proyeka

Ini isi 1

$
```
![24](images/3/37.1.png)

Untuk kembali ke perubahan pada saat yang sudah lama, yang perlu dilakukan adalah melakukan ```git revert <posisi>``` kemudian mengedit secara manual kemudian push ke repo.

```$ cat README.md
# My Awesome Project

Ini isi proyeka

Ini isi 1

Ini isi 2

Ini isi 3
$ git log --oneline
b14810f (HEAD -> master, origin/master, origin/HEAD) Add: isi 3
a7615fb Add: isi 2
f800ced Revert "Add: contents - 2"
fed7e79 Add: contents - 2
c55fd06 Add: contents
7e546b0 Merge pull request #1 from oldstager/edit-readme-1
032d079 (origin/edit-readme-1) Add: isi README.md
2ab2e28 Add: README.md
8dd68d4 Initial commit
$ git revert a7615fb
error: could not revert a7615fb... Add: isi 2
hint: after resolving the conflicts, mark the corrected paths
hint: with 'git add <paths>' or 'git rm <paths>'
hint: and commit the result with 'git commit'
$
```
![24](images/3/38.1.png)

Setelah itu, jika dilihat pada file, akan muncul tambahan untuk memudahkan meng-edit. File ini harus di-resolve terlebih dahulu, setelah itu baru di add dan commit:

```$ cat README.md
# My Awesome Project

Ini isi proyeka

Ini isi 1

<<<<<<< HEAD
Ini isi 2

Ini isi 3
=======
>>>>>>> parent of a7615fb... Add: isi 2
```
Edit file tersebut, setelah itu simpan.
```
$ vim README.md 
$ cat README.md
# My Awesome Project

Ini isi proyek

Ini isi 1

Ini isi 2 setelah revert

Rni isi 3
$ git status
On branch master
Your branch is up to date with 'origin/master'.

You are currently reverting commit a7615fb.
  (fix conflicts and run "git revert --continue")
  (use "git revert --abort" to cancel the revert operation)

Unmerged paths:
  (use "git reset HEAD <file>..." to unstage)
  (use "git add <file>..." to mark resolution)

	both modified:   README.md

no changes added to commit (use "git add" and/or "git commit -a")
$
```
Setelah itu, lanjutkan proses revert. Saat git revert --continue isikan pesan revert.
```
$ git add README.md
$ git status
On branch master
Your branch is up to date with 'origin/master'.

You are currently reverting commit a7615fb.
  (all conflicts fixed: run "git revert --continue")
  (use "git revert --abort" to cancel the revert operation)

Changes to be committed:
  (use "git reset HEAD <file>..." to unstage)

	modified:   README.md

$ git revert --continue
[master dec93e1] Revert "Add: isi 2"
 1 file changed, 3 insertions(+), 3 deletions(-)
$ git status
On branch master
Your branch is ahead of 'origin/master' by 1 commit.
  (use "git push" to publish your local commits)

nothing to commit, working tree clean
$ git push origin master
Username for 'https://github.com': oldstager
Password for 'https://oldstager@github.com': 
Enumerating objects: 5, done.
Counting objects: 100% (5/5), done.
Delta compression using up to 8 threads
Compressing objects: 100% (3/3), done.
Writing objects: 100% (3/3), 401 bytes | 401.00 KiB/s, done.
Total 3 (delta 0), reused 0 (delta 0)
To https://github.com/oldstager/awesome-project
  b14810f..dec93e1  master -> master
$
```
![24](images/3/39.1.png)

## Mengelola Repo Sendiri di Organisasi

Repo yang dibuat bisa diletakkan pada account kita maupun berada pada suatu organisasi. Organisasi bisa kita buat sendiri maupun kita dimasukkan menjadi anggota organisasi. Pada dasarnya, bagian ini sama dengan bagian sebelumnya, hanya saja, pada saat membuat repo Owner dari repo adalah organisasi seperti berikut ini:

![01](images/3/43.png)


### Membuat repository dengan command line

```
echo "# playground" >> README.md
git init
git add README.md
git commit -m "first commit"
git branch -M main
git remote add origin https://github.com/UTDI-TC/playground.git
git push -u origin main
```

![0](images/3/44.png)

### Push repository dari command line

```
git remote add origin https://github.com/UTDI-TC/playground.git
git branch -M main
git push -u origin main
```

![0](images/3/45.png)

### Clone

![0](images/3/46.png)
