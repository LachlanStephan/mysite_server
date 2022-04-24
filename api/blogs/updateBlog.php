<?php
    require("../../index.php");
    require("./functions.php");
    require("../utils/sanitize.php");

    $res = [
        'status' => 500,
        'msg' => 'failed',
    ];

    if (!empty($_POST)) {
        $form_data = [
            'content' => ($_POST['content']),
            'description' => sanitizeData($_POST['description']),
            'title' => sanitizeData($_POST['title']),
            'id' => sanitizeData($_POST['db_id']),
        ];
        
        if (updateBlog($form_data)) {
            $res = [
                'status' => 202,
                'msg' => 'success',
            ];
        }
    }
    
    echo json_encode($res);
?>