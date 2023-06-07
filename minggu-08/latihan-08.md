# Getting Started - Docker Compose

## Langkah 1: Membuat Depedensi Aplikasi
1. Buat diretori untuk projek 

    ![1](8/1.png)
    
2. Buat file bernama app.py dalam folder projek dan paste code ini :
    
    ![2](8/2.png)
    
3. Buat file lain bernama requirements.txt pada projek 

    ![3](8/3.png)
 
## Langkah 2: Buat Dockerfile

Dockerfile digunakan untuk membuat image yang berhubungan dengan Python. Buat file Docker file

![4](8/4.png)

## Langkah 3: Definisikan Service dalam Compose File
Buatlah file bernama docker-compose.yml dalam project 

![5](8/5.png)

## Langkah 4: Bangun dan Jalankan Aplikasi dengan Compose

1. Dari lokasi direktori projek jalankan perintah ```docker compose up```

    ![6](8/6.png)
    ![7](8/7.png)
    
2. Jalankan URL [http://localhost:8000/](http://localhost:8000) pada browser dan lihat apa yang ditampilkan

    ![8](8/8.png)
    
3. Refresh browser dan angka yang durampilkan berubah mengikuti perhitungan

    ![9](8/9.png)
    
4. Pindah ke terminal yang berbeda untuk menjalankan perintah ```docker image ls```

    ![10](8/10.png)   
    
5. Hentikan aplikasi dengan perintah ```docker compose down``` atau dengan klik CTRL + C pada terminal

    ![11](8/12.png)
    
## Langkah 5: Edit File Compose dan Tambahkan Bind Mount
Edit file docker-coompose.yml

![12](8/11.png)

## Langkah 6: Re-Build dan jalankan Aplikasi dengan Compose
1. Jalankan aplikasi dengan mengetikkan perintah ```docker compose up```

    ![13](8/13.png)
    
2. Cek pada Hello World pada browser, refresh dan lihat jumlah dari increment

    ![14](8/14.png)
    
 ## Langkah 7: Update Aplikasi
 
 1. Ganti salam pada file app.py dari ```Hello World``` menjadi ```Hello From Docker```
 
     ![15](8/15.png)
     
 2. Jalankan kembali browser dengan cara refersh dan akan terlihat perubahan yang sudah dilakukan

      ![16](8/16.png)
      
## Langkah 8: Percobaan dengan Beberapa Perintah
1. Perintah ```docker compose run``` akan mengizinkan untuk run-one-off dari servis. Contoh :

    ![17](8/17.png)
    
2. Jika menjalankan Compose dengan perintah ```docker compose up -d```, sedangkan untuk menghentikan servis menggunakan perintah ```docker compose stop```

    ![18](8/18.png)
    
 3. Untuk menghapus container menggunakan perintah ```down``` dilanjutkan dengan ```--volumes```

    ![19](8/19.png)
