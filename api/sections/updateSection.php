
<?php
    require("../../index.php");
    require("./functions.php");

    $res = [
        'status' => 500,
        'msg' => 'failed',
    ];

    echo json_encode($res);

?>