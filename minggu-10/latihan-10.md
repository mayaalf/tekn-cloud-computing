# Docker Networking Hands-on Lab

## Bagian 1 - Dasar Networking
### Step 1: Perintah Docker Network
Configurasi dengan perintah ```docker netwrok``` dan jalankan perintah tersebut
```
docker netwrok
```
![1](10/1.png)
### Step 2: Daftar Network
Jalankan perintah untuk melihat container yang terbentuk pada Docker Host dengan perintah
```
docker network ls
```
![2](10/2.png)
### Step 3: Memeriksa Network
Perintah untuk melihat detail dari nama, ID, IPAM dan conntainer yang terhubung menggunakan perintah 
```
docker network inspect bridge
```
![3](10/3.png)
### Step 4: Daftar Network Driver Plugins
Menampilkan banyak informasi tentang Docker installasi dengan perintah
```
docker info
```
![4](10/4.png)
## Bagian 2 - Penghubung Dengan Network
### Step 1: Dasar
1. Jalankan perintah ```docker network ls``` untuk memeriksa Docker bridge

   ![5](10/5.png)

   Hasil keluaran dari **bridge** terletak pada penyimpanan lokal
2. Install ```brctl``` command, hal ini dugunakan untuk mengatur, memelihara, memeriksa internet configurasi pada Linux kernel. untuk mendapatkan hal tersebut jalankan perintah
   ```
   apk update
   apk add bridge
   ```
   ![6](10/6.png)
3. Menmapilkan daftar bridge yang ada dalam Docker yang berjalan denga perintah ```brctl show```

   ![7](10/7.png)

4. Jalankan perintah ```ip a``` untuk melihar detail dari docker0 bridge

   ![8](10/8.png)
### Step 2: Menghubungkan dengan Container
1. Membuaat cintainer baru dengan bridge network
   Jalankan perintah ```docker run -dt ubuntu sleep infinity```

   ![9](10/9.png)
2. Jalankan perintah ```docker ps``` untuk melihat container yang berasal dari ubuntu:latelest image berjalan

   ![10](10/10.png)
3. 
