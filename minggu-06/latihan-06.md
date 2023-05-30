# Installasi Software

## Install GO Programming Language

1. Masuk pada [URL download GO](https://go.dev/dl/), lalu pilih fitur yang sesuai dengan spesifikasi laptop yang dimiliki

    ![1](6/1.png)

2. Setelah proses download berhasil, lalu lakukan proses installasi. Klik **Next**

    ![2](6/2.png)
    
3. Klik _'I accept the terms in the License Agreement'_. Klik **Next**

    ![3](6/3.png)
    
4. Pilihlah lokasi penyimpanan untuk install GO. Klik **Next**

    ![4](6/4.png)

5. Klik **Install** lalu tunggulah proses installasi sampai selesai. Klik Finish

    ![5](6/5.png)
    
    ![6](6/6.png)
    
    ![7](6/7.png)
        
## Install MySQL

1. Masuk pada [URL download MySQL](https://www.mysql.com/products/community/), lalu klik _MySQL Intsaller Windows_ dan kemudian pilih fitur yang sesuai dengan spesifikasi laptop yang dimiliki

    ![8](6/8.png)
    
    ![9](6/9.png)
    
    ![10](6/10.png)

2. Setelah proses download berhasil, lalu lakukan proses installasi. Klik **Next**

    ![11](6/12.png)
    
3. Pilihlah lokasi penyimpanan untuk install GO. Klik **Next**

    ![12](6/13.png)

5. Klik **Next** lalu klik **Execute** tunggulah sampai proses download paket selesai. Klik Finish

    ![13](6/14.png)
    
## Install MongoDB

1. Masuk pada [URL download MongoDB](https://www.mongodb.com/try/download/community), lalu pilih fitur yang sesuai dengan spesifikasi laptop yang dimiliki

    ![14](6/15.png)

2. Setelah proses download berhasil, lalu lakukan proses installasi. Klik **Next**

   ![15](6/16.png)
    
3. Klik _'I accept the terms in the License Agreement'_. Klik **Next**

    ![16](6/17.png)
    
4. Pilihlah tipe Complete untuk konfigurasi MongoDB. Klik **Next**

    ![17](6/18.png)
    
    ![18](6/19.png)
    
5. Klik **Install** lalu tunggulah proses installasi sampai selesai. Klik Finish

    ![19](6/20.png)
    
    ![20](6/21.png)
   

# Program Koneksi dan Baca Data Pada MySQL dan MongoDB 
## Program Koneksi dan Baca Data Pada MySQL
 
1. Buatlah folder kerja untuk membuat file kerja koneksi MySQL bernama **pcloud6**
2. Masuk pada cmd untuk insialisasi folder projek dengan mengetikkan perintah seperti berikut :

    ![22](6/23.png)

3. Ketikkan perintah untuk mengunduh library mysql 

    ![23](6/27.png)
 
4. Ketikkan perintah mod tidy untuk melakukan validasi dependensi.

    ![24](6/28.png)
    
5. Membuat database bernama **test** dengan nama tabel **mahasiswa** 

    ![25](6/25.png)
    
6. Ketikkan code dengan menggunakan library dari database mysql untuk nantinya dapat dibaca data yang telah tersimpan dalam database.

    ![26](6/33.png)
    
7. Jalankan perintah untuk membaca file database yang sudah dibuat

    ![27](6/26.png)
    
    
## Program Koneksi dan Baca Data Pada MongoDB 

1. Membuat inisailisas folder projek kerja yang bernama cloudMongo

    ![28](6/31.png)
   
2. Ketikkan perintah untuk mengunduh library mongoDB

    ![29](6/24.png)
    
3. Ketikkan code dengan menggunakan library dari database mongoDB untuk nantinya dapat dibaca data yang telah diinputkan untuk tersimpan pada database

    ![30](6/29.png)

4. Jalankan perintah untuk membaca file database yang sudah dibuat

    ![31](6/30.png)

5. Cek kembali pada cmd mongoBD untuk memastikan database sudah terbuat dengan baik dengan perintah ```show dbs```

    ![32](6/32.png)
    

