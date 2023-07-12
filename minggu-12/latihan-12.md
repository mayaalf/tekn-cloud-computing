# Docker Orchestration Hands-on Lab

## Bagian 1 : Configure Swam Mode

Membuat container baru pada node 1 dengan menjalankan perintah
```
docker run -dt ubuntu sleep infinity
```
![1](12/1.png)

Melihat container yang berjalan pada node1

![2](12/2.png)

### Step 2.1 Create Manager Node
Inisialisasi Swarm baru, di gabungkan dengan single worker node, jalankan perintah pada node 1
```
docker swarm init --advertise-addr $(hostname -i)
```

![3](12/3.png)

Jalankan ```docker info``` untuk verifikasi bahwa node 1 sudah sukses terfkonfigurasi dengan swam manager node

![4](12/4.png)

### Step 2.2 Join Wokrer Nodes To The Swarm

Copy command terakhir yang tertampil pada verivikasi pada docker info pada node 1 ke node 2 dan node 3

![5](12/5.png)

Node 2

![6](12/6.png)

Node 3

![7](12/7.png)

Lihat status node pada node 1, apakah node 2 dan 3 sudah berhasil join.

![8](12/8.png)

## Bagian 2 : Deploy Applications Across Multiple Hosts 
### Step 2.1 Deploy the application components as Docker services

Jalankan pada node 1 perintah sebagai berikut
```
docker service create --name sleep-app ubuntu sleep infinity
```

![9](12/9.png)

Jalankan perintah untuk verifikasi service create sudah diterima oleh Swarm Manager

![10](12/10.png)

##Bagian 3 : Scale the Application

Pada node 1 jalankan perintah ``` docker service update --replicas 7 sleep-app``` untuk scale nomor container di sleep-app menjadi no 7

![10](12/11.png)

Jalankan perintah untuk melihat container secara real time

![11](12/12.png)

```Terlihat sekarang container pada node 1 berada di urutan no 7```

Ubah service untuk menghapus container 4 sleep-app

![13](12/13.png)

Jalankan perintah untuk menampilkan service sleep-app

![14](12/14.png)

## Bagian 4 Drain a node and reschedule the containers
Jalankan perintah untuk melihat status node pada node1

![15](12/15.png)

Jalankan pada node 2 untuk melihat container yang berjalan

![16](12/16.png)

Balik ke node1 untuk menghapus service pada node 2 

![16](12/17.png)

Gunakan ID node 2 dan jalankan perintah pada node 1

![16](12/29.png)

Cek status

![16](12/20.png)

Balik pada node 2 dan lihat apa yang terjadi

![16](12/21.png)
Container yang berjalan tidak ada pada node 2

Pastikan lagi di node 1 dengan perintah

![16](12/22.png)

## Cleaning UP

Jalankan docker service rm sleep-app pada node 1 untuk menghapus service

![16](12/23.png)

Jalankan perintah untuk lihat container yang berjalan

![24](12/24.png)

Remove semua node 1,2,3 dari swarm dengan perintah

![16](12/25.png)
