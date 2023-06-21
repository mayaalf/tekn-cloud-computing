<img width="334" alt="image" src="https://github.com/mayaalf/tekn-cloud-computing/assets/106806229/cb61498d-4fdb-4d00-a6b2-2e78068ada28"># Docker Networking Hands-on Lab

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
3. Setelah container terhubung dan dirambahkan pada bridge network, jalankan perintah ```brctl show```

   ![11](10/11.png)

4. Memeriksa bridge network kembali dengan menjalankan perintah ```docker network inspect bridge```

   ![12](10/12.png)
### Step 3: Test Koneksi Nework
1. Lakukan perintah ```docker network inspect``` dengan ping koneksi dengan perintah ```ping -c5 <IPv4 Address>```

   ![13](10/13.png)
2. Menjalankan container dengan perintah ```docker ps```
   
   ![14](10/14.png)
3. Menjalankan ubuntu container dengan menjalankan ```docker exec -it <CONTAINER ID> /bin/bash```

   ![15](10/15.png)
4. Kemudian jalankan perintah untuk install ping program dengan perintah ```apt-get update && apt-get install -y iputils-ping```

   ![16](10/16.png)
5. Jalankan perintah ping ```ping -c5 www.github.com```

   ![17](10/17.png)
6. Lakukan pemutusan koneksi dengan containner dengan perintah exit
   
   ![18](10/18.png)
