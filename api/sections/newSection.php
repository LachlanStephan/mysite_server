<?php
    require("../../index.php");
    require("./functions.php");
    require("../utils/sanitize.php");

    $res = [
        'status' => 500,
        'msg' => 'rejected'
    ];

    if (!empty($_POST)) {
        $form_data = [
            'title' => sanitizeData($_POST['title_input_section']),
            'description' => sanitizeData($_POST['desc_input_section']),
        ];

        if (insertNewSection($form_data)) {
            $res['status'] = 202;
            $res['msg'] = 'success';
        }
    }
    
    echo json_encode($res);
?>