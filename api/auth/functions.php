<?php

function validatePass($pass) 
{

    global $conn;

    $u = 'Lach';
    $sql = '
        SELECT user_name, password 
        FROM user 
        WHERE user_name = :u
    ';

    $stmt = $conn->prepare($sql);
    $stmt->execute([
        ':u' => $u
    ]);

    $row = $stmt->fetch();
    if ($row['password'] === $pass) {
        return true;
    } else {
        return false;
    }
}

?>