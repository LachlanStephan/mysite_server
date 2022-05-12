<?php
    require("../../index.php");
    require("./functions.php");
    require("../utils/sanitize.php");

    $res = [
        'status' => 401,
        'msg' => 'Rejected',
        'read_time' => 0,
    ];

    if(!empty($_POST)) {
        $form_data = [
            'title' => sanitizeData($_POST['title_input_blog']),
            'description' => sanitizeData($_POST['desc_input_blog']),
            'content' => $_POST['content_blog'],
        ];

        $read_time = workOutReadTime($form_data['content']);

        $form_data['title'] = $form_data['title'] . " " . "-" . " " . $read_time;

        if (insertNewBlog($form_data)) {
            $res['status'] = 202;
            $res['msg'] = "Success";
            $res['read_time'] = $read_time;
        }
    }

    echo json_encode($res);
?>