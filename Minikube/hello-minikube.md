# Kubernetes - Intro
## Install minikube
1. Download dan jalankan intsaller rilis versi terakhir atu dengan menggunakan command 
pada powershell

    ![1](13/1.png)

    ![2](13/3.png)

2. Tambahkan minikube.exe pada path, jalankan pada powershell administrator

    ![3](13/2.png)

## Membuat Minikube Cluster dengan menjalankan perintah ```minikube start```

![4](13/4.png)

 ## Buka Dashboard

![4](13/5.png)

Direct ke web browser

![4](13/6.png)

Penjelasan:
Dashboard ini digunakan untuk mengaktifkan add-on dashboard dan membuka proxy di 
browser.

## Create Deployment, gunakan powershell baru dengan tidak menghentikan layanan yang sebelumnya dijalankan
1. Gunakan perintah kubectl create untuk menjalankan create deployment yang 
memanajemen POD. POD menjalankan container berdasarkan Docker image.

    ![4](13/7.png)

2. Melihat deployment

    ![4](13/8.png)
   
3. View the Pod

     ![4](13/9.png)
   
4. View cluster event
    
    ![4](13/10.png)

5. View the kubectl configuration

    ![4](13/11.png)
## Membuat Layanan
1. Expose Pod ke internet publik dengan perintah kubectl expose

   ![4](13/12.png)
Kode aplikasi di dalam gambar uji hanya mendengarkan pada port TCP 8080. Jika Anda 
biasa mengekspos port yang berbeda, klien tidak dapat terhubung ke por lain
2. View layanan yang dibuat

    ![4](13/13.png)
3. Jalankan perintah untuk run service

     ![4](13/14.png)

     ![4](13/15.png)
Penjelasan :
Perintah ini akan menjalankan broweser yang menjalankan layanan aplikasi yang 
dibuat.
## Enable addons
1. List addons yang sedang dijalankan

   ![4](13/16.png)
   
2. Enable addon dengan metric-server

    ![4](13/17.png)
   
5. Lihat Pod dan layanan yang dibuat dengan menginstall addon

     ![4](13/18.png)

6. Disable matrics-server

     ![4](13/19.png)

 ## Clean Up
1. Bersihkan sumber yang dibuat dari cluster

   ![4](13/20.png)
   
2. Stop minikube cluster

   ![4](13/21.png)
   
3. Hapus minikube VM

   ![4](13/22.png)
