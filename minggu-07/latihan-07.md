# Installasi Docker

1. Masuk pada halaman [Docker](https://docs.docker.com/get-docker/) untuk mendapatkan aplikasi Docker sesuai dengan spesifikasi yang dimiliki.
   ![1](7/1.png)

2. Jika file sudah selesai terdownload lanjutkan proses installasi sampai selesai 

   ![2](7/2.png)
  
    ![3](7/3.png)

3. Jika proses installasi sudah selesai klik _'Close and restart'_ agar Docker dapat bekerja dengan baik dalam komputer.

   ![4](7/4.png)

# Get Started With Docker

1. Lakukan proses login Docker

   ![5](7/5.png)
  
    ![6](7/6.png)
  
2. Update subsistem windows untuk dipasang sistem Linux

    ![7](7/8.png)
    
3. Proses Login berhasil dilakukan, Docker dapat digunakan

    ![8](7/9.png)
  
4. Clone repository getting-started dengan perintah sebagai berikut 

    ![9](7/10.png)
 
 5. Masuk pada folder getting-started/app dan akan terlihat file package.json dan 2 subdirectory bernama src dan spec

    ![10](7/12.png)
  
6. Masuk pada directory getting-started/app kemudian ketikkan perintah ```type nul > Dockerfile``` untuk membuat file kosong bernama Dockerfile

    ![11](7/11.png)
  
 7. Gunakan editor untuk menambahkan kode seperti berikut 
      ```
      # syntax=docker/dockerfile:1

      FROM node:18-alpine
      WORKDIR /app
      COPY . .
      RUN yarn install --production
      CMD ["node", "src/index.js"]
      EXPOSE 3000
      ```

    ![12](7/13.png)
8. Membuat container image menggunakan perintah sebagai berikut 
    ```
    docker build -t getting-started .
    ```
    
    ![13](7/14.png)
    
## Menjalankan App Container

1. Jalankan Docker dengan perintah ```docker run```  dan spesifik nama image yang dibuat 
    ```
    docker run -dp 3000:3000 getting-started
    ```
    
    ![14](7/16.png)

2. Setelah itu buka browser dan masuk pada http://localhost:3000

    ![15](7/17.png)
    
3. Cobalah untuk mengisikan data
    
    ![16](7/18.png)
    
 4. Jalankan perintah ```docker ps``` pada terminal untuk mengetahui list container yang telah dibuat

    ![17](7/19.png)
    
    
## Update Aplikasi

1. Update aplikasi dengan mengubah beberapa line code pada file app.js

    ![18](7/20.png)

2. Buatlah updated versi dari image dengan nama docker build dengan perintah 
    ```
    docker build -t getting-started .
    ```
   
    ![19](7/22.png)
    
3. Jalankan container baru. Akan terlihat pesan error pada ID yang berbeda. Kenapa terjadi error? hal ini dikarenakan container lama yang masih berjalan sehingga tidak dapat menjalankan container baru.
    
    ![20](7/21.png)
    
## Menghapus Container Lama

1. Jalankan perintah ```docker ps``` 

   ![21](7/24.png)

2. Gunakan perintah ```docker stop``` untuk menghentikan container

   ![22](7/23.png)

3. Hapus container dengan perintah ```docker rm```

    ![23](7/25.png)
    
    
## Jalankan App Container Terbaru

1. Jalankan perintah docker ```docker run``` untuk menjalankan container terbaru

    ![24](7/26.png)
    
2. Refresh browser untuk menjalankan kembali https:localhost:3000

    ![25](7/27.png)
