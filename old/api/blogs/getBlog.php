<?php
    require("../../index.php");
    require("./functions.php");

    $res = [
        'blogs' => [],
        'status' => 500,
    ];

    $blogs = getBlogs();

    if (!empty($blogs)) {
        $res['blogs'] = $blogs;
        $res['status'] = 202;
    }

    echo json_encode($res);

?>