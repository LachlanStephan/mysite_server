<?php
    require("../../index.php");
    require("./functions.php");
    require("../utils/sanitize.php");

    $res = [
        'msg' => 'failed to send',
        'status' => 500,
    ];

    if (!empty($_POST)) {
        $form_data = [
            'subject' => sanitizeData($_POST['subject']),
            'message' => sanitizeData($_POST['message']),
        ];
        
        if (emailContactMessage($form_data)) {
            $res = [
                'status' => 202,
                'msg' => 'sent',
            ];
        }
    }

    echo json_encode($res);

?>
