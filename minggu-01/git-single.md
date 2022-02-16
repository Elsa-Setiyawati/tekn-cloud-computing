## _215611104 - Elsa Setiyawati_

# Latihan

## 1. Instalasi GIT di Windows

Sebelum install maka download terlebih dahulu file master Git.

1. Setelah itu, jalankan file master tersebut dengan cara klik double filenya. Maka akan muncul gambar di bawah ini lisensi git. Klik **Next** .

![01](images/Git_Install/1.png)

2. Kemudian, pilih lokasi instalasi. Secara default akan terisi C:\Program Files\Git. Ganti lokasi jika memang anda menginginkan lokasi lain. Klik **Next** .

![02](images/Git_Install/2.png)

3. Klik **Next** pada pilih komponen, karena pengaturan sudah default.

![03](images/Git_Install/3.png)

4. Setting nama shortcut, klik **Next** untuk melanjutkan atau ganti nama sesuai keinginan.

![04](images/Git_Install/4.png)

5. Pilih editor untukk Git. Klik **Next**.

![05](images/Git_Install/5.png)

6. Menambahkan path environment untuk Git, klik **Next**.

![06](images/Git_Install/6.png)

7. Setting HTTPS yang akan digunakan, klik **Next**.

![07](images/Git_Install/7.png)

8. Setting pengaturan text pada editor untuk Git. Klik **Next**.

![08](images/Git_Install/8.png)

9. Setting konfigurasi terminal emulator dengan menggunakan Git Bash. klik **Next** .

![09]((images/Git_Install/9.png)

10. Memilih default untuk git pull. klik **Next**.

![10](images/Git_Install/10.png)

11. Konfigurasi fitur untuk dienable. klik **Next**.

![11](images/Git_Install/11.png)

12. Lalu, klik Install

![12](images/Git_Install/12.png)

13. Proses instalasi berjalan, tunggu hingga selesai.

![13](images/Git_Install/13.png)

14. Proses instalasi telah selesai, pilih yes jika ingin restart komputer sekarang, atau pilih no jika belum akan merestart komputer. klik **Next**

![14](images/Git_Install/14.png)

15. Mengecek apakah git sudah terinstall dan mengetahui versinya. dengan perintah "git --version".

![15](images/Git_Install/15.png)

## 2. Konfigurasi GIT

konfigurasi git dilakukan dengan perintah berikut ini :

1. Berikan perintah berikut, dengan nama serta email yang digunakan untuk mendaftar di GitHub.

```sh
$ git config --global user.name "Nama Anda di GitHub"
$ git config --global user.email email@domain.tld
```

2. untuk melihat hasil konfigurasi dengan perintah

```sh
$ git config --list
```

untuk hasilnya seperti ini

![20](images/Git_Konfigurasi/1.png)

## 3. Mengelola Repo Sendiri di Account Sendiri

Berikut ini langkah - langkah untuk membuat repository sendiri:

1. Klik tanda + pada bagian atas setelah login, pilih **New repository**

![20](images/Git_3/1.png)

2. Isikan nama, lalu memilih mau public atua prive repo yang akan dibuat. klik **Create repository**.

![20](images/Git_3/2.png)

3. Hasil repository yang telah dibuat.

![20](images/Git_3/3.png)

## 3. Mengelola Repo Sendiri di Organisasi

Untuk membuat repository organisasi, masuk pada Organisasi Kalian lalu Klik tanda + pada bagian atas setelah login, pilih **New repository** , lalu memilih mau public atua prive repo yang akan dibuat. klik **Create repository**. lalu, hasilnya seperti gambar berikut ini.

![20](images/Git_3/4.png)
