# Application Containerization and Microservice Orchestration

## Setup
Mulai dengan cloning repository demo

![1](11/1.png)

## Step 0: Basic Link Extractor Script
1. Melihat list pada cabang step0

   ![2](11/2.png)
   
2. Terdapat file linkextractor.py, untuk melihat isi file gunakan perintah
  
   ![3](11/3.png)

3. Menjalankan file

   ![4](11/4.png)
  
4. Ketika mencoba menjalankannya sebagai skrip, kami mendapat kesalahan. Mari kita periksa izin saat ini pada file

   ![5](11/5.png)
   
5. Menjalankannya sebagai program Python

   ![6](11/6.png)

## Step 1: Containerized Link Extractor Script
1. Periksa daftar pada step1

   ![7](11/7.png)

2. Melihat isi dockerfile
  
   ![8](11/8.png)
   
3. Membuat image pada step1

   ![9](11/9.png)

4. Setelah build image berhasil lakukan untuk melihat isi image

   ![10](11/10.png)

5. Menjalankan container dengan image dan web extract link
  
   ![11](11/11.png)

6. Mencoba web page dengan banyak link

   ![12](11/12.png)

## Step 2: Link Extractor Module with Full URI and Anchor Text
1. Periksa daftar pada step2

   ![13](11/13.png)

2. Melihat isi file linkextractor.py terbaru
  
   ![14](11/14.png)
   
3. Membuat image baru pada step2

   ![15](11/15.png)

4. Setelah build image berhasil lakukan untuk melihat isi image

   ![16](11/16.png)

5. Menjalankan container dengan image dan web extract link
  
   ![17](11/17.png)

6. Menjalankan container dengan image sebelumnya

   ![18](11/18.png)

## Step 3: Link Extractor API Service
1. Periksa daftar pada step3

   ![19](11/19.png)

2. Melihat isi file dockerfile terbaru
  
   ![20](11/20.png)
   
3. Menambah file baru bernama main.py

   ![21](11/21.png)

4. Membuat image baru

   ![22](11/22.png)

5. Menjalankan container dengan port 5000

   ![23](11/23.png)
   
6. Setelah build image berhasil lakukan untuk melihat isi image

   ![24](11/24.png)

7. Memembuat permintaan HTTP dalam bentuk untuk berbicara dengan server ini dan mengambil respons yang berisi tautan yang diekstraksi
  
   ![25](11/25.png)

8. Melihat catatan yang terjadi pada container menggunakan log

   ![26](11/26.png)

9. Menonaktifkan dan menghapus container

   ![27](11/27.png)

## Step 4: Link Extractor API and Web Front End Services
1. Periksa daftar pada step4

   ![28](11/28.png)

2. Melihat isi file docker-compose.yml
  
   ![29](11/29.png)
   
3. Melihat script yang ada pada tampilan pengguna

   ![30](11/30.png)

4. Membawa layanan dalam mode terpisah menggunakan utilitas docker-compose

   ![31](11/31.png)

5. Memeriksa daftar kontainer yang berjalan menegaskan bahwa kedua layanan tersebut memang berjalan

   ![32](11/32.png)
   
6. meminta izin pada API Service untuk akses pada interface web [link Extractor](http://ip172-18-0-78-cig2de4snmng00fa3omg-80.direct.labs.play-with-docker.com/?url=https%3A%2F%2Fgithub.com%2Fmayaalf%3Ftab%3Drepositories)

   ![33](11/33.png)

7. Modifikasi www/index.php dengan mengganti semua kejadian dari Link Extractor dengan Super Link Extractor
  
   ![34](11/34.png)

   Klik link ini [click here](http://ip172-18-0-78-cig2de4snmng00fa3omg-80.direct.labs.play-with-docker.com/?url=https%3A%2F%2Fgithub.com%2Fmayaalf%3Ftab%3Drepositories)
   
   ![35](11/35.png)

8. Mengembalikan perubahan ini sekarang untuk membersihkan pelacakan Git

   ![36](11/36.png)

9. Sebelum pindah pada step berikutnya, lakukan shut service down untuk membantu menjaga Docker Compose

    ![37](11/37.png)

## Step 5: Redis Service for Caching
1. Periksa daftar pada step5

   ![38](11/38.png)

2. Melihat penambahan baru pada dockerfile 
  
   ![39](11/39.png)
   
3. Selanjutnya, melihat file server API tempat kita menggunakan cache Redis:

   ![40](11/40.png)

4. Melihat perubahan pada docker-compose.yml

   ![41](11/41.png)

5. Melakukan boot layanan dengan perintah ini
  
   ![41](11/42.png)

6. Clik akses interface dengan [Link Extractor](http://ip172-18-0-78-cig2de4snmng00fa3omg-80.direct.labs.play-with-docker.com/?url=https%3A%2F%2Fgithub.com%2Fmayaalf%3Ftab%3Drepositories)

   ![43](11/43.png)

7. Gunakan docker-compose exec untuk monitor Redis CLI Command

   ![44](11/44.png)

   Menjalankan link extractor
   
   ![45](11/45.png)

8. Verifikasi perubahan yang dibuat dalam lokal tidak mencerminkan dalam layananan yang sedang berjalan

   ![46](11/46.png)

9. Mencerminkan layanan yang berjalan

   ![47](11/47.png)

10. Shut down service

     ![48](11/48.png)

## Step 6: Swap Python API Service with Ruby
1. Periksa daftar pada step6

   ![49](11/49.png)   

2. Melihat penambahan baru
   
   ![50](11/50.png)
   
3. melihat api pada dockefile

  ![51](11/51.png)   

4. Selanjutnya, melihat file server API Ruby pada Docker-compose.yml

   ![52](11/52.png)

6. Membuat build baru layanan

   ![53](11/53.png)

7. Melakukan akses APIdengan port terbaru
  
   ![54](11/54.png)
   
8. Clik akses interface dengan [Link Extractor](http://ip172-18-0-78-cig2de4snmng00fa3omg-80.direct.labs.play-with-docker.com/?url=https%3A%2F%2Fgithub.com%2Fmayaalf%3Ftab%3Drepositories)

   ![55](11/55.png)

9. Gunakan tail dengan perintah -f untuk melihat output
    
   ![56](11/56.png)

   Menjalankan link extractor
   
   ![57](11/57.png)

10. Verifikasi perubahan yang dibuat dalam lokal tidak mencerminkan dalam layananan yang sedang berjalan

   ![58](11/58.png)

11. Karena telah mempertahankan log, log tersebut harus tetap tersedia setelah layanan hilang

   ![59](11/59.png)
