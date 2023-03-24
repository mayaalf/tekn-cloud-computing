# Signup Heroku
1. Masuk dalam [URL Heroku](https://signup.heroku.com/) untuk membuat account

    
    ![01](m3/1.png)

2. Setelah masuk pada link tersebut maka akan ditampilkan form pembuatan akun. Isikan data yang diminta lalu klik **create an account**

    
    ![02](m3/2.png)

3. Buat akun akan diminta verifikasi email, bukalah gmail anda 

    
    ![03](m3/3.png)

4. Buat password yang memenuhi 3 kriteria yang diminta oleh Heroku

    
    ![04](m3/4.png)

5. Akun siap digunakan, klik **Click Here to Proceed*

    
    ![05](m3/5.png)

# Getting Started on Heroku With Python

1. Login Heroku menggunakan command 

    
    ![06](m3/6.png)

    
    ![06.1](m3/6.1.png)

## Prepare The App

1. Clone contoh aplikasi pada penyimpanan lokal yang akan di deploy pada Heroku

    
    ![07](m3/9.png)

## Deploy The App

1. Membuat aplikasi pada Heroku

    
    ![08](m3/10.png)

    
    Saat membuat aplikasi pada git remote 'heroku' akan terbentuk nama aplikasi secara random yang bernama 'sleepy-beyond-27600'.

2. Deploy 

    
    ![09](m3/11.png)

    
    Aplikasi sudah berhasil dideploy.

3. Membuka aplikasi pada website

    
    ![10](m3/12.png)

    
    ![11](m3/13.png)

```Hasil : ```

    
   ![12](m3/13.1.png)


## View Logs

Melihat infromasi aplikasi dengan perintah logs ```heroku logs --tail```

![13](m3/14.png)

## Define Procfile

Procfile adalah file teks pada direktori root pada aplikasi. Contoh deploy aplikasi yang memiliki isi Procfile :

![14](m3/15.png)

#### Microsoft Windows
Lokal deploy pada Microsoft Windows terletak pada file Procfile.windows

![15](m3/15.1.png)

## Scale The App

Menjalankan web tungak Dyno dengan perintah ps ```heroku ps```

![16](m3/16.png)

Skala aplikasi pada Heroku akan berubah jika dynos scale pada web menjadi 0

![17](m3/17.png)

Kemudian ketikkan perintah ```heroku open```, lihatlah perubahan yang terjadi pada loading web browser.

![18](m3/17.1.png)

Sekarang kembalikan pengaturan pada scale web 1

![19](m3/18.png)

## Install App Dependencies Locally

Heroku mengenali aplikasi seperti aplikasi Python dengan melihat kata kunci file termasuk ```requirements.txt``` pada direktori root yang berisikan :

![20](m3/19.png)

Heroku membaca file ini dengan perintah ```pip install -r```

![21](m3/20.png)

Install pada local dengan membuat *"Virtual Environment (venv)"* 

![22](m3/21.png)

Pada akhir dapat menginstall pada environment baru dengan perintah :

![23](m3/22.png)


## Run The App Locally

Aplikasi hampir siap untuk dimulai secara local. Menggunakan Django sehingga perlu running ```collectstatic```, respon dengan Y or yes.

![24](m3/23.2.png)

Hentikan program dengan ```ctrl``` + ```C```.

Jalankan dengan [https://localhost:5000] pada browser

![25](m3/23.1.png)

## Push Local Changes

Modifikasi pada file ```requirements.txt``` dengan menambahkan ```requests```

![26](m3/24.png)

Update dengan perintah pip install

![27](m3/24.1.png)

Ubahlah file views.py dalam direktori hello dengan modul requests

![28](m3/24.2.png)
