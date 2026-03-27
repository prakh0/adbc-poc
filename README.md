# ADBC MySQL Data Ingestion (Go)

A high-performance Go application that uses Apache Arrow ADBC to stream data from a MySQL table and ingest it into another table efficiently.

---

##  Features

- Uses Apache Arrow ADBC for fast data transfer
- Streams large datasets with configurable batch sizes
- Efficient ingestion using `IngestStream`
- Environment-based configuration (no hardcoded secrets)

---

##  Tech Stack

- Go
- Apache Arrow ADBC
- MySQL

---

##  Project Structure
```
├── main.go
├── go.mod
├── go.sum
├── .env.example
├── .gitignore
└── README.md
```
---

##  Setup

###  Clone the repository

```bash
git clone https://github.com/your-username/adbc-mysql-ingest.git
cd adbc-mysql-ingest
```
###  Set environment variables

Copy the example file:
```
cp .env.example .env
```
Update .env with your database credentials:
```
DB_URI=username:password@tcp(host:3306)/database
```
### 3. Install dependencies
```
go mod tidy
```
### 4. Run the application
```
go run main.go
```
---

##  Security Notes

- Never commit `.env` files or real credentials  
- Use a non-root database user with limited permissions  
- Rotate credentials regularly  
- Restrict database access by IP if possible  

---

##  Configuration

You can tweak performance using:

- `adbc.statement.batch_size` → controls read batch size  
- `adbc.statement.ingest.batch_size` → controls ingest batch size  

---

## 🧠 Future Improvements

- Add CLI flags for dynamic table names  
- Support multiple databases  
- Add logging & metrics  
- Dockerize the application  
