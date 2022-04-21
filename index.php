<?php
    header('Access-Control-Allow-Origin: *');
    header('Access-Control-Allow-Methods: GET, POST');
    header("Access-Control-Allow-Headers: X-Requested-With");

    
    function isLocal() 
    {
        $local = [
            '127.0.0.1',
            '::1'
        ];

        if (in_array($_SERVER['REMOTE_ADDR'], $local)) {
            return true;
        } return false;
    }

    if (isLocal()) {
        require(__DIR__ . '/env.php');
        $host = $vars['host'];
        $database = $vars['database'];
        $username = $vars['username'];
        $password = $vars['password'];
    } else {
        $host  = getenv('host');
        $database = getenv('database');
        $username = getenv('username');
        $password = getenv('password');
    }

    try {
        $conn = new PDO("mysql:host=$host;dbname=$database", $username, $password);

        //set attributes 
        $conn->setAttribute(PDO::ATTR_ERRMODE, PDO::ERRMODE_EXCEPTION);
        $conn->setAttribute(PDO::ATTR_EMULATE_PREPARES, false);
    } catch (PDOException $e) {
        $error_message = $e->getMessage();
        echo $error_message;
        exit();
    }

?>