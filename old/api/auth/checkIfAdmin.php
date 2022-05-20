<?php
    require("../../index.php");
    require("./functions.php");

    $res = [
        'status' => 403,
        'msg' => 'invalid',
    ];

    if (!empty($_POST)) {
        $pass = $_POST['pass'];
        if (validatePass($pass)) {
            $res['status'] = 202;
            $res['msg'] = 'valid';
        }
    }

    echo json_encode($res);

?>