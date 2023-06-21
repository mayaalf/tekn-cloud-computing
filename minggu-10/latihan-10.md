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
