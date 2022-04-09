<?php
    require("../../index.php");
    require("./functions.php");

    $res = [
        'sections' => [],
        'status' => 500,
    ];

    $sects = getSections();

    if (!empty($sects)) {
        $res['sections'] = $sects;
        $res['status'] = 202;
    }

    echo json_encode($res);

?>