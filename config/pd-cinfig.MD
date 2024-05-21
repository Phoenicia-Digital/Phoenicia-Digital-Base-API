## Phoenicia Digital Project - PostgreSQL Configuration Guide

This guide outlines the recommended PostgreSQL configuration for consistency across Phoenicia Digital projects.

**Recommended Configuration:**

* User: `phoeniciadigital`
* Password: `pdsoftware` (**Please note:** This is a placeholder, use a strong password in production)
* Database: `pd_database`

**Steps:**

1. **Connect to PostgreSQL as the `postgres` user:**
```bash
psql -h localhost -U postgres
```

2. **Create a new user and grant database creation privileges (replace `'pdsoftware'` with your strong password):**
```bash
CREATE ROLE phoeniciadigital WITH LOGIN PASSWORD 'pdsoftware';
ALTER ROLE phoeniciadigital CREATEDB;
```

3. **Verify user creation (optional):**
```bash
\du
```

4. **Exit psql:**
```bash
\q
```

5. **Connect to PostgreSQL as the new user:**
```bash
psql -h localhost -U phoeniciadigital
```

6. **Create a database:**
```bash
CREATE DATABASE pd_database;
```

7. **Verify database creation (optional):**
```bash
\l
```

8. **Connect to the newly created database:**
```bash
\connect pd_database;
```