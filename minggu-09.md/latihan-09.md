# Docker for Beginners - Linux

## Task 0: Prerequisites
### Clone repositori 

![1](9/1.png)

## Task 1: Menjalankan beberapa container sederhana Docker
### Menjalankan satu tugas dalam container Alpine Linux
1. Jalankan perintah dalam konsol

    ![2](9/2.png)    

2. Menampilkan docker container yang berjalan selama proses dimulai

    ![3](9/3.png)
    
   **Nota**: ID kontainer adalah nama host yang ditampilkan kontainer. Dalam contoh di atas adalah *8114cf7109dd*
   
### Menjalankan kontainer Ubuntu interaktif
1. Menjalankan vontainer docker dan akses shell

    ![4](9/4.png)
 
2. jalankan perintah berikut dalam container

    ![5](9/5.png)
   
3. Ketik **exit** untuk menghentikan sesi shell
 
    ![6](9/6.png)
    
4.Menjalankan untuk memeriksa versi VM host

   ![7](9/7.png)
   
### Menjalankan Background Container MySQL 
1. Menjalankan container MySQL dengan perintah

    ![8](9/8.png)
 
2. Mencantumkan container yang sedang berjalan

    ![9](9/9.png)
    
3. Memeriksa log dari container MySQL Docker

    ![10](9/10.png)   
    
    Jalankan perintah untuk melihat proses yang berjalan pada container
    
    ![11](9/17.png)
    
4.  Cantumkan versi MySQL dengan menggunakan perintah dibahawh ini
    
    ![12](9/11.png)
    
5. Menyambungkan proses shell dalam container yang sudah berjalan
    
    ![13](9/12.png)
    
6. Periksa nomor versi dengan menjalankan perintah yang sama dalam sesi shell baru dalam container
    
    ![14](9/13.png)  
    
7. Ketik **exit** untuk menghentikan sesi shell
 
    ![15](9/14.png)    
    
## Task 2: Mengemas dan menjalankan aplikasi kustom menggunakan Docker
### Membuat gambar situs web sederhana
1. Pastikan sudah berada pada direktori **linux_tweet_app**

    ![16](9/15.png)
    
2.  Menampilkan konten yang ada dalam Dockerfile

    ![17](9/16.png)  
    
3. Menjalankan perintah ```export DOCKERID=<your docker id>``` untuk mengekspor variabel lingkungan yang ada dalam DockerID
    
    ![18](9/18.png)
    
4. Mencetak nilai variabel untuk memastikan disimpan dengan benar

    ![19](9/19.png) 
    
5. Menggunakan perintah untuk membuat image Docker baru dengan menggunakan perintah berikut

   ![20](9/20.png)    
   
6. Perintah menjalankan container baru

   ![21](9/21.png) 
   
7. Klik [website](http://ip172-19-0-3-ci4jndggftqg00ehj120-80.direct.labs.play-with-docker.com/) untuk memastikan web sudah berjalan   
    
    ![22](9/22.png) 
    
8. Matikan dan hapus container
    
    ![23](9/23.png) 
    
## Task 3: Mengubah situs web yang sedang berjalan    
### Memulai aplikasi web dengan bind mount
1. Memulai aplikasi web dan pasang direktori. Pastikan berjalan pada direktori **linux_tweet_app**

    ![24](9/24.png) 
    
2. Menjalankan [situs web](http://ip172-19-0-3-ci4jndggftqg00ehj120-80.direct.labs.play-with-docker.com/) dan ini harus berjalan
    
    ![25](9/25.png) 
    
### Ubah situs yang sedang berjalan
1. Salin ke dalam container index.html

    ![26](9/26.png) 
    
2. Buka [situs](http://ip172-19-0-3-ci4jndggftqg00ehj120-80.direct.labs.play-with-docker.com/) dan segarkan halaman.

    ![27](9/25.png) 
    
3. Hentikan dan hapus container
    
    ![28](9/27.png) 
    
4. Jalankan kembali dengan versi saat ini tanpa bind mount

    ![29](9/28.png)
    
5. uka [situs](http://ip172-19-0-3-ci4jndggftqg00ehj120-80.direct.labs.play-with-docker.com/) dan segarkan halaman, tampilan akan kembali versi awal   

    ![30](9/29.png) 
    
6. Hentikan dan menghapus kontainer saat ini

    ![31](9/27.png)    
    
### Perbarui Image    
