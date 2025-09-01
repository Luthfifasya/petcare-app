# petcare-app

-Menggunakan Gin FrameWork 

-Menggunakan payment gateaway midtrans

-Menggunakan Gorm sebagai Relational Object 

-Menggunakan database Postgres 

-Menggunakan Middleware JWT

USERS
post /api/users/login --> untuk login dan dapatkan token jwt 
post /api/users/register --> untuk register akun 
get /api/users --> untuk melihat akun yang sudah didaftarkan/yg ada di database

contoh body - raw - json di postman {
    "email": "luthfi@example.com", 
    "password": "admin" 
}

PETS
==> get /api/pets Method yang digunakan adalah get. Api ini digunakan untuk menampilkan seluruh pet

==> post /api/pets Method yang digunakan adalah post. Api ini digunakan untuk menambahkan pet.

==> get /api/pets/:id Method yang digunakan adalah get. Api ini digunakan untuk menampilkan detail pet by id.

==> put /api/pets/:id Method yang digunakan adalah put. Api ini digunakan untuk mengedit pet by id.

==> delete /api/pets/:id Method yang digunakan adalah delete. Api ini digunakan untuk mendelete pet.

contoh body - raw - json di postman 
{
  "user_id": 1,
  "name": "Milo",
  "species": "Dog",
  "breed": "Golden Retriever",
  "age": 3
}

TREATMENT
==> get /api/treatment Method yang digunakan adalah get. Api ini digunakan untuk menampilkan seluruh treatment

==> post /api/books Method yang digunakan adalah post. Api ini digunakan untuk menambahkan treatment

==> get /api/treatment/:id Method yang digunakan adalah get. Api ini digunakan untuk menampilkan detail treatment by id.

==> put /api/treatment/:id Method yang digunakan adalah put. Api ini digunakan untuk mengedit treatment by id.

==> delete /api/treatment/:id Method yang digunakan adalah delete. Api ini digunakan untuk mendelete treatment by id.

contoh body - raw - json di postman 
{
  "name": daycare,
  "description": "jasa penitipan pet",
  "price": 70000,
}

APPOINTMENT
==> get /api/appointment Method yang digunakan adalah get. Api ini digunakan untuk menampilkan seluruh appointment

==> post /api/appointment Method yang digunakan adalah post. Api ini digunakan untuk menambahkan appointment

==> get /api/appointment/:id Method yang digunakan adalah get. Api ini digunakan untuk menampilkan detail appointment by id.

==> put /api/appointment/:id Method yang digunakan adalah put. Api ini digunakan untuk mengedit appointment by id.

==> delete /api/appointment/:id Method yang digunakan adalah delete. Api ini digunakan untuk mendelete appointment by id.

contoh body - raw - json di postman
{
  "pet_id": 5,
  "treatment_id": 1,
  "date": "2025-09-23T10:00:00+07:00"
}


