<?php
    require(__DIR__ . '/env.php');
    require(__DIR__ . '/checkEnv.php');

    header('Access-Control-Allow-Origin: *');
    header('Access-Control-Allow-Methods: GET, POST');
    header("Access-Control-Allow-Headers: X-Requested-With");

    $env = isLocal();

    if ($env['isDev']) {
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