CREATE TABLE IF NOT EXISTS courses(
    id SERIAL PRIMARY KEY,
    course_name VARCHAR(50), 
    create_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);