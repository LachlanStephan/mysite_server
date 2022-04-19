<?php 
    $res = [
        'isProd' => false,
        'isDev' => false
    ];

    function isLocal() 
    {
        $local = [
            '127.0.0.1',
            '::1'
        ];

        if (in_array($_SERVER['REMOTE_ADDR'], $local)) {
            $res['isDev'] = true;
        } else {
            $res['isProd'] = true;
        }
    }

    echo json_encode($res);
?>