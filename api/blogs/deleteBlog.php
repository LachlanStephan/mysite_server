<?php
    require("../../index.php");
    require("./functions.php");
    require("../utils/sanitize.php");

    $res = [
        'status' => 500,
        'msg' => 'failed',
    ];

    if (!empty($_POST)) {
        $delete_id = sanitizeData($_POST['del_id']);
    }

    if (deleteBlog($delete_id)) {
        $res['status']  = 202;
        $res['msg'] = 'success';
    }

    echo json_encode($res);

