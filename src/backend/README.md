# DNA Test - Backend

*Backend* dari Tugas Besar 3 Strategi Algoritma

## *Run* di Lokal

 1. Pastikan `DATABASE_URL` sudah dijadikan *environment variable*. ([detail](#DATABASE_URL))
 2. lakukan `go build`, lalu `./dna-matcher`.
 3. *Backend* akan jalan di lokal (Biasanya `localhost:8080`)

## Deploy

 1. Di *Dashboard Heroku*, pergi ke *tab* *Deploy*.
 2. Di *Manual Deploy*, pilih *branch* **backend**, lalu klik *deploy*.
 3. Bisa dicoba *Automatic Deploy*, tapi takutnya *chaos* kalau banyak *commit* sekaligus.

## `DATABASE_URL`

 1. Di *Dashboard Heroku*, pergi ke *tab* *Resources*.
 2. Di *Add-ons*, pilih *Heroku Postgres*.
 3. Di *Heroku Postgres*, pilih *Settings* lalu pilih *Database Credentials*.
 4. `DATABASE_URL` ada di `URI`.

Untuk menjadikan `DATABASE_URL` sebagai *environment variable*:

```
export DATABASE_URL=[Data URI dari langkah diatas]
```

## Connect ke Database

Disini menggunakan `psql` sebagai *front-end* *PostgreSQL*.

```
psql $DATABASE_URL
```

`$DATABASE_URL` merupakan *environment variable* yang berlaku sebagai *credentials* untuk mengakses database.
