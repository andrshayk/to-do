   CREATE TABLE tasks (
     id SERIAL PRIMARY KEY,
     title TEXT NOT NULL,
     content TEXT,
     done BOOLEAN DEFAULT FALSE
   );