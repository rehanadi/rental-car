-- Create users table
CREATE TABLE users (
  user_id SERIAL PRIMARY KEY,
  name VARCHAR(100),
  email VARCHAR(100) UNIQUE NOT NULL,
  password VARCHAR(255) NOT NULL,
  deposit_amount DECIMAL(10, 2) DEFAULT 0.00
);

-- Create payments table
CREATE TABLE payments (
  payment_id SERIAL PRIMARY KEY,
  user_id INT REFERENCES users(user_id),
  amount DECIMAL(10, 2) NOT NULL,
  payment_method VARCHAR(100) NOT NULL,
  status VARCHAR(100) DEFAULT 'pending',
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create categories table
CREATE TABLE categories (
  category_id SERIAL PRIMARY KEY,
  name VARCHAR(100) UNIQUE NOT NULL,
  description TEXT
);

-- Create cars table
CREATE TABLE cars (
  car_id SERIAL PRIMARY KEY,
  category_id INT REFERENCES categories(category_id),
  name VARCHAR(100) UNIQUE NOT NULL,
  description TEXT,
  rental_cost DECIMAL(10, 2) NOT NULL,
  stock INT NOT NULL
);

-- Create rentals table
CREATE TABLE rentals (
  rental_id SERIAL PRIMARY KEY,
  user_id INT REFERENCES users(user_id),
  car_id INT REFERENCES cars(car_id),
  rental_cost DECIMAL(10, 2) NOT NULL,
  rental_days INT NOT NULL,
  subtotal_cost DECIMAL(10, 2) NOT NULL,
  tax_cost DECIMAL(10, 2) NOT NULL,
  total_cost DECIMAL(10, 2) NOT NULL,
  status VARCHAR(100) DEFAULT 'ongoing',
  rented_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  expired_at TIMESTAMP,
  returned_at TIMESTAMP
);

-- insert into categories table
INSERT INTO categories (name, description) VALUES
('SUV', 'Sport Utility Vehicle'),
('Sedan', 'A car with a closed body (i.e., a fixed roof) which is longer than a hatchback'),
('Truck', 'A motor vehicle designed to transport cargo'),
('Coupe', 'A car with a fixed roof and two doors');

-- insert into cars table with select category name and rupiah currency
INSERT INTO cars (category_id, name, description, rental_cost, stock) VALUES
((SELECT category_id FROM categories WHERE name = 'SUV'), 'Toyota Fortuner', 'Toyota Fortuner is a mid-size SUV', 500000, 5),
((SELECT category_id FROM categories WHERE name = 'SUV'), 'Mitsubishi Pajero', 'Mitsubishi Pajero is a mid-size SUV', 450000, 5),
((SELECT category_id FROM categories WHERE name = 'Sedan'), 'Toyota Camry', 'Toyota Camry is a mid-size sedan', 400000, 5),
((SELECT category_id FROM categories WHERE name = 'Sedan'), 'Honda Civic', 'Honda Civic is a mid-size sedan', 350000, 5),
((SELECT category_id FROM categories WHERE name = 'Truck'), 'Isuzu Elf', 'Isuzu Elf is a light truck', 600000, 5),
((SELECT category_id FROM categories WHERE name = 'Truck'), 'Mitsubishi Colt Diesel', 'Mitsubishi Colt Diesel is a light truck', 550000, 5),
((SELECT category_id FROM categories WHERE name = 'Coupe'), 'Toyota 86', 'Toyota 86 is a sports car', 700000, 5),
((SELECT category_id FROM categories WHERE name = 'Coupe'), 'Mazda MX-5', 'Mazda MX-5 is a sports car', 650000, 5);
