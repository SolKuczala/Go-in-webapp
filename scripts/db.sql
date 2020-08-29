CREATE DATABASE soku;
USE soku;

CREATE TABLE user_auth(
    ID INT NOT NULL AUTO_INCREMENT,
    email VARCHAR(256) NOT NULL,
    pass PASSWORD(256) NULL,
    google_token VARCHAR(256),

    PRIMARY KEY (ID)
);

CREATE TABLE user_info(
    ID INT NOT NULL AUTO_INCREMENT,
    user_auth_id INT NOT NULL,
    full_name VARCHAR(256),
    address VARCHAR(256),
    phone VARCHAR(256),

    PRIMARY KEY (ID)
    FOREIGN KEY (user_auth_id) REFERENCES user_auth(id)
);

/* create new user (sign up using password) */
INSERT INTO user_auth(email, pass, google_token)
VALUES ("pepe@pepe","123456789", NULL);

/* create new user (sign up using google) */
INSERT INTO user_auth(email, pass, google_token)
VALUES ("pepe@pepe",NULL, "a231d3asd3ass132a46s5d46sad");

/* get user auth for login */
SELECT id, email
FROM user_auth
WHERE email="input email" AND pass="input password";

/* get user profile info after login */
SELECT full_name, address, phone
FROM user_info
WHERE id={EL ID DEL LOGIN};

/* add new user profile info */
INSERT INTO user_info(user_auth_id, full_name, address, phone)
VALUES (1, "pepito papito", "blabla 123", "44444444");

/* update existing user profile info */
/*START TRANSACTION*/
    UPDATE user_info
    set full_name="new nombre",
        address="new address",
        phone="new phone"
    WHERE user_auth_id=1;

    UPDATE user_auth
    set email="new email"
    WHERE id=1;
/*COMMIT*/

/* password reset */
UPDATE user_auth
    set pass="new password"
    WHERE id=1;




