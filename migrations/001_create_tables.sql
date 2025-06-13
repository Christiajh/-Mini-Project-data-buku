CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  username VARCHAR(100),
  password VARCHAR(255),
  created_at TIMESTAMP DEFAULT now()
);

CREATE TABLE categories (
  id SERIAL PRIMARY KEY,
  name VARCHAR(100),
  created_at TIMESTAMP DEFAULT now()
);

CREATE TABLE books (
  id SERIAL PRIMARY KEY,
  title VARCHAR(255),
  description TEXT,
  image_url TEXT,
  release_year INT CHECK (release_year BETWEEN 1980 AND 2024),
  price INT,
  total_page INT,
  thickness VARCHAR(10),
  category_id INT REFERENCES categories(id),
  created_at TIMESTAMP DEFAULT now()
);
