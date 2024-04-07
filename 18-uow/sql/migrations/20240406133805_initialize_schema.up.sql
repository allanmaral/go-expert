CREATE TABLE categories (
    id          int     PRIMARY KEY AUTO_INCREMENT,
    name        text    NOT NULL
);

CREATE TABLE courses (
    id          int     PRIMARY KEY AUTO_INCREMENT,
    name        text    NOT NULL,
    category_id int     NOT NULL,
    FOREIGN KEY (category_id) REFERENCES categories(id)
);

