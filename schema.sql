CREATE TABLE IF NOT EXISTS user (
    user_id int(11) NOT NULL AUTO_INCREMENT,
    user_name varchar(255),
    password varchar(255),
    PRIMARY KEY (user_id)
);

CREATE TABLE IF NOT EXISTS blog (
    blog_id int(11) NOT NULL AUTO_INCREMENT,
    title varchar(255),
    description varchar(255),
    content varchar(255),
    PRIMARY KEY (blog_id)
);

CREATE TABLE IF NOT EXISTS section (
    section_id int(11) NOT NULL AUTO_INCREMENT,
    title varchar(255),
    content varchar(255),
    PRIMARY KEY (section_id)
);
